FROM golang

# Copy the current directory contents into the container at /app
COPY . /go/src/github.com/mygotest/workspace/webdemo2/
WORKDIR /go/src/github.com/mygotest/workspace/webdemo2/

RUN go build

EXPOSE 9100 9101

ENV NAME World

# Run app.py when the container launches
CMD ["./webdemo2"]

# docker build . -t web2:1
# docker run -p 9100:9100 -p 9101:9101 -it -d -v /crmlogs:/go/src/github.com/mygotest/workspace/webdemo2/logs/ a5cd