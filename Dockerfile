# syntax=docker/dockerfile:1
# 1)BUILD API
FROM golang:1.18-alpine AS build-go
WORKDIR /app
COPY backend/go.mod ./
COPY backend/go.sum ./
RUN go mod download && go mod verify
COPY backend/*.go ./
COPY backend/todos ./todos
COPY backend/users ./users
COPY backend/db ./db
RUN go build -o /backend

#2) BUILD AG
FROM node:16-alpine AS build-ag
WORKDIR /usr/src/app
COPY fontend/package.json ./
RUN npm install -g -f yarn && yarn install
COPY fontend ./
RUN yarn build


FROM nginx:alpine
RUN apk update
WORKDIR /usr/share/nginx/html
COPY --from=build-ag /usr/src/app/dist/mytodo ./
COPY --from=build-go /backend /app/
COPY nginx.conf /etc/nginx/nginx.conf
COPY commands.sh ./
ENTRYPOINT [ "/bin/sh", "./commands.sh" ]
