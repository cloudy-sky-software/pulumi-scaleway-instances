## 0.1.1

Upgrade pulumi-provider-framework to fix a bug with validating response codes for DELETE calls.

## 0.1.0

`servers.getPlacementGroupServer` was renamed to `servers.getPlacementGroupServers`.

## 0.0.7

- Prepare for public release.

## 0.0.6

- Update the OpenAPI spec to mark name as required which will translate to an auto-name capable property in Pulumi.

## 0.0.5

- Mark the `state` property in `Volume` resource as an output-only property.

## 0.0.4

- Restore `id` property in the OpenAPI spec for resources. The schema generator in `pulschema` will exclude it from input properties for a resource.

## 0.0.3

- Pluck outputs from the top-level property for create/read/update responses

## 0.0.1 - 0.0.2

Initial release
