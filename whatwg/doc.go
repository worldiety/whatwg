package whatwg

import (
	"strconv"
	"strings"
)

type Printer struct {
	lines        []string
	indent       int
	maxWidth     int
	prefix       string
	indentPrefix string
}

func (p *Printer) IndentRight() {
	p.indent += 2
}

func (p *Printer) IndentLeft() {
	p.indent -= 2
}

func (p *Printer) String() string {
	sb := &strings.Builder{}
	for _, line := range p.lines {
		sb.WriteString(p.prefix)
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	return sb.String()
}

func (p *Printer) SetMaxWidth(w int) {
	p.maxWidth = w
}

func (p *Printer) Write(str ...string) {
	for _, s := range str {
		lines := strings.Split(s, "\n")
		for _, line := range lines {
			p.Append(line)
		}
	}
}

func (p *Printer) Append(str string) {
	if len(p.lines) == 0 {
		p.lines = append(p.lines, "")
	}
	lastLine := p.lines[len(p.lines)-1]
	if strings.HasSuffix(lastLine, " ") && strings.HasPrefix(str, " ") {
		str = trimPrefixSpace(str)
	}
	lastLine += str
	p.lines = p.lines[0 : len(p.lines)-1]
	maxLen := p.maxWidth - p.indent
	indent := &strings.Builder{}
	for i := 0; i < p.indent; i++ {
		indent.WriteString(" ")
	}
	strIndent := indent.String() // the indent as white space

	for len(lastLine) > maxLen {
		if !strings.HasPrefix(lastLine, strIndent) {
			lastLine = strIndent + lastLine //ensure lastLine always starts with the correct indent
		}

		// now break and insert lines as required
		breakWord := maxLen
		for i := maxLen; i >= 0; i-- {
			if lastLine[i] == ' ' {
				breakWord = i
				break
			}
		}
		filledLine := lastLine[0:breakWord]
		nextLine := trimPrefixSpace(lastLine[breakWord:])
		p.lines = append(p.lines, filledLine)
		lastLine = nextLine
	}
	if !strings.HasPrefix(lastLine, strIndent) {
		lastLine = strIndent + lastLine //ensure lastLine always starts with the correct indent
	}
	p.lines = append(p.lines, lastLine)
}

func trimPrefixSpace(s string) string {
	for strings.HasPrefix(s, " ") {
		s = s[1:]
	}
	return s
}

func (p *Printer) NewLine() {
	p.lines = append(p.lines, "")
	if len(p.lines) >= 3 { //ever have a single empty line
		eof := len(p.lines) - 1
		if isWhitespace(p.lines[eof]) && isWhitespace(p.lines[eof-1]) && isWhitespace(p.lines[eof-2]) {
			p.lines = p.lines[0:eof]
		}
	}
}

func isWhitespace(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

type Node interface {
	Parent() NodeList
	SetParent(list NodeList)
	Print(sb *Printer)
}

type NodeList interface {
	Node
	Add(node Node)
}

type Doc struct {
	Nodes  []Node
	parent NodeList
}

func (d *Doc) Parent() NodeList {
	return d.parent
}

func (d *Doc) SetParent(list NodeList) {
	d.parent = list
}

func (d *Doc) Add(node Node) {
	d.Nodes = append(d.Nodes, node)
}

func (d *Doc) Print(sb *Printer) {
	for _, n := range d.Nodes {
		n.Print(sb)
	}
}

type InlineText struct {
	Text   string
	parent NodeList
}

func (p *InlineText) String() string {
	return p.Text
}

func (p *InlineText) Parent() NodeList {
	return p.parent
}

func (p *InlineText) SetParent(list NodeList) {
	p.parent = list
}

func (p *InlineText) Print(sb *Printer) {
	sb.Write(p.Text)
}

type Paragraph struct {
	Nodes  []Node
	parent NodeList
	indent bool
}

func (p *Paragraph) Parent() NodeList {
	return p.parent
}

func (p *Paragraph) SetParent(list NodeList) {
	p.parent = list
}

func (p *Paragraph) Print(sb *Printer) {

	if p.indent {
		sb.IndentRight()
	}
	for _, n := range p.Nodes {
		n.Print(sb)
	}
	if p.indent {
		sb.IndentLeft()
	}
	sb.NewLine()
	sb.NewLine()
}

func (p *Paragraph) flatten() (r []Node) {
	for _, node := range p.Nodes {
		if otherP, ok := node.(*Paragraph); ok {
			r = append(r, otherP.flatten()...)
		} else {
			r = append(r, node)
		}
	}
	return
}

func (p *Paragraph) Add(node Node) {
	if len(p.Nodes) > 0 {
		last, okLast := p.Nodes[len(p.Nodes)-1].(*InlineText)
		this, okThis := node.(*InlineText)
		if okLast && okThis {
			last.Text = trimSuffixSpace(last.Text)
			this.Text = trimPrefixSpace(this.Text)
			if len(this.Text) == 0 {
				return
			}
			if hasPunctuationPrefix(this.Text) {
				last.Text = last.Text + this.Text
				return
			}

			if hasPunctuationPrefix(this.Text) && hasLowerCasePrefix(this.Text) {
				last.Text = last.Text + this.Text
				return
			}

			last.Text = last.Text + " " + this.Text
			return
		}
	}
	p.Nodes = append(p.Nodes, node)
}

func hasLowerCasePrefix(s string) bool {
	return len(s) > 0 && s[0] >= 'a' && s[0] <= 'z'
}

func hasPunctuationPrefix(s string) bool {
	p := []string{".", ",", ":", ")"}
	for _, c := range p {
		if strings.HasPrefix(s, c) {
			return true
		}
	}
	return false
}

func trimSuffixSpace(s string) string {
	for strings.HasSuffix(s, " ") {
		s = s[:len(s)-1]
	}
	return s
}

type Heading struct {
	Text string
}

func (p *Heading) Print(sb *Printer) {
	sb.Write(p.Text)
}

type OrderedList struct {
	Items  []Node
	parent NodeList
}

func (d *OrderedList) Parent() NodeList {
	return d.parent
}

func (d *OrderedList) SetParent(list NodeList) {
	d.parent = list
}

func (d *OrderedList) Add(node Node) {
	d.Items = append(d.Items, node)
}

func (d *OrderedList) Print(sb *Printer) {
	for i, n := range d.Items {
		sb.IndentRight()
		sb.Write(strconv.Itoa(i+1), ". ")
		n.Print(sb)
		sb.IndentLeft()
	}
}

type UnorderedList struct {
	Items  []Node
	parent NodeList
}

func (d *UnorderedList) Parent() NodeList {
	return d.parent
}

func (d *UnorderedList) SetParent(list NodeList) {
	d.parent = list
}

func (d *UnorderedList) Add(node Node) {
	d.Items = append(d.Items, node)
}

func (d *UnorderedList) Print(sb *Printer) {
	for _, n := range d.Items {
		sb.IndentRight()
		sb.Write("* ")
		n.Print(sb)
		sb.IndentLeft()
	}
}
