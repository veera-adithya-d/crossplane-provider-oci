# Architecture Overview

This document captures how the OCI Crossplane provider is wired today and how it should evolve to surpass the AWS, GCP, and Azure Upjet providers.

## Scope & Philosophy
- The provider embeds Oracle's Terraform schema (`config/provider.go`) and emits Crossplane-native CRDs under `oci.upbound.io`, matching OCI nouns (VCN, Vault, OKE, etc.).
- A single codebase powers both the monolithic binary and prospective sub-packages via the `SUBPACKAGES` Makefile plumbing, so every controller shares the same reconciliation primitives, metrics, and image build.
- Knowledge bases live outside the binary (.work cache, Terraform docs, Oracle docs, Crossplane docs) and are meant to be cited in every architectural decision.

## Wiring Blueprint (`config/provider.go`)
1. **Provider seed** – `GetProvider()` loads `schema.json`, `provider-metadata.yaml`, the `resourcePrefix` (`oci`), and the root group `oci.upbound.io`, ensuring every CRD is rooted in the same API domain.
2. **Resource shaping** – include lists combine hand-tuned resources with a wildcard, while `ProblematicResources()` gates anything unstable. Default options inject group/kind overrides, external-name policies, and automatic OCID handling, so even unconfigured resources behave predictably.
3. **Reference graph** – both the Upjet dynamic injector and `NewStaticReferenceInjector()` populate `[*]Ref`/`[*]Selector` fields, aligning dependencies (VCN→Subnet→Instance) across services.
4. **Service modules** – each package in `config/<service>` (identity, objectstorage, kms, etc.) pushes overrides before `ConfigureResources()` locks in the catalog. This mirrors the OCI product map and keeps reconciliation logic close to Terraform semantics.

## Resource & Delivery Flow
1. `make generate.init` downloads the pinned Terraform OCI provider and docs, storing them in `.work`.
2. `make generate.resolve` rewrites resolver hints so service modules can inject Crossplane references without hand-editing every CRD.
3. `make build` (or `make build SUBPACKAGES=...`) compiles controllers into `cmd/provider`, emitting CRDs under `package/crds` and examples in `examples/`.
4. `make local-deploy` and `make uptest` provide repeatable smoke and end-to-end coverage, wiring the generated providers into a kind cluster.

## Feature & Process Inventory
- **Controller surface** – `internal/controller/<service>` holds reconcilers, sharing helper packages for retry/late-init logic; APIs reside in `apis/<group>/<version>` with generated clients.
- **Packaging** – `config/` and `package/` define the XPKG manifests, while `internal/features` gates experimental toggles. Registry orgs default to `xpkg.upbound.io/upbound` but are overridable.
- **Security posture** – `ProviderConfig` secrets flow through Crossplane references, and `kustomize-crds` (invoked by `make build.init`) keeps published CRDs aligned whenever schemas change.
- **Knowledge enforcement** – every controller change should cite `.work/oracle/oci/website/docs`; fall back to the Terraform Registry (v7.27.0), Oracle service docs, then Crossplane docs when clarifying behavior.

## Benchmarks from Other Upjet Providers
- **AWS provider playbook** – Upbound's AWS provider already exposes quick-start flows, marketplace API references, plus sizing and monitoring guides that help operators estimate resource use. That level of operational collateral is absent here and should be emulated. citeturn3search0
- **Family architecture** – Installing an AWS service provider (e.g., S3) automatically pulls `provider-family-aws`, centralizing ProviderConfig/auth governance for every AWS service. OCI still ships only a monolith; introducing a family-level package would mirror this proven pattern. citeturn8search5
- **Release rigor** – Recent AWS releases prioritize compatibility with Crossplane v2, namespace-scoped managed resources, and Terraform v6+ alignment—showing a cadence of breaking-change communication we should match when uplifting OCI schemas. citeturn3search1
- **GCP provider focus** – The GCP Upjet README highlights a streamlined quickstart (install provider, configure ProviderConfig, apply compositions) and points users at the Upbound Marketplace for resource docs; this improves customer onboarding compared with our sparse docs. citeturn11view0
- **Azure provider focus** – Azure's provider offers a similar quickstart plus marketplace links, reinforcing that hyperscaler parity demands curated install flows and published API catalogs. citeturn12view0

## Optimization Plan
1. **Family packages for OCI** – Introduce `provider-family-oci` as a meta-package that installs shared auth/config controllers plus per-service “child” packages (compute, networking, storage). This reduces blast radius and mirrors AWS's proven model. citeturn8search5
2. **Operational collateral** – Publish sizing and monitoring guides patterned after the AWS documentation set so platform teams can plan cluster resources and query controller metrics without guesswork. citeturn3search0
3. **Release discipline** – Adopt semver’d changelog slices similar to AWS v2.1.x releases (call out Crossplane version compatibility, Terraform schema bumps, and deprecated OCI services) so upgrades stay predictable. citeturn3search1
4. **Guided quickstarts** – Build GCP/Azure-style quickstart docs (install → configure ProviderConfig → deploy sample compositions) tied to the Upbound Marketplace entries for every OCI service to cut customer onboarding time. citeturn11view0turn12view0
5. **Example parity tests** – Extend `make uptest` with composable scenarios that match AWS sizing guides (e.g., baseline networking + compute stacks) to validate cross-resource wiring with every release. citeturn3search0

## Next-Generation Architecture Concepts
- **Composable control planes** – Split the monolith into OCI family packages that share a core ProviderConfig and install only the controllers a customer needs (similar to AWS service packages), but add Crossplane compositions that span OCI + third-party SaaS for a differentiated story. citeturn8search5
- **Autonomous documentation sync** – Stream docs from `.work/oracle/oci/website/docs` into the Upbound Marketplace so OCI services show the same polished API references AWS/GCP customers already expect. citeturn3search0turn11view0turn12view0
- **Observability-first runtime** – Ship built-in dashboards and alerts (controller-runtime metrics, reconciliation SLIs) to overtake hyperscalers that currently rely on external sizing/monitoring PDFs; integrate these with Crossplane metrics APIs mentioned in AWS guidance to leapfrog operator UX. citeturn3search0
