FROM golang:alpine

COPY . /app

WORKDIR /app

EXPOSE 8080

RUN ["go", "build", "-o", "server", "/app/cmd/server/"]

RUN go test -cover ./...

CMD [ "./server" ]