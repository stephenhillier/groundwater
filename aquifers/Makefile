build:
	protoc -I ../ -I ${GOPATH}/src --go_out=plugins=grpc:$(GOPATH)/src/github.com/stephenhillier/groundwater/ \
	  aquifers/proto/aquifers/aquifers.proto
