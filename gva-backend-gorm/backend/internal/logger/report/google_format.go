package report

import "fmt"

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
