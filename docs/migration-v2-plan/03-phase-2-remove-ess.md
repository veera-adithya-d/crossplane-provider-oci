# Phase 2: Remove ESS

## Goal

Remove External Secret Store (`StoreConfig`, `publishConnectionDetailsTo`) usage.

## Key Files

- `apis/v1alpha1/types.go`
- `apis/v1alpha1/register.go`
- `hack/main.go.tmpl`
- generated `cmd/provider/*/zz_main.go`
- generated `internal/controller/**/zz_controller.go`
- generated `package/crds/*.yaml`

## Steps

1. Delete `StoreConfig` API type and registration.
2. Remove `enable-external-secret-stores` flag path from mains/template.
3. Remove `SecretStoreConfigGVK` wiring from controller options.
4. Regenerate.
5. Verify:
```bash
rg -n "StoreConfig|publishConnectionDetailsTo|SecretStoreConfigGVK" apis cmd internal package
```

## Exit Criteria

- [ ] No `StoreConfig` CRD
- [ ] No `publishConnectionDetailsTo` in CRDs
- [ ] No `SecretStoreConfigGVK` in controllers
