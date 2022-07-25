# 2) BUILD UI
### STAGE 1: Build ###
FROM node:16-alpine AS build-ag
WORKDIR /usr/src/app
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build

### STAGE 2: Run ###
FROM nginx:1.23.0-alpine
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=build-ag /usr/src/app/dist/mytodo /usr/share/nginx/html

# Expose port 80
EXPOSE 80