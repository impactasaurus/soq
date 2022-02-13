FROM golang:1.17 as source

WORKDIR /go/src/github.com/impactasaurus/soq
COPY . .

FROM source as vendor

RUN go mod download

FROM vendor as runner

RUN go install github.com/impactasaurus/soq/cmd/http
CMD /go/bin/http