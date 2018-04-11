##  my kafka location
```
cd /Users/li/go/logs/kafka_2.11-1.1.0
```

## open zookeeper
```
bin/zookeeper-server-start.sh config/zookeeper.properties
```

## open kafka
```
bin/kafka-server-start.sh config/server.properties

```


## create first topic
```
bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test
```


## send msg to test topic
### check topics 
```
bin/kafka-topics.sh --list --zookeeper localhost:2181
```

### publisher 
```
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test
```

## consumer

```
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
```

## check topic status
```
bin/kafka-topics.sh --describe --zookeeper localhost:2181 --topic test
```
