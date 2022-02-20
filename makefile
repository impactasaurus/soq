.PHONY: test
test:
	docker build --target source -t soq-unit-tester .
	docker run soq-unit-tester go test -v `go list ./... | grep -v tools`

.PHONY: build
build:
	docker build --target source -t soq-builder .
	docker run -v ${PWD}/build_output/:/output/ -e "GOOS=linux" -e "GOARCH=amd64" soq-builder go build -o /output/lamdba ./cmd/lambda