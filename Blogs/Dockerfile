FROM golang:alpine as build_container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o blogsservice

FROM alpine
COPY --from=build_container /app/blogsservice /usr/bin
EXPOSE 8082
ENTRYPOINT ["blogsservice"]