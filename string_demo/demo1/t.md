go tool compile -N -l -S main.go  
go tool objdump main.o 

go build -gcflags -S main.go
 GOOS=linux GOARCH=amd64 go tool compile -S main.go   
GOSSAFUNC=hello go build main.go  

dlv debug main.go  
b main.main
b runtime.rawstring
b string.go:262
b string.go:322

b proc.go:559
bp
clear 2

b string.go:322
bp
condition 2 size==9


b main.main
b runtime.gostringnocopy
b string.go:476
