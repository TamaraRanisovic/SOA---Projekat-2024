FROM golang:alpine as build_container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o usermanagementservice

FROM alpine
COPY --from=build_container /app/usermanagementservice /usr/bin
EXPOSE 8081
ENTRYPOINT ["usermanagementservice"]