package golang

import (
	"fmt"
	"github.com/worldiety/whatwg/spec"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Emitter struct {
	rootDir    string
	moduleName string
}

func NewEmitter(rootDir string, moduleName string) *Emitter {
	return &Emitter{rootDir: rootDir, moduleName: moduleName}
}

func (g *Emitter) Emit(api *spec.API) error {
	err := g.emitNamespace(filepath.Dir(g.rootDir), api.Root) //one step up, emitNamespaces goes down anyway
	if err != nil {
		return fmt.Errorf("failed to emit api: %w", err)
	}
	return nil
}

func (g *Emitter) emitNamespace(parentDir string, ns *spec.Namespace) error {
	dir := filepath.Join(parentDir, ns.Name)
	a := NewAppender()

	a.Line("// DO NOT EDIT.")
	a.Line("// Generated by github.com/worldiety/whatwg")
	a.Line()
	a.Line("// +build js,wasm")
	a.Line()
	a.Line("package ", ns.Name)
	a.Line()
	a.Line(`import "syscall/js"`)
	a.Line()

	for _, typeDef := range ns.Types {
		g.emitType(a, typeDef)
	}
	for _, cns := range ns.Namespaces {
		err := g.emitNamespace(dir, cns)
		if err != nil {
			return fmt.Errorf("failed to emitNamespace %s: %w", ns.Name, err)
		}
	}

	_ = os.MkdirAll(dir, os.ModePerm)
	err := ioutil.WriteFile(filepath.Join(dir, "impl.go"), []byte(a.String()), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to emitNamespace %s: %w", ns.Name, err)
	}
	return nil
}

func (g *Emitter) Import(modName string, funcOrType string) string {
	switch modName {
	case "syscall/js":
		return "js." + funcOrType
	default:
		panic("TODO")
	}
}

func (g *Emitter) emitType(a *Appender, t *spec.Type) {

	switch t.Stereotype {
	case spec.Interface:
		g.emitInterface(a, t)
		g.emitInterfaceImplWasmWrapper(a, t)
	case spec.Dictionary:
		g.emitStruct(a, t)
	default:
		fmt.Println("emitter not yet implemented for stereotype", t.Stereotype)
	}
	a.Line()
}

func (g *Emitter) ImportRef(t *spec.TypeRef) string {
	return t.String() // TODO
}

type Appender struct {
	sb *strings.Builder
}

func NewAppender() *Appender {
	return &Appender{sb: &strings.Builder{}}
}

func (a *Appender) W(strings ...string) {
	for _, s := range strings {
		a.sb.WriteString(s)
	}
}

func (a *Appender) Line(strings ...string) {
	a.W(strings...)
	a.NewLine()
}

func (a *Appender) NewLine() {
	a.sb.WriteString("\n")
}

func (a *Appender) String() string {
	res, err := format.Source([]byte(a.sb.String()))
	if err != nil {
		fmt.Println(err)
		return a.sb.String()
	}

	return string(res)
}
