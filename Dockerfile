ARG ALPINE_VERSION=3.20

FROM node:20.18-alpine${ALPINE_VERSION} AS web
WORKDIR /web
COPY web/package*.json ./
RUN npm install -g npm@latest && npm install
COPY web .
RUN npm run build

FROM golang:1.23-alpine${ALPINE_VERSION} AS server
WORKDIR /server
COPY server/go.mod server/go.sum ./
RUN go mod download
COPY server .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o server

FROM alpine:${ALPINE_VERSION} AS app
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
