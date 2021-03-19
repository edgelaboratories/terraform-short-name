# terraform-terraform-short-name

Generate short, unique, sanitized terraform variable names.

This module helps truncating long names whilst ensuring the name remains unique.
It also sanitizes the name by replacing `_` with `-`.

## How to use

In your Terraform file, instanciate the module with:

```hcl
module "my_unique_name" {
  source = "git@github.com:edgelaboratories/terraform-modules.git//short-name?ref=v17"

  name          = "my-long-we?rd/.non-unique-name"
  max_length    = 20
  suffix_length = 4
  upper         = false
}
```

This will result in the sanitized, unique, name "my-long-we-rd-9C4S" that you can get with

```hcl
locals {
  unique_name = module.my_unique_name.name
}
```
