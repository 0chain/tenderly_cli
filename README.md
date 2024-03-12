# NftTracker

This repository contains the implementation of NFT event tracker node, used to scan NFT related events for the upcoming
transaction, scan all NFT related events for the given block range and have a standardized data representation.

## Setup

### Single node launch

It's required to execute the following command to build the node:

```bash
make build
```

After a successful build process, the following command needs to be executed to start a single node:

```shell
make start
```

The following table represents all the parameters that can be redefined using **GNU Make** build configuration file:

| Name | Default value |
-------|--------------|
graphHost  | localhost    |
graphPort | 9096         |
listenerHost  | localhost    |
listenerPort | 9097         |
registratorHost  | localhost    |
registratorPort | 9098         |
databaseHost  | localhost    |
sync | false        |

### Documentation

![start.png](code/go/0chain.net/nft_tracker/docs/start.png)

For detailed guidelines on scalable deployment processes and comprehensive documentation on the internal workflow of the nft_tracker node, please refer to the 0chain project GitBook.

## Tests

To run tests, the appropriate **GNU Make** target must be executed using the following command:

```shell
make test
```
