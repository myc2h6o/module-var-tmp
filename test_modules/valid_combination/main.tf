resource "azurerm_resource_group" "test" {
  # A valid variable name is same as property name or property name with prefix
  name = var.resource_group_name
  location = var.location
  tags = var.test_tags
}

resource "azurerm_var_validation" "test" {
  # Direct usage
  name = var.test_name

  # Nested
  location = var.nested.location

  # List
  list = var.test_list[2]

  # Nested list
  region = var.nested_list[1].region

  # Nested property
  address {
    street = var.test_street
  }

  # Vars inside List are excluded
  ip_addresses = ["0.0.0.0", var.ip_addr_1]

  # Vars inside TypeMap are excluded
  tags = {
    a = var.inner_type_map_test
  }

  # Vars in string template are excluded
  string_template = "Test_${var.str_template_test}"

  # Vars in function call are excluded
  sub_string = substr(var.sub_str_test, 0, 1)

  # For conditional expression, vars in condition part are excluded. Only vars in true/false value part are included, and excluded pattern like string template are still excluded.
  condition_prop = var.condition_flag != "" ? var.true_condition_prop : "Test_${false_prop}"

  # Built-in properties are excluded
  # https://www.terraform.io/language/resources/syntax#meta-arguments and https://www.terraform.io/language/resources/syntax#operation-timeouts
  depends_on = var.depends
  timeouts {
    read = var.timeout_read_1
  }
  lifecycle {
    ignore_changes = var.ignored_props
  }
}
