.PHONY: lua shared run

all: build-shared-lib build-lua-image

build-shared-lib:
	go build -buildmode=c-shared -o lua/pdf.so pkg/pdf.go
	go tool cgo -exportheader lua/pdf.h pkg/pdf.go

build-lua-image:
	cd lua && docker build -t ebcom/luago .

run:
	docker run --name=lua ebcom/luago && docker cp lua:/app/invoice.pdf . && docker rm lua

clean:
	rm -rf _obj
	docker rm lua


