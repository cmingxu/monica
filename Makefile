PROTOS := $(shell find proto -name *.proto)

gen:
	@for PROTO in $$PROTOS; do
		relative_dir=$(dirname $(PROTO) | sed  "s/\.\/proto//" )
		protoc --go_out=monica/proto_go/$(relative_dir) --proto_path=proto/common  --proto_path=proto $(PROTO)
	done

