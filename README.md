# Service Weaver Performance Testing

This repo is used to test the performance of the Service Weaver framework.

To get started run either `start_trad.sh` or `start_weaver.sh` in the scripts folder.

To benchmark run `bensh.sh trad` or `bench.sh weaver`.

# Prerequisites

## Running the server

You need to have atleast Go 1.20 installed.
Check the instructions here: https://go.dev/doc/install

Quick steps for myself:

```shell
wget https://go.dev/dl/go1.20.4.linux-amd64.tar.gz
```

```shell
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.4.linux-amd64.tar.gz
```

Add to .bashrc:

```shell
export PATH=$PATH:/usr/local/go/bin:~/go/bin
```

You need to install the `weaver-gke-local` command

```shell
go install github.com/ServiceWeaver/weaver-gke/cmd/weaver-gke-local@latest
```

and the `weaver` command

```shell
go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
```

## Running the benchmark

You need Apache Bench (ab) installed.

For debian:

```shell
apt install apache2-utils
```
