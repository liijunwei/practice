defer01: format clean
	@go build -o a.out defer01.go
	@./a.out | sort > tmp.md && cat tmp.md

clean:
	@rm -rf tmp*

format:
	@gofmt -w *.go
