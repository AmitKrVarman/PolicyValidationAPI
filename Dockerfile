FROM golang:1.8
MAINTAINER Samuel Cozannet <samnco@gmail.com>

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
