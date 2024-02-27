FROM golang:latest as build
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o weatherapp .

FROM scratch
WORKDIR /app
COPY --from=build /app/weatherapp .
ENTRYPOINT ["./weatherapp"]