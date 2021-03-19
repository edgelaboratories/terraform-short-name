variable "name" {
  description = "Original name."
}

variable "max_length" {
  description = "Maximum length the output name can be."
}

variable "suffix_length" {
  description = "Length of the random suffix used to produce a unique name."
  default     = 4
}

variable "upper" {
  description = "Include uppercase alphabet characters in random suffix."
  default     = true
}
