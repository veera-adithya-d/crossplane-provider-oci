# Phase 0: Preflight

## Goal

Create a safe branch and baseline signals.

## Steps

1. Create branch.
```bash
git checkout -b feat/upjet-v2-migration
```

2. Capture baseline.
```bash
git status
go version
go test ./cmd/provider -run TestNonExistent -count=0
```

3. Capture generated scope counts.
```bash
rg -l "scope: Cluster" package/crds | wc -l
rg -l "scope: Namespaced" package/crds | wc -l
rg -l "publishConnectionDetailsTo" package/crds | wc -l
```

## Exit Criteria

- [ ] Branch created
- [ ] Baseline command outputs saved
- [ ] Scope and ESS counts saved
