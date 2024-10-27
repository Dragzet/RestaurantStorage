FROM golang:1.23-alpine AS build

WORKDIR /RestaurantAPI

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o app ./cmd/RestaurantStorage

FROM alpine:latest
WORKDIR /root/

COPY --from=build /RestaurantAPI/app /RestaurantAPI/app
COPY ./logs ./logs
COPY ./config/config.yaml ./config/config.yaml

EXPOSE 8080

CMD ["/RestaurantAPI/app"]