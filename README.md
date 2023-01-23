# openapi-generator-go

This is a Go CLI generator used for generating 5GC api server and client code from the [3GPP OpenAPI specs](https://github.com/jdegre/5GC_APIs). 
It relies on [OpenAPI Generator project](https://openapi-generator.tech/).
We directly use the [OpenAPI Generator API](https://api.openapi-generator.tech) to generate the code from the remote 3GPP OpenAPI specs.

> Note: This project is still under development and not ready for production use.
> Note2: At the moment, code generation is not totally working. We are working on it.
> Note3: Spec files must be remote files on github. We are working on a solution to make run the generator with local files and other sources.


## Installation

To install the generator, run:

```bash
go get -u github.com/5GCoreNet/api-generator-go
```

## Usage

### Generate server

```bash
Usage:
  api-generator-go server [flags]

Flags:
  -h, --help   help for server

Global Flags:
      --exit-on-failure           Exit on failure (default true)
      --output string             Output directory for the generated code
      --path string               Path to the OpenAPI specification folder under the repository (default ".")
      --ref string                Commit hash or branch name of the OpenAPI specification repository (default "Rel-18")
      --repository string         URL to the OpenAPI specification repository (default "5GC_APIs")
      --repository-owner string   Owner of the OpenAPI specification repository (default "jdegre")
```

### Generate client

```bash
Usage:
  api-generator-go client [flags]

Flags:
  -h, --help   help for client

Global Flags:
      --exit-on-failure           Exit on failure (default true)
      --output string             Output directory for the generated code
      --path string               Path to the OpenAPI specification folder under the repository (default ".")
      --ref string                Commit hash or branch name of the OpenAPI specification repository (default "Rel-18")
      --repository string         URL to the OpenAPI specification repository (default "5GC_APIs")
      --repository-owner string   Owner of the OpenAPI specification repository (default "jdegre")
```

## Build from source

To build the generator from source, run:

```bash
git clone https://github.com/5GCoreNet/api-generator-go.git
cd api-generator-go
make build
```

## Contributing
Feel free to open issues and submit pull requests.