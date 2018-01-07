FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/app
COPY . .
RUN make

# runtime image
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/app/app /app
EXPOSE 9999
ENTRYPOINT ["/app"]