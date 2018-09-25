build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/stephenhillier/groundwater/eventstore \
	  proto/events/events.proto
