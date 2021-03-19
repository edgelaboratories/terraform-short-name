output "name" {
  description = "Shortened sanitized name."
  value       = local.name_is_too_long ? format("%s-%s", local.truncated_name, random_string.suffix.result) : local.sanitized_name
}
