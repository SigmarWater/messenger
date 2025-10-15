FROM alpine:3.14

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD *.sql migrations-local/
ADD migrations-local.sh .
ADD local.env .

RUN chmod +x migration-local.sh

ENTRYPOINT ["bash", "migrations-local.sh"]