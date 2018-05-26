
cd /Users/li/go/src/github.com/mygotest/gomicrodemo/demo2/
protoc --proto_path=$GOPATH/src/github.com/mygotest/gomicrodemo/demo2/proto:. --micro_out=. --go_out=. greeter.proto

或者进入
cd /Users/li/go/src/github.com/mygotest/gomicrodemo/demo2/proto

protoc --proto_path=.:. --micro_out=. --go_out=. greeter.proto