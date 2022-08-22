
resource "azurerm_var_validation" "test" {
  // plain text
  # annotation 0
  name = var.test_name

  // nested
  location = var.nested_test.location

  // list
  prop = var.test_prop[2]

  // nested prop
  address {
    street = var.test_street
  }

  // invalid var which has suffix
  invalid = var.invalid_2

  // direct assign TypeMap is included
  tags = var.tags

  // composed TypeMap like tags are excluded
  tags2 = {
    a = var.tags2_a
  }

  // Composed string is excluded
  string_template = "Test_${var.test_str}"
}

data "azurerm_var_valication" "test" {

}