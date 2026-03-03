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

## Codex Context Infrastructure (Tiered)
- Tier 1 (always-on): [AGENTS.md](AGENTS.md) + [CODEX.md](CODEX.md)
- Tier 2 (path-scoped specialists):
  - `config/AGENTS.md`
  - `apis/AGENTS.md`
  - `internal/controller/AGENTS.md`
  - `examples/AGENTS.md`
  - `docs/AGENTS.md`
  - `package/AGENTS.md`
- Tier 3 (on-demand deep context):
  - `.codex/context/architecture-map.md`
  - `.codex/context/generation-lifecycle.md`
  - `.codex/context/resource-onboarding-playbook.md`
  - `.codex/context/customer-enablement.md`

When tasks are architectural or cross-module, load Tier 3 docs explicitly before implementing.

<!-- BEGIN BEADS INTEGRATION -->
## Issue Tracking with bd (beads)

**IMPORTANT**: This project uses **bd (beads)** for ALL issue tracking. Do NOT use markdown TODOs, task lists, or other tracking methods.

### Why bd?

- Dependency-aware: Track blockers and relationships between issues
- Git-friendly: Dolt-powered version control with native sync
- Agent-optimized: JSON output, ready work detection, discovered-from links
- Prevents duplicate tracking systems and confusion

### Quick Start

**Check for ready work:**

```bash
bd ready --json
```

**Create new issues:**

```bash
bd create "Issue title" --description="Detailed context" -t bug|feature|task -p 0-4 --json
bd create "Issue title" --description="What this issue is about" -p 1 --deps discovered-from:bd-123 --json
```

**Claim and update:**

```bash
bd update <id> --claim --json
bd update bd-42 --priority 1 --json
```

**Complete work:**

```bash
bd close bd-42 --reason "Completed" --json
```

### Issue Types

- `bug` - Something broken
- `feature` - New functionality
- `task` - Work item (tests, docs, refactoring)
- `epic` - Large feature with subtasks
- `chore` - Maintenance (dependencies, tooling)

### Priorities

- `0` - Critical (security, data loss, broken builds)
- `1` - High (major features, important bugs)
- `2` - Medium (default, nice-to-have)
- `3` - Low (polish, optimization)
- `4` - Backlog (future ideas)

### Workflow for AI Agents

1. **Check ready work**: `bd ready` shows unblocked issues
2. **Claim your task atomically**: `bd update <id> --claim`
3. **Work on it**: Implement, test, document
4. **Discover new work?** Create linked issue:
   - `bd create "Found bug" --description="Details about what was found" -p 1 --deps discovered-from:<parent-id>`
5. **Complete**: `bd close <id> --reason "Done"`

### Auto-Sync

bd automatically syncs via Dolt:

- Each write auto-commits to Dolt history
- Use `bd dolt push`/`bd dolt pull` for remote sync
- No manual export/import needed!

### Important Rules

- ✅ Use bd for ALL task tracking
- ✅ Always use `--json` flag for programmatic use
- ✅ Link discovered work with `discovered-from` dependencies
- ✅ Check `bd ready` before asking "what should I work on?"
- ❌ Do NOT create markdown TODO lists
- ❌ Do NOT use external issue trackers
- ❌ Do NOT duplicate tracking systems

For more details, see README.md and docs/QUICKSTART.md.

<!-- END BEADS INTEGRATION -->

## Landing the Plane (Session Completion)

**When ending a work session**, you MUST complete ALL steps below. Work is NOT complete until `git push` succeeds.

**MANDATORY WORKFLOW:**

1. **File issues for remaining work** - Create issues for anything that needs follow-up
2. **Run quality gates** (if code changed) - Tests, linters, builds
3. **Update issue status** - Close finished work, update in-progress items
4. **PUSH TO REMOTE** - This is MANDATORY:
   ```bash
   git pull --rebase
   bd sync
   git push
   git status  # MUST show "up to date with origin"
   ```
5. **Clean up** - Clear stashes, prune remote branches
6. **Verify** - All changes committed AND pushed
7. **Hand off** - Provide context for next session

**CRITICAL RULES:**
- Work is NOT complete until `git push` succeeds
- NEVER stop before pushing - that leaves work stranded locally
- NEVER say "ready to push when you are" - YOU must push
- If push fails, resolve and retry until it succeeds
