variable "name" {
  description = "Original name."
  type        = string
}

variable "max_length" {
  description = "Maximum length the output name can be."
  type        = number
}

variable "suffix_length" {
  description = "Length of the random suffix used to produce a unique name."
  default     = 4
  type        = number
}

variable "upper" {
  description = "Include uppercase alphabet characters in random suffix."
  default     = true
  type        = bool
}
