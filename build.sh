set -e

cd idl
antlr -Dlanguage=Go -o parser WebIDL


rm -rf whatwg/dom.bs
curl https://raw.githubusercontent.com/whatwg/dom/master/dom.bs --output whatwg/dom.bs

