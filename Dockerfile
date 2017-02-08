FROM ubuntu:14.04

RUN apt-get update && apt-get install -y \
	golang-go \
	bash \
	curl
ADD . /app
RUN chmod a+x /app/counter.sh
WORKDIR /app
RUN go build -o http
ENV PORT 8000

VOLUME /app/config
VOLUME /app/data

EXPOSE 8000

CMD ["/app/http"]
