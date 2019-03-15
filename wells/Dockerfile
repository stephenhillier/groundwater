FROM golang:1.12
ENV GO111MODULE=on
WORKDIR /go/src/github.com/stephenhillier/groundwater/wells/
ADD . /go/src/github.com/stephenhillier/groundwater/wells/
RUN go install
ENTRYPOINT [ "/go/bin/wells" ]
EXPOSE 7777
