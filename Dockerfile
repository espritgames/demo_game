FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN make build-alpine

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=0 /go/src/app/bin/main .

EXPOSE 8080

CMD ["./main"]
