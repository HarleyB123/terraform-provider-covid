# Covid Terraform Provider

Heavily inspired by the work on the [McBroken Terraform Provider](https://github.com/circa10a/terraform-provider-mcbroken) - This Terraform provider provides a data source for COVID Cases stats across the World and in a chosen Country.

## Getting Started

```
make build-mac && \
cd examples && \
terraform init && \
terraform apply
```

Replace build-mac with build-linux if appropriate.

You can also lint:

```
make lint
```

and run tests:

```
make test
```

## Potential TODO

- Skip first line in JSON output for all countries and make world stats a seperate data source
- Some sort of Fuzzy Finder (Eg: Did you mean 'Slovakia'?), through [this?](https://github.com/agnivade/levenshtein) ✅
