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
		for _, attr := range typeDef.Attributes {
			r.mapBuildInType(attr.TypeRef)
			attr.Name = r.mapInvalidIdentifier(attr.Name)
		}
		for _, field := range typeDef.Fields {
			r.mapBuildInType(field.TypeRef)
		}

		for _, fun := range typeDef.FuncSpecs {
			fun.Name = r.mapInvalidIdentifier(fun.Name)
			for _, arg := range fun.Arguments {
				r.mapBuildInType(arg.TypeRef)
				arg.Name = r.mapInvalidIdentifier(arg.Name)
			}
		}

		for _, con := range typeDef.Constructors {
			con.Name = r.mapInvalidIdentifier(con.Name)
			for _, arg := range con.Arguments {
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
