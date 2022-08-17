variable "test_name" {
  type = "string"
}

variable "invalid_2" {
  type = "string"
}

variable "nested_test" {
  type = object({
    location    = string
  })
}

variable "test_street" {
  type = "string"
}

variable "test_str" {
  type = "string"
}

variable "tags2_a" {
  type = "string"
}

variable "tags" {
  type = map(string)
}