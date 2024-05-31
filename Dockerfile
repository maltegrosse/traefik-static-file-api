ARG ARCH=amd64
# build stage
FROM docker.io/golang:1.22 AS builder
RUN mkdir -p /go/src/traefikapi
WORKDIR /go/src/traefikapi
COPY . ./
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$ARCH go build -a -o /app .


# final stage
FROM docker.io/alpine:latest
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /app ./
RUN chmod +x ./app
ENTRYPOINT ["./app"]
EXPOSE 8080