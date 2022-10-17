resource "azurerm_resource_group" "test" {
  # A variable name with property name plus suffix is invalid
  name = var.name_1

  # A variable name with different value from property is invalid
  location = var.address

  # A variable name with only partial property name is invalid
  tags = var.tag
}

resource "azurerm_var_validation" "test" {
  address {
    street = var.street_test
  }

  # condition_flag is ignored, testing only condition_prop_true and condition_prop_false
  condition_prop = "" != "" ? var.condition_prop_true : var.condition_prop_false
}
