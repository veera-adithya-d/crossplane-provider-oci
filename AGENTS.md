# Repository Guidelines

This project ships the Crossplane provider for Oracle Cloud Infrastructure (OCI). Use this guide as the rules of engagement for both developer-facing architecture work and customer-facing enablement.

## Architecture-First Workflow (developers start here)
- Map your change to the correct module: `cmd/provider` (entrypoints), `internal/controller` (reconcilers/shared helpers), `apis` (CRDs + clients), `config`/`package` (xpkg manifests), `build` (tooling), and `examples`/`docs` (usage narratives).
- Understand control-plane topology before coding; consult per-service READMEs plus `docs/` diagrams to keep controller layering consistent.
- Follow the knowledge-base order: `.work/oracle/oci/website/docs` → Terraform Registry for oracle/oci v7.27.0 → Oracle Cloud service docs → Crossplane docs. Quote the relevant source when explaining design choices.
- Keep versions centralized in `Makefile`; do not bump Terraform, Go, or OCI provider pins ad hoc.
- Preserve resolver/reference annotations when editing generated types; re-run `make generate.init` (schema/docs) and `make generate.resolve` (resolver sync) after structural edits.

## Customer Resource Catalog (answer user questions with this)
- Service families shipped today include: config (provider-family-oci), compute, networking (VCN, load balancers, firewalls), storage (block, object, file), identity, security (KMS, Vault, certificates), data & messaging (streaming, events, ons), observability (logging, monitoring, health checks), container & developer tools (OKE, functions, artifacts), and additional platform surfaces (dns, networkloadbalancer, certificatesmanagement, filestorage, networkfirewall).
- Each family has controllers in `internal/controller/<service>` backed by APIs under `apis/<group>/<version>` and example manifests in `examples/<service>`; reference these when guiding customers on creating resources.
- Customers interact through Crossplane composite or managed resources; document required `ProviderConfig`, CRD fields, and dependency chains (e.g., networking before compute) in your PRs and user docs.

## Build & Test Essentials
- Local build paths: `make build` (monolith), `make build SUBPACKAGES=a,b` (targeted), `make build.dev` (host-only quick loop).
- Validation: `go test ./...`, `make test.subpackage.<svc>` for service suites, `make uptest` for full-stack scenarios, `make cobertura` for CI XML. Always finish with `make generate.resolve` and `make build` before opening a PR.

## Contribution & Support Expectations
- Commits: imperative, ≤72 chars, include service prefix when relevant (`compute: tighten retry`). Reference issues (`Fixes #123`) and describe controller/API impact plus customer-visible behavior.
- Security: never commit credentials; rely on env vars and `ProviderConfig` secrets. After CRD or metadata updates, run `kustomize-crds` via `make build.init` so package artifacts stay aligned.
- Support readiness: highlight new services, controllers, or API surfaces in release notes and `docs/`, ensuring customers know which manifests to apply and which knowledge-base references justify the design.
