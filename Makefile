test: generate
	go mod tidy
	go test -v .
	@echo "Done."

generate: tictac.pb.go
.PHONY: generate

tictac.pb.go: tictac.proto
	protoc --go_out=./ --go-grpc_out=./ $<

clean:
	rm -f tictac.pb.goa
