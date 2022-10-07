GRPC_PATH = internal/adapters/framework/left/grpc
PROTO_PATH = $(GRPC_PATH)/proto/*.proto
PROTO_OUT_SERVER = $(GRPC_PATH)/pb

gen-proto:
	protoc \
	--go_out=$(PROTO_OUT_SERVER) \
	--go-grpc_out=$(PROTO_OUT_SERVER) \
	--validaate_out="lang=go:$(PROTO_OUT_SERVER)" \
	$(PROTO_PATH)
