# Kubemq Redis Target Connector

Kubemq redis target connector allows services using kubemq server to access redis server functions such `set`, `get` and `delete`.

## Prerequisites
The following are required to run the redis target connector:

- kubemq cluster
- redis v5.0.0 (or later)
- kubemq-target-connectors deployment

## Configuration

Redis target connector configuration properties:

| Properties Key | Required | Description                  | Example          |
|:---------------|:---------|:-----------------------------|:-----------------|
| host           | yes      | redis address                | "localhost:6379" |
| password       | no       | redis server password        | ""               |
| enable_tls     | no       | connect to redis in tls mode | "false"          |

Example:

```yaml
bindings:
  - name: kubemq-query-redis
    source:
      kind: source.kubemq.query
      name: kubemq-query
      properties:
        host: "localhost"
        port: "50000"
        client_id: "kubemq-query-redis-connector"
        auth_token: ""
        channel: "query.redis"
        group:   ""
        concurrency: "1"
        auto_reconnect: "true"
        reconnect_interval_seconds: "1"
        max_reconnects: "0"
    target:
      kind: target.cache.redis
      name: target-redis
      properties:
        host: "localhost:6379"
        password: ""
        enable_tls: "false"
```

## Usage

### Get Request

Get request metadata setting:

| Metadata Key | Required | Description      | Possible values |
|:-------------|:---------|:-----------------|:----------------|
| key          | yes      | redis key string | any string      |
| method       | yes      | get              | "get"           |


Example:

```json
{
  "metadata": {
    "key": "your-redis-key",
    "method": "get"
  },
  "data": null
}
```

### Set Request

Set request metadata setting:

| Metadata Key | Required | Description      | Possible values |
|:-------------|:---------|:-----------------|:----------------|
| key          | yes      | redis key string | any string      |
| method       | yes      | set              | "set"           |
| etag         | no       | set etag version | "0"             |
| concurrency  | no       | set concurrency  | ""              |
|              |          |                  | "first-write"   |
|              |          |                  | "last-write"    |
|              |          |                  |                 |
| consistency  | no       | set consistency  | ""              |
|              |          |                  | "strong"        |
|              |          |                  | "eventual"      |

Set request data setting:

| Data Key | Required | Description                   | Possible values     |
|:---------|:---------|:------------------------------|:--------------------|
| data     | yes      | data to set for the redis key | base64 bytes array |

Example:

```json
{
  "metadata": {
    "key": "your-redis-key",
    "method": "set",
    "etag": "0",
    "concurrency": "",
    "consistency": ""
  },
  "data": "c29tZS1kYXRh" 
}
```
### Delete Request

Delete request metadata setting:

| Metadata Key | Required | Description      | Possible values |
|:-------------|:---------|:-----------------|:----------------|
| key          | yes      | redis key string | any string      |
| method       | yes      | delete           | "delete"        |


Example:

```json
{
  "metadata": {
    "key": "your-redis-key",
    "method": "delete"
  },
  "data": null
}
```