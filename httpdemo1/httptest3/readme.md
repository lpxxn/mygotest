Normally, we’d build this using: `go build`. But that would build it for the local architecture. To build it for the Raspberry Pi, we use: env `GOOS=linux GOARCH=arm GOARM=7 go build`

That line configures the target OS as Linux, the architecture as `ARM` and the `ARM version as 7`, which is good for the Raspberry Pi 2 and 3 boards. For other versions of the Pi – A, A+, B, B+ or Zero – you’d using `GOARM=6.`

build raspberry 
```
GOOS=linux GOARCH=arm GOARM=7 go build

scp ./httptest3 pi@192.168.11.130:/home/pi/tmp/httptest3
```