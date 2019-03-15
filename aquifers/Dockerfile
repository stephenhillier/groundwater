FROM golang:1.12
ENV GO111MODULE=on
WORKDIR /go/src/github.com/stephenhillier/groundwater/aquifers/
ADD . /go/src/github.com/stephenhillier/groundwater/aquifers/
RUN go install
ENTRYPOINT [ "/go/bin/aquifers" ]
EXPOSE 7778
