FROM golang:1.8
MAINTAINER Samuel Cozannet <samnco@gmail.com>

WORKDIR /go/src/app
COPY . .

LABEL "com.example.vendor"="ACME Incorporated"
LABEL name="pva"
LABEL com.example.label-with-value="foo"
LABEL version="1.0"
LABEL description="This text illustrates \
that label-values can span multiple lines."

EXPOSE 8260 

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
