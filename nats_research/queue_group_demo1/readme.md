pwd
```
/Users/lipeng/go/src/github.com/nats-io/nats.go/examples
```

queue    
https://docs.nats.io/nats-concepts/queue    
 Using queue subscribers will balance message delivery across a group of subscribers which can be used to provide application fault tolerance and scale workload processing.     
multi terminal 
on client will receive msg
```
go run nats-qsub/main.go foo my-queue
```

multi terminal    
all client receive msg
```
go run nats-sub/main.go foo
```

run 
```
go run nats-pub/main.go foo "Hello NATS Again!"
```
