FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

ADD staticfiles ./staticfiles

ADD go.mod ./go.mod

RUN go mod download

ADD . .

RUN go test ./...

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main /
ADD staticfiles ./staticfiles

ENTRYPOINT ["/main"]
