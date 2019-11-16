package golang

import "github.com/worldiety/whatwg/spec"

func (g *Emitter) emitInterfaceImplWasmWrapper(a *Appender, t *spec.Type) {
	implTypeName := lowCase(t.Name) + "Wrapper"
	this := implTypeName[0:1]

	a.Line("type ", implTypeName, " struct {")
	a.Line("val ", g.Import("syscall/js", "Value"))
	a.Line("}")

	hasConstructor := len(t.Constructors) > 0
	if hasConstructor {
		con := t.Constructors[0]
		a.W("func New"+upCase(t.Name), "(")
		for _, arg := range con.Arguments {
			a.W(arg.Name, " ", g.ImportRef(arg.TypeRef), ",")
		}
		a.Line(")", upCase(t.Name), "{")
		a.W("val := ", g.Import("syscall/js", "Global()"), ".New(\"", t.Name, "\",")
		for _, arg := range con.Arguments {
			a.W(arg.Name, ",")
		}
		a.Line(")")
		a.Line("return ", implTypeName, "{val:val}")
		a.Line("}")
	}

	a.Line("func Wrap"+upCase(t.Name), "(val ", g.Import("syscall/js", "Value"), ")", upCase(t.Name), "{")
	a.Line("return ", implTypeName, "{val:val}")
	a.Line("}")

	for _, attr := range t.Attributes {
		if attr.Read {
			a.Line("func (", this, " ", implTypeName, ")", upCase(attr.Name), "() ", g.ImportRef(attr.TypeRef), "{")
			a.Line("}")
			a.Line()
		}

		if attr.Write {
			a.Line("func (", this, " ", implTypeName, ") Set", upCase(attr.Name), "(val ", g.ImportRef(attr.TypeRef), ") ", "{")
			a.Line("}")
			a.Line()
		}
	}
}
