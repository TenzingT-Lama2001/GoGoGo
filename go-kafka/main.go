package main

func main() {

}

/**
	What is Kafka?
	- Kafka is a distributed streaming platform that is used publish and subscribe to streams of records.
	- Kafka is used for building real-time data pipelines and streaming apps.
	- Kafka is horizontally scalable, fault-tolerant, wicked fast, and runs in production in thousands of companies.
	- is anevent streaming platform
	- event indicates a change in state or a change in data or something happenend

	1. Key: uuid
	2. Value: Task was created with XYZ details
	3. Timestamp: time.Now()

	-Kafka stores, publishes, subscribes and processes events

	Why is it popular?

	1. Scalability: The scalability of a system is determined by how well it can maintain its
		performance when exposed to hanges in application and processing demands. Apahe Kafka
		has a distributed architeture capable of handling incoming messages with higher
		volume and veloity. As a resut Kafka is highly scalable without any downtime impact.

	2. High throughput: Apache kafka is able to handle thousands of essages per second. Messages coming
		in at a high volume or a high velocity or both will not affect the performance of kafka.

	3. Low latency: Latency refers to the amount of time taken for a system to processa single event.
		Apache Kafka has a low latency, meaning it is able to process a single event in a very short
		amount of time (low as ten milliseconds).

	4. Fault tolernace: Fault tolerance is the ability of a system to continue operating without any
		downtime even when some of its components fail. Apache Kafka is fault tolerant and is able to
		operate without any downtime even when some of its brokers fail. This is because Kafka replicates
		its messages across multiple brokers. If one broker fails, the other brokers will continue to
		operate and process messages. When the failed broker comes back online, it will be able to
		synchronize with the other brokers and continue processing messages. This is possible because
		Kafka stores messages in a distributed commit log. This commit log is partitioned and replicated
		across multiple brokers.

	5. Reliablity
	6. Durability: Data present on th Kafka luster is allowed to remain persistent more on the cluster
		than on the disk. This ensures that Kafka's data remains durable.

	7. Ability to handle real time data: Kafka supports real-time data handling and is an excellent choice when data
		has to be processed in real-time.


	What is a Kafka cluster?

	- A Kafka cluster is a group of brokers that work together to process and store data.
	- A Kafka cluster can be expanded without downtime. This means that new brokers can be added to the cluster
		and Kafka will automatically start to use them. Similarly, existing brokers can be removed from the cluster
		and Kafka will automatically stop using them.
	- Akafka cluster is a distributed system composed of multiple kafka brokersworking together to handle the storage
		and processing of real-time streaming data,

	WHat is topics?

	- A topic is a category or feed name to which records are published. Topics in Kafka are always multi-subscriber;
		that is, a topic can have zero, one, or many consumers that subscribe to the data written to it.

	- A stream of messages that are a part of a specific category or feed name is referred to as a Kafka topic.
		In kafka, data is stored in the form of topis. Producers write their data to topics and onsumers read the data
		from these topics


	What is Brokers?

	- A Kafka cluster is composed of multiple brokers. A broker is a single server that stores and processes data.
		Each broker may have zero or more partitions per topic. Assuming that we have a topic with a replication factor
		of 3, each partition will have two other replicas. Each partition will have one leader and two followers. The
		leader handles all read and write requests for the partition while the followers passively replicate the leader.
		If the leader fails, one of the followers will automatically become the new leader. Each broker is identified
		with its ID (integer). Each broker has a property called broker.id. The broker.id is required and must be set
		in the broker config file. The broker.id is 0 by default. As mentioned earlier, messages are written to topics
		and topics are partitioned. Each partition is replicated across a configurable number of servers for fault tolerance.
		Each partition has one server which acts as the "leader" and zero or more servers which act as "followers".

	- A kafka luster comprises one or more servers that are known as brokers. In Kafka, a broker works asa a
		container that an hold multiple topics with different partitions. A unique integer ID is used to identify brokers
		in the kafka cluster. COnnection with any one of the kafka brokers in the cluser implies a connection
		with the whole cluster. If there is more than one broker in a cluster, the brokers need not contain the oplete
		data associated with a particular topic.


	Who are onsumers and consumer groups

	- A consumer is a client that reads data from Kafka topics. The data to be read by the consumers has to be pulled from the broker
		when the consumer is ready to receive the message. A consumer group in Kafka referes to a number of
		consumers that pull data from the same topi or same set of topics.


	What is a producer?

	- Producers in Kafka publish messages to one or more topics THey send data to the Kafka cluster. Whenver a
		Kafka producer publishes a message to kafka, the broker receives the message and appends it to a particular partition.
		Producers are given a choice to publish messages to a partition of their choice.


	What is a partition?

	- Topics in  Kafka are divided into a configurable number of aprts, which are known as partitions.
		Partitions allow several consumers to read data from a particualr topic in parllel. Each partition is an ordered,
		Partitions are separated in order. The number of partitions is specified when configuring a topic, but this number an be changed later on.
		The partitions omprising a topi are distributed across servers in the Kafka cluster. Each server in the cluster
		handles data and requests for a share of the partitions. Messages are sent to the broker along with a key.
		The key can be used to determine which partition that particular message will go to.
		All messages whihch have the same key go to the same partition. If the key is not specified, then the partition
		will be decided in a round-robin fashion.


	What is a partition offset?

	- Messages or records in Kafka are assigned to a partition. To specify the position of the records within the
		partition, each reord is provided with an offset. A record can be uniquely identified within its partition using
		the offset value associated with it. A partition offset carries meaning only within that particular partition.
		Older reords  have lower offset values since records are added to the ends of partitions.


	What are replicas?

	- Replias are like backups for partitions in kafka. Theya re used to ensure that there is no data loss i the event of a
		failure or a planned shutdown. Partitions of a topic are published across multiple servers in a Kafka cluster.
		cOpies of the partition are known as Replicas.

	Who is a leader and follower





**/
