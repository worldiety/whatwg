package golang

import "github.com/worldiety/whatwg/spec"

func (g *Emitter) emitInterface(a *Appender, t *spec.Type) {
	a.Line("type ", t.Name, " interface {")
	for _, attr := range t.Attributes {
		if attr.Read {
			a.W("// ", upCase(attr.Name), " returns the value of ", attr.Name)
			if attr.TypeRef.Optional {
				a.W(" or nil")
			}
			a.Line(".")
			a.Line(upCase(attr.Name), "() ", g.ImportRef(attr.TypeRef))
			a.Line()
		}

		if attr.Write {
			a.Line("// Set", upCase(attr.Name), " updates the ", attr.Name, " value.")
			if attr.TypeRef.Optional {
				a.Line("// The value may also be nil.")
			}
			a.W("Set"+upCase(attr.Name), "(")
			a.W("value ", g.ImportRef(attr.TypeRef))
			a.Line(")")
			a.Line()
		}

	}
	a.Line("}")
}
