# kubemq-targets

## Concept

## Supported Targets

### Standalone Services

| Category   | Target         | Kind                         | Configuration                                 |
|:-----------|:---------------|:-----------------------------|:----------------------------------------------|
| Cache      |                |                              |                                               |
|            | Redis          | target.cache.redis           | [Usage](targets/cache/redis/README.md)        |
|            | Memcached      | target.cache.memcached       | [Usage](targets/cache/memcached/README.md)     |
| Stores/db  |                |                              |                                               |
|            | Postgres       | target.stores.postgres       | [Usage](targets/stores/postgres/README.md)     |
|            | Mysql          | target.stores.mysql          | [Usage](targets/stores/mysql/README.md)        |
|            | MSSql          | target.stores.mssql          | [Usage](targets/stores/mssql/README.md)        |
|            | MongoDB        | target.stores.mongodb        | [Usage](targets/stores/mongodb/README.md)      |
|            | Elastic Search | target.stores.elastic-search | [Usage](targets/stores/elastic/README.md)      |
|            | Cassandra      | target.stores.cassandra      | [Usage](targets/stores/cassandra/README.md)    |
|            | Couchbase      | target.stores.couchbase      | [Usage](targets/stores/couchbase/README.md)    |
| Messaging  |                |                              |                                               |
|            | Kafka          | target.messaging.kafka       | [Usage](targets/messaging/kafka/README.md)     |
|            | RabbitMQ       | target.messaging.rabbitmq    | [Usage](targets/messaging/rabbitmq/README.md)  |
|            | MQTT           | target.messaging.mqtt        | [Usage](targets/messaging/mqtt/README.md)      |
|            | ActiveMQ       | target.messaging.activemq    | [Usage](targets/messaging/postgres/README.md)  |
| Storage    |                |                              |                                               |
|            | Minio/S3       | target.storage.minio         | [Usage](targets/storage/minio/README.md)       |
| Serverless |                |                              |                                               |
|            | OpenFaas       | target.serverless.openfaas   | [Usage](targets/serverless/openfass/README.md) |
| Http       |                |                              |                                               |
|            | Http           | target.http                  | [Usage](targets/http/README.md)                 |


### Google Cloud Platform (GCP)

| Category   | Target         | Kind                         | Configuration                                 |
|:-----------|:---------------|:-----------------------------|:----------------------------------------------|
| Cache      |                |                              |                                               |
|            | Redis   https://cloud.google.com/memorystore       | target.gcp.cache.redis           | [Usage](targets/gcp/memorystore/redis/README.md)        |
|            | Memcached  https://cloud.google.com/memorystore    | target.gcp.cache.memcached       | [Usage](targets/gcp/memorystore/memcached/README.md)     |
| Stores/db  |                |                              |                                               |
|            | [Postgres](https://cloud.google.com/sql)       | target.gcp.stores.postgres       | [Usage](targets/gcp/sql/postgres/README.md)     |
|            | [Mysql] (https://cloud.google.com/sql)           | target.gcp.stores.mysql          | [Usage](targets/gcp/sql/mysql/README.md)        |
|            | [BigQuery](https://cloud.google.com/bigquery)      | target.gcp.bigquery        | [Usage](targets/gcp/bigquery/README.md)      |
|            | [BigTable](https://cloud.google.com/bigtable)       | target.gcp.bigtable        | [Usage](targets/gcp/bigtable/README.md)      |
|            | [Firestore](https://cloud.google.com/firestore)     | target.gcp.firestore        | [Usage](targets/gcp/firestore/README.md)      |
|            | [Spanner](https://cloud.google.com/spanner)     | target.gcp.spanner        | [Usage](targets/gcp/spanner/README.md)      |
| Messaging  |                |                              |                                               |
|            | Pub/Sub          | target.gcp.pubsub       | [Usage](targets/gcp/pubsub/README.md)     |
| Storage    |                |                              |                                               |
|            | Storage       | target.gcp.storage         | [Usage](targets/gcp/storage/README.md)       |
| Serverless |                |                              |                                               |
|            | Functions       | target.gcp.cloudfunctions   | [Usage](targets/gcp/cloudfunctions/README.md) |






## Installation


## Configuration


