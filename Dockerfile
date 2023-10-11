FROM golang:1.21.1-alpine AS build

WORKDIR /app
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/app /app/app

EXPOSE 8092

CMD ["./app"]
