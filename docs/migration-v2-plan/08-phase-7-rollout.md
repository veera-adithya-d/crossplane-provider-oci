# Phase 7: Rollout

## Goal

De-risk rollout by piloting a small set first.

## Pilot Order

1. `config`
2. one service provider (`objectstorage` suggested)
3. full monolith + full subpackages

## Steps

1. Build pilot:
```bash
make build SUBPACKAGES=config,objectstorage
```
2. Package pilot:
```bash
make build-subpackages SUBPACKAGES_FOR_BATCH=config,objectstorage BATCH_PLATFORMS=linux_amd64
```
3. Deploy in test cluster and verify:
   - provider health
   - providerconfig resolution (legacy + modern)
   - one namespaced MR lifecycle
4. Expand to all packages.

## Exit Criteria

- [ ] Pilot stable
- [ ] Full package generation stable
- [ ] Release notes include migration notes and breaking changes
