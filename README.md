# Deploy Multiple VMs with Orka

This project serves to demonstrate how to deploy multiple Orka virtual machines concurrently by wrapping calls to the Orka API using a simple Golang client. This code is in no way intended to be production ready and is for demonstration purposes only. This project may be used as a model for how to integrate calls to the Orka API into CI/CD production pipelines, but is not intended to be used as-is.

## Prerequisites

In order to build this project and try the demo, you will need:

* [Golang v1.17](https://golang.org/dl/) installed
* An [Orka cluster](https://orkadocs.macstadium.com/docs) running >= v1.6.0
* An established [VPN connection](https://orkadocs.macstadium.com/docs/vpn-connect) to the Orka environment from your local machine
* The [Orka CLI](https://orkadocs.macstadium.com/docs/downloads) installed

## How to Build

From the current working directory, simply run the command:

```sh
go build .
```

To build for a specific architecture or OS, set the `GOARCH` or `GOOS` environment variables respectively.

## How to Run

For the purposes of this demo, a token is needed for the client to authenticate to the Orka API. A token may be generated and cached locally using the Orka API:

1. [Configure the Orka CLI](https://orkadocs.macstadium.com/docs/quick-start) by running the command `orka config`
1. [Create a new user](https://orkadocs.macstadium.com/docs/users) and login. Alternatively, you may login with an existing user by running the command `orka login`

When the above requirements have been met, the demo can by run by invoking the binary compiled from the previous build step. On *nix systems (including macOS) this can be done by running the command `./deploy-demo` from the current working directory. On Windows simply run the executable `deploy-demo.exe`.

### Supported Commandline Arguments

The follows flags may be passed when running the demo:

| Flag | Description | Default Value
|---|---|---|
| `--base-image` | base image to use for virtual machine configuration | 90GBigSurSSH.img |
| `--cpu-count` | cpu core count to use for virtual machine configuration | 3 |
| `--deploy` | number of virtual machines to deploy | 4 |

The `--help` or `-h` flag may be passed to view help text.

> NOTE: The base image specified must be present on cluster storage!

### Examples

The following example will deploy 4 virtual machines using the base image `90GBigSurSSH.img`:

```sh
./deploy-demo
```

The following example deploys 3 virtual machines each with 6 vcpu using the custom image `monterey-prod-v1.img`:

```sh
./deploy-demo --deploy=3 --base-image=monterey-prod-v1.img --cpu-count=6
```
