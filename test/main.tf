provider "azurerm" {
  features {}
}

data "azurerm_subscription" "current" {
}

resource "azurerm_role_assignment" "aks-agentpool-rg-acr-acrpull" {
  name                 = "${var.name_prefix}_${var.aks_name}"
  scope                = "/subscriptions/${data.azurerm_subscription.current.subscription_id}"
  role_definition_name = "AcrPull"
  principal_id         = var.role_assignment_principal_id

  conditional_prop = var.flag ? var.my_conditional_prop : "b"

  nested_var_prop = var.props.my_nested_var_prop
  array_prop = var.my_array_prop[1]

  tag = {
    "a" = var.tag_a
    "b" = var.tag_b
  }
  
  nested_prop {
    n_a = var.n_a
	n_b = 1
	n_c = var.nested_prop_n_c
	n_d = var.nested_prop_n_d_invalid
  }
}
