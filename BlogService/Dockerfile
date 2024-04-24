FROM golang:alpine as build_container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o blogservice

FROM alpine
COPY --from=build_container /app/blogservice /usr/bin
COPY --from=build_container /app/static /app/static
EXPOSE 8084
ENTRYPOINT ["blogservice"]
