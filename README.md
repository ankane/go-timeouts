# Go Timeouts

An unresponsive service can be worse than a down one. It can tie up your entire system if not handled properly. **All network requests should have a timeout.**

Here’s how to add timeouts for popular Go packages. **[All have been tested](test)**. The default is no timeout, unless otherwise specified. Enjoy!

Also available for [Ruby](https://github.com/ankane/the-ultimate-guide-to-ruby-timeouts), [Python](https://github.com/ankane/python-timeouts), [Node](https://github.com/ankane/node-timeouts), [PHP](https://github.com/ankane/php-timeouts), and [Rust](https://github.com/ankane/rust-timeouts)

[![Build Status](https://github.com/ankane/go-timeouts/actions/workflows/build.yml/badge.svg)](https://github.com/ankane/go-timeouts/actions)

## Packages

Standard library

- [net](#net)
- [net/http](#nethttp)
- [net/smtp](#netsmtp)
- [os/exec](#osexec)

github.com

- [aws/aws-sdk-go](#awsaws-sdk-go)
- [aws/aws-sdk-go-v2](#awsaws-sdk-go-v2)
- [bradfitz/gomemcache](#bradfitzgomemcache)
- [couchbase/gocb](#couchbasegocb)
- [elastic/go-elasticsearch](#elasticgo-elasticsearch)
- [emersion/go-smtp](#emersiongo-smtp)
- [gocql/gocql](#gocqlgocql)
- [go-pg/pg](#go-pgpg)
- [go-redis/redis](#go-redisredis)
- [go-sql-driver/mysql](#go-sql-drivermysql)
- [opensearch-project/opensearch-go](#opensearch-projectopensearch-go)

Other

- [go.mongodb.org/mongo-driver/mongo](#gomongodborgmongo-drivermongo)

## Standard Library

### net

```go
conn, err := net.DialTimeout(network, address, time.Second)
if err != nil {
    // handle error
}
err = conn.SetDeadline(time.Now().Add(time.Second))
```

### net/http

```go
client := http.Client{
    Timeout: time.Second,
}
```

### net/smtp

```go
conn, err := net.DialTimeout("tcp", address, time.Second)
if err != nil {
    // handle error
}
err = conn.SetDeadline(time.Now().Add(time.Second))
if err != nil {
    // handle error
}
client, err = smtp.NewClient(conn, host)
```

### os/exec

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
err := exec.CommandContext(ctx, cmd).Run()
```

## github.com

### aws/aws-sdk-go

```go
sess := session.Must(session.NewSession(&aws.Config{
    HTTPClient: &http.Client{Timeout: time.Second},
}))
```

### aws/aws-sdk-go-v2

```go
cfg, err := config.LoadDefaultConfig(context.TODO(),
    config.WithHTTPClient(&http.Client{Timeout: time.Second}),
)
```

### bradfitz/gomemcache

```go
mc := memcache.New(host)
mc.Timeout = time.Second
```

### couchbase/gocb

```go
cluster, err := gocb.Connect(connStr, gocb.ClusterOptions{
    TimeoutsConfig: gocb.TimeoutsConfig{
        ConnectTimeout:    time.Second,
        KVTimeout:         time.Second,
        ViewTimeout:       time.Second,
        QueryTimeout:      time.Second,
        AnalyticsTimeout:  time.Second,
        SearchTimeout:     time.Second,
        ManagementTimeout: time.Second,
    },
})
```

See [the docs](https://docs.couchbase.com/go-sdk/current/ref/client-settings.html#timeout-options) for defaults

### elastic/go-elasticsearch

```go
cfg := elasticsearch.Config{
    Transport: &http.Transport{
        DialContext: (&net.Dialer{
            Timeout: time.Second,
        }).DialContext,
        ResponseHeaderTimeout: time.Second,
    },
}
es, err := elasticsearch.NewClient(cfg)
```

### emersion/go-smtp

```go
conn, err := net.DialTimeout("tcp", address, time.Second)
if err != nil {
    // handle error
}
err = conn.SetDeadline(time.Now().Add(time.Second))
if err != nil {
    // handle error
}
client, err = smtp.NewClient(conn, host)
```

### gocql/gocql

```go
cluster := gocql.NewCluster(host)
cluster.ConnectTimeout = time.Second
cluster.Timeout = time.Second
```

Default: 600ms connect timeout, 600ms read timeout

### go-pg/pg

```go
db := pg.Connect(&pg.Options{
    DialTimeout:  time.Second,
    ReadTimeout:  time.Second,
    WriteTimeout: time.Second,
})
```

### go-redis/redis

```go
rdb := redis.NewClient(&redis.Options{
    DialTimeout:  time.Second,
    ReadTimeout:  time.Second,
    WriteTimeout: time.Second,
})
```

### go-sql-driver/mysql

```go
cfg := mysql.Config{
    Timeout:      time.Second,
    ReadTimeout:  time.Second,
    WriteTimeout: time.Second,
}
db, err := sql.Open("mysql", cfg.FormatDSN())
```

### opensearch-project/opensearch-go

```go
cfg := opensearch.Config{
    Transport: &http.Transport{
        DialContext: (&net.Dialer{
            Timeout: time.Second,
        }).DialContext,
        ResponseHeaderTimeout: time.Second,
    },
}
client, err := opensearch.NewClient(cfg)
```

## Other

### go.mongodb.org/mongo-driver/mongo

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
```

## Don’t see a library you use?

[Let us know](https://github.com/ankane/go-timeouts/issues/new). Even better, [create a pull request](https://github.com/ankane/go-timeouts/pulls) for it.

## Running the Tests

```sh
git clone https://github.com/ankane/go-timeouts.git
cd go-timeouts
go mod tidy
```

To run all tests, use:

```sh
go test ./... -v
```

To run individual tests, use:

```sh
go test test/helper_test.go test/net_http_test.go -v
```
