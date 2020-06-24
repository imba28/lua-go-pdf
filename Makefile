shared:
	go build -buildmode=c-shared -o lua/pdf.so pkg/pdf.go
	go tool cgo -exportheader lua/pdf.h pkg/pdf.go
