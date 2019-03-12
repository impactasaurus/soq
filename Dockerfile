FROM golang:1.12.0

CMD /go/bin/http
WORKDIR /go/src/github.com/impactasaurus/soq

RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.* ./
RUN dep ensure -vendor-only

COPY . .
RUN go install github.com/impactasaurus/soq/cmd/http