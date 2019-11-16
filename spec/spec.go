package spec

import (
	"encoding/json"
	"strings"
)

// An API contains everything to know about the DOM, ... APIs
type API struct {
	Root *Namespace // the contained root namespaces. The root namespace has always an empty/ignored name
}

// ToJson serializes into json
func (a *API) ToJson() string {
	b, err := json.Marshal(a)
	if err != nil {
		panic(err) // should not happen because of our closed model
	}
	return string(b)
}

// ToJson serializes into json
func (a *API) String() string {
	b, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		panic(err) // should not happen because of our closed model
	}
	return string(b)
}

// A Qualifier is just a helper type of a string array with some helper functions
type Qualifier []string

// NewQualifier wraps the given string array
func NewQualifier(s ...string) Qualifier {
	return s
}

// Parent returns empty slice or the parent part
func (p Qualifier) Parent() Qualifier {
	if len(p) == 0 {
		return nil
	}
	return p[:len(p)-1]
}

// Qualifier returns the last element or empty string
func (p Qualifier) Name() string {
	if len(p) == 0 {
		return ""
	}
	return p[len(p)-1]
}

func (p Qualifier) NewName(str string) Qualifier {
	return append(p[:len(p)-1], str)
}

// Child concates the name and returns a new Qualifier
func (p Qualifier) Child(name string) Qualifier {
	return append(p, name)
}

// A Namespace contains other namespaces and types
type Namespace struct {
	Name       string       // Qualifier of this namespace
	Types      []*Type      `json:",omitempty"` // Types within this namespace
	Namespaces []*Namespace `json:",omitempty"` // Root namespaces
}

// AddNamespace appends another namespace
func (n *Namespace) AddNamespace(ns *Namespace) {
	n.Namespaces = append(n.Namespaces, ns)
}

// AddType appends another types
func (n *Namespace) AddType(ns *Type) {
	n.Types = append(n.Types, ns)
}

// LookupNamespace returns the denoted namespace or nil. The empty path returns this namespace
func (n *Namespace) LookupNamespace(path []string) *Namespace {
	if n == nil || len(path) == 0 {
		return n
	}
	if len(path) == 0 {
		return nil
	}
	for _, child := range n.Namespaces {
		if child.Name == path[0] {
			return child.LookupNamespace(path[1:])
		}
	}
	return nil
}

// A TypeRef describes the usage of a type in a definition, e.g. as a parameter or generic
type TypeRef struct {
	Qualifier Qualifier  // Qualifier is an array containing the full qualified name by providing all parent namespaces.
	Optional  bool       `json:",omitempty"` // Optional indicates if the nil value is an acceptable value for this type reference
	Generics  []*TypeRef `json:",omitempty"` // Generics contains other type references to narrow the actual reference
}

// NewTypeRef parses the given string and tries to detect also the qualifier path and generics from a blub.Name<a, b> kind of text
func NewTypeRef(str string) *TypeRef {
	var generics []string
	if strings.Contains(str, "<") {
		// a generic spec
		tokens := strings.Split(str, "<")
		str = tokens[0]
		generics = strings.Split(tokens[1][:len(tokens[1])-1], ",")
		for i, g := range generics {
			generics[i] = strings.TrimSpace(g)
		}
	}
	// str is now either blub.Name or just Name
	q := NewQualifier(strings.Split(str, ".")...)
	ref := &TypeRef{
		Qualifier: q,
		Optional:  false,
	}

	for _, g := range generics {
		ref.Generics = append(ref.Generics, NewTypeRef(g))
	}
	return ref
}

func (t *TypeRef) String() string {
	return strings.Join(t.Qualifier, ".")
}

// A Stereotype defines the kind of the type like interface, dictionary, enum etc.
type Stereotype string

const (
	Interface  Stereotype = "interface"
	Dictionary Stereotype = "dictionary"
)

// An Attribute describes a setter/getter contract for a specific field or member
type Attribute struct {
	Description string   `json:",omitempty"` // Description or documentation for this attribute
	Name        string   // Qualifier of the attribute
	TypeRef     *TypeRef // TypeRef refers to the actual type
	Read        bool     // Read flag indicates a read access to the attribute
	Write       bool     // Write flag indicates a write access to the attribute
	Deprecated  bool     `json:",omitempty"` // Deprecated indicates a historical or deprecated attribute which should not be used anymore
}

// A field describes a member of a struct or map like data structure
type Field struct {
	Description string   `json:",omitempty"` // Description or documentation for this attribute
	Name        string   // Qualifier of the attribute
	TypeRef     *TypeRef // TypeRef refers to the actual type
	Deprecated  bool     `json:",omitempty"` // Deprecated indicates a historical or deprecated attribute which should not be used anymore
}

// A Parameter describes a FuncSpec variable
type Parameter struct {
	Description string   `json:",omitempty"` // Description contains details about the Type itself
	Name        string   // Qualifier of the parameter
	TypeRef     *TypeRef // TypeRef refers to the actual type
}

// A FuncSpec describes a function or method
type FuncSpec struct {
	Description string       // Description or documentation for this function
	Name        string       // Name of the function
	Arguments   []*Parameter // Arguments of the function
	Returns     []*Parameter // Returns contains the result parameters
}

// AddReturns appends another parameter as a return parameter
func (f *FuncSpec) AddReturn(p *Parameter) {
	f.Returns = append(f.Returns, p)
}

// AddReturns appends another parameter as a return parameter
func (f *FuncSpec) AddArgument(a *Parameter) {
	f.Arguments = append(f.Arguments, a)
}

// A Type defines a named type or contract
type Type struct {
	Description  string       `json:",omitempty"` // Description contains details about the Type itself
	Stereotype   Stereotype   // Stereotype specifies the kind of type
	Name         string       // Qualifier of this Type in the name
	Attributes   []*Attribute `json:",omitempty"` // Attributes contains all defined attributes for this type
	Fields       []*Field     `json:",omitempty"` // Fields contains all defined fields or members for this type
	Constructors []*FuncSpec  `json:",omitempty"` // Constructors contains all defined function signatures
	FuncSpecs    []*FuncSpec  `json:",omitempty"` // FuncSpecs contains all defined function signatures

}

// AddAttribute appends another attribute
func (t *Type) AddAttribute(a *Attribute) {
	t.Attributes = append(t.Attributes, a)
}

// AddField appends another field
func (t *Type) AddField(f *Field) {
	t.Fields = append(t.Fields, f)
}

// AddFuncSpec appends another signature
func (t *Type) AddFuncSpec(f *FuncSpec) {
	t.FuncSpecs = append(t.FuncSpecs, f)
}

// AddConstructor appends another constructor
func (t *Type) AddConstructor(f *FuncSpec) {
	t.Constructors = append(t.Constructors, f)
}
