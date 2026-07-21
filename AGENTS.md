# talos-metal-agent — agent guide

This is the ever-growing knowledge base for the project.
Maintain it as you go: whenever you learn something durable about how this repo works, add it here.
It is not append-only, so fix or delete anything that becomes wrong or outdated.
The goal is for this file to keep growing and keep getting more correct over time.
Only capture timeless, project-general knowledge here, never in-flight work or an individual's local setup.

Markdown note: this file is linted (`markdownlint` with the `sentences-per-line` rule), so keep one sentence per line and the first line as the single top-level heading.

`CLAUDE.md` and `GEMINI.md` just import this file, so this is the one place to edit.

## What this repo is

This repo builds the metal agent binary and its container image.
The agent is then packaged in two further ways.
It becomes a Talos system extension in `siderolabs/extensions` under `guest-agents/metal-agent`, which packages the agent binary as a container.
It is also packed into a boot-assets image (`ghcr.io/siderolabs/talos-metal-agent-boot-assets`) built by this repo's own CI.

The agent runs on a bare-metal machine booted into Talos "agent mode" and talks to the Omni bare-metal infra provider (`siderolabs/omni-infra-provider-bare-metal`, the "BM provider").

## The two consumption paths

The BM provider boots machines into agent mode by one of two paths, and it is worth knowing both before changing boot or version logic.

The image factory path is the production default.
The provider asks the image factory for a schematic that lists the extension as `siderolabs/metal-agent` with no version, for the provider's `AgentModeTalosVersion`.
The factory resolves that name against its per-Talos-version official-extensions catalog and errors if the extension is not published for that Talos version.
So the agent version served this way is whatever the extensions catalog pins for `AgentModeTalosVersion`, and advancing it means getting `AgentModeTalosVersion` onto a Talos version whose catalog pins the desired agent version, by waiting for such a catalog or bumping the provider setting to one.

The local boot-assets path is for development and airgap use and is enabled by the provider's `--use-local-boot-assets` flag.
The boot-assets image is baked into the provider image at build time, and the flag only decides whether those already-baked files get served.
When that flag is set the provider's agent-mode Talos version has no effect, because the agent binary comes from the baked-in initramfs rather than the factory.

## The boot-assets image

The image is built by `hack/boot-assets/build.sh`, wired through the kres custom steps `image-boot-assets` and `push-boot-assets` in `.kres.yaml`.

Its tag scheme is `<IMAGER_TAG>-agent-<git-describe>`.
On a release tag it is clean (`<imager>-agent-vX`), and on other main commits it carries a `-N-gsha` suffix.

The build does not require the metal-agent extension to be published anywhere first.
It clones `siderolabs/extensions` at the pinned `EXTENSIONS_REF`, overrides the extension's version and image prefix, and builds the extension from source against the just-built agent image.
It still uses the published imager and pulls the firmware extensions from the `ghcr.io/siderolabs/extensions` catalog for the same Talos version.

`IMAGER_TAG` is the Talos version the local boot assets are built against.
It lives in `.kres.yaml` in two places, the `image-boot-assets` and `push-boot-assets` steps, and they must stay in sync (the file has a note about this).
Bump it during release preparation rather than in a follow-up commit, so the clean release-tag boot-assets image is already the newest.
Track the latest Talos patch, since the imager publishes patch tags, and keep it at or above the Talos commit where agent mode was introduced (see the comment on the variable).
The boot-assets image is mainly for local development and airgap use.

## The extensions repo coupling

In `siderolabs/extensions`, the metal-agent extension version is templated from `TALOS_METAL_AGENT_VERSION`.
That variable is bumped by hand when a new agent release should ship, and the bump is sometimes backported to older release branches.
The extension version is pinned per Talos version in each catalog, meaning each Talos release tag pins whatever agent version was current then, and the extensions default branch can lag its release branches.
So a new metal-agent version reaches the image factory path only through the next Talos catalog, gated by that hand bump and the Talos release cadence, not through a standalone extensions release.

## The "two releases" reality

On this repo's side it is essentially one release.
A metal-agent tag builds both the agent image and the boot-assets image in a single CI run.
The extensions bump is a separate hand-authored change in another repo and rides the next Talos catalog, which is out of this repo's hands.
Any habit of pushing a second tag or follow-up commit is really an `IMAGER_TAG` follow-up, which should be folded into release preparation instead.

## Test mode

The agent has a test mode, enabled by the `metal.provider.test.mode` kernel argument (or the `--test-mode` flag), which the BM provider sets when it runs with its own agent test mode flag.
In test mode the agent reports API-based power management instead of reading the BMC address and setting up an IPMI user from inside the OS.
This backs the emulated bare-metal setup used for development and integration testing, where QEMU machines expose a small HTTP power API through the Talos provision machinery (the same one behind `talosctl cluster create`).

## Repo conventions and gotchas

- Always branch from the latest upstream default branch, since a stale base can predate current files and produce confusing diffs.
- Never hand-edit generated files, among them `Makefile`, `Dockerfile`, `.github/workflows/*`, `.dockerignore`, `.gitignore`, `.golangci.yml`, and `lefthook.yml`.
  The list is not exhaustive, and a kres-generated file identifies itself with a generated-file comment near its top.
  To change them, edit `.kres.yaml` and run `make rekres`.
- `make rekres` pulls the latest kres image, so a newer kres than last-generated brings tooling churn, which is expected in a bump-deps or rekres commit.
- kres auto-detects tracked root-level markdown and wires it into the Dockerfile markdown-lint stage, which is why every committed root markdown file, including this one, must pass markdownlint.
  Personal untracked files, for example a git-ignored `CLAUDE.local.md`, are left out.
- The boot-assets `EXTENSIONS_REF` pin in `.kres.yaml` carries an explanatory comment and is deliberate, so do not bump it during routine dependency bumps.
- The committed version tag file under `internal/version` is written by the release step, not by bump commits, so revert it if a stray generate step changes it during a dependency bump.

## Dependency bumps

Bump direct dependencies only.
List the direct, non-main dependencies with `go list -m -f '{{if and (not .Main) (not .Indirect)}}{{.Path}}@upgrade{{end}}' all`, feed them to `go get`, and then run `go mod tidy`.
Never use `-u` or `all`, since those drag indirect dependencies past what the direct ones require.
Keep the `go` directive in `go.mod` matching the Makefile Go version, which a `talos/pkg/machinery` bump can raise on its own.

Verify with the make targets, which are authoritative: `make lint-fmt`, `make lint` (includes govulncheck and markdownlint), `make generate`, and `make unit-tests`.

## Release process

Releases are cut with the repo's own tooling (`hack/release.sh` and `hack/release.toml`) in two phases.
First a release pull request sets `hack/release.toml` `previous` to the latest tag, regenerates the version files, runs `hack/release.sh changelog vX` to update the changelog, and creates the `release(vX): prepare release` commit via `hack/release.sh commit vX` (the script adds a DCO sign-off, and commits are GPG-signed when git is configured to sign).
Then, after that merges, a signed tag is pushed, which triggers CI to build and push the release image and the boot-assets image and to draft a GitHub release.
The dependency bump and the release are separate pull requests.
This repo does not sign its images, so there is no image-signing step.

## Commit and pull request conventions

Siderolabs pull requests often contain a single commit, but multiple commits are fine for separate atomic logical changes (conform sets `maximumOfOneCommit: false`).
A single commit is usually still preferable, since the pull request title and body then equal the commit title and body minus the sign-off trailer.
Either way, keep pull request titles and bodies simple like a commit message, without fancy markdown, and not long like a documentation page.
The commit title follows the conform rules: a `type[(scope)]: imperative summary` line, with types such as `feat`, `fix`, `chore`, `refactor`, `test`, `docs`, and `release`.
Commits carry a `Signed-off-by` trailer and are GPG-signed.
Describe behavior rather than code identifiers, and write issue references inline as sentences, for example `Closes siderolabs/talos-metal-agent#<n>.`

## Reference pointers

- The agent's protocol and behavior live under `api/agent` and `internal/agent`.
- The boot-assets build and its extensions coupling are in `hack/boot-assets` and `.kres.yaml`.
- The BM provider that consumes this agent is `siderolabs/omni-infra-provider-bare-metal`, whose own `AGENTS.md` documents the end-to-end bare-metal lifecycle.
