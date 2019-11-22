package idl

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/worldiety/whatwg/idl/parser"
	"github.com/worldiety/whatwg/spec"
	"strings"
)

const debug = true

type apiBuilder struct {
	parser.BaseWebIDLListener
	spec        *spec.API
	namespace   *spec.Namespace // currently on top or nil
	path        spec.Qualifier  // currently on top or nil
	typeDef     *spec.Type      // currently on top or nil
	attr        *spec.Attribute // currently on top or nil
	field       *spec.Field     // currently on top or nil
	function    *spec.FuncSpec  // currently on top or nil
	functionArg *spec.Parameter // currently on top or nil
}

func dbg(str ...interface{}) {
	if !debug {
		return
	}
	fmt.Println(str...)
}

// NewAPIBuilder creates a new builder with an empty API
func NewAPIBuilder() *apiBuilder {
	b := &apiBuilder{
		spec: &spec.API{Root: &spec.Namespace{}},
	}
	b.namespace = b.spec.Root
	return b
}

// Parse takes the string as an WebIDL definition and parses it into a data structure
func Parse(str string) *spec.API {
	// Setup the input
	is := antlr.NewInputStream(str)

	// Create the Lexer
	lexer := parser.NewWebIDLLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewWebIDLParser(stream)

	// Finally parse the expression
	builder := NewAPIBuilder()
	antlr.ParseTreeWalkerDefault.Walk(builder, p.WebIDL())
	return builder.spec
}

// EnterNamespace lookups or creates a namespace, adds it to the api and makes it current
func (s *apiBuilder) EnterNamespace(name string) {
	dbg("EnterNamespace", name)

	s.path = s.path.Child(name)
	ns := s.spec.Root.LookupNamespace(s.path)
	if ns == nil {
		ns = &spec.Namespace{Name: name}
		s.spec.Root.LookupNamespace(s.path.Parent()).AddNamespace(ns)
	}
	s.namespace = ns
}

// ExitNamespace leaves the current namespace and makes the parent
func (s *apiBuilder) ExitNamespace(name string) {
	dbg("ExitNamespace", name)

	if s.path.Name() != name {
		panic("expected " + name + " but got " + s.path.Name())
	}
	s.path = s.path.Parent()
}

func (s *apiBuilder) EnterDictionary(c *parser.DictionaryContext) {
	dbg("EnterDictionary", c.GetText())

	s.typeDef = &spec.Type{}
	s.typeDef.Name = c.IDENTIFIER_WEBIDL().GetText()
	s.typeDef.Stereotype = spec.Dictionary

	s.namespace.AddType(s.typeDef)
}

func (s *apiBuilder) EnterDictionaryMember(ctx *parser.DictionaryMemberContext) {
	dbg("EnterDictionaryMember", ctx.GetText())

	s.field = &spec.Field{}
	s.field.Name = ctx.IDENTIFIER_WEBIDL().GetText()
	s.field.TypeRef = &spec.TypeRef{
		Qualifier: spec.NewQualifier(ctx.Type().GetText()),
		Optional:  false,
		Generics:  nil,
	}
	s.typeDef.AddField(s.field)
}

func (s *apiBuilder) ExitDictionaryMember(ctx *parser.DictionaryMemberContext) {
	dbg("ExitDictionaryMember", ctx.GetText())

	s.field = nil
}

func returnTypeText(ctx parser.IReturnTypeContext) string {
	if ctx == nil {
		return ""
	}
	return ctx.GetText()
}

func (s *apiBuilder) EnterOperation(ctx *parser.OperationContext) {
	dbg("EnterOperation", ctx.GetText())

	s.function = &spec.FuncSpec{}
	s.function.AddReturn(&spec.Parameter{
		Description: "",
		Name:        "",
		TypeRef:     spec.NewTypeRef(returnTypeText(ctx.ReturnType())),
	})
	if returnTypeText(ctx.ReturnType()) == "constructor" {
		s.typeDef.AddConstructor(s.function)
	} else {
		s.typeDef.AddFuncSpec(s.function)
	}
}

func (s *apiBuilder) EnterOperationRest(c *parser.OperationRestContext) {
	dbg("EnterOperationRest", c.GetText())

	s.function.Name = c.OptionalIdentifier().GetText()
}

func (s *apiBuilder) EnterOptionalOrRequiredArgument(c *parser.OptionalOrRequiredArgumentContext) {

	dbg("EnterOptionalOrRequiredArgument", c.GetText())

	s.functionArg = &spec.Parameter{
		Description: "",
		Name:        c.ArgumentName().GetText(),
		TypeRef: &spec.TypeRef{
			Qualifier: spec.NewQualifier(c.Type().GetText()),
			Optional:  false,
			Generics:  nil,
		},
	}
	s.function.AddArgument(s.functionArg)
}

func (s *apiBuilder) EnterPartialInterface(ctx *parser.PartialInterfaceContext) {
	dbg("EnterPartialInterface", ctx.GetText())

	s.typeDef = &spec.Type{}
	s.typeDef.Name = ctx.IDENTIFIER_WEBIDL().GetText()
	s.typeDef.Stereotype = spec.Interface

	s.namespace.AddType(s.typeDef)
}

func (s *apiBuilder) EnterInterface_(ctx *parser.Interface_Context) {
	dbg("EnterInterface_", ctx.GetText())

	s.typeDef = &spec.Type{}
	s.typeDef.Name = ctx.IDENTIFIER_WEBIDL().GetText()
	s.typeDef.Stereotype = spec.Interface

	s.namespace.AddType(s.typeDef)
}

func (s *apiBuilder) ExitInterface_(ctx *parser.Interface_Context) {

}

func (s *apiBuilder) EnterReadWriteAttribute(ctx *parser.ReadWriteAttributeContext) {
	s.attr = &spec.Attribute{}
	s.attr.Read = true
	s.attr.Write = true
	s.typeDef.AddAttribute(s.attr)
}

func (s *apiBuilder) EnterReadonlyMember(c *parser.ReadonlyMemberContext) {
	dbg("EnterReadonlyMember", c.GetText())

	s.attr = &spec.Attribute{}
	s.attr.Read = true
	s.attr.Write = false
	s.typeDef.AddAttribute(s.attr)
}

func (s *apiBuilder) EnterStringifier(ctx *parser.StringifierContext) {
	dbg("EnterStringifier", ctx.GetText())

	s.attr = &spec.Attribute{}
	s.attr.Read = true
	s.attr.Write = false
	s.typeDef.AddAttribute(s.attr)
}

func (s *apiBuilder) ExitReadOnly(c *parser.ReadOnlyContext) {
	dbg("ExitReadOnly", c.GetText())

	//s.attr = nil
}

func (s *apiBuilder) ExitReadWriteAttribute(c *parser.ReadWriteAttributeContext) {
	dbg("ExitReadWriteAttribute", c.GetText())

	//s.attr = nil
}

func (s *apiBuilder) EnterAttributeName(c *parser.AttributeNameContext) {
	dbg("EnterAttributeName", c.GetText())

	s.attr.Name = c.GetText()

}

func (s *apiBuilder) EnterSingleType(ctx *parser.SingleTypeContext) {
	dbg("EnterSingleType", ctx.GetText())

	if s.attr != nil {
		optional := false
		typeName := ctx.GetText()
		if strings.HasSuffix(typeName, "?") {
			optional = true
			typeName = typeName[0 : len(typeName)-1]
		}

		typeRef := &spec.TypeRef{
			Qualifier: spec.NewQualifier(typeName),
			Optional:  optional,
			Generics:  nil,
		}
		s.attr.TypeRef = typeRef
		return
	}
	//fmt.Println("hello ",ctx.GetText())
}
