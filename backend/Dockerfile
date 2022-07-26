FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY users ./users
COPY db ./db

RUN go build -o /goapi

EXPOSE 1234

CMD [ "/goapi" ]