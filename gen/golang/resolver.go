package golang

import (
	"github.com/worldiety/whatwg/spec"
	"strings"
)

type Resolver struct {
	api *spec.API
}

func NewResolve(api *spec.API) *Resolver {
	return &Resolver{api: api}
}

// Resolve modifies and rewrites the given api so that it can be used to generate go code. This involves also
// a rewrite of type names and adds idiomatic things.
func (r *Resolver) Resolve(moduleName string) error {
	names := strings.Split(moduleName, "/")
	r.api.Root.Name = names[len(names)-1]
	r.mapBuildInTypes(r.api.Root)
	return nil
}

func (r *Resolver) mapBuildInTypes(ns *spec.Namespace) {
	for _, typeDef := range ns.Types {
		typeDef.Gen = &spec.Gen{}
		typeDef.Gen.Name = r.mapInvalidIdentifier(typeDef.Name)
		for _, attr := range typeDef.Attributes {
			r.mapBuildInType(attr.TypeRef)
			attr.Gen = &spec.Gen{}
			attr.Gen.Name = r.mapInvalidIdentifier(attr.Name)
		}
		for _, field := range typeDef.Fields {
			field.Gen = &spec.Gen{}
			field.Gen.Name = r.mapInvalidIdentifier(field.Name)
			r.mapBuildInType(field.TypeRef)
		}

		for _, fun := range typeDef.FuncSpecs {
			fun.Gen = &spec.Gen{}
			fun.Gen.Name = r.mapInvalidIdentifier(fun.Name)
			for _, arg := range fun.Arguments {
				arg.Gen = &spec.Gen{}
				arg.Gen.Name = r.mapInvalidIdentifier(arg.Name)

				r.mapBuildInType(arg.TypeRef)
				arg.Name = r.mapInvalidIdentifier(arg.Name)
			}
		}

		for _, con := range typeDef.Constructors {
			con.Gen = &spec.Gen{}
			con.Gen.Name = r.mapInvalidIdentifier(con.Name)
			for _, arg := range con.Arguments {
				arg.Gen = &spec.Gen{}
				arg.Gen.Name = r.mapInvalidIdentifier(arg.Name)

				r.mapBuildInType(arg.TypeRef)
				arg.Name = r.mapInvalidIdentifier(arg.Name)
			}
		}
	}
	for _, c := range ns.Namespaces {
		r.mapBuildInTypes(c)
	}
}

func (r *Resolver) mapBuildInType(t *spec.TypeRef) {
	switch t.Qualifier.Name() {
	case "USVString":
		fallthrough
	case "DOMString":
		t.Qualifier = spec.NewQualifier("string")
	case "void":
		t.Qualifier = nil
	case "constructor":
		t.Qualifier = nil
	case "boolean":
		t.Qualifier = spec.NewQualifier("bool")
	case "unsignedshort":
		t.Qualifier = spec.NewQualifier("uint16")
	case "long":
		t.Qualifier = spec.NewQualifier("int64")
	case "double":
		t.Qualifier = spec.NewQualifier("float64")
	case "unsignedlong":
		t.Qualifier = spec.NewQualifier("uint64")
	}
	for _, g := range t.Generics {
		r.mapBuildInType(g)
	}
}

func (r *Resolver) mapInvalidIdentifier(name string) string {
	switch name {
	case "type":
		return "type_"
	default:
		return name
	}

}
