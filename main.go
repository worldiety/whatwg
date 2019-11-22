package main

import (
	"fmt"
	"github.com/worldiety/whatwg/docfiles"
	"github.com/worldiety/whatwg/gen/golang"
	"github.com/worldiety/whatwg/idl"
	"github.com/worldiety/whatwg/spec"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	apis, err := docfiles.APIs(filepath.Join(dir, "docfiles"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("found %d WebIDL specifications\n", len(apis))

	modName := "github.com/worldiety/dom"

	var singleApi *spec.API
	for i, api := range apis {
		fmt.Println(api.Idl)
		spec := idl.Parse(api.Idl)

		err = golang.NewResolve(spec).Resolve(modName)
		if err != nil {
			panic(err)
		}

		if singleApi == nil {
			singleApi = spec
		} else {
			singleApi.Merge(spec)
		}

		if i==53{
			break
		}
		fmt.Printf("success no %d\n\n", i)
	}

	genDir := filepath.Join(filepath.Dir(dir), "dom")
	emitter := golang.NewEmitter(genDir, modName)
	err = emitter.Emit(singleApi)
	if err != nil {
		panic(err)
	}

}
