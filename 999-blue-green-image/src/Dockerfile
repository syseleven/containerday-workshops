FROM golang:alpine3.17 AS build
ADD . /app
WORKDIR /app
RUN go mod init whoami_example \
    && go mod tidy \
    && go build -o http

FROM alpine:3.17
WORKDIR /app
ENV PORT 8080
EXPOSE 8080
COPY --from=build /app/http /app
CMD ["/app/http"]
