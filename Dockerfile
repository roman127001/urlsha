FROM golang:1.22.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/urlsha_app cmd/*.go

FROM alpine:3.19
RUN apk --no-cache add ca-certificates

RUN addgroup -S userapp && adduser -S userapp -G userapp
COPY --from=builder --chown=userapp /app/urlsha_app .
USER userapp

EXPOSE 8080
CMD ["./urlsha_app"]
