```
go list -e -json -compiled=true -test=false -export=false -deps=true -find=false -tags=wireinject

go list -e -json -compiled=true -test=false -export=false -deps=true -find=false -tags=wireinject -- .

go list -e -json -compiled=true -test=false -export=false -deps=true -find=false -tags=wireinject -- /Users/lipeng/go/src/github.com/mygotest/wire/demo1/

go list -json
```
run
```
GO111MODULE=auto GOPACKAGESDEBUG=true wire
```

goland debug 的时候 
```
GOPACKAGESDEBUG=true;GO111MODULE=auto
```
https://github.com/google/wire
