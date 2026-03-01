# Phase 1: Baseline Fix + V2 Dependencies

## Goal

Make repo compile, then move to `crossplane-runtime/v2` and `upjet/v2`.

## Key Files

- `go.mod`
- `hack/main.go.tmpl`
- `cmd/provider/*/zz_main.go` (generated)
- `cmd/generator/main.go`
- `Makefile`

## Steps

1. Fix compile issues in template-generated subpackage mains (`kingpin`, `feature` wiring).
2. Update module deps to v2 lines.
3. Replace imports:
```text
github.com/crossplane/crossplane-runtime/ -> github.com/crossplane/crossplane-runtime/v2/
github.com/crossplane/upjet/ -> github.com/crossplane/upjet/v2/
```
4. Update generator pipeline call site after config split placeholder is ready.
5. Run:
```bash
go mod tidy
go test ./cmd/generator -run TestNonExistent -count=0
go test ./cmd/provider/config -run TestNonExistent -count=0
```

## Diagram

![Current Wiring](./diagrams/01-current-wiring.png)

## Exit Criteria

- [ ] No v1 runtime/upjet imports
- [ ] `go.mod` on v2 modules
- [ ] representative subpackage mains compile
