#!/bin/bash

cd ../../kafka_2.12-2.4.0

echo "STARTING KAFKA ZOOKEEPER"
bin/zookeeper-server-start.sh config/zookeeper.properties &
sleep 2

echo "STARTING KAFKA SERVER"
bin/kafka-server-start.sh config/server.properties &
sleep 2

echo "CREATING KAFKA TOPICS"
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic coinbase-spot
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic binance-spot
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic binance-fut
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic bitmex-perp
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic kraken-spot
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic deribit-perp
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic ftx-perp
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic bitfinex-spot

echo "Active Brokers"
bin/zookeeper-shell.sh localhost:2181 ls /brokers/ids