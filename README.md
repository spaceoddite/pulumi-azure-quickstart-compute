# Azure-Quickstart-Compute Pulumi Component Provider (Go)





An example of using the `Virtual Machine` component in TypeScript is in `examples/simple`.

Note that the generated provider plugin (`pulumi-resource-azure-quickstart-compute`) must be on your `PATH` to be used by Pulumi deployments. 

$ export PATH=$PATH:$PWD/bin

If creating a provider for distribution to other users, you should ensure they install this plugin to their `PATH`.

## Prerequisites

- Go 1.16
- Pulumi CLI
- Node.js (to build the Node.js SDK)
- Yarn (to build the Node.js SDK)
- Python 3.6+ (to build the Python SDK)
- .NET Core SDK (to build the .NET SDK)

## Build and Test

```bash
# Build and install the provider (plugin copied to $GOPATH/bin)
make install_provider

# Regenerate SDKs
make generate

# Test Node.js SDK
$ make install_nodejs_sdk
$ cd examples/simple
$ yarn install
$ yarn link @pulumi/xyz
$ pulumi stack init test
$ pulumi config set aws:region us-east-1
$ pulumi up


# Test Python SDK
$ make build_python_sdk
$ cd examples/python
$ python -m pip install -e location\to\pulumi-azure-quickstart-compute\sdk\python
$ pulumi stack init test
$ pulumi up
```



