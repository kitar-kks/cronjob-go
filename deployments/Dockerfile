# build stage
FROM golang:alpine AS builder

ENV GO111MODULE=on

# install git.
RUN apk update && apk add --no-cache git

RUN mkdir -p /go/src/xx.com/yyy/smart-schedule
WORKDIR /go/src/xx.com/yyy/smart-schedule
COPY . .

# Using go mod with go 1.11
RUN go mod vendor

# With go ≥ 1.10
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/smart-schedule cmd/smart-schedule/main.go

# small image
FROM alpine:3.7

WORKDIR /app
COPY --from=builder /go/src/xx.com/yyy/smart-schedule/configuration/production.yml /app/configuration/production.yml
COPY --from=builder /go/bin/smart-schedule .

ENV TZ=Asia/Bangkok
RUN echo "Asia/Bangkok" > /etc/timezone

# run binary.
ENTRYPOINT ["/app/smart-schedule", "-env", "production"]