FROM golang:1.12
ENV GO111MODULE=on
WORKDIR /go/src/github.com/stephenhillier/groundwater/wells/wellsapi
ADD . /go/src/github.com/stephenhillier/groundwater/wells/wellsapi
RUN go install
ENTRYPOINT [ "/go/bin/wellsapi" ]
EXPOSE 8000
