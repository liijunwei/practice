test:
	go test -v ./...

CommitHash = $(shell git rev-parse --short HEAD)
CommitTime = $(shell git log main -n1 --format='%cd' --date='iso-strict')

demo_pass_parameter_while_build:
	go build -o a.out -ldflags "-X main.CommitHash=$(CommitHash) -X main.CommitTime=$(CommitTime)" cmd/build_example/go_build_demo.go
	./a.out

demo_graph_based_scheduling:
	go build -o a.out cmd/graph_based_scheduling/main.go
	./a.out
