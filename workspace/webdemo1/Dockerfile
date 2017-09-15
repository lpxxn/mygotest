# Use an official Python runtime as a parent image
FROM golang

# Copy the current directory contents into the container at /app
COPY . /go/src/github.com/lpxxn/webdemo/vendor/github.com/mygotest/workspace/webdemo1/ 

WORKDIR /go/src/github.com/lpxxn/webdemo/vendor/github.com/mygotest/workspace/webdemo1/ 

RUN go build

# Make port 9065 available to the world outside this container
EXPOSE 9065

# Define environment variable
ENV NAME World

# Run app.py when the container launches
CMD ["./webdemo1"]