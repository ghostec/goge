setup:
	@go get -u github.com/kisielk/godepgraph
	@go get -u github.com/gopherjs/gopherjs

build:
	@cd dist && gopherjs build github.com/ghostec/goge --verbose --minify --localmap --tags js

depgraph:
	@godepgraph -s github.com/ghostec/goge | dot -Tpng -o goge.png

watch:
	@cd dist && gopherjs build github.com/ghostec/goge --watch --verbose --tags js
