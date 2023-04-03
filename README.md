# go-kafka

```check&&create topic
kafka-topics --bootstrap-server=localhost:9092 --list
kafka-topics --bootstrap-server=localhost:9092 --topic=sing --create
```

## producer publish ส่งข้อมูล

```producer :: publish

````producer :: publish

```producer :: publish
<!-- kafka-console-producer --broker-list=localhost:9092 --topic=sing -->
kafka-console-producer --bootstrap-server=localhost:9092 --topic=sing
````

## consumer subscribe topic รับข้อมูล

```consumer :: subscribe
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=sing
```

```consumer group
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=sing --gr
oup=mygroup
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=sing --gr
oup=noti
```

```consumer :: subscribe multi topic
kafka-console-consumer --bootstrap-server=localhost:9092 --include="sing|kp" --group=sing
```
