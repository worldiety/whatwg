#whatwg
Tracks parts of https://github.com/whatwg/ like the DOM API and provides an XML representation to make code generation 
easier. 

Please note that this project is neither developed by, nor endorsed by WHATWG (Apple, Google, Mozilla, Microsoft).
WHATWG and its content is Copyright © 2018 WHATWG (Apple, Google, Mozilla, Microsoft).

## specification
by example
```yaml
- typedef
  
```

```xml
<api>
        <define>
            <doc>An Event object is simply named an event. It allows for
                 signaling that something has occurred, e.g., that an image has completed downloading.</doc>
            <type>interface</type>
            <namespace>dom</namespace>
            <name>Event</name>
            <attribute>
                <doc>historical alias of .stopPropagation</doc>
                <read>true</read>
                <write>true</write>
            </attribute>
        </define>

</api>
```

## license
The derived work is licensed under the same conditions as the original work from the authors (Creative Commons 
Attribution 4.0 International License), as required by the license. Please note, that you have to comply and choose 
a compatible software license when using the original or any transformed data. Remember, that you have to respect the 
attributions.