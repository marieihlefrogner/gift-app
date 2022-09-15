# Build backend
FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
RUN mkdir /build
ADD ./backend /build/
WORKDIR /build
RUN go get -d -v
RUN go build -o gift-app .

# Build frontend
FROM node:alpine AS node_builder
RUN mkdir /build
ADD ./frontend /build/
WORKDIR /build
RUN npm install
RUN npm run build

# Combine them
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/ /app/
COPY --from=node_builder /build/build /app/public
WORKDIR /app
CMD ["./gift-app"]
