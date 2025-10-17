package report

import "context"

func WithPayload(p *Payload) GoogleOption {
	return func(c *GoogleChat) {
		c.payload = p
	}
}

func WithTitle(title string) GoogleOption {
	return func(c *GoogleChat) {
		c.payload.Title = title
	}
}

func WithIcon(icon string) GoogleOption {
	return func(c *GoogleChat) {
		c.payload.Icon = icon
	}
}

func AddTag(tags ...string) GoogleOption {
	return func(c *GoogleChat) {
		c.payload.Tags = append(c.payload.Tags, tags...)
	}
}

func AddMessage(key string, value string) GoogleOption {
	return func(c *GoogleChat) {
		c.payload.Message = append(c.payload.Message, MessageEntry{
			Key:   key,
			Value: value,
		})
	}
}

func AddMention(mentions ...string) GoogleOption {
	return func(c *GoogleChat) {
		c.payload.Mentions = append(c.payload.Mentions, mentions...)
	}
}

func WithUrl(url string) GoogleOption {
	return func(gc *GoogleChat) {
		gc.url = url
	}
}

func WithMode(mode string) GoogleOption {
	return func(gc *GoogleChat) {
		gc.payload.Mode = mode
	}
}

func Send(ctx context.Context, opts ...GoogleOption) error {
	gc := NewGoogleChat(opts...)
	if gc.payload.Mode == "dev" {
		return nil
	}

	text := gc.payload.String()
	return gc.SendRaw(ctx, text)
}
