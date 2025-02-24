package report

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"strings"
	"time"
)

type GoogleChat struct {
	url     string
	payload *Payload
}

func NewGoogleChat(opts ...ReportOption) *GoogleChat {
	gc := &GoogleChat{
		url: os.Getenv("GOOGLE_CHAT_WEBHOOK"),
		payload: &Payload{
			Title: "default",
			Mode:  "dev",
		},
	}

	for _, opt := range opts {
		opt(gc)
	}

	return gc
}

func (gc *GoogleChat) SendRaw(text string) error {
	if gc.url == "" {
		return errors.New("empty Google Chat webhook URL")
	}

	// Prepare the payload
	payload := map[string]string{"text": text}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, gc.url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to send message to Google Chat")
	}

	return nil
}

type Payload struct {
	Title    string         `json:"title"`
	Message  []MessageEntry `json:"message"`
	Mentions []string       `json:"mentions"`
	Tags     []string       `json:"tags"`
	Mode     string         `json:"mode"`
	Icon     string         `json:"icon"`
}

type MessageEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Bold(text string) string {
	return "*" + text + "*"
}

func Italicize(text string) string {
	return "_" + text + "_"
}
func Strikethrough(text string) string {
	return "~" + text + "~"
}

func InlineCodeBlock(text string) string {
	return fmt.Sprintf("`%v`", text)
}

func MultiLineCodeBlock(text string) string {
	return "```" + text + "```"
}

/*
Format option	Markdown example

	Bold	*Text*
	Italicize	_Text_
	Strikethrough	~Text~
	Inline code block	`Text`
	Multi-line code block	```Text```
*/
func (p *Payload) String() string {
	if p.Icon == "" {
		p.Icon = "♣"
	}

	text := p.Icon + " " + Bold(strings.ToTitle(p.Title))
	text += "\n"
	text += time.Now().Local().Format("2006-01-02 03:04:05PM -0700")
	text += "\n"
	text += InlineCodeBlock(p.Mode) + " "
	if len(p.Tags) > 0 {
		list := make([]string, len(p.Tags))
		for i, tag := range p.Tags {
			list[i] = InlineCodeBlock(tag)
		}
		text += strings.Join(list, " ")
	}

	if len(p.Mentions) > 0 {
		text += "\n "

		list := make([]string, len(p.Mentions))
		for i, mn := range p.Mentions {
			list[i] = fmt.Sprintf("<users/%v>", mn)
		}

		text += strings.Join(list, ", ")
	}

	text += "\n"
	for _, entry := range p.Message {
		if entry.Key == "Stack" {
			if len(entry.Value) > 3000 {
				entry.Value = entry.Value[:3000]
			}
		}
		if entry.Key != "" {
			text += Bold(entry.Key) + ": \n"
		}
		if entry.Value != "" {
			text += MultiLineCodeBlock(entry.Value + "\n")
		}
	}

	return text
}

type ReportOption func(*GoogleChat)

func WithPayload(p *Payload) ReportOption {
	return func(c *GoogleChat) {
		c.payload = p
	}
}

func WithIcon(icon string) ReportOption {
	return func(c *GoogleChat) {
		c.payload.Icon = icon
	}
}

func AddTag(tags ...string) ReportOption {
	return func(c *GoogleChat) {
		c.payload.Tags = append(c.payload.Tags, tags...)
	}
}

func AddMessage(key string, value string) ReportOption {
	return func(c *GoogleChat) {
		c.payload.Message = append(c.payload.Message, MessageEntry{
			Key:   key,
			Value: value,
		})
	}
}

func AddMention(mentions ...string) ReportOption {
	return func(c *GoogleChat) {
		c.payload.Mentions = append(c.payload.Mentions, mentions...)
	}
}

func WithScrapUrl() ReportOption {
	url := os.Getenv("GOOGLE_CHAT_SCRAP_WEBHOOK")
	if url == "" {
		url = os.Getenv("GOOGLE_CHAT_WEBHOOK")
	}

	return WithUrl(url)
}

func WithUrl(url string) ReportOption {
	return func(gc *GoogleChat) {
		gc.url = url
	}
}

func WithMode(mode string) ReportOption {
	return func(gc *GoogleChat) {
		gc.payload.Mode = mode
	}
}

func Send(title string, opts ...ReportOption) error {
	gc := NewGoogleChat(opts...)
	gc.payload.Title = title

	if gc.payload.Mode == "dev" {
		return nil
	}

	text := gc.payload.String()

	return gc.SendRaw(text)
}
