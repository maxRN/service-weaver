# Service Weaver Performance Testing

This repo is used to test the performance of the Service Weaver framework.

To get started run either `start_trad.sh` or `start_weaver.sh` in the scripts folder.

To benchmark run `bensh.sh trad` or `bench.sh weaver`.

## Prerequisites

You need to have atleast Go 1.20 installed.
You need to install the `weaver-gke-local` command.
```shell
go install github.com/ServiceWeaver/weaver-gke/cmd/weaver-gke-local@latest
```
