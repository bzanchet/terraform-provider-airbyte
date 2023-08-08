---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_cassandra Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationCassandra DataSource
---

# airbyte_destination_cassandra (Data Source)

DestinationCassandra DataSource

## Example Usage

```terraform
data "airbyte_destination_cassandra" "my_destination_cassandra" {
  destination_id = "...my_destination_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_id` (String)

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `address` (String) Address to connect to.
- `datacenter` (String) Datacenter of the cassandra cluster.
- `destination_type` (String) must be one of ["cassandra"]
- `keyspace` (String) Default Cassandra keyspace to create data in.
- `password` (String) Password associated with Cassandra.
- `port` (Number) Port of Cassandra.
- `replication` (Number) Indicates to how many nodes the data should be replicated to.
- `username` (String) Username to use to access Cassandra.

