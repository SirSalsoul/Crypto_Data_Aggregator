#!/bin/bash

cd ../../kafka_2.12-2.4.0

echo "Remove Topics"
bin/kafka-topics.sh --zookeeper 127.0.0.1:2181 --delete --topic coinbase-spot
bin/kafka-topics.sh --zookeeper 127.0.0.1:2181 --delete --topic binance-spot
bin/kafka-topics.sh --zookeeper 127.0.0.1:2181 --delete --topic binance-fut
bin/kafka-topics.sh --zookeeper 127.0.0.1:2181 --delete --topic bitmex-perp
bin/kafka-topics.sh --zookeeper 127.0.0.1:2181 --delete --topic kraken-spot
bin/kafka-topics.sh --zookeeper 127.0.0.1:2181 --delete --topic deribit-perp
bin/kafka-topics.sh --zookeeper 127.0.0.1:2181 --delete --topic ftx-perp
bin/kafka-topics.sh --zookeeper 127.0.0.1:2181 --delete --topic bitfinex-spot

echo "Bring down kafka server"
bin/kafka-server-stop.sh
sleep 2

echo "Bring down zookeeper"
bin/zookeeper-server-stop.sh
sleep 2
