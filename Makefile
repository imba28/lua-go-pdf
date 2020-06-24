.PHONY: lua shared run

all: build-shared-lib build-lua-image

build-shared-lib:
	go build -buildmode=c-shared -o lua/pdf.so pdf.go
	go tool cgo -exportheader lua/pdf.h pdf.go

build-lua-image:
	cd lua && docker build -t imba28/lua-go-pdf .

run:
	docker run --name=lua-go-pdf imba28/lua-go-pdf && docker cp lua-go-pdf:/app/invoice.pdf . && docker rm lua-go-pdf

clean:
	rm -rf _obj
	docker rm lua-go-pdf

build-dll:
	go build -buildmode=c-archive pdf.go