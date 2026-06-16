config {
  format = "compact"
}

// As a module, we allow all versions.
rule "terraform_required_version" {
  enabled = false
}

// Same here, the versions are set outside the module.
rule "terraform_required_providers" {
  enabled = false
}
