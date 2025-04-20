APP_NAME=github.com/CuriousHet/Notify
PROTO_DIR=proto
PROTO_FILES=$(PROTO_DIR)/post.proto

proto:
	protoc --go_out=. --go-grpc_out=. $(PROTO_FILES)

run:
	go run main.go

clean:
	rm -rf proto/postpb

.PHONY: proto run clean
