

cd idl
antlr -Dlanguage=Go -o parser WebIDL

set -e
cd ..
#rm -rf whatwg/dom.bs
#curl https://raw.githubusercontent.com/whatwg/dom/master/dom.bs --output whatwg/dom.bs

rm -rf docfiles/html.bs
curl https://raw.githubusercontent.com/whatwg/html/master/source --output docfiles/html.bs