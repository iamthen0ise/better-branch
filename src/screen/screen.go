package screen

import "strings"

const (
	Reset           = "\u001b[0m"
	SavePosition    = "\033[s"
	RestorePosition = "\033[u"
	EraseLineToEnd  = "\033[K"

	Magenta = "\u001b[36m"
	Yellow  = "\u001b[33m"
	//BRIGHT_BLACK = "\u001b[30;1m"
)

func Colorize(text *string, color string) string {
	var sb strings.Builder

	sb.WriteString(color)
	sb.WriteString(*text)
	sb.WriteString(Reset)

	return sb.String()

}
