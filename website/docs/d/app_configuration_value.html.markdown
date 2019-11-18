---
subcategory: ""
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_app_configuration_value"
sidebar_current: "docs-azurerm-datasource-app-configuration-value"
description: |-
  Gets information about an existing App Configuration value.

---

# Data Source: azurerm_app_configuration_value

Use this data source to access an existing App Configuration value.

~> **Note:** All arguments will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

## Example Usage

```hcl
data "azurerm_app_configuration_value" "db_url" {
  name         = "my-app/db_url"
  app_conf_id = resource.azurerm_app_configuration.existing.id
}

output "db_url_value" {
  value = data.azurerm_app_configuration_value.db_url.value
}
```

## Argument Reference

The following arguments are supported:

* `key` - (Required) Specifies the key of the App Configuration value.

* `app_conf_id` - (Required) Specifies the ID of the App Configuration instance where the value resides, available on the `azurerm_app_configuration` Resource.

* `label` - (Optional) Specifies the label of the App Configuration value.

**NOTE:** The App Configuration must be in the same subscription as the provider. If the App Configuration is in another subscription, you must create an aliased provider for that subscription.

## Attributes Reference

The following attributes are exported:

* `id` - The  Etag of this App Configuration value.

* `value` - The actual value.

* `content_type` - The content type for this App Configuration value.

* `tags` - Any tags assigned to this App Configuration value.
