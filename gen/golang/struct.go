package golang

import "github.com/worldiety/whatwg/spec"

func (g *Emitter) emitStruct(a *Appender, t *spec.Type) {
	a.Line("type ", t.Name, " struct {")
	for _, field := range t.Fields {
		a.Line(upCase(field.Name), " ", g.ImportRef(field.TypeRef))
	}
	a.Line("}")
}
