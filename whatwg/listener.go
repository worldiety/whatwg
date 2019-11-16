package whatwg

import (
	"fmt"
	"github.com/worldiety/whatwg/idl/parser"
	"go/format"
	"strings"
)

type AttrType int

const (
	None AttrType = iota
	ReadWrite
	Read
)

type myListener struct {
	parser.BaseWebIDLListener
	sb       *strings.Builder
	typ      AttrType
	typeName string
}

func (s *myListener) EnterInterface_(ctx *parser.Interface_Context) {
	s.w("type ", ctx.IDENTIFIER_WEBIDL().GetText(), " interface {\n")
}

func (s *myListener) ExitInterface_(ctx *parser.Interface_Context) {
	s.w("}\n")
}

func (s *myListener) EnterReadWriteAttribute(ctx *parser.ReadWriteAttributeContext) {
	s.typ = ReadWrite
}

func (s *myListener) EnterReadonlyMember(c *parser.ReadonlyMemberContext) {
	s.typ = Read
}
func (s *myListener) ExitReadOnly(c *parser.ReadOnlyContext) {
	s.typ = None
}

func (s *myListener) ExitReadWriteAttribute(c *parser.ReadWriteAttributeContext) {
	s.typ = None
}

func (s *myListener) EnterAttributeName(c *parser.AttributeNameContext) {
	setterName := "Set" + upCase(c.GetText())
	getterName := upCase(c.GetText())

	resolvedTypeName := resolveType(s.typeName)
	switch s.typ {
	case ReadWrite:

		if isNilable(s.typeName) {
			s.w("// ", setterName, " sets the value which may be nil.\n")
		}
		s.w(setterName, "(", strings.ToLower(string(resolvedTypeName[0])), " ", resolvedTypeName, ")\n\n")

		if isNilable(s.typeName) {
			s.w("// ", setterName, " returns the value which may be nil.\n")
		}
		s.w(getterName, "()", resolvedTypeName, "\n\n")
	case Read:
		if isNilable(s.typeName) {
			s.w("// ", setterName, " returns the value which may be nil.\n")
		}
		s.w(getterName, "()", resolvedTypeName, "\n\n")
	}

}

func (s *myListener) EnterSingleType(ctx *parser.SingleTypeContext) {
	s.typeName = ctx.GetText()

}


func (s *myListener) w(str ...string) {
	for _, t := range str {
		s.sb.WriteString(t)
	}
}

func (s *myListener) String() string {
	src, err := format.Source([]byte(s.sb.String()))
	if err != nil {
		fmt.Println(s.sb.String())
		panic(err)
	}

	return string(src)
}
