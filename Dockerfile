FROM golang:1.13-alpine3.10 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o semver cmd/semver/main.go

FROM alpine:3
COPY --from=builder /build/semver /bin/
RUN apk update && apk add --no-cache git bash
CMD ["./bash"]
