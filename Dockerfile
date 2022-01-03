FROM golang:1.17-alpine

ENV META_S_HOST="0.0.0.0"
ENV META_S_PORT="3333"

RUN apk update && apk upgrade &&\
    apk add \
    bash \
    exiftool

WORKDIR /home/server/build

ADD src /home/server/build

RUN go build -o /home/server/meta-server . &&\
    rm -rf /home/server/build

WORKDIR /home/server

CMD [ "sh", "-c", "/home/server/meta-server -h ${META_S_HOST} -p ${META_S_PORT} -release" ]
