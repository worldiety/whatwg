package whatwg

import (
	"bytes"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/worldiety/dom/idl/parser"
	"golang.org/x/net/html"
	"strings"
)

type API struct {
	Idl     string
	IdlNode *html.Node
	Doc     string
}

func parseIdl(str string) {
	// Setup the input
	is := antlr.NewInputStream(str)

	// Create the Lexer
	lexer := parser.NewWebIDLLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewWebIDLParser(stream)

	// Finally parse the expression
	goEmitter := &myListener{sb: &strings.Builder{}}
	antlr.ParseTreeWalkerDefault.Walk(goEmitter, p.WebIDL())
	fmt.Println(goEmitter.String())
}

func Children(n *html.Node) (r []*html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r = append(r, c)
	}
	return r
}

func Text(n *html.Node) string {
	sb := &strings.Builder{}

	Visit(n, func(c *html.Node) bool {
		if c.Type == html.TextNode {
			sb.WriteString(c.Data)
		}
		return true
	})
	return sb.String()
}

func Visit(root *html.Node, cb func(*html.Node) bool) {
	if !cb(root) {
		return
	}
	for _, c := range Children(root) {
		Visit(c, cb)
	}
}

func Visit2(root *html.Node, before func(*html.Node), cb func(*html.Node) bool, after func(*html.Node)) {
	before(root)
	if !cb(root) {
		return
	}
	for _, c := range Children(root) {
		Visit2(c, before, cb, after)
	}
	after(root)
}

func Select(n *html.Node, path string) (r []*html.Node) {
	tokens := strings.Split(path, "/")
	root := n
	for i, token := range tokens {
		last := i == len(tokens)-1
		found := false
		for _, child := range Children(root) {
			if child.Data == token {
				root = child
				found = true
				if last {
					r = append(r, root)
				}
			}
		}
		if !found {
			return
		}
	}

	return r
}

func Attr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func grabDocuNodes(root *html.Node) (r []*html.Node) {
	for root.NextSibling != nil && Attr(root.NextSibling, "class") != "idl" {
		root = root.NextSibling
		r = append(r, root)
	}
	return
}

func Parse(data []byte) []*API {
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	apis := make([]*API, 0)

	for _, c := range Select(doc, "html/body/pre") {
		idl := false
		if Attr(c, "class") == "idl" {
			idl = true
		}
		if idl {
			docuNodes := grabDocuNodes(c)

			api := &API{
				Idl:     Text(c.FirstChild),
				IdlNode: c,
				Doc:     AsGoDoc(docuNodes...),
			}
			apis = append(apis, api)
			fmt.Println("=========")
			fmt.Println(api.Idl)
			//fmt.Println(api.Doc)
			fmt.Println("=======")

			parseIdl(api.Idl)
		}
	}
	return apis
}

// An Event object is simply named an event. It allows for signaling that something has occurred, e.g., that an image has
// completed downloading.
//
// A potential event target is null or an EventTarget object.
//
// An event has an associated target (a potential event target). Unless stated otherwise it is null.
//
// An event has an associated relatedTarget (a potential event target). Unless stated otherwise it is null.
//
// Other specifications use relatedTarget to define a relatedTarget attribute. [[UIEVENTS]]
//
// An event has an associated touch target list (a list of zero or more potential event targets). Unless stated otherwise
// it is the empty list.
//
// The touch target list is for the exclusive use of defining the TouchEvent interface and related interfaces.
// [[TOUCH-EVENTS]]
//
// An event has an associated path. A path is a list of structs. Each struct consists of an invocation target (an
// EventTarget object), an invocation-target-in-shadow-tree (a boolean), a shadow-adjusted target (a potential event
// target), a relatedTarget (a potential event target), a touch target list (a list of potential event targets), a
// root-of-closed-tree (a boolean), and a slot-in-closed-tree (a boolean). A path is initially the empty list.
//
//   event = new Event(type [, eventInitDict])Returns a new event whose  Event.type attribute value is set to  type. The
//   eventInitDict argument  allows for setting the Event.bubbles}} and  {{Event.cancelable attributes via object
//   members of the same name.
//
//   event . Event.typeReturns the type of event, e.g.  "click", "hashchange", or  "submit".
//
//   event . Event.targetReturns the object to which event is dispatched (its  target).
//
//   event . Event.currentTargetReturns the object whose event listener's callback is currently being  invoked.
//
//   event . Event.composedPath()Returns the invocation target objects of event's  path (objects on which listeners will
//   be invoked), except for any  nodes in shadow trees of which the shadow root's  mode is "closed" that are not
//   reachable from event's  Event.currentTarget.
//
//   event . Event.eventPhaseReturns the event's phase, which is one of Event.NONE}},  {{Event.CAPTURING_PHASE}},
//   {{Event.AT_TARGET}}, and  {{Event.BUBBLING_PHASE.
//
//   event . stopPropagation()When dispatched in a tree, invoking this method prevents  event from reaching any objects
//   other than the current object.
//
//   event . stopImmediatePropagation()Invoking this method prevents event from reaching  any registered event listeners
//   after the current one finishes running and, when  dispatched in a tree, also prevents event from reaching any  other
//   objects.
//
//   event . Event.bubblesReturns true or false depending on how event was initialized. True if  event goes through its
//   target's ancestors in reverse  tree order, and false otherwise.
//
//   event . Event.cancelableReturns true or false depending on how event was initialized. Its return  value does not
//   always carry meaning, but true can indicate that part of the operation  during which event was dispatched, can be
//   canceled by invoking the  Event.preventDefault() method.
//
//   event . preventDefault()If invoked when the Event.cancelable attribute value is true, and while executing a
//   listener for the event with AddEventListenerOptions.passive set to false, signals to  the operation that caused
//   event to be dispatched that it needs to be canceled.
//
//   event . Event.defaultPreventedReturns true if Event.preventDefault() was invoked successfully to indicate
//   cancelation,  and false otherwise.
//
//   event . Event.composedReturns true or false depending on how event was initialized. True if  event invokes listeners
//   past a ShadowRoot node that is the  root of its target, and false otherwise.
//
//   event . Event.isTrustedReturns true if event was  dispatched by the user agent, and  false otherwise.
//
//   event . Event.timeStampReturns the event's timestamp as the number of milliseconds measured relative to  the time
//   origin.
//
// The type attribute must return the value it was initialized to. When an event is created the attribute must be
// initialized to the empty string.
//
// The target attribute's getter, when invoked, must return the context object's target.
//
// The srcElement attribute's getter, when invoked, must return the context object's target.
//
// The currentTarget attribute must return the value it was initialized to. When an event is created the attribute must be
// initialized to null.
//
// The composedPath() method, when invoked, must run these steps:
//
//   1. Let composedPath be an empty list.
//
//   2. Let path be the context object's path.
//
//   3. If pathis empty, then return composedPath.
//
//   4. Let currentTarget be the context object's Event.currentTarget  attribute value.
//
//   5. AppendcurrentTarget to composedPath.
//
//   6. Let currentTargetIndex be 0.
//
//   7. Let currentTargetHiddenSubtreeLevel be 0.
//
//   8. Let index be path's size − 1.
//
//   9. While index is greater than or equal to 0:
//
//     1. If path[index]'s root-of-closed-tree is true,    then increase currentTargetHiddenSubtreeLevel by 1.
//
//     2. If path[index]'s invocation target is    currentTarget, then set currentTargetIndex to index and    break.
//
//     3. If path[index]'s slot-in-closed-tree is true,    then decrease currentTargetHiddenSubtreeLevel by 1.
//
//     4. Decrease index by 1.
//
//   10. Let currentHiddenLevel and maxHiddenLevel be  currentTargetHiddenSubtreeLevel.
//
//   11. Set index to currentTargetIndex − 1.
//
//   12. While index is greater than or equal to 0:
//
//     1. If path[index]'s root-of-closed-tree is true,    then increase currentHiddenLevel by 1.
//
//     2. If currentHiddenLevel is less than or equal to maxHiddenLevel, then    prependpath[index]'s    invocation
//     target to composedPath.
//
//     3. If path[index]'s slot-in-closed-tree is true,     then:
//
//       1. Decrease currentHiddenLevel by 1.
//
//       2. If currentHiddenLevel is less than maxHiddenLevel, then set      maxHiddenLevel to currentHiddenLevel.
//
//     4. Decrease index by 1.
//
//   13. Set currentHiddenLevel and maxHiddenLevel to  currentTargetHiddenSubtreeLevel.
//
//   14. Set index to currentTargetIndex + 1.
//
//   15. While index is less than path's size:
//
//     1. If path[index]'s slot-in-closed-tree is true,    then increase currentHiddenLevel by 1.
//
//     2. If currentHiddenLevel is less than or equal to maxHiddenLevel, then    appendpath[index]'s    invocation
//     target to composedPath.
//
//     3. If path[index]'s root-of-closed-tree is true,     then:
//
//       1. Decrease currentHiddenLevel by 1.
//
//       2. If currentHiddenLevel is less than maxHiddenLevel, then set      maxHiddenLevel to currentHiddenLevel.
//
//     4. Increase index by 1.
//
//   16. Return composedPath.
//
//
type event struct {
}
