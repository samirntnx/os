package-converter: tools/package-converter/main.go
	cd tools/package-converter && go build .

rewriter: tools/rewriter/main.go
	cd tools/rewriter && go build .

generate:
	./tools/package-converter/package-converter -package io/ioutil > io/ioutil/ioutil.go
	goimports -w io/ioutil/ioutil.go
	mockgen -source=io/ioutil/ioutil.go -package=ioutil -destination io/ioutil/mock_ioutil.go
	./tools/package-converter/package-converter -package os > fs.go
	goimports -w fs.go
	mockgen -source=fs.go -package=os -destination mock_fs.go
	./tools/package-converter/package-converter -package os/exec > exec/exec.go
	goimports -w exec/exec.go
	mockgen -source=exec/exec.go -package=exec -destination exec/mock_exec.go

