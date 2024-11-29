FROM node:20.18-alpine3.20 AS web
WORKDIR /web
COPY web .
RUN npm install && npm run build

FROM golang:1.23-alpine3.20 AS server
RUN apk update && apk add gcc musl-dev
WORKDIR /server
COPY server .
RUN CGO_ENABLED=1 go build -ldflags="-w -s" -o server

FROM alpine:3.20 AS app
WORKDIR /app
VOLUME [ "/data" ]
COPY --from=web /web/dist /app/web
COPY --from=server /server/server /app/server
ENV DEBUG=false
ENV WEB_FILES_PATH=/app/web
ENV SERVER_PORT=80
ENV DB_PATH=/data/system.sqlite3
ENV BOT_DB_PATH=/data/bot.sqlite3
EXPOSE ${SERVER_PORT}
ENTRYPOINT ["/app/server"]