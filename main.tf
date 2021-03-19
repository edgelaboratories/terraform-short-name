locals {
  # Replace every "weird" characters with dashes and remove ending dashes
  sanitized_name = replace(replace(lower(var.name), "/[^a-z\\-0-9]/", "-"), "/-*$/", "")
  # Truncate sanitized name and remove ending dashes
  truncated_name   = replace(substr(local.sanitized_name, 0, var.max_length - 1 - var.suffix_length), "/-*$/", "")
  name_is_too_long = length(local.sanitized_name) > var.max_length
}

resource "random_string" "suffix" {
  keepers = {
    name = local.sanitized_name
  }

  special = false
  upper   = var.upper
  length  = var.suffix_length
}
