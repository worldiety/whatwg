package idl

import (
	"fmt"
	"testing"
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

func TestBuilder(t *testing.T) {
	spec := Parse(idl0)
	fmt.Println(spec.String())
	if len(spec.Root.Types) != 2 {
		t.Fatalf("expected %d but got %d", 2, len(spec.Root.Types))
	}

	if len(spec.Root.Types[0].Attributes) != 13 {
		t.Fatalf("expected %d but got %d", 13, len(spec.Root.Types[0].Attributes))
	}

	if len(spec.Root.Types[1].Fields) != 3 {
		t.Fatalf("expected %d but got %d", 3, len(spec.Root.Types[1].Fields))
	}
}
