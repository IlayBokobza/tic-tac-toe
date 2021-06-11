FROM golang:1.16

WORKDIR /app

ADD ./server /app

RUN go build -o server

CMD [ "/app/server" ]