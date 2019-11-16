package main

import (
	"github.com/worldiety/whatwg/gen/golang"
	"github.com/worldiety/whatwg/idl"
	"os"
	"path/filepath"
)

const idl0 = `
[Exposed=(Window,Worker,AudioWorklet)]
interface Event {
  constructor(DOMString type, optional EventInit eventInitDict = {});

  readonly attribute DOMString type;
  readonly attribute EventTarget? target;
  readonly attribute EventTarget? srcElement; // historical
  readonly attribute EventTarget? currentTarget;
  sequence<EventTarget> composedPath();

  const unsigned short NONE = 0;
  const unsigned short CAPTURING_PHASE = 1;
  const unsigned short AT_TARGET = 2;
  const unsigned short BUBBLING_PHASE = 3;
  readonly attribute unsigned short eventPhase;

  void stopPropagation();
           attribute boolean cancelBubble; // historical alias of .stopPropagation
  void stopImmediatePropagation();

  readonly attribute boolean bubbles;
  readonly attribute boolean cancelable;
           attribute boolean returnValue;  // historical
  void preventDefault();
  readonly attribute boolean defaultPrevented;
  readonly attribute boolean composed;

  [Unforgeable] readonly attribute boolean isTrusted;
  readonly attribute DOMHighResTimeStamp timeStamp;

  void initEvent(DOMString type, optional boolean bubbles = false, optional boolean cancelable = false); // historical
};

dictionary EventInit {
  boolean bubbles = false;
  boolean cancelable = false;
  boolean composed = false;
};
`

func main() {
	modName := "github.com/worldiety/dom"
	spec := idl.Parse(idl0)
	err := golang.NewResolve(spec).Resolve(modName)
	if err != nil {
		panic(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	genDir := filepath.Join(filepath.Dir(dir), "dom")
	emitter := golang.NewEmitter(genDir, modName)
	err = emitter.Emit(spec)
	if err != nil {
		panic(err)
	}

}
