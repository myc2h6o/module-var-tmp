provider "azurerm" {
  features {}
}

data "azurerm_subscription" "current" {
}

resource "azurerm_role_assignment" "aks-agentpool-rg-acr-acrpull" {
  scope                = "/subscriptions/${data.azurerm_subscription.current.subscription_id}"
  role_definition_name = "AcrPull"
  principal_id         = var.role_assignment_principal_id

  conditional_prop = var.bool_conditional_var ? "a" : "b"
  
  tag = {
    "a" = var.tag_a
  }
  
  nested_prop {
    n_a = var.n_a
	n_b = 1
	n_c = var.nested_prop_n_c
	n_d = var.nested_prop_n_d_invalid
  }
}
