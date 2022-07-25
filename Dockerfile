# syntax=docker/dockerfile:1
# 1)BUILD API
FROM golang:1.18-alpine AS build-go
WORKDIR /app
COPY goapi/go.mod ./
COPY goapi/go.sum ./
RUN go mod download && go mod verify
COPY goapi/*.go ./
COPY goapi/todos ./todos
COPY goapi/users ./users
COPY goapi/db ./db
RUN go build -o /goapi

#2) BUILD AG
FROM node:16-alpine AS build-ag
WORKDIR /usr/src/app
COPY agapp/package.json ./
RUN npm install -g -f yarn && yarn install
COPY agapp ./
RUN yarn build
# 2) FINAL BUILD
# FROM alpine:latest
# RUN apk update && apk add --no-cache supervisor nginx
# COPY supervisord.conf /etc/supervisord.conf
# COPY nginx.conf /etc/nginx/nginx.conf
# COPY --from=build-ag /usr/src/app/dist/mytodo /usr/share/nginx/html
# COPY --from=build-go /goapi /app/
# EXPOSE 80
# EXPOSE 1234
# CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]

FROM nginx:alpine
RUN apk update
WORKDIR /usr/share/nginx/html
COPY --from=build-ag /usr/src/app/dist/mytodo ./
COPY --from=build-go /goapi /app/
COPY nginx.conf /etc/nginx/nginx.conf
COPY commands.sh ./
ENTRYPOINT [ "/bin/sh", "./commands.sh" ]
