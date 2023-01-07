# InfluxDB Web Admin Interface

[![EN doc](https://img.shields.io/badge/document-English-blue.svg)](https://github.com/influxtsdb/influxdb-admin/blob/master/README.md)
[![Go Report Card](https://goreportcard.com/badge/influxtsdb/influxdb-admin)](https://goreportcard.com/report/influxtsdb/influxdb-admin)
[![LICENSE](https://img.shields.io/github/license/influxtsdb/influxdb-admin.svg)](https://github.com/influxtsdb/influxdb-admin/blob/master/LICENSE)
[![Releases](https://img.shields.io/github/v/release/influxtsdb/influxdb-admin.svg)](https://github.com/influxtsdb/influxdb-admin/releases)
![GitHub stars](https://img.shields.io/github/stars/influxtsdb/influxdb-admin.svg?label=github%20stars&logo=github)
[![Docker pulls](https://img.shields.io/docker/pulls/influxtsdb/influxdb-admin.svg)](https://hub.docker.com/r/influxtsdb/influxdb-admin)

This is the built-in admin interface in [InfluxDB v1.2.4](https://github.com/influxdata/influxdb/tree/v1.2.4/services/admin),
which was removed from InfluxDB v1.3 and is a completely independent service from now.

Related official documentation: [Web Admin Interface](https://archive.docs.influxdata.com/influxdb/v1.2/tools/web_admin/)

## Quickstart

### Quickstart by docker

```sh
$ docker run -d --name influxdb-admin -p 8083:8083 influxtsdb/influxdb-admin
```

### Quickstart by pre-built releases

Download one of the [pre-built releases](https://github.com/influxtsdb/influxdb-admin/releases).

```sh
$ ./influxdb-admin
```

### Quickstart by kubernetes & helm chart

Download [InfluxDB Admin Helm chart](https://github.com/influxtsdb/helm-charts/tree/master/charts/influxdb-admin):

```sh
$ helm install influxdb-admin ./influxdb-admin
```

## Usage

```sh
$ ./influxdb-admin -h
Usage of ./influxdb-admin:
  -bind-address string
    	The default bind address used by the admin service (default ":8083")
  -https-certificate string
    	The SSL certificate used when HTTPS is enabled
  -https-enabled
    	Whether the admin service should use HTTPS
```

## Building

Please refer to [README.md](admin/README.md).

## Build release

```sh
$ # build current platform
$ make build
$ # build linux amd64
$ make linux
$ # cross-build all platforms
$ make release
```

## Related projects

- [InfluxDB](https://github.com/influxdata/influxdb/tree/master-1.x): An Open-Source Time Series Database from InfluxData
- [InfluxDB Cluster](https://github.com/chengshiwen/influxdb-cluster): An Open-Source Distributed Time Series Database, Open Source Alternative to InfluxDB Enterprise
- [InfluxDB Proxy](https://github.com/chengshiwen/influx-proxy): A Layer to InfluxDB with High Availability and Consistent Hash
