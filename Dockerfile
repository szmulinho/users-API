FROM golang:1.21.1-alpine AS build

WORKDIR /users
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o users

FROM alpine:latest

WORKDIR /users
COPY --from=build /users/users /users/users

EXPOSE 8092

CMD ["./users"]
