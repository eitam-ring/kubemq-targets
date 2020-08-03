# Kubemq storage target Connector

Kubemq gcp-storage target connector allows services using kubemq server to access google storage server.

## Prerequisites
The following required to run the gcp-storage target connector:

- kubemq cluster
- gcp-storage set up
- kubemq-source deployment

## Configuration

storage target connector configuration properties:

| Properties Key | Required | Description                                | Example                         |
|:---------------|:---------|:-------------------------------------------|:--------------------------------|
| credentials    | yes      | gcp credentials files                      | "<google json credentials"      |


Example:

```yaml
bindings:
  - name: kubemq-query-gcp-storage
    source:
      kind: source.query
      name: kubemq-query
      properties:
        host: "localhost"
        port: "50000"
        client_id: "kubemq-query-gcp-storage-connector"
        auth_token: ""
        channel: "query.gcp.storage"
        group:   ""
        concurrency: "1"
        auto_reconnect: "true"
        reconnect_interval_seconds: "1"
        max_reconnects: "0"
    target:
      kind: target.gcp.storage
      name: target-gcp-storage
      properties:
        credentials: 'json'
```

## Usage

### Create bucket 

Create bucket metadata settings:

| Metadata Key        | Required | Description                            | Possible values          |
|:--------------------|:---------|:---------------------------------------|:-------------------------|
| method              | yes      | type of method                         | "create_bucket"          |
| bucket              | yes      | bucket name                            | "bucket name"            |
| storage_class       | yes      | gcp-storage_class                      | "storage_class"          |
| project_id          | yes      | gcp storage project_id                 | "<googleurl>/myproject"  |
| location            | yes      | gcp storage valid location             | "gcp-supported locations"|


Example:

```json
{
  "metadata": {
    "method": "create_bucket",
    "bucket": "myBucketName",
    "storage_class": "COLDLINE",
    "project_id": "MyID",
    "location": "us"
  },
  "data": null
}
```


### Upload file 

Upload file metadata settings:

| Metadata Key | Required | Description                             | Possible values                                   |
|:-------------|:---------|:---------------------------------------|:---------------------------------------------------|
| method       | yes      | type of method                         | "upload"                                           |
| bucket       | yes      | bucket name                            | "bucket name"                                      | 
| object       | yes      | object name to save the file under     | "anyString"                                        |
| path         | yes      | path to the file to upload             | "<absolute or relative path to file/filename.type>"|


Example:

```json
{
  "metadata": {
    "method": "upload",
    "bucket": "myBucketName",
    "object": "MyFile",
    "path": "./myFile.yaml"
  },
  "data": null
}
```

### Delete file 

Delete file metadata settings:

| Metadata Key | Required | Description                             | Possible values |
|:-------------|:---------|:----------------------------------------|:----------------|
| method       | yes      | type of method                         | "delete"         |
| bucket       | yes      | bucket name                            | "bucket name"    |  
| object       | yes      | object name                            | "anyString"      |


Example:

```json
{
  "metadata": {
    "method": "delete",
    "bucket": "myBucketName",
    "object": "MyFile"
  },
  "data": null
}
```

### Download file 

Download file metadata settings:

| Metadata Key | Required | Description                            | Possible values |
|:------------ |:---------|:---------------------------------------|:----------------|
| method       | yes      | type of method                         | "download"      |
| bucket       | yes      | bucket name                            | "bucket name"   |  
| object       | yes      | object name                            | "anyString"     |

Example:

```json
{
  "metadata": {
    "method": "download",
    "bucket": "myBucketName",
    "object": "MyFile"
  },
  "data": null
}
```


### Rename file 

Rename file metadata settings:

| Metadata Key | Required | Description                             | Possible values |
|:-------------|:---------|:----------------------------------------|:----------------|
| method                  | yes      | type of method               | "rename"        |
| bucket                  | yes      | bucket name                  | "bucket name"   |  
| object                  | yes      | old object name              | "anyString"     |
| rename_object           | yes      | new object name              | "anyString"     |

Example:

```json
{
  "metadata": {
    "method": "rename",
    "bucket": "myBucketName",
    "object": "MyOldFile",
    "rename_object": "MyNewFile"
  },
  "data": null
}
```


### Copy file 

Copy file metadata settings:

| Metadata Key         | Required | Description                            | Possible values   |
|:---------------------|:---------|:---------------------------------------|:------------------|
| method               | yes      | type of method                         | "copy"            |
| bucket               | yes      | old bucket name                        | "bucket name"     |  
| dst_bucket           | yes      | new bucket name(can be the same)       | "bucket name"     |  
| object               | yes      | old object name                        | "anyString"       |
| rename_object        | yes      | new object name(can be the same)       | "anyString"       |

Example:

```json
{
  "metadata": {
    "method": "copy",
    "bucket": "myOldBucketName",
    "dst_bucket": "myNewBucketName",
    "object": "MyOldFile",
    "rename_object": "MyNewFile"
  },
  "data": null
}
```


### Move file 

Move file metadata settings:

| Metadata Key        | Required | Description                            | Possible values          |
|:--------------------|:---------|:---------------------------------------|:-------------------------|
| method              | yes      | type of method                         | "move"                   |
| bucket              | yes      | old bucket name                        | "bucket name"            |  
| dst_bucket          | yes      | new bucket name(can be the same)       | "bucket name"            |  
| object              | yes      | old object name                        | "anyString"              |
| rename_object       | yes      | new object name(can be the same)       | "anyString"              |

Example:

```json
{
  "metadata": {
    "method": "move",
    "bucket": "myOldBucketName",
    "dst_bucket": "myNewBucketName",
    "object": "MyOldFile",
    "rename_object": "MyNewFile"
  },
  "data": null
}
```

### List files

List files metadata settings:

| Metadata Key | Required | Description                            | Possible values          |
|:-------------|:---------|:---------------------------------------|:-------------------------|
| method       | yes      | type of method                         | "list"                   |
| bucket       | yes      | old bucket name                        | "bucket name"            |  

Example:

```json
{
  "metadata": {
    "method": "list",
    "bucket": "myBucketName"
  },
  "data": null
}
```