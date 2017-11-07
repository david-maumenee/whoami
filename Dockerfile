FROM golang:alpine3.6 AS binary
ADD . /app
WORKDIR /app
RUN go build -o http

FROM alpine:3.6
RUN apk --no-cache add curl
WORKDIR /app
ENV PORT 8000
EXPOSE 8000
COPY --from=binary /app/http /app
COPY --from=binary /app/http /app

ADD ./config /app/config
ADD ./data /app/data
ADD ./counter.sh /app


VOLUME /app/config
VOLUME /app/data

CMD ["/app/http"]
