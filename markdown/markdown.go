// package markdown implements a routine to translate markdown document to html
package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	re1 := regexp.MustCompile(`__(.*?)__`)
	re2 := regexp.MustCompile(`_(.*?)_`)
	re3 := regexp.MustCompile(`(#+) (.*)`)
	output := []string{}
	for _, line := range strings.Split(markdown, "\n") {
		line = string(re1.ReplaceAll([]byte(line), []byte("<strong>$1</strong>")))
		line = string(re2.ReplaceAll([]byte(line), []byte("<em>$1</em>")))
		header := re3.FindStringSubmatch(line)
		if len(header) > 0 {
			n := len(header[1])
			output = append(output, fmt.Sprintf("<h%d>%s</h%d>", n, header[2], n))
		} else if line[:2] == "* " {
			n := len(output)
			if len(output) > 0 && output[n-1] == "</ul>" {
				output = output[:n-1]
			} else {
				output = append(output, "<ul>")
			}
			output = append(output, fmt.Sprintf("<li>%s</li>", line[2:]))
			output = append(output, "</ul>")
		} else {
			output = append(output, fmt.Sprintf("<p>%s</p>", line))
		}
	}
	return strings.Join(output, "")
}
