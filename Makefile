setup:
	@go get -u github.com/gopherjs/gopherjs

buildjs:
	@cd dist && gopherjs build github.com/ghostec/goge
