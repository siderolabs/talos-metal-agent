# talos-metal-agent

This repository contains the metal agent extension for Talos.

This repository builds the agent binary and its container image - the binary gets packaged into an extension
in the [extensions repository](https://github.com/siderolabs/extensions/tree/main/guest-agents/metal-agent).

## Development

When developing the agent, after doing your changes, run the following command to build Talos boot assets with the agent containing your changes:

1. Set up a `buildx` builder instance with access to the host network, if you don't have one already:

   ```bash
   docker buildx create --driver docker-container --driver-opt network=host --name local1 --buildkitd-flags '--allow-insecure-entitlement security.insecure' --use
   ```

2. Start a local image registry if you don't have one running already:

   ```bash
   docker run -d -p 5005:5000 --restart always --name local registry:2
   ```

3. Build Talos boot assets with the agent containing your changes:

   ```bash
   make image-boot-assets \
   OUTPUT_REGISTRY_AND_USERNAME=127.0.0.1:5005/siderolabs
   ```

   Hint: see `.kres.yaml` for more customization options.
   For example:
   - You can build against a different Talos version using the `IMAGER_TAG` variable.
   - You can customize the extensions repo references using `EXTENSIONS_*` variables.

   This command will build a container image with a tag like `127.0.0.1:5005/siderolabs/talos-metal-agent-boot-assets:v1.9.0-alpha.2-agent-198cabf-dirty`, with the following structure:

   ```text
   Permission     UID:GID       Size  Filetree
   -rw-r--r--         0:0      187 B  ├── cmdline-metal-amd64
   -rw-r--r--         0:0      203 B  ├── cmdline-metal-arm64
   -rw-r--r--         0:0      97 MB  ├── initramfs-metal-amd64.xz
   -rw-r--r--         0:0      78 MB  ├── initramfs-metal-arm64.xz
   -rw-r--r--         0:0      19 MB  ├── kernel-amd64
   -rw-r--r--         0:0      22 MB  └── kernel-arm64
   ```

4. Push this image to your local registry:

   ```bash
   make push-boot-assets \
   OUTPUT_REGISTRY_AND_USERNAME=127.0.0.1:5005/siderolabs
   ```

5. See the README in [Omni Bare-Metal Infra Provider](https://github.com/siderolabs/omni-infra-provider-bare-metal) repository to build a provider image containing these boot assets.
