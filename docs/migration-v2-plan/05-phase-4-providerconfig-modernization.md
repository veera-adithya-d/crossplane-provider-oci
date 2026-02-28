# Phase 4: ProviderConfig Modernization

## Goal

Support modern namespaced MR config refs and keep legacy cluster behavior.

## Key Files

- `apis/namespaced/v1beta1/types.go`
- `apis/namespaced/v1beta1/register.go`
- `internal/controller/*/providerconfig/config.go`
- `internal/clients/oci.go`

## Steps

1. In namespaced API group:
   - Use `TypedProviderConfigUsage`
   - Add namespaced `ProviderConfig`
   - Add cluster-scoped `ClusterProviderConfig`
2. Keep legacy `ProviderConfig.oci.upbound.io` for cluster MRs.
3. In client setup:
   - Read typed `providerConfigRef.kind/name`
   - Route to legacy vs modern kinds by MR group/scope
   - Normalize into common config struct
4. Ensure local secret refs remain local for namespaced MRs.

## Diagram

![ProviderConfig Resolution](./diagrams/03-providerconfig-resolution.png)

## Exit Criteria

- [ ] Namespaced group has `ProviderConfig`, `ClusterProviderConfig`, typed usage
- [ ] Kind-aware provider config resolution is implemented
- [ ] Legacy cluster MRs still resolve legacy provider config
