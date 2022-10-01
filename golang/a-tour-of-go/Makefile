.PHONY: interfaces1 interfaces2 interfaces3 fmt

interfaces1: fmt
	@go run interfaces01.go

interfaces2: fmt
	@go run interfaces02.go

interfaces3: fmt
	@go run interfaces03.go

interfaces4: fmt
	@go run interfaces04.go

interfaces5: fmt
	@go run interfaces05.go

interfaces6: fmt
	@go run interfaces06.go

interfaces7: fmt
	@go run interfaces07.go

interfaces8: fmt
	@go run interfaces08.go

interfaces9: fmt
	@go run interfaces09.go

interfaces10: fmt
	@go run interfaces10.go

# interfaces11: fmt
# 	@go run interfaces11.go

goroutines01: fmt
	@go build -o a.out goroutines01.go
	@watch -n2 -d "./a.out"

channels01: fmt
	@go build -o a.out channels01.go
	@./a.out

channels02: fmt
	@go build -o a.out channels02.go
	@./a.out

channels03: fmt
	@go build -o a.out channels03.go
	@./a.out

select01: fmt
	@go build -o a.out select01.go
	@./a.out

select02: fmt
	@go build -o a.out select02.go
	@./a.out

fmt:
	@gofmt -w *.go
