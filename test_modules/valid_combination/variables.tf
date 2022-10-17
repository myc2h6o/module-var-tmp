variable "resource_group_name" {
  type = "string"
}

variable "location" {
  type = "string"
}

variable "test_tags" {
  type = map(string)
}

variable "test_name" {
  type = "string"
}

variable "nested" {
  type = object({
    location = string
  })
}

variable "test_list" {
  type = list(string)
}

variable "nested_list" {
  type = list(object({
    region = string
  }))
}

variable "test_street" {
  type = "string"
}

variable "ip_addr_1" {
  type = "string"
}

variable "inner_type_map_test" {
  type = "string"
}

variable "str_template_test" {
  type = "string"
}

variable "sub_str_test" {
  type = "string"
}

variable "condition_flag" {
  type = "string"
}

variable "true_condition_prop" {
  type = "string"
}

variable "depends" {
  type = list(string)
}

variable "timeout_read_1" {
  type = "string"
}

variable "ignored_props" {
  type = list(string)
}