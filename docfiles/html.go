package docfiles

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

type API struct {
	Idl     string
	IdlNode *html.Node
	Doc     string
}

func Children(n *html.Node) (r []*html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r = append(r, c)
	}
	return r
}

func Text(n *html.Node) string {
	sb := &strings.Builder{}

	Visit(n, func(c *html.Node) bool {
		if c.Type == html.TextNode {
			sb.WriteString(c.Data)
		}
		return true
	})
	return sb.String()
}

func Visit(root *html.Node, cb func(*html.Node) bool) {
	if !cb(root) {
		return
	}
	for _, c := range Children(root) {
		Visit(c, cb)
	}
}

func Visit2(root *html.Node, before func(*html.Node), cb func(*html.Node) bool, after func(*html.Node)) {
	before(root)
	if !cb(root) {
		return
	}
	for _, c := range Children(root) {
		Visit2(c, before, cb, after)
	}
	after(root)
}

func Select(n *html.Node, path string) (r []*html.Node) {
	tokens := strings.Split(path, "/")
	root := n
	for i, token := range tokens {
		last := i == len(tokens)-1
		found := false
		for _, child := range Children(root) {
			if child.Data == token {
				root = child
				found = true
				if last {
					r = append(r, root)
				}
			}
		}
		if !found {
			return
		}
	}

	return r
}

func Attr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func grabDocuNodes(root *html.Node) (r []*html.Node) {
	for root.NextSibling != nil && Attr(root.NextSibling, "class") != "idl" {
		root = root.NextSibling
		r = append(r, root)
	}
	return
}

func parseBikeshed(data []byte) ([]*API, error) {
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	apis := make([]*API, 0)

	Visit(doc, func(node *html.Node) bool {
		if Attr(node, "class") == "idl" {
			_ = grabDocuNodes(node)

			api := &API{
				Idl:     Text(node),
				IdlNode: node,
			}
			apis = append(apis, api)
		}
		return true
	})

	return apis, nil
}
