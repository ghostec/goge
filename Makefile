setup:
	@go get -u github.com/kisielk/godepgraph
	@go get -u github.com/gopherjs/gopherjs

buildjs:
	@cd dist && gopherjs build github.com/ghostec/goge

depgraph:
	@godepgraph -s github.com/ghostec/goge | dot -Tpng -o goge.png
