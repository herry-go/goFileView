FROM golang:1.12.9-alpine

MAINTAINER herry<herry0412@163.com>

ENV kpdir /go/src/perview

RUN mkdir -p ${kpdir}

ADD . ${kpdir}/

WORKDIR ${kpdir}

RUN go build -v

EXPOSE 9090

ENTRYPOINT ["./perview"]