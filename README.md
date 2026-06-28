# Stout Curated Workflows

Curated **build-from-source** workflows for the Stout catalog factory.

Each file under [`curated/`](./curated) is a `CuratedWorkflow` — a vetted mapping
from a package to its **canonical source** (repo + ref + build directory) so Stout
can **rebuild it from source**, scan the result, and emit signed SLSA provenance.
This is the catalog-driven half of the build lane: an operator triggers a curated
workflow from the admin console and Stout runs it as a `strategy=build` verification
run (clone → build → validate → publish + provenance).

> **Why a separate repo?** The curated catalog is config, not code in the registry.
> Keeping it here means the set of "things we rebuild from source" is reviewable,
> versioned, and PR-gated independently of the platform — and the platform syncs it
> read-only.

## Layout

| Path | What |
|---|---|
| `curated/*.yaml` | one `CuratedWorkflow` per package |
| `curated/index.yaml` | manifest listing the active workflows (what the admin console syncs) |
| `schema/curated-workflow.schema.json` | JSON Schema for a workflow file |
| `samples/` | **self-contained, dependency-free** sample packages used for end-to-end validation (so a build never depends on a flaky external fetch) |

## Workflow spec

```yaml
apiVersion: stout.io/v1
kind: CuratedWorkflow
metadata:
  name: hello-go                       # unique slug
  description: Minimal Go module.
spec:
  target:
    name: hello-go                     # registry package name
    version: 0.1.0                     # version to build
    packageType: go                    # go|npm|python|... (a Stout-supported build ecosystem)
  source:
    repo: https://github.com/stout-io/curated-workflows
    gitRef: main                       # branch/tag/commit to pin
    buildDir: samples/hello-go         # subdir within the repo to build
  strategy: build                      # build-from-source (vs mirror)
```

The admin console maps a `CuratedWorkflow` to a verification run:
`TriggerRun(strategy=build, target_{name,version,package_type}, submitted_source_repo=spec.source.repo, submitted_git_ref=spec.source.gitRef, build_dir=spec.source.buildDir)`.
Stout resolves the commit (write-once), builds in `buildDir`, runs the validation
battery, and on a clean verdict publishes the artifact with a signed provenance
statement. The build runs **isolated, with use-once credentials** (the scan-runner
model) in production.

## Adding a workflow

1. Add `curated/<name>.yaml` (validate against the schema).
2. Add its `name` to `curated/index.yaml`.
3. Open a PR. Merged workflows are synced into the catalog on the next admin sync.

Source repos referenced here should be ones we trust to rebuild; `verifiedUpstream`
in the platform's `package_sources` marks a curator's confirmation that the repo is
the authentic upstream.
