resource "azurerm_var_validation" "test" {
  condition_prop = "" != "" ? var.condition_prop_true : ""
}
