# Stage 1: Build Stage
FROM golang:1.23.2-alpine3.20 AS build

WORKDIR /src

# build
COPY go.* .
COPY rest-api/ rest-api/

RUN go mod download && go build  -o /api rest-api/cmd/api/main.go

# Stage 2: Run Stage
FROM golang:1.23.2-alpine3.20 AS app
EXPOSE 8080
WORKDIR /app
COPY --from=build /api /api
ENTRYPOINT ["/api"]

