FROM golang:1.21.1-alpine3.18 as build
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.18
RUN apk add bash
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/development.env .
COPY --from=build /app/migrate ./migrate
COPY --from=build /app/db/migration ./migration
COPY --from=build /app/start.sh .
COPY --from=build /app/wait-for.sh .




EXPOSE 8888
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]