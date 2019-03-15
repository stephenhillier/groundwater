build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/stephenhillier/groundwater/wells \
	  proto/wells/wells.proto

gw:
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:$(GOPATH)/src/github.com/stephenhillier/groundwater/wells \
		proto/wells/wells.proto

proxy:
	protoc -I/usr/local/include -I. \
  -I$(GOPATH)/src \
  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  proto/wells/wells.proto