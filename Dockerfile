FROM golang:1.16

WORKDIR /app

ADD ./server /app

EXPOSE 3000

RUN go build -o server

CMD [ "/app/server" ]