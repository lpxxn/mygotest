FROM golang

COPY . /go/src/github.com/mygotest/httpdemo1/httptest4

WORKDIR /go/src/github.com/mygotest/httpdemo1/httptest4

RUN go build

EXPOSE 10001

ENV MYNAME LiPeng

CMD ["./httptest4"]