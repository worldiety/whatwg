package whatwg

import (
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

// e.g. {{Event/initEvent()}} or  {{Event}}
var bracesDeclaration = regexp.MustCompile(`{{(\w*\W*)*}}`)

// e.g. \n \r
var br = regexp.MustCompile(`\x{000D}\x{000A}|[\x{000A}\x{000B}\x{000C}\x{000D}\x{0085}\x{2028}\x{2029}]`)

func AsGoDoc(nodes ...*html.Node) string {
	doc := &Paragraph{}
	var root NodeList = doc
	for _, dn := range nodes {
		Visit2(dn, func(node *html.Node) {
			switch node.Type {
			case html.ElementNode:
				var newRoot NodeList
				switch node.Data {
				case "p":
					newRoot = &Paragraph{}
				case "li":
					newRoot = &Paragraph{}
				case "dd":

					newRoot = &Paragraph{indent: true}
				case "ol":

					newRoot = &OrderedList{}
				}
				if newRoot == nil {
					return
				}
				newRoot.SetParent(root)
				root.Add(newRoot)
				root = newRoot
			}
		}, func(node *html.Node) bool {
			switch node.Type {
			case html.TextNode:
				text := formatText(node.Data)
				if len(strings.TrimSpace(text)) > 0 {
					root.Add(&InlineText{Text: text})
				}

			}
			return true
		},
			func(node *html.Node) {
				switch node.Type {
				case html.ElementNode:
					switch node.Data {
					case "p":
						fallthrough
					case "dd":
						fallthrough
					case "li":
						fallthrough
					case "ol":
						root = root.Parent()
					}

				}

			})
	}
	printer := &Printer{}
	printer.SetMaxWidth(120)
	printer.prefix = "// "
	doc.Print(printer)
	return printer.String()
}

func formatText(in string) string {
	return convertBraceNotation(removeBreaks(in))
}

func removeBreaks(in string) string {
	return br.ReplaceAllStringFunc(in, func(s string) string {
		return " "
	})
}

// convertBraceNotation converts {{Event/initEvent()}} things to Event.initEvent()
func convertBraceNotation(in string) string {
	return bracesDeclaration.ReplaceAllStringFunc(in, func(s string) string {
		return strings.ReplaceAll(s, "/", ".")[2 : len(s)-2]
	})
}
