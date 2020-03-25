# Terraform Module: Get Kubernetes Secret Value

> This repository is a [Terraform](https://terraform.io/) Module to use Kubernetes secrets as data.

## Table of Contents

- [Requirements](#requirements)
- [Dependencies](#dependencies)
- [Usage](#usage)
  - [Module Variables](#module-variables)
- [Contributing](#contributing)
- [Maintainers](#maintainers)

## Requirements

This module requires Terraform version `0.10.x` or newer and a valid out-of-cluster Kubernetes configuration at the default path (`$HOME/.kube/config`) or the path defined by `$KUBECONFIG`.

## Usage

Add the module to your Terraform resources:

```hcl
module "rds-postgres-password" {
  source    = "github.com/jobteaser/terraform-kubernetes-get-secret?ref=v0.2.3"

  namespace = "default"
  name = "terraform"
  key = "database-password"
  context = "<env>.jt"
}
```

and load the module using `terraform get`.

The binary won't fail if you don't pass the context but terraform will.

### Module Variables

Available variables are listed below, along with their default values:

| variable    | description                      |
|-------------|----------------------------------|
| `namespace` | The kubernetes namespace         |
| `name`      | The kubernetes secret name       |
| `key`       | The kubernetes secret key to get |
| `context`   | The kubernetes context to use    |

### Module outputs

Available outputs are listed below, along with their description:

| output    | description                   |
|-----------|-------------------------------|
| `result`  | A string of the secret value. |

## Contributing

### Requirements

- fully installed and configured `go` environment
- fully installed and configured [golang/dep](https://github.com/golang/dep)

### Contributing code

To build binaries, use the following commands:

```
$ git clone git@github.com:jobteaser/terraform-kubernetes-get-secret.git
$ cd terraform-kubernetes-get-secret
$ make
```

## Maintainers

This module is a fork of [gearnode/terraform-kubernetes-get-secret](https://github.com/gearnode/terraform-kubernetes-get-secret).


