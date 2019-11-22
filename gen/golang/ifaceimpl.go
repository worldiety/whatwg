package golang

import "github.com/worldiety/whatwg/spec"

func (g *Emitter) emitInterfaceImplWasmWrapper(a *Appender, t *spec.Type) {
	implTypeName := lowCase(t.Gen.Name) + "Wrapper"
	this := implTypeName[0:1]

	a.Line("type ", implTypeName, " struct {")
	a.Line("val ", g.Import("syscall/js", "Value"))
	a.Line("}")

	hasConstructor := len(t.Constructors) > 0
	if hasConstructor {
		con := t.Constructors[0]
		a.W("func New"+upCase(t.Gen.Name), "(")
		for _, arg := range con.Arguments {
			a.W(arg.Gen.Name, " ", g.ImportRef(arg.TypeRef), ",")
		}
		a.Line(")", upCase(t.Name), "{")
		a.W("val := ", g.Import("syscall/js", "Global()"), ".New(\"", t.Gen.Name, "\",")
		for _, arg := range con.Arguments {
			a.W(arg.Gen.Name, ",")
		}
		a.Line(")")
		a.Line("return ", implTypeName, "{val:val}")
		a.Line("}")
	}

	a.Line("func Wrap"+upCase(t.Gen.Name), "(val ", g.Import("syscall/js", "Value"), ")", upCase(t.Gen.Name), "{")
	a.Line("return ", implTypeName, "{val:val}")
	a.Line("}")

	for _, attr := range t.Attributes {
		if attr.Read {
			a.Line("func (", this, " ", implTypeName, ")", upCase(attr.Gen.Name), "() ", g.ImportRef(attr.TypeRef), "{")
			a.Line("return ", g.getJSAttrGetter(this, attr.Name, attr.TypeRef))
			a.Line("}")
			a.Line()
		}

		if attr.Write {
			a.Line("func (", this, " ", implTypeName, ") Set", upCase(attr.Gen.Name), "(val ", g.ImportRef(attr.TypeRef), ") ", "{")
			a.Line("}")
			a.Line()
		}
	}
}

func (g *Emitter) getJSAttrGetter(receiverName string, attrName string, ref *spec.TypeRef) string {
	switch ref.Qualifier.Name() {
	case "bool":
		return receiverName + ".val.Get(\"" + attrName + "\").Bool()"
	case "uint16":
		return "uint16(" + receiverName + ".val.Get(\"" + attrName + "\").Int())"
	case "string":
		return receiverName + ".val.Get(\"" + attrName + "\").String()"
	case "int64":
		return "int64(" + receiverName + ".val.Get(\"" + attrName + "\").Int())"
	case "uint64":
		return "uint64(" + receiverName + ".val.Get(\"" + attrName + "\").Int())"
	case "float64":
		return receiverName + ".val.Get(\"" + attrName + "\").Float()"
	default:
		// not correct but expect unknown types to be a custom type
		return "Wrap" + ref.Qualifier.Name() + "(" + receiverName + ".val.Get(\"" + attrName + "\"))"
	}
	return "not implemented"
}
