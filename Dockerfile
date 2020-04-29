FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN apk add git
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a -installsuffix cgo \
    -o poc-helm ./cmd/poc-helm-server/main.go

##############

FROM alpine
WORKDIR /app
COPY --from=builder /app/ /app/
ENV POC_HELM_TEST_VAL "TEST via docker env"
ENV POC_HELM_TEST_CONFIG "TEST CONFIG via docker env"
EXPOSE 8080
CMD ["./poc-helm"]
