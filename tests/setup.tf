variable "input_name" {}

variable "max_length" {}

variable "suffix_length" {}

variable "upper" {}

module "my_unique_name" {
  source = "../"

  name          = var.input_name
  max_length    = var.max_length
  suffix_length = var.suffix_length
  upper         = var.upper
}

output "unique_name" {
  value = module.my_unique_name.name
}
