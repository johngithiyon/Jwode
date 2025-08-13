FROM golang:1.22-alpine

WORKDIR /app

COPY . .

EXPOSE 8001

RUN go build -o file-service

CMD [ "./file-service" ]

