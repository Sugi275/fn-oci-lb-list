FROM golang:1.11.5 as builder
WORKDIR /go/src/funcdir
COPY Gopkg.toml .
COPY src .
RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o func .

FROM alpine:3.9
WORKDIR /function/
RUN apk update && \
	apk add --no-cache ca-certificates
COPY --from=builder /go/src/funcdir/func .
ENTRYPOINT ["/function/func"]
