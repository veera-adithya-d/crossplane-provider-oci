# Phase 6: Regenerate + Validate + Docs

## Goal

Regenerate everything, then align tests/docs/examples.

## Key Files

- `apis/generate.go`
- `package/crds/*`
- `examples-generated/*`
- `cluster/test/setup.sh`
- `docs/quickstart.md`
- `argocd-examples/argocd/config/argocd-crossplane-config.yaml`

## Steps

1. Regenerate:
```bash
make generate
```
2. Validate counts:
```bash
rg -l "scope: Namespaced" package/crds | wc -l
rg -l "scope: Cluster" package/crds | wc -l
rg -l "publishConnectionDetailsTo" package/crds | wc -l
```
3. Update test setup default to modern `ClusterProviderConfig.<group>.m.<domain>`.
4. Update docs/examples:
   - namespaced MR `metadata.namespace`
   - modern `providerConfigRef.kind/name`
   - local secret refs (no namespace in local ref fields)
5. Update Argo CD health mapping for new provider config groups.

## Diagram

![Generation Flow](./diagrams/05-generation-flow.png)

## Exit Criteria

- [ ] namespaced CRDs generated
- [ ] no ESS fields
- [ ] docs/tests/examples updated for modern group/kinds
