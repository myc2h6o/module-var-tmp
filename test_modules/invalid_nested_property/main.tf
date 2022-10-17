resource "azurerm_var_validation" "test" {
  address {
    street = var.street_test
  }
}
