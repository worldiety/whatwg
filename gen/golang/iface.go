package golang

import "github.com/worldiety/whatwg/spec"

func (g *Emitter) emitInterface(a *Appender, t *spec.Type) {
	a.Line("type ", t.Gen.Name, " interface {")
	for _, attr := range t.Attributes {
		if attr.Read {
			a.W("// ", upCase(attr.Gen.Name), " returns the value of ", attr.Gen.Name)
			if attr.TypeRef.Optional {
				a.W(" or nil")
			}
			a.Line(".")
			a.Line(upCase(attr.Gen.Name), "() ", g.ImportRef(attr.TypeRef))
			a.Line()
		}

		if attr.Write {
			a.Line("// Set", upCase(attr.Gen.Name), " updates the ", attr.Gen.Name, " value.")
			if attr.TypeRef.Optional {
				a.Line("// The value may also be nil.")
			}
			a.W("Set"+upCase(attr.Gen.Name), "(")
			a.W("value ", g.ImportRef(attr.TypeRef))
			a.Line(")")
			a.Line()
		}

	}
	a.Line("}")
}
