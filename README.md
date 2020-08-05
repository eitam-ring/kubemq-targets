# KubeMQ Targets

KubeMQ Targets connects KubeMQ Message Broker with external systems and cloud services.

KubeMQ Targets allows to build a message-base microservices architecture on Kubernetes with minimal efforts and without developing connectivity interfaces between KubeMQ Message Broker and external systems such databases, cache, messaging and REST-base APIs.

**Key Features**:

- **Runs anywhere**  - Kubernetes, Cloud, on-prem , anywhere
- **Stand-alone** - small docker container / binary
- **Single Interface** - One interface all the services
- **Any Service** - Support all major services types (databases, cache, messaging, serverless, HTTP etc)
- **Plug-in Architecture** Easy to extend, easy to connect
- **Middleware Supports** - Logs, Metrics, Retries and Rate Limiters
- **Easy Configuration** - simple yaml file builds your topology

## Concepts

KubeMQ Targets building blocks are:
 - Binding
 - Source
 - Target
 - Request/Response


### Binding

Binding is a 1:1 connection between Source and Target. Every Binding runs independently.


![binding](.github/assets/binding.jpeg)

### Target

Target is an external service which expose an API allowing to interact and serve his functionalists with other services.

Targets can be Cache systems such as Redis and Memcached, SQL Databases such Postgres and MySql and event an HTTP generic Rest interface.

KubeMQ Targets integrate each one of the supported targets and serve requests based on the request data.

A list of supported targets is below.

#### Standalone Services

| Category   | Target                                                              | Kind                         | Configuration                                  |
|:-----------|:--------------------------------------------------------------------|:-----------------------------|:-----------------------------------------------|
| Cache      |                                                                     |                              |                                                |
|            | [Redis](https://redis.io/)                                          | target.cache.redis           | [Usage](targets/cache/redis)         |
|            | [Memcached](https://memcached.org/)                                 | target.cache.memcached       | [Usage](targets/cache/memcached)     |
| Stores/db  |                                                                     |                              |                                                |
|            | [Postgres](https://www.postgresql.org/)                             | target.stores.postgres       | [Usage](targets/stores/postgres)     |
|            | [Mysql](https://www.mysql.com/)                                     | target.stores.mysql          | [Usage](targets/stores/mysql)        |
|            | [MSSql](https://www.microsoft.com/en-us/sql-server/sql-server-2019) | target.stores.mssql          | [Usage](targets/stores/mssql)        |
|            | [MongoDB](https://www.mongodb.com/)                                 | target.stores.mongodb        | [Usage](targets/stores/mongodb)      |
|            | [Elastic Search](https://www.elastic.co/)                           | target.stores.elastic-search | [Usage](targets/stores/elastic)      |
|            | [Cassandra](https://cassandra.apache.org/)                          | target.stores.cassandra      | [Usage](targets/stores/cassandra)    |
|            | [Couchbase](https://www.couchbase.com/)                             | target.stores.couchbase      | [Usage](targets/stores/couchbase)    |
| Messaging  |                                                                     |                              |                                                |
|            | [Kafka](https://kafka.apache.org/)                                  | target.messaging.kafka       | [Usage](targets/messaging/kafka)     |
|            | [RabbitMQ](https://www.rabbitmq.com/)                               | target.messaging.rabbitmq    | [Usage](targets/messaging/rabbitmq)  |
|            | [MQTT](http://mqtt.org/)                                            | target.messaging.mqtt        | [Usage](targets/messaging/mqtt)      |
|            | [ActiveMQ](http://activemq.apache.org/)                             | target.messaging.activemq    | [Usage](targets/messaging/postgres)  |
| Storage    |                                                                     |                              |                                                |
|            | [Minio/S3](https://min.io/)                                         | target.storage.minio         | [Usage](targets/storage/minio)       |
| Serverless |                                                                     |                              |                                                |
|            | [OpenFaas](https://www.openfaas.com/)                               | target.serverless.openfaas   | [Usage](targets/serverless/openfass) |
| Http       |                                                                     |                              |                                                |
|            | Http                                                                | target.http                  | [Usage](targets/http)                |



#### Google Cloud Platform (GCP)

| Category   | Target                                                              | Kind                       | Configuration                                        |
|:-----------|:--------------------------------------------------------------------|:---------------------------|:-----------------------------------------------------|
| Cache      |                                                                     |                            |                                                      |
|            | [Redis](https://cloud.google.com/memorystore)                       | target.gcp.cache.redis     | [Usage](targets/gcp/memorystore/redis)     |
|            | [Memcached](https://cloud.google.com/memorystore)                   | target.gcp.cache.memcached | [Usage](targets/gcp/memorystore/memcached) |
| Stores/db  |                                                                     |                            |                                                      |
|            | [Postgres](https://cloud.google.com/sql)                            | target.gcp.stores.postgres | [Usage](targets/gcp/sql/postgres)          |
|            | [Mysql](https://cloud.google.com/sql)                               | target.gcp.stores.mysql    | [Usage](targets/gcp/sql/mysql)             |
|            | [BigQuery](https://cloud.google.com/bigquery)                       | target.gcp.bigquery        | [Usage](targets/gcp/bigquery)              |
|            | [BigTable](https://cloud.google.com/bigtable)                       | target.gcp.bigtable        | [Usage](targets/gcp/bigtable)              |
|            | [Firestore](https://cloud.google.com/firestore)                     | target.gcp.firestore       | [Usage](targets/gcp/firestore)             |
|            | [Spanner](https://cloud.google.com/spanner)                         | target.gcp.spanner         | [Usage](targets/gcp/spanner)               |
|            | [Firebase](https://firebase.google.com/products/realtime-database/) | target.gcp.firebase        | [Usage](targets/gcp/firebase)              |
| Messaging  |                                                                     |                            |                                                      |
|            | [Pub/Sub](https://cloud.google.com/pubsub)                          | target.gcp.pubsub          | [Usage](targets/gcp/pubsub)                |
| Storage    |                                                                     |                            |                                                      |
|            | [Storage](https://cloud.google.com/storage)                         | target.gcp.storage         | [Usage](targets/gcp/storage)               |
| Serverless |                                                                     |                            |                                                      |
|            | [Functions](https://cloud.google.com/functions)                     | target.gcp.cloudfunctions  | [Usage](targets/gcp/cloudfunctions)        |

#### Amazon Web Service (AWS)


#### Microsoft Azure
WIP

### Source

Source is a KubeMQ connection (in subscription mode) which listen to requests from services and route them to the appropriate target for action, and return back a response if needed.

KubeMQ Targets supports all KubeMQ's messaging patterns: Queue, Events, Events-Store, Command and Query


| Type                                                                              | Kind                | Configuration                           |
|:----------------------------------------------------------------------------------|:--------------------|:----------------------------------------|
| [Queue](https://docs.kubemq.io/learn/message-patterns/queue)                      | source.queue        | [Usage](sources/queue/README.md)        |
| [Events](https://docs.kubemq.io/learn/message-patterns/pubsub#events)             | source.events       | [Usage](sources/events/README.md)       |
| [Events Store](https://docs.kubemq.io/learn/message-patterns/pubsub#events-store) | source.events-store | [Usage](sources/events-store/README.md) |
| [Command](https://docs.kubemq.io/learn/message-patterns/rpc#commands)             | source.command      | [Usage](sources/command/README.md)      |
| [Query](https://docs.kubemq.io/learn/message-patterns/rpc#queries)                | source.query        | [Usage](sources/query/README.md)        |


### Request / Response

![concept](.github/assets/concept.jpeg)

#### Request

Request is an object that send to a designated target with metadata and data fields which contains the needed information to perform the requested data.

##### Request Object Structure

| Field  | Type | Description                |
|:-------|:---------|:---------------------------|
| metadata | string,string object      | contains metadata information for action           |
| data  | bytes array      | contains raw data for action |

##### Exmaple

Request to get a data from Redis cache for the key "log"
```json
{
  "metadata": {
    "method": "get",
    "key": "log"
  },
  "data": null
}
```
#### Response
Response is an object that send back as a result of executing an action in the target


##### Response Object Structure

| Field    | Type                 | Description                                     |
|:---------|:---------------------|:------------------------------------------------|
| metadata | string,string object | contains metadata information result for action |
| data     | bytes array          | contains raw data result                        |
| is_error | bool                 | indicate if the action ended with an error      |
| error    | string               | contains error information if any               |


##### Example

Response received on request to get the data stored in Redis for key "log"
```json
{
  "metadata": {
    "result": "ok",
    "key": "log"
  },
  "data": "SU5TRVJUIElOVE8gcG9zdChJRCxUSVRMRSxDT05URU5UKSBWQUxVRVMKCSAgICAgICAgICAgICAgICAgICAgICA"
}
```

## Installation

### Kubernetes

An example of kubernetes deployment for redis target connectors can be find below:

1. Run Redis Cluster deployment yaml

```bash
kubectl apply -f ./redis-example.yaml -n kubemq
```

2. Run KubeMQ Targets deployment yaml

```bash
kubectl apply -f ./deployment-example.yaml
```

### Binary (Cross-platform)

Download the appropriate version for your platform from KubeMQ Targets Releases. Once downloaded, the binary can be run from anywhere.

Ideally, you should install it somewhere in your PATH for easy use. /usr/local/bin is the most probable location.

Running KubeMQ Targets

```bash
kubemq-targets --config config.yaml
```


## Configuration

### Structure

Config file structure:

```yaml

apiPort: 8080 # kubemq targets api and health end-point port
bindings:
  - name: clusters-sources # unique binding name
    properties: # Bindings properties such middleware configurations
      log_level: error
      retry_attempts: 3
      retry_delay_milliseconds: 1000
      retry_max_jitter_milliseconds: 100
      retry_delay_type: "back-off"
      rate_per_second: 100
    source:
      kind: source.query # source kind
      name: name-of-sources # source name 
      properties: # a set of key/value settings per each source kind
        .....
    target:
      kind: target.cache.redis # target kind
      name: name-of-target # targets name
      properties: # a set of key/value settings per each target kind
        - .....
```

### Properties

In bindings configuration, KubeMQ targets supports properties setting for each pair of source and target bindings.

These properties contain middleware information settings as follows:

#### Logs Middleware

KubeMQ targets supports level based logging to console according to as follows:

| Property  | Description       | Possible Values        |
|:----------|:------------------|:-----------------------|
| log_level | log level setting | "debug","info","error" |
|           |                   |  "" - indicate no logging on this bindings |

An example for only error level log to console:

```yaml
bindings:
  - name: sample-binding 
    properties: 
      log_level: error
    source:
    ......  
```

#### Retry Middleware

KubeMQ targets supports Retries' target execution before reporting of error back to the source on failed execution.

Retry middleware settings values:


| Property                      | Description                                           | Possible Values                             |
|:------------------------------|:------------------------------------------------------|:--------------------------------------------|
| retry_attempts                | how many retries before giving up on target execution | default - 1, or any int number              |
| retry_delay_milliseconds      | how long to wait between retries in milliseconds      | default - 100ms or any int number           |
| retry_max_jitter_milliseconds | max delay jitter between retries                      | default - 100ms or any int number           |
| retry_delay_type              | type of retry delay                                   | "back-off" - delay increase on each attempt |
|                               |                                                       | "fixed" - fixed time delay                  |
|                               |                                                       | "random" - random time delay                |

An example for 3 retries with back-off strategy:

```yaml
bindings:
  - name: sample-binding 
    properties: 
      retry_attempts: 3
      retry_delay_milliseconds: 1000
      retry_max_jitter_milliseconds: 100
      retry_delay_type: "back-off"
    source:
    ......  
```

#### Rate Limiter Middleware

KubeMQ targets supports Rate Limiting of target executions.

Rate Limiter middleware settings values:


| Property        | Description                                    | Possible Values                |
|:----------------|:-----------------------------------------------|:-------------------------------|
| rate_per_second | how many executions per second will be allowed | 0 - no limitation              |
|                 |                                                | 1 - n integer times per second |

An example for 100 executions per second:

```yaml
bindings:
  - name: sample-binding 
    properties: 
      rate_per_second: 100
    source:
    ......  
```

### Source

Source section contains source configuration for binding as follows:

| Property    | Description                                       | Possible Values                                               |
|:------------|:--------------------------------------------------|:--------------------------------------------------------------|
| name        | sources name (will show up in logs)               | string without white spaces                                   |
| kind        | source kind type                                  | source.queue                                                  |
|             |                                                   | source.query                                                  |
|             |                                                   | source.command                                                |
|             |                                                   | source.events                                                 |
|             |                                                   | source.events-store                                           |
| properties | an array of keu/value setting for source connection| see above               |


### Target

Target section contains target configuration for binding as follows:

| Property    | Description                                       | Possible Values                                               |
|:------------|:--------------------------------------------------|:--------------------------------------------------------------|
| name        | targets name (will show up in logs)               | string without white spaces                                   |
| kind        | source kind type                                  | target.type-of-target                                                  |
| properties | an array of keu/value setting for target connection | see above              |





