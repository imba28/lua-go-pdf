shared:
	go build -buildmode=c-shared -o lua/pdf.so pkg/pdf.go
	go tool cgo -exportheader lua/pdf.h pkg/pdf.go

lua:
	cd lua && docker build -t ebcom/luago .

run:
	docker run --name=lua ebcom/luago && docker cp lua:/app/invoice.pdf . && docker rm lua

.PHONY: lua