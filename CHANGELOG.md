## [talos-metal-agent 0.1.7](https://github.com/siderolabs/talos-metal-agent/releases/tag/v0.1.7) (2026-09-24)

Welcome to the v0.1.7 release of talos-metal-agent!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/talos-metal-agent/issues.

### Contributors

* Andrey Smirnov
* Noel Georgi
* Mateusz Urbanek
* Maja Bojarska
* Utku Ozdemir
* Dmitrii Sharshakov
* Orzelius
* Erwan Leboucher
* Kevin Tijssen
* Mickaël Canévet
* Aleksei Sviridkin
* Edward Sammut Alessi
* buckaroo
* immanuwell
* kastakhov
* Andras Elso
* Artem Chernyshev
* Benoît Knecht
* Brian Topping
* Christian Korneck
* David Orman
* Dharsan Baskar
* Dima Aratin
* Dmitriy Matrenichev
* Dmitry Sharshakov
* Filip Boye-Kofi
* Fritz Schaal
* Immanuel Tikhonov
* Ivan Demchuk
* Jaakko Sirén
* Jonny
* Justin Garrison
* Konstantin Nesterov
* Loki San
* Louis Deconinck
* Mario Cole
* Max Makarov
* Maxime Bertin
* Nico Berlee
* Noel
* Oscar Wieman
* Pranav Patil
* Raphaël DUCOM
* Sacha Weatherstone
* Spencer Smith
* Zadkiel AHARONIAN
* dadbravo
* leppeK
* scmtble

### Changes
<details><summary>3 commits</summary>
<p>

* [`105b460`](https://github.com/siderolabs/talos-metal-agent/commit/105b4608499f569fcbafe416fe139854ffc00fdc) chore: bump deps, rekres
* [`3042207`](https://github.com/siderolabs/talos-metal-agent/commit/304220743addef3d7fd1b0d6ffc90a83df0d8841) fix: destroy MD arrays before wiping disks
* [`9691a62`](https://github.com/siderolabs/talos-metal-agent/commit/9691a62cc91df32c0459a053ae5b8bc032a0e342) docs: add agent guide
</p>
</details>

### Changes from siderolabs/talos
<details><summary>659 commits</summary>
<p>

* [`2f86b9d2a`](https://github.com/siderolabs/talos/commit/2f86b9d2a29b413deddd7122a8420b8913813615) release(v1.14.1): prepare release
* [`65f704ee1`](https://github.com/siderolabs/talos/commit/65f704ee16a95fa65cb6fe230db007be31aaf02a) chore: pass kernel version down to VEX generator
* [`63101af5b`](https://github.com/siderolabs/talos/commit/63101af5b6f9c7272c025340fe6a34e707f388f3) fix: prevent sandboxd signal dispositions leaking into services
* [`676276f72`](https://github.com/siderolabs/talos/commit/676276f72e9fe69dff4669502675b067a7b8feca) test: peer passively with the MetalLB speaker
* [`2b8b46dcd`](https://github.com/siderolabs/talos/commit/2b8b46dcd791d8a0e8587c6efac656aa21d8c743) test: stop ARP flux breaking the BGP VRF test
* [`357d6006a`](https://github.com/siderolabs/talos/commit/357d6006ac28b7addf5e1c4fd6358cd24dd9b5c1) fix: rebuild the BGP server when its VRF is recreated
* [`db0b5a17c`](https://github.com/siderolabs/talos/commit/db0b5a17c413d369eb307dfaaf5dc666bebe4f94) fix: guard against nil config document slices
* [`094741ff3`](https://github.com/siderolabs/talos/commit/094741ff34895ff4a7e2f74ea45978bce3ce0594) docs: clarify the kube-apiserver extra args and new config
* [`a13b560f4`](https://github.com/siderolabs/talos/commit/a13b560f4ddcad756e081054955330aecc98135c) docs: correct the UnattendedInstallConfig name in the schema
* [`166c4070c`](https://github.com/siderolabs/talos/commit/166c4070cea0cbcd1727e59380a03d50c4495f5b) fix: make --insecure reachable for talosctl meta subcommands
* [`aac106867`](https://github.com/siderolabs/talos/commit/aac10686725927d24fc8fa67bb82ddad22948dfd) fix: ignore apply config dry-run for try mode
* [`7a2c4e8cf`](https://github.com/siderolabs/talos/commit/7a2c4e8cf3d9b2d392bafab3a175a9f6ad8b0d45) fix: tighten the validation of v1alpha1 configs vs. migration
* [`aef64fa38`](https://github.com/siderolabs/talos/commit/aef64fa389d6b1ad3baa084e6a7a3ccd85487de7) fix: reconnect the WireGuard over gRPC tunnel after a failure
* [`d07a21ad2`](https://github.com/siderolabs/talos/commit/d07a21ad21f23e78e459442a2514332f44fc0b30) fix: drop logical links if they no longer declare as logical
* [`9996bc871`](https://github.com/siderolabs/talos/commit/9996bc87114670a3d3f7e1ba9d3109c523e702f2) fix: create GRUB bootloader ISOs only for BIOS
* [`0604432de`](https://github.com/siderolabs/talos/commit/0604432deb864399ee224d2c5c5f456774099bff) feat: add NixOS OVMF search path
* [`5cb44dcab`](https://github.com/siderolabs/talos/commit/5cb44dcab35ca17a59047203eaf6134269d12fbb) fix: wait for USB settle explicitly
* [`20dcd515a`](https://github.com/siderolabs/talos/commit/20dcd515ab48b54d963da94dde954cf9409e992a) fix: empty searchdomains dropped on merge
* [`1e3e3fe50`](https://github.com/siderolabs/talos/commit/1e3e3fe503c983c9e97c38434fb1932a137e19d9) chore: support correctly various disk types for the system disk
* [`7c2e0b113`](https://github.com/siderolabs/talos/commit/7c2e0b113a2313cb7b7c113533fa5f50f836a295) fix: notify about link alias changes
* [`e9a67e163`](https://github.com/siderolabs/talos/commit/e9a67e16380e06a5f2aeb6b7c4af5819d89c1345) chore: use the host page cache for the QEMU cluster disks
* [`0afebca70`](https://github.com/siderolabs/talos/commit/0afebca70d017f6dbc0a910ed9bc108a27ae6994) fix: use the final config version in upgrade-k8s
* [`0ad18bb55`](https://github.com/siderolabs/talos/commit/0ad18bb555d4be85eb3c93b2b7396ec81177ae5e) feat: bring in containerd 2.3.5
* [`04c49d8a3`](https://github.com/siderolabs/talos/commit/04c49d8a34e3a59ccf9c8fc164ede7d84eaafe52) feat: allow generating an ECDSA service account key in secrets bundles
* [`bb2cb91fc`](https://github.com/siderolabs/talos/commit/bb2cb91fc95b3e6f477cf716b3a7c04c6bf06352) fix(security): define the permissions the 6.18 kernel expects in the classes
* [`bf31b2811`](https://github.com/siderolabs/talos/commit/bf31b281115bd6207ba039f9f07c3795227fada3) fix: improve resilience of the action tracker against dropped conns
* [`00a0ea03c`](https://github.com/siderolabs/talos/commit/00a0ea03cd6a90309bd25408e091461722f2ec97) fix: set TCP keealive and user timeout on apid proxied connections
* [`6c065607e`](https://github.com/siderolabs/talos/commit/6c065607e4d05f25b871f8294cf8614f0bea1fbd) test: revert disabling PS/2 in QEMU
* [`9841e0b47`](https://github.com/siderolabs/talos/commit/9841e0b470f6302c4564e993dba3495b03ac4fae) docs: fix containerconfig.dependson examples
* [`a11a260e6`](https://github.com/siderolabs/talos/commit/a11a260e605175deb6e1d3d770a85cf4ad0b9492) feat: add USB LAN78XX drivers to the rootfs
* [`9f8277434`](https://github.com/siderolabs/talos/commit/9f82774346d4dda2962d2ed4a090307fa65e8ef6) fix: harden the code around kubelet's client certificate handling
* [`3260b1e1f`](https://github.com/siderolabs/talos/commit/3260b1e1f983eef144dfe620bbbbe55b9258a4ec) fix: resolve volume devices in shared selector helper
* [`084152592`](https://github.com/siderolabs/talos/commit/0841525924ac2ecb48b55ef58699e7860b58a21a) fix: create LVM physical volumes on the decrypted device
* [`7dabdeb43`](https://github.com/siderolabs/talos/commit/7dabdeb43cb7bbd76080a80513e973af2280bc86) feat: add xfrm interface module
* [`63963f7da`](https://github.com/siderolabs/talos/commit/63963f7daa28657c8f0b39cf1967b5524d4d1641) feat: sync pkgs/tools
* [`09681e895`](https://github.com/siderolabs/talos/commit/09681e895fe692c1a81e0e970d5a008dbe7cd518) fix: correct the bug with overlay assets in ESP being dropped
* [`9abd05af4`](https://github.com/siderolabs/talos/commit/9abd05af449ebf9cb1827648298291afce18d714) release(v1.14.0): prepare release
* [`ff772d64f`](https://github.com/siderolabs/talos/commit/ff772d64f1c1923abdfbbfea21429a552453af13) chore: mark release as stable
* [`5424dac73`](https://github.com/siderolabs/talos/commit/5424dac73757e908788e78b37343ea1661de261f) test: relax OOM test conditions
* [`9f6719422`](https://github.com/siderolabs/talos/commit/9f6719422ade07fd423dc3adeb0bb7d85278ae35) fix: normalize image reference when passing to image verification flow
* [`6ce58d308`](https://github.com/siderolabs/talos/commit/6ce58d308974e8e0d4ec24c1b16abd1164e953d3) fix: name diskSelector in UnattendedInstallConfig validation errors
* [`af4e11b50`](https://github.com/siderolabs/talos/commit/af4e11b501f3bd5c7645a6027010140a91151528) test: isolate base OCI spec test in BGP CLOS runs
* [`ff56b6d66`](https://github.com/siderolabs/talos/commit/ff56b6d667101d133c898e08cd8f172269cb733e) chore: update tools and pkgs
* [`8c8b9b45f`](https://github.com/siderolabs/talos/commit/8c8b9b45f3e53dab889da1ae3eb83df13989e845) fix: filter the output in talosctl
* [`9777a1bb4`](https://github.com/siderolabs/talos/commit/9777a1bb473d8beebfd2157b9b255f22d642e1af) feat: talos containers run with imagegccontroller
* [`dd566c9c7`](https://github.com/siderolabs/talos/commit/dd566c9c791885fec7284f5e277fce15d7967b7b) fix: container mount rshared propagation
* [`2f8931768`](https://github.com/siderolabs/talos/commit/2f8931768f5e1a8974f3c867d0c60763197e982d) refactor: don't skip container tests in short mode
* [`2f0fdd11e`](https://github.com/siderolabs/talos/commit/2f0fdd11e9a45085c15f8e4540ae15ce967aceb1) fix: ctr machinedaccss allowed roles
* [`64aef9053`](https://github.com/siderolabs/talos/commit/64aef9053187846f72990f8f4bd3fed8c5c674c6) feat: containers can connect to machined
* [`ce2148aa0`](https://github.com/siderolabs/talos/commit/ce2148aa0a5b039510e2c45cff9973bb09bf1fbc) feat: impl. container mount controller
* [`9105c5e1e`](https://github.com/siderolabs/talos/commit/9105c5e1eabaf76ffbd9bb726d04fbe89ee8eca4) feat: talos containers support for talosctl
* [`36fd44677`](https://github.com/siderolabs/talos/commit/36fd44677244feff117364a8e5931cf9e4295bf8) feat: impl. container runtime controller
* [`afc09d418`](https://github.com/siderolabs/talos/commit/afc09d41831be7dfab8b6ba2c0f35c09048cf202) refactor: make runner.runner context-native
* [`eff936cf0`](https://github.com/siderolabs/talos/commit/eff936cf0c36ef56a98dba69e3b21a248ed6cab2) feat: impl ContainerInstanceController
* [`95abc8bff`](https://github.com/siderolabs/talos/commit/95abc8bffc8bad9f4906ae8e65ab6427cd7cd864) fix: filter out passed metadata in API proxying
* [`10c4f1a8b`](https://github.com/siderolabs/talos/commit/10c4f1a8b3236e108bec5c69f132687c1430a72e) fix: treat desired roles empty as error in Talos API access
* [`807b692a1`](https://github.com/siderolabs/talos/commit/807b692a1bf399503256fd062a5a6fbfea68d501) fix: allow CSI volumes to be mounted with an SELinux context
* [`6249f3f4e`](https://github.com/siderolabs/talos/commit/6249f3f4e7675b371b02325b03fe4f42db99e838) test: skip filemap heap test under race detector
* [`87a031b7f`](https://github.com/siderolabs/talos/commit/87a031b7f0975c16ea9e26708a97efb75956fa72) test: libvirtd extensions
* [`e8b376365`](https://github.com/siderolabs/talos/commit/e8b376365146f99e1ff2db86eb17a001c10041b9) fix: validate received kubeconfig
* [`abce4c937`](https://github.com/siderolabs/talos/commit/abce4c937a23d5953bf498d449d2235bcc682dec) fix: use os.Root in the talosctl extract path
* [`1deae73dd`](https://github.com/siderolabs/talos/commit/1deae73ddcafafaa4ebfdadabfc3d394bb2300e4) fix: escape output in the talosctl dashboard
* [`5b153233c`](https://github.com/siderolabs/talos/commit/5b153233ca5bf9051bff02df161d33c1200f736d) chore: speed up unit-tests
* [`3715ce908`](https://github.com/siderolabs/talos/commit/3715ce908cf2ff78a09c5c665f52e713665a9864) fix: add checks for meta key in the API path
* [`138958053`](https://github.com/siderolabs/talos/commit/13895805315a5a8f2b96806f0d4100ceb9d598c6) feat: update COSI to 1.16.3
* [`bdc7f3053`](https://github.com/siderolabs/talos/commit/bdc7f30537a9447cfc22b7a832f1d4e4942b356b) fix: preserve shared libs symlinks
* [`7366edc90`](https://github.com/siderolabs/talos/commit/7366edc90103a8fd3444a62caa1442aa42b6e91f) chore: reduce verbosity of the grype scan
* [`c027a947f`](https://github.com/siderolabs/talos/commit/c027a947fc697fbf6e936eefe3a2d655f927ccbf) fix: exclude scheduler config from apiserver config version
* [`544fc52cb`](https://github.com/siderolabs/talos/commit/544fc52cbc4de3fd892219263fe904b354e81f97) feat: add host extension pre-shutdown hooks
* [`1d911581d`](https://github.com/siderolabs/talos/commit/1d911581d07da2dea08c9c0ae29a73ef1e721353) fix: route creation churning every 100ms
* [`baca68d0d`](https://github.com/siderolabs/talos/commit/baca68d0db21ce4c75b3301a0a2af25f16319c41) fix: drop dead legacy registries init from worker generation
* [`dec9dc3dc`](https://github.com/siderolabs/talos/commit/dec9dc3dc580276bf586516e2afb0f0736fd323f) chore: bump sigstore/rekor to v1.5.4
* [`dc9e7fea9`](https://github.com/siderolabs/talos/commit/dc9e7fea91f09c63dc245ef53796981f13875b1e) chore: add SPD5118 driver to squashfs
* [`ae4d8228e`](https://github.com/siderolabs/talos/commit/ae4d8228e5294e0a87fe4b91f111c9251617f4c9) fix: keep the time sync boot timeout across a syncer restart
* [`755900fcc`](https://github.com/siderolabs/talos/commit/755900fcc206b0b2a84c3921c4deee4e629af7d0) fix: preserve special modes when extracting files
* [`6c927a0ff`](https://github.com/siderolabs/talos/commit/6c927a0ff9f47f025fd1d475e1fd5b565b663c8b) fix: hostns etc mount
* [`3cdd0e748`](https://github.com/siderolabs/talos/commit/3cdd0e7482c2e2b036b2098e5baa8bd4f0ac0600) feat: bump kubernetes to 1.37.0
* [`03c567929`](https://github.com/siderolabs/talos/commit/03c567929b98b925528e111ef9829d52fb0cd077) refactor: log filter stays enabled on enter press
* [`b690b7ad5`](https://github.com/siderolabs/talos/commit/b690b7ad5f340353811a31669416ded7c8cca58a) fix: write the uploaded etcd snapshot atomically
* [`25a366162`](https://github.com/siderolabs/talos/commit/25a36616206260b7799f38704c150aff8ce1726e) chore: update multipath notes
* [`866db94f1`](https://github.com/siderolabs/talos/commit/866db94f1a855c65a5b35461c0bf73b537a76c5c) fix: drop gRPC client keepalive aggressive settings in the tracker
* [`74e88bd57`](https://github.com/siderolabs/talos/commit/74e88bd57361d2a7715fc3cb74121ea4fd8c072f) feat: support host extension services
* [`e828fe80a`](https://github.com/siderolabs/talos/commit/e828fe80a7d82274b77336ae21ae1a1192b07b37) fix: skip selinux label for read-only/detached/external mounts
* [`1be6787a5`](https://github.com/siderolabs/talos/commit/1be6787a5ed005e2f3e6ba7cb8208c52ad3dc2aa) fix: capture packets correctly on VLANs
* [`d16d765f7`](https://github.com/siderolabs/talos/commit/d16d765f7db1618429bf1558557f31f55878b6e7) docs: update CRI customization example
* [`1b31e1213`](https://github.com/siderolabs/talos/commit/1b31e12137165c260e0d5b224c6f76528401cb75) fix: use os.Root in the untar path
* [`91a6af032`](https://github.com/siderolabs/talos/commit/91a6af032bacd02473291b1d971b67de2557d2b2) fix: enable additional SELinux permissions for Cilium
* [`322de8bf2`](https://github.com/siderolabs/talos/commit/322de8bf2974b529ef676fd6be1746aaf2c3a74a) fix: cache filemap layers on disk
* [`414a1d463`](https://github.com/siderolabs/talos/commit/414a1d46333bc4f7484381cf513b94fde52a0636) release(v1.14.0-rc.2): prepare release
* [`a740329b9`](https://github.com/siderolabs/talos/commit/a740329b9de0de336737fc1167ebb2954e96d0d9) feat: bump kernel, containerd and go
* [`0048cd3f4`](https://github.com/siderolabs/talos/commit/0048cd3f4ada8933fb13a7d3784b5d6dfbe560ad) chore: bump vulncheck dates
* [`acc89cbef`](https://github.com/siderolabs/talos/commit/acc89cbeff95cb661e4344be76c9b1c10512fd3e) fix: use default terminal theme colors in talosctl dashboard
* [`d78c61c82`](https://github.com/siderolabs/talos/commit/d78c61c82ecbc1f0a5bc03b68879a0d5d9133c7f) fix: don't create new client in dry-run mode
* [`7276d54ef`](https://github.com/siderolabs/talos/commit/7276d54ef6b135450c3936684bc35ef37f1ed33e) fix: preserve selected sd-boot entry on upgrade
* [`68a436656`](https://github.com/siderolabs/talos/commit/68a436656733c4c59451878f01b77afc60b58d61) fix: use the UKI command line when the config has no install section
* [`e32a266d9`](https://github.com/siderolabs/talos/commit/e32a266d99b8cf9d4a3cb83ed15a4ed9a7b4e603) fix: persist in-memory meta on fresh install
* [`35c8f172c`](https://github.com/siderolabs/talos/commit/35c8f172cffb602a3c024db983987dc54c0c2449) fix: drop lockdown=confidentiality default for 1.14+
* [`977199548`](https://github.com/siderolabs/talos/commit/9771995485e9389b78aad271e14673b9b8ee005e) fix: reduce stalls in the etcd member promotion cycle
* [`4c381dfae`](https://github.com/siderolabs/talos/commit/4c381dfaecd733bdcf79045fb135e390ac4f5b4a) feat: update CoreDNS to 1.14.7
* [`95135f804`](https://github.com/siderolabs/talos/commit/95135f8042c48752e4b863463ed34fade4b27d2f) feat: update Kubernetes to 1.37.0-rc.1
* [`7ac3cad5e`](https://github.com/siderolabs/talos/commit/7ac3cad5e6319f9132cf97f210ae7873931b1e72) fix: watch IPv6 route changes in RouteSpecController
* [`d31a66599`](https://github.com/siderolabs/talos/commit/d31a665992d6d0d6cb5ee74efd9d9ce2b982a759) fix: enable SELinux to work with overlays
* [`3cc04997f`](https://github.com/siderolabs/talos/commit/3cc04997fe2eb5788a053634c5170a1ab377bd11) fix: move the spike information field of the time.Status resource
* [`52e874785`](https://github.com/siderolabs/talos/commit/52e8747858416265e8c37ec992430e3352b58ec9) feat: log peer address in gRPC request log
* [`5f0005c25`](https://github.com/siderolabs/talos/commit/5f0005c258d126ab3880371997ee993e7bfc151f) feat: talosctl dashboard log filtering
* [`5d13cc0e7`](https://github.com/siderolabs/talos/commit/5d13cc0e7c7863f8a46906903c3fea524fdc6c21) fix: truncate files replaced by system extensions
* [`afc8d952d`](https://github.com/siderolabs/talos/commit/afc8d952d1787886d4f1da44cfc861124ae12c4c) fix: skip target settings for external volume mounts
* [`6e83aece4`](https://github.com/siderolabs/talos/commit/6e83aece435a14d453178d30f57e9c7ed0dd0478) test: use tiny hostns debug image
* [`2f371abd5`](https://github.com/siderolabs/talos/commit/2f371abd502877e4a12eb27c1d9b5ebcb5ac2e7d) fix: support try mode apply without prior config
* [`d3544c2e3`](https://github.com/siderolabs/talos/commit/d3544c2e3befb23b315adc6006ac18bf17a298be) fix: adjust cluster size for VFAT on ISO
* [`3909ca9b2`](https://github.com/siderolabs/talos/commit/3909ca9b238fc4a11cf1f61ebbd7278802c39c0d) feat: impl ContainerImageController
* [`c563615ed`](https://github.com/siderolabs/talos/commit/c563615ed7da05a6639f1e4d31c50ebc1fa39d26) test: add cachefilesd to the test matrix
* [`d36a20e84`](https://github.com/siderolabs/talos/commit/d36a20e842687ece37781950ce12cf476baf0ac0) fix: apply directory user volume mount security
* [`78efbb413`](https://github.com/siderolabs/talos/commit/78efbb413863a902319c2f0b3668e310280cc7e3) fix: install conntrack handler in accept ingress firewall mode
* [`7a84d742b`](https://github.com/siderolabs/talos/commit/7a84d742b3d53451cbaa1eb91d9b12403e80161c) release(v1.14.0-rc.1): prepare release
* [`89ea1af12`](https://github.com/siderolabs/talos/commit/89ea1af1238b5356df49e815d0bfea8fa4b6ec33) chore: ci uses extensions release-1.14
* [`b881ccee1`](https://github.com/siderolabs/talos/commit/b881ccee11eff97ce5079cfea93970d0d99a1386) chore: backport go 1.26.6
* [`38a88d7a5`](https://github.com/siderolabs/talos/commit/38a88d7a586d0b081997ca6f4d085f4d8da0ed39) fix: share IPC namespace with the host for extension services
* [`26d4d389e`](https://github.com/siderolabs/talos/commit/26d4d389e26ad3790f31fabc49099aafd082bdbf) fix: use v1.13 config to test downgrade failure
* [`6b6a4cc01`](https://github.com/siderolabs/talos/commit/6b6a4cc01f785bb36bf610b04197945615ba7e31) fix: provide read-only random seed in the ISO
* [`250865dec`](https://github.com/siderolabs/talos/commit/250865decc548ac1271d6165df9f219e9bd89c47) chore: bump go deps
* [`c0613dfe0`](https://github.com/siderolabs/talos/commit/c0613dfe0f745124a6fbb61578232263e33ed3b5) chore: rekres
* [`18e26bbb0`](https://github.com/siderolabs/talos/commit/18e26bbb04865779a22d49f1222ad7d2082d0a0a) chore: bump tools and pkgs to v1.14.0
* [`b2262db3b`](https://github.com/siderolabs/talos/commit/b2262db3b0f8bef887c9fb802ba8c9243ba85662) fix: respect authentication-config extra arg for legacy config
* [`1407a242e`](https://github.com/siderolabs/talos/commit/1407a242eebb742d9480c2f2c6db6e2cdf066604) test: restore Talos 1.13 ephemeral policy skip
* [`16a147dc7`](https://github.com/siderolabs/talos/commit/16a147dc76ae11a7fc3b0f73c54603fa876b7466) feat: allow passing extra QEMU arguments per node
* [`d6db2fd44`](https://github.com/siderolabs/talos/commit/d6db2fd449ec7db6698d055f51ba44a2d3c27e5d) fix: render absolute CRI registry TLS paths
* [`ee18fb424`](https://github.com/siderolabs/talos/commit/ee18fb4240a8744277c416c8772f55c98821a926) fix: data race in Never condition closures
* [`dc77862dc`](https://github.com/siderolabs/talos/commit/dc77862dcc53d1a778e7904025df39c8ee08c899) fix: show installer output on upgrade failure
* [`7fbe57f8c`](https://github.com/siderolabs/talos/commit/7fbe57f8c9e83c9343f18eea55433f18d915a17c) fix: build native custom linter for lint targets
* [`51f96d6bb`](https://github.com/siderolabs/talos/commit/51f96d6bb8bcf92996bfe07fbaba2d9975b75cd6) fix: rework bootloader install and image generation
* [`82fe416a4`](https://github.com/siderolabs/talos/commit/82fe416a4a380555a5801bc8b9785b189c8e9cd8) test: fix ephemeral check for talos < 1.14
* [`c96fdc764`](https://github.com/siderolabs/talos/commit/c96fdc7643498c17bb3524ee0d4d371631d9ce2c) chore: dependency updates 2026-08-11
* [`f86ad4d77`](https://github.com/siderolabs/talos/commit/f86ad4d77d4ef090f9eccd8de338598b218b2f9a) chore: bump flannel to 0.28.9
* [`a23c6b9f5`](https://github.com/siderolabs/talos/commit/a23c6b9f5ee42407ee4527defac33b60595bbf01) test: retry k8s node discovery
* [`2666f13dc`](https://github.com/siderolabs/talos/commit/2666f13dce752ab90a796c0fab1068b455429480) fix: flag all devices backing system disk, not just top one
* [`c166e8863`](https://github.com/siderolabs/talos/commit/c166e8863366a637233f2cc946eda1c4ad8bf1d9) feat: run full md boot integration suite
* [`a81e32c97`](https://github.com/siderolabs/talos/commit/a81e32c971325dc9bd2e1ac95f6645b3561e0e3a) feat: add alibabacloud platform
* [`c14b43a9b`](https://github.com/siderolabs/talos/commit/c14b43a9bd1191bf6bbaed7798ff3a5493729767) fix: use less memory on the install path
* [`87bfa703b`](https://github.com/siderolabs/talos/commit/87bfa703bddb71a23e646873035c964fe158d63d) fix: size the receive/send buffers for nftables netlink
* [`83c132e6a`](https://github.com/siderolabs/talos/commit/83c132e6afb7a0f05cc37d485c5226b50eee7588) docs: update volume mount secure options
* [`cd0359d94`](https://github.com/siderolabs/talos/commit/cd0359d94cd80ce20d8bc81b823de745e498205d) feat: impl ContainerConfigController
* [`54b11fd9c`](https://github.com/siderolabs/talos/commit/54b11fd9c72364cfa9903dab356e9afc23c49c99) test: fix the flakiness in image pull in provision-3 pipeline
* [`0303f3181`](https://github.com/siderolabs/talos/commit/0303f3181c64445634bda155c7abc4ce7d5aba3f) fix: preserve connected prefixes in BGP advertisements
* [`6fa811a0d`](https://github.com/siderolabs/talos/commit/6fa811a0d426b431e958f6f9ec66556573eb2508) fix: drop `noexec` for KUBELET, EPHEMERAL and CRI
* [`25d8c0a51`](https://github.com/siderolabs/talos/commit/25d8c0a51ed4cc9c1f5176d8ca0cf1313b719615) feat: update Kubernetes to 1.37.0-rc.0
* [`63ef4df99`](https://github.com/siderolabs/talos/commit/63ef4df995ef4d08cbe9c6c4d58c34d7cbf012af) fix: keep host DNS enabled for partial machine config
* [`b00c06b35`](https://github.com/siderolabs/talos/commit/b00c06b358671586cd4aaa1a227811cffc962c5d) fix: support image factory URLs with explicit port
* [`825844afd`](https://github.com/siderolabs/talos/commit/825844afd31a4ba21975f5c408cb39826e0ca546) chore: build custom-gcl for the host OS/arch
* [`54673711f`](https://github.com/siderolabs/talos/commit/54673711fd6c40e4791acbb46a06e78ca321a4a0) feat: tag published cloud images with a build type
* [`3abe89e00`](https://github.com/siderolabs/talos/commit/3abe89e00020ab4a31fcde3c95c480a87c7f84ba) fix: avoid small panics
* [`c75361127`](https://github.com/siderolabs/talos/commit/c75361127143c5b081e7e39d9b15d44be6d54012) test: wait for CRI runtime spec overrides
* [`b0b77bcae`](https://github.com/siderolabs/talos/commit/b0b77bcae6ec78266a1a19454fb95cdea933fff4) fix: recover router advertisement sender panics
* [`6e3d0c55c`](https://github.com/siderolabs/talos/commit/6e3d0c55cb733b463d890d4c6914dfab7b256a60) fix: image pull via the API should not have timeout or retries
* [`a150503d5`](https://github.com/siderolabs/talos/commit/a150503d5cd533fc0d111327f6f7a5e998cffb9c) fix: collapse machined/apid logs with authz messages
* [`30ae29b1b`](https://github.com/siderolabs/talos/commit/30ae29b1b1333be6fdf947fa5a6a1062a77ea3eb) test: skip iptables compatibility test in enforcing
* [`8ad52d6dd`](https://github.com/siderolabs/talos/commit/8ad52d6dddeb15bc05a588d7f188dbf5fa61a846) fix: wait for router advertisement senders on shutdown
* [`a0b021e36`](https://github.com/siderolabs/talos/commit/a0b021e363653dc53b8852ff0e31ed494bf213aa) chore: update go-talos-support to 0.3.1
* [`85e97a55f`](https://github.com/siderolabs/talos/commit/85e97a55f2cf2ba0c9796424cb1c5ffda0dfc914) fix: panic when KubeProxy is disabled without image override
* [`969098c91`](https://github.com/siderolabs/talos/commit/969098c915c4cd147b96fd107972a4bf69e7dcf3) fix: bring in fixed Linux kernel with iptables xt modules
* [`7d01fc936`](https://github.com/siderolabs/talos/commit/7d01fc936bcbc5e1f8c70e7b94463c7e2f80b3b2) fix: ignore unmanaged address flags in AddressSpecController
* [`9ffa772ba`](https://github.com/siderolabs/talos/commit/9ffa772ba58b035171a1fc53137cfd0c467f4d2a) feat: support experimental k8s-less and etcd-less mode
* [`aab940f6a`](https://github.com/siderolabs/talos/commit/aab940f6aef58f913120c4bb3af20d24c4970b3f) chore: update kernel to 6.18.42
* [`6e45d0520`](https://github.com/siderolabs/talos/commit/6e45d05204c6a9c0b9cd3dbb4e7490ce64177340) fix: ignore HostDNS IPv6 address in node addresses
* [`4b89c911f`](https://github.com/siderolabs/talos/commit/4b89c911f0048a2c1fd1b3b65291eaf48f2a40d7) feat: add support for static VLAN configuration to the dashboard
* [`e225ff060`](https://github.com/siderolabs/talos/commit/e225ff060d6263803218da47d8642a923f5084dd) fix: keep host dns enabled during bootstrap
* [`02c87ba96`](https://github.com/siderolabs/talos/commit/02c87ba9668bc14e6d6426902997d0fae8420977) fix: record PID properly when under sandboxd
* [`fe2b5b430`](https://github.com/siderolabs/talos/commit/fe2b5b4301df0f6a5879d58576b319473793ae14) test: apply correctly hydrophone timeouts
* [`ab42416e0`](https://github.com/siderolabs/talos/commit/ab42416e0c9f836710874b15251b0f2076b81ef7) test: use new multi-doc external manifest
* [`f3974dc4b`](https://github.com/siderolabs/talos/commit/f3974dc4bc408e3880d4d99460e02a9c82676fca) fix: a condition when unattended install status can flip to installed
* [`e20509bae`](https://github.com/siderolabs/talos/commit/e20509bae0fcea526111cb0832a1985a7d8dba46) test: update airgapped patches
* [`0d1f6e576`](https://github.com/siderolabs/talos/commit/0d1f6e576bfb078e7ea29be99adb43968a82ba82) release(v1.14.0-beta.1): prepare release
* [`1ab88f743`](https://github.com/siderolabs/talos/commit/1ab88f743c46e8ccedbb76ca8c94aec9a1e0df87) fix: load the raid1 module for configured MD arrays
* [`9e8568d26`](https://github.com/siderolabs/talos/commit/9e8568d2680f961d7b8075433000bd9dfdae6f11) fix: use inmem containerd for installs/upgrades
* [`8cf28da5f`](https://github.com/siderolabs/talos/commit/8cf28da5f658b500b0f99f899a5ca9b9b3086ec5) fix: preserve kmsg reconciliation after config changes
* [`311b6fde9`](https://github.com/siderolabs/talos/commit/311b6fde9773ec07338794c05f170c5492ae3f47) fix: make reset during boot phase more robust
* [`945d1cdab`](https://github.com/siderolabs/talos/commit/945d1cdab48b6386defff30ab142725da82c3565) feat: bump kernel to 6.18.41
* [`4e77d40e8`](https://github.com/siderolabs/talos/commit/4e77d40e8c535d7c845df61ea112504c4dc33817) fix: restore the systemd-timesync best-sample check in spike detection
* [`4444a187d`](https://github.com/siderolabs/talos/commit/4444a187d6a8e58cbdcbdc08e85dc0f3a6ff1dfd) test: inspect host mount state from the current thread
* [`b7c0497b0`](https://github.com/siderolabs/talos/commit/b7c0497b09fdfb0acf3c6c304bacdc2bc5b7d23f) fix: race between vol.cfg. and vol.mgr. ctrl
* [`ef9a091ec`](https://github.com/siderolabs/talos/commit/ef9a091ec5c1bc2099bfbe25c63ffc63a8cede7f) feat: allow attaching extra disks to controlplane machines
* [`a418c0e1e`](https://github.com/siderolabs/talos/commit/a418c0e1ea2e3fc9c7e9efad82eae78eef28a0b8) test: fix CRI restart event assertion race
* [`a9bfdbdd4`](https://github.com/siderolabs/talos/commit/a9bfdbdd4b0d03f3b8190e136ea4feef25648a71) fix: ignore insecure-only imager assets
* [`54e3b20e8`](https://github.com/siderolabs/talos/commit/54e3b20e849c0f2bf0ca2ce0e6d6e96358fad7cd) fix: hold the darwin vmnet bridge open for the network lifetime
* [`bc59389fa`](https://github.com/siderolabs/talos/commit/bc59389fa2d43c70c83169af1d84eea90215d233) fix: drop the OOM config rule about overall system memory PSI
* [`69be56ea9`](https://github.com/siderolabs/talos/commit/69be56ea93a0bbcb67750a3e626211258e495daf) chore: add some initial set of libvirt SELinux policies
* [`6170ad8b9`](https://github.com/siderolabs/talos/commit/6170ad8b9e0c96d71bd4c73ad31b9e123a53a099) fix: validate kubespan & discovery config correctly for multi-doc
* [`6e58c6d5a`](https://github.com/siderolabs/talos/commit/6e58c6d5a37652e3c3e5a7163ee5f0cfbb162d53) test: fix ded. vol. plumbing in integration tests
* [`b644d1640`](https://github.com/siderolabs/talos/commit/b644d1640c4c6822a2d6f7fcd018e5348cb8ba00) fix: allow directory-backed volumes in reset api
* [`28e7a8742`](https://github.com/siderolabs/talos/commit/28e7a87424ab5ede581d38263f5bb6b1af99a98e) feat: bump etcd to v3.7.1
* [`2c657c224`](https://github.com/siderolabs/talos/commit/2c657c2243706462d54731624547bc77ba184b35) feat: bump kernel to 6.18.40
* [`b1bdc8c07`](https://github.com/siderolabs/talos/commit/b1bdc8c071bef8928cecc9f8891239763b0fd0b2) test: restart qemu process several times on startup failures
* [`7533057a7`](https://github.com/siderolabs/talos/commit/7533057a7df431c4f6854ed71e284311a31e006f) feat: support route imports between BGP instances
* [`a94783704`](https://github.com/siderolabs/talos/commit/a9478370497fff42a181753be7d1c26f30ab732c) docs: remove duplicated docs in the markdown for CLI
* [`6f17c5033`](https://github.com/siderolabs/talos/commit/6f17c5033b9edc722adc0dc85c0380080097281a) fix: verify the public key signed images correctly
* [`570fe34f0`](https://github.com/siderolabs/talos/commit/570fe34f085e58d76c58d413cd7b0ee5eeb12dc0) fix: generate backwards compatible etcd encryption config
* [`f7790816e`](https://github.com/siderolabs/talos/commit/f7790816ef42b4399715fdaed85a8a3ba0090473) fix: use context without cancelation for etcd locks
* [`67e61ef30`](https://github.com/siderolabs/talos/commit/67e61ef30967f94578aa5ac2349bc12d057a339e) feat: add the fs_scrub controller
* [`1c156458a`](https://github.com/siderolabs/talos/commit/1c156458a822356587440ebbd530a6e132baf91a) fix: override DHCP search domains via explicit ResolverConfig domains
* [`fd8dbd8a1`](https://github.com/siderolabs/talos/commit/fd8dbd8a103c2ae460843eb427f6b51fedb99c1a) fix: skip pod check if desired number of pods is zero
* [`9aede5429`](https://github.com/siderolabs/talos/commit/9aede542993dec2a6567719a23b7e86aefba6efb) feat: add kubeimportlinter for versioned k8s imports
* [`ae93d1462`](https://github.com/siderolabs/talos/commit/ae93d14624927a802042ed7217875417b9938291) fix: redact resource specs in the merge controllers
* [`fc5743cd0`](https://github.com/siderolabs/talos/commit/fc5743cd044749ce8bc9c852d2f6c9f7894bce0c) feat: add GrubUseUKICmdline install option
* [`6bba77724`](https://github.com/siderolabs/talos/commit/6bba7772421238113998ce4022069e4c47d859a8) feat: add named native BGP instances
* [`c5ab22f1d`](https://github.com/siderolabs/talos/commit/c5ab22f1da569e1a2ead20cc15010445af0f1e8e) feat: move Talos API access from Kubernetes config to multi-doc
* [`b1abd9c03`](https://github.com/siderolabs/talos/commit/b1abd9c0334b77f5a684d0d154b286b1db7b278d) fix: split the up/finished service events
* [`644ecbc66`](https://github.com/siderolabs/talos/commit/644ecbc66c7c8c810edc321b2dcbf49091d56bbf) feat: add discovered volumes status controller
* [`6be2b1384`](https://github.com/siderolabs/talos/commit/6be2b138439b92b87710cfff2176784c549dc933) feat: add veth pair support
* [`7514401de`](https://github.com/siderolabs/talos/commit/7514401def9cd58d29aadb3d9508fcfc15a55cc8) fix: drop the controlplane static pod change
* [`9a521f667`](https://github.com/siderolabs/talos/commit/9a521f66799f0bd8613b533da956ce21db1ae73e) fix: preserve container tasks across CRI restarts
* [`9048d4157`](https://github.com/siderolabs/talos/commit/9048d41577272e35036517975d9dc14bfea0b54b) fix: fix a nil-map assignment panic in configpatcher
* [`45eaf2037`](https://github.com/siderolabs/talos/commit/45eaf20377ca60de3c1c61477c0aa1a7a0ea3e82) release(v1.14.0-beta.0): prepare release
* [`7e58e0442`](https://github.com/siderolabs/talos/commit/7e58e0442fcd2eb17c9e0d229b55d19c6ac54f36) feat: add dedicated CRI configuration documents
* [`076c38136`](https://github.com/siderolabs/talos/commit/076c381362ae202f2c9f15cadf82107c894cf7fd) fix: race with PCR extensions and volume unlock
* [`88884194c`](https://github.com/siderolabs/talos/commit/88884194cd2dbe98650a71f1abe6699e3c6ad77a) fix: teardown ephemeral mount request during reset
* [`c793bcbf5`](https://github.com/siderolabs/talos/commit/c793bcbf567a776ddce8e57d714846d66be1be0a) fix: configure bonds during initial link creation
* [`9b3bf6e51`](https://github.com/siderolabs/talos/commit/9b3bf6e5170d95d86f7f10d8e688bd5ffff8cf38) fix(talosctl): prevent duplicate QEMU config server ports
* [`fa6cd1ca8`](https://github.com/siderolabs/talos/commit/fa6cd1ca86c95ee3d22747080912f602461e8e5d) fix(machined): preserve health when services reach running
* [`9d5554e69`](https://github.com/siderolabs/talos/commit/9d5554e6978b731d1726191be51f6e4656bb7fb1) fix(machined): wait for host namespace commands through reaper
* [`fc08533bf`](https://github.com/siderolabs/talos/commit/fc08533bfa954651a449c7385b9a31603ed6f9f0) chore: update dependencies
* [`c08863cdd`](https://github.com/siderolabs/talos/commit/c08863cddf3dcaa97bd4a3227c8f1eb52fc4f4ee) feat: provide different heuristics for xfs allocation groups
* [`e955d9bd7`](https://github.com/siderolabs/talos/commit/e955d9bd7c0834296459f40b148dd9a5850fb5b2) feat: update CoreDNS to 1.14.6
* [`c3f757f9e`](https://github.com/siderolabs/talos/commit/c3f757f9e19723e8c676ee421462b54001a2a12a) feat: update Flannel to 0.28.8
* [`fada0d960`](https://github.com/siderolabs/talos/commit/fada0d960cbc907407241cfa14fec1a37dd0c184) fix: provide non-sensitive KubeletStatus resource
* [`c68085286`](https://github.com/siderolabs/talos/commit/c6808528628840b36c860813a26974f49dd7120c) fix: volume mount race (third attempt) around service restart
* [`b185752e5`](https://github.com/siderolabs/talos/commit/b185752e57d3ee1a26fc4cb2a00e94fa7a286e6a) feat: refactor KubePrism config into multidoc
* [`499d4ebf9`](https://github.com/siderolabs/talos/commit/499d4ebf92e4ace7ebd218b3d8223c7e43cd28ae) test: update Calico in canal reset test
* [`5b6ed0068`](https://github.com/siderolabs/talos/commit/5b6ed00687c8c5031253c1086188a9be01f21597) test: add a test for kata-qemu runtime class
* [`1a075383a`](https://github.com/siderolabs/talos/commit/1a075383a2af988c5e25207efd6720649557e710) feat: allow "duplicate" kinds in the config patches
* [`06943be9e`](https://github.com/siderolabs/talos/commit/06943be9ee88b2a95bfadf2b6c87b646191ae5b7) feat: update Kubernetes to 1.37.0-beta.0
* [`01f2a1423`](https://github.com/siderolabs/talos/commit/01f2a1423290f4e48c89f8daac8e5770f7edeadf) fix: preserve trailing rate-limited trigger events
* [`46fab8057`](https://github.com/siderolabs/talos/commit/46fab8057449dcbdfe04fb6a354b666c69538f0c) test: stabilize AWS readiness and Talos 1.13 QEMU config
* [`a26ac746d`](https://github.com/siderolabs/talos/commit/a26ac746da68b83c7c86452ca5ff9e71992637c1) feat: move static pods and manifests into multi-doc
* [`67464cbef`](https://github.com/siderolabs/talos/commit/67464cbefc67d878c310f04bbb9432ae0854fe5f) fix: update the vulnerability dates and description
* [`4920ee06f`](https://github.com/siderolabs/talos/commit/4920ee06fbacdcddb0632b8a83ed03d1a368fbce) feat: update Linux to 6.18.39
* [`286fa8006`](https://github.com/siderolabs/talos/commit/286fa8006f7c78cddde9f78409e9f1ce563ec0be) feat: include CA into kube-apiserver serving certificate
* [`6d65e223b`](https://github.com/siderolabs/talos/commit/6d65e223b36c3d87346940f18869e26326ede0fe) feat: drop kubernetes flexvolume mounts
* [`4935e9452`](https://github.com/siderolabs/talos/commit/4935e94523f8ea63744b5524c8324e6108182d41) feat: refactor kubelet's config into `KubeletConfig`
* [`241bd0ff1`](https://github.com/siderolabs/talos/commit/241bd0ff1913a6f044b270c34ff0939241e328e3) feat: custom cfg for system volumes (cri, kubelet, etcd)
* [`ea9557816`](https://github.com/siderolabs/talos/commit/ea95578160e631cc0130e0a8f2771c003ef0723e) fix: talosctl build
* [`c2b763608`](https://github.com/siderolabs/talos/commit/c2b763608d66cf90cfac9c51fe63f9788207d4a2) feat: add UFSHC and some other modules
* [`2193b5781`](https://github.com/siderolabs/talos/commit/2193b57813da84ecec672506baded1fe1cf8e2e4) feat: native BGP support via embedded GoBGP
* [`2e42c5900`](https://github.com/siderolabs/talos/commit/2e42c590031da01039926bb7ff5815629b5040af) fix: add ca-certificates to talosctl
* [`0f55e1f05`](https://github.com/siderolabs/talos/commit/0f55e1f055ee7f149c12fced5aaa60865721a7e8) feat: refactor Kubernetes configs into `KubeNodeConfig`
* [`6efdc8f71`](https://github.com/siderolabs/talos/commit/6efdc8f71444b8245116e8a496d51eafe0ac53c9) fix: zero MD superblock via block wipe on destroy
* [`f78f5e5a1`](https://github.com/siderolabs/talos/commit/f78f5e5a12996d0d2461c38c218a83ebbfe0667a) fix: vrf sorting
* [`77385181a`](https://github.com/siderolabs/talos/commit/77385181ac4f5a71eaa0a40113fe90452005923c) fix: oom podruntime protection
* [`c1184d38e`](https://github.com/siderolabs/talos/commit/c1184d38ef11aa16f4dbc0b7709f024286cddbe6) feat: update to runc 1.5.1
* [`4bff7eb90`](https://github.com/siderolabs/talos/commit/4bff7eb90cfd7997b072b1eaec246501975778b7) feat: support reboot and sync for remote provisioner
* [`c791fa8c0`](https://github.com/siderolabs/talos/commit/c791fa8c03451050003465d8fc0d3a844a831802) feat: add host-namespace debug profile
* [`e370e40b7`](https://github.com/siderolabs/talos/commit/e370e40b7eb134e116d06c93f2bf00b33e04a39c) feat: implement KubeClusterConfig
* [`37c78bfc0`](https://github.com/siderolabs/talos/commit/37c78bfc053df5aea78299da2408e40045b2b407) fix(ci): skip ephemeral noexec test on 1.13
* [`0ab6695e6`](https://github.com/siderolabs/talos/commit/0ab6695e6c4a794633db6e4160eb5243983f7562) feat: update Kubernetes to 1.37.0-alpha.3
* [`443ca17e1`](https://github.com/siderolabs/talos/commit/443ca17e1b8a04fff19a90861ae325a14415ed26) test: bump test dependencies
* [`c4242088b`](https://github.com/siderolabs/talos/commit/c4242088b718f93c1c99c42b34e56e74e65663cc) fix: enable `noexec` for EPHEMERAL only for new machines
* [`fc9f72648`](https://github.com/siderolabs/talos/commit/fc9f726484764dcb181bb569985e5e717cacfc36) feat: bump CoreDNS, Flannel
* [`352b1bdeb`](https://github.com/siderolabs/talos/commit/352b1bdeb70c451bb8fc6db947efe0f575906113) fix: use symlinks for init aliases
* [`883775a9e`](https://github.com/siderolabs/talos/commit/883775a9ef30e58e7a4fa28027e39edcdedbe218) fix: move sandboxd into a separate cgroup
* [`099a2ceda`](https://github.com/siderolabs/talos/commit/099a2ceda72ab5a214d3f8682564dffc3dc74c71) fix: remote provisioner name
* [`ff67aaf32`](https://github.com/siderolabs/talos/commit/ff67aaf3254680e46e12a17616c8a227e2df7f4a) feat: bump go dependencies
* [`79c0c5414`](https://github.com/siderolabs/talos/commit/79c0c5414e36455ca953345f5e15a4146cc1a7f2) feat: add iommufd as a kernel module
* [`f34e93fe2`](https://github.com/siderolabs/talos/commit/f34e93fe255583d744f9d7d436e60813cc1c8752) fix: do proper backoff for NTP Kiss-of-Death responses
* [`a3e644d8d`](https://github.com/siderolabs/talos/commit/a3e644d8dd4e61019219d6f1860bb7ae30b59b30) chore: bump tools and pkgs
* [`efa88f2f6`](https://github.com/siderolabs/talos/commit/efa88f2f626597468b72c1a944e12dae7154025b) fix: flaky tests
* [`17a134711`](https://github.com/siderolabs/talos/commit/17a134711d9ba1a674ed20a09d6183b546cf734b) feat: add ignoreRoutes option to DHCPv4 config document
* [`2519bf231`](https://github.com/siderolabs/talos/commit/2519bf231a8a0bfb35b6dbf50139c430f983baef) fix: make audit restartable
* [`54b4bbc03`](https://github.com/siderolabs/talos/commit/54b4bbc03eff42f3391929edacc5c3da62e92169) fix: provide correct handler for Ctrl-Alt-Delete sequence
* [`87e126ab7`](https://github.com/siderolabs/talos/commit/87e126ab75b88687645bd7b4f74aaedaace6a8f3) feat: isolate cri, kubelet and pods in a sandbox namespace
* [`3fb8f4e9e`](https://github.com/siderolabs/talos/commit/3fb8f4e9eec29d259a2de087ac8f9bca014b3f85) fix: avoid image cache mount request churn
* [`9753fc27f`](https://github.com/siderolabs/talos/commit/9753fc27fc2c3ff68c08fb135bee2aa9cd180cf3) fix: e2e test flakes
* [`f756ff232`](https://github.com/siderolabs/talos/commit/f756ff232ba33e6df3d7e8f899700002e0d16b65) feat: kubenetworkconfig supports per-node pod cidr configuration
* [`b42c42976`](https://github.com/siderolabs/talos/commit/b42c429764ccb5b2b952d9087f8dbdd72be7ce44) fix(ci): fix more flaky tests
* [`5d97eccdf`](https://github.com/siderolabs/talos/commit/5d97eccdf32ad1950f8b37703db697d96d3a2d8c) feat: bring in ifb.ko module
* [`6769a1d5c`](https://github.com/siderolabs/talos/commit/6769a1d5c310a280d11e623b1f99324ca85f2afc) fix: terminate log persistence a bit harder
* [`98cce792f`](https://github.com/siderolabs/talos/commit/98cce792f6250363191bc9a264a92cba73ccaf9e) fix(ci): extensions test
* [`057d554d2`](https://github.com/siderolabs/talos/commit/057d554d2f5408246a821e5f1d307a830300273e) test: assert dm transport for device-mapper disks
* [`9fd16a21e`](https://github.com/siderolabs/talos/commit/9fd16a21e3b8462f53017c06aa17610324796c54) feat: bump etcd to 3.7.0
* [`3048eeb23`](https://github.com/siderolabs/talos/commit/3048eeb23e6ecb1641fd2a375bc50c65c7959918) feat: support booting from MD RAID1 array
* [`e1fc7a4a1`](https://github.com/siderolabs/talos/commit/e1fc7a4a129f40a6d9735166f41c86d6c2fad573) fix: do not block volume lifecycle teardown on failed user volumes
* [`147dea148`](https://github.com/siderolabs/talos/commit/147dea148b19c7984cba6e9ae870e46faac4493d) feat: add --no-reboot flag to upgrade cmd
* [`1b23b11fc`](https://github.com/siderolabs/talos/commit/1b23b11fc33fad9db309653b5a923ff7325e0025) chore: update pkgs and tools
* [`bfa9fb4e8`](https://github.com/siderolabs/talos/commit/bfa9fb4e8bc49e8fba6a4397868e7c390beb751c) fix: flaky tests
* [`a1ede48cb`](https://github.com/siderolabs/talos/commit/a1ede48cb950acb44f0b01242d8594277a17cb4f) test: fix testremovemember etcd integration flake
* [`ea90e690d`](https://github.com/siderolabs/talos/commit/ea90e690dbbd8dd76711d70649b51102dd568de4) feat: add MD RAID gRPC service and reconcile controllers
* [`74486ef6d`](https://github.com/siderolabs/talos/commit/74486ef6d53432fea5869f03edffc27754c990d0) chore: update deps
* [`f59c3ccad`](https://github.com/siderolabs/talos/commit/f59c3ccadd2036d2b3227d0605c736b369c95190) feat: implement service account configuration
* [`baff2d3f9`](https://github.com/siderolabs/talos/commit/baff2d3f919efba90854298681bde603780b6123) test: fix some test flakiness
* [`5450ec303`](https://github.com/siderolabs/talos/commit/5450ec3030b46dda6fd10bbe68fd93c659921685) fix: use a forked version of secure-io/siv-go
* [`33fac3f85`](https://github.com/siderolabs/talos/commit/33fac3f85dab9b1fe8c39c1fda674ad7cb776525) test: stabilize netapp trident csi fio runs
* [`afdde2a8f`](https://github.com/siderolabs/talos/commit/afdde2a8fe717868f8591ed351054ac9a870aa50) chore(ci): add netapp trident csi integration tests
* [`21eca156f`](https://github.com/siderolabs/talos/commit/21eca156f1c6c2fb00d5436f1261a0370f67e1c2) fix: print link status changes
* [`210f4e369`](https://github.com/siderolabs/talos/commit/210f4e369a1857785980b1cdf71cd1452d6945f3) fix: shutdown/reboot via usermode helpers
* [`d193f278d`](https://github.com/siderolabs/talos/commit/d193f278dc91ee6d381874b5209d2420eb0eff38) test: fix cilium test config patching
* [`e06898069`](https://github.com/siderolabs/talos/commit/e068980690de64f91a2d5f77ab21b193f3621f81) fix: flaky tests
* [`b7398ec00`](https://github.com/siderolabs/talos/commit/b7398ec004d936eb269bbc9761c2902ded95de75) feat: move kernel module config into multi-doc
* [`55bc643af`](https://github.com/siderolabs/talos/commit/55bc643af53f70231926bcd8b0bf378eaf5abc1f) fix: flaky serviceaccount suite test
* [`dced7d570`](https://github.com/siderolabs/talos/commit/dced7d570d4127838596ab208301b7e910ad1516) fix: correctly treat guaranteed QoS pods in the OOM handler
* [`f783f6636`](https://github.com/siderolabs/talos/commit/f783f6636b33750af25e37bd5d6c79fc2698acc9) feat: implement controlplane only config validation
* [`d0291bb0b`](https://github.com/siderolabs/talos/commit/d0291bb0b3c2bad6f9001316ad84c67f79537768) feat: extract Kubernetes CA config into a separate document
* [`97ed958a8`](https://github.com/siderolabs/talos/commit/97ed958a8385fd02be1d81bf9a39cff8be0ea4b8) chore: use lefthook globs to skip noop jobs
* [`a145c6356`](https://github.com/siderolabs/talos/commit/a145c6356f4783c3e890d4f66b9a0ab3b2b06f76) chore: lefthook USERNAME env, post-commit hook
* [`f836707ad`](https://github.com/siderolabs/talos/commit/f836707ada73515976c4fbb1e750f28fb3c632f4) fix: use UnattendedInstallConfig for extensions
* [`67293c809`](https://github.com/siderolabs/talos/commit/67293c809803087c31be78fad73dea32deec3dae) chore: add lefthook.yml
* [`726ea8fc2`](https://github.com/siderolabs/talos/commit/726ea8fc21f900aecda3dfff75045dd979765d2e) chore: switch v1alpha1 validation to use cluster config struct
* [`d1d848022`](https://github.com/siderolabs/talos/commit/d1d84802297684050f7e2ee5cd44b37fc0916e50) feat: add mdadm tooling and udev rules
* [`020de3f51`](https://github.com/siderolabs/talos/commit/020de3f514d59960906e0e7747e1df4501843355) chore: update go dependencies
* [`ae84f56a0`](https://github.com/siderolabs/talos/commit/ae84f56a0fba16dc69f3a33081fcd17449153a3c) chore: remove orphaned unattendedinstall.md
* [`416073748`](https://github.com/siderolabs/talos/commit/416073748bb55cbd69d304128eefabf146836fa3) feat: add UnattendedInstall config and controller
* [`4e5b4c6a7`](https://github.com/siderolabs/talos/commit/4e5b4c6a79c89769ab1fbc526361c0c4b8690e3b) feat: extract clusterid and clustersecret to discoveryidentityconfig
* [`0a641f268`](https://github.com/siderolabs/talos/commit/0a641f2683f7e638901a16d02b4ae6132c26ea20) refactor: simplify device status controller
* [`99da7f27f`](https://github.com/siderolabs/talos/commit/99da7f27fb3d75e908e26db7686553c02084764c) fix: data race in manifest sync
* [`54ac1cbd6`](https://github.com/siderolabs/talos/commit/54ac1cbd63e45d4c5e0452b5de3aa063587c352f) fix: provide cooldown period for the QoS trigger
* [`788562586`](https://github.com/siderolabs/talos/commit/788562586c8e1b11bfdd3d2314be50de02218ee4) feat: udevd controller and udev rules config document
* [`6e34da25c`](https://github.com/siderolabs/talos/commit/6e34da25c03ea132ba276dc79bf331a5cdaeead4) feat: delegate drain ops to go-kubernetes/nodedrain
* [`e9e027c63`](https://github.com/siderolabs/talos/commit/e9e027c6317aef55b8eb3463993ae6e856dcfc1c) fix: kubelet stuck restarting
* [`6f481b420`](https://github.com/siderolabs/talos/commit/6f481b420c9494918ecd1885ac218d0b31ac9866) fix: decode extraArgs list values correctly
* [`c8bdcc252`](https://github.com/siderolabs/talos/commit/c8bdcc252bb8feb241e242c89545340770b06407) feat: update runc to 1.5.0
* [`eae11ab0c`](https://github.com/siderolabs/talos/commit/eae11ab0cf22368a0886aba866385297285a466e) feat: allow user managed etc files
* [`47d4bd87e`](https://github.com/siderolabs/talos/commit/47d4bd87e687bdf4288873962c5db14cedd5bb50) feat: set user-agent for Kubernetes client
* [`ba926c6ce`](https://github.com/siderolabs/talos/commit/ba926c6ceb5024cddac8600d84b57cf52823f6a7) chore: update golangcilint config
* [`45497bd5b`](https://github.com/siderolabs/talos/commit/45497bd5b3155cbd37a6a083e76002f64ea1d0e4) feat: bring systemd 261.1
* [`8d9ecec93`](https://github.com/siderolabs/talos/commit/8d9ecec931f0db57ee2b50cef94990fe8119a326) refactor: improve stability for process_test.go
* [`31221e7ee`](https://github.com/siderolabs/talos/commit/31221e7ee978d64ec76edb4c48edd450b464bb5c) refactor: talosctl running tasks are yellow
* [`b268a6b08`](https://github.com/siderolabs/talos/commit/b268a6b08d29afe40a2f68356ebe719ec39bd541) feat: refactor CoreDNS config into multi-doc
* [`416d5fe4b`](https://github.com/siderolabs/talos/commit/416d5fe4b059d0cfd9e9eead9e179f16dbba0a33) fix: race in etcd member add
* [`c244e4c46`](https://github.com/siderolabs/talos/commit/c244e4c4655423a36bc02657572b6208ff1fc901) fix: building integration test binary on darwin
* [`b15a64b31`](https://github.com/siderolabs/talos/commit/b15a64b317ece19c74ff1d063f4f31ae0023716d) chore: bump rekor for GHSA-47q9-m4ww-924m
* [`cd8b0fe39`](https://github.com/siderolabs/talos/commit/cd8b0fe394351efa7f965174bc123fc57b7c5997) release(v1.14.0-alpha.2): prepare release
* [`917820cb3`](https://github.com/siderolabs/talos/commit/917820cb3ebe74e58bde1f78b74e1ae55ba111bf) chore: sync pkgs/tools
* [`b34be14e9`](https://github.com/siderolabs/talos/commit/b34be14e90f81f3bccb539d589185e757a02e55a) fix: cli.md codeblock generation
* [`25abcc6b5`](https://github.com/siderolabs/talos/commit/25abcc6b5973e9485af996c0d866ff28ca126b65) docs: update kubespanconfig to match discoveryserviceconfig
* [`742589f50`](https://github.com/siderolabs/talos/commit/742589f50ef7ea0ef9b0fc69d0ee80dc24369933) feat: support multiple discovery service configs
* [`fc3f27d79`](https://github.com/siderolabs/talos/commit/fc3f27d796cc33ea3dce9e0ed202fbb007b2922b) chore: enrich the SBOM with Go module licenses
* [`47d5c3351`](https://github.com/siderolabs/talos/commit/47d5c33514b05454ec67c488e3b2f0997596dac8) fix: handle image cache being disabled
* [`1a965aec3`](https://github.com/siderolabs/talos/commit/1a965aec37ea9e9551237255e11a0a33667f1876) test: disable LongHorn ublk test and add more cores
* [`6d03b3f61`](https://github.com/siderolabs/talos/commit/6d03b3f611de4ca06a6ee0587b91f8957e1e867f) fix: align documented image cache partition label
* [`6447d854f`](https://github.com/siderolabs/talos/commit/6447d854f20de6ee712f541f84cdbac9c2b33b85) fix(talosctl): use aio threads on darwin
* [`f856d1808`](https://github.com/siderolabs/talos/commit/f856d18084a6bfb119aaccc6af36b2b80e2f5b7c) fix: image verification with referrers
* [`11a7fbe4c`](https://github.com/siderolabs/talos/commit/11a7fbe4c6cba823f69205ed9ccfc36b7ab027bb) feat: extract kube-apiserver config into multi-doc configs
* [`337654d2b`](https://github.com/siderolabs/talos/commit/337654d2b8f0468507b40f53d03beb9a799a8c43) test: fix rook-ceph tests
* [`e33a86825`](https://github.com/siderolabs/talos/commit/e33a868254098659ce069d73570fb69ff4f148c6) feat: add AMD XGBE driver to initramfs
* [`bd2d6242a`](https://github.com/siderolabs/talos/commit/bd2d6242a3b0b5c51e6a56a0accc9137bcb7372e) fix: revert coredns to 1.14.2
* [`7c4e644f8`](https://github.com/siderolabs/talos/commit/7c4e644f810678e91a00656c124e886e38b0f12c) feat: update Linux to 6.18.36
* [`6e23a5c2f`](https://github.com/siderolabs/talos/commit/6e23a5c2f6d915bf9ddb8e23c0738119cccc8001) chore: refactor bare opentree_clone into a mount helper
* [`dfbd30959`](https://github.com/siderolabs/talos/commit/dfbd30959bb51d92f15cf4d4d43942591b9829f5) fix(talosctl): prevent appending type 11 smbios values on restart
* [`5926dd70d`](https://github.com/siderolabs/talos/commit/5926dd70d3df61ced4be38db604248d2dd4d53d6) test: support running integration test against remote provisioner
* [`f146c6a18`](https://github.com/siderolabs/talos/commit/f146c6a18334ea35cc5640f47e9756aaadadb153) feat: refactor /etc mounts
* [`ebe364117`](https://github.com/siderolabs/talos/commit/ebe364117cb32a227d9f995c59e75aea20a9e75a) chore: bump containerd to 2.3.2
* [`bc30c61a1`](https://github.com/siderolabs/talos/commit/bc30c61a14a785ee979d163c28a531247231845c) chore: bump deps (go, k8s, docker)
* [`00d739d0a`](https://github.com/siderolabs/talos/commit/00d739d0a92a334b3b619d161e68adedf09bb7d1) test: skip fstrim default schedule on cloud tests
* [`d9c6edf01`](https://github.com/siderolabs/talos/commit/d9c6edf01e2072971610f5eecf5601092649289f) fix: bump number of open files for etcd
* [`990c5395c`](https://github.com/siderolabs/talos/commit/990c5395c6fb6272ff4b5361f09d354f047c534a) chore: update tools and pkgs 2026-06-17
* [`325be7cd8`](https://github.com/siderolabs/talos/commit/325be7cd821128cc3e0b32ecf6244ee1b41ec20a) refactor: config generate uses multi-doc sysctlconfig
* [`d6930633b`](https://github.com/siderolabs/talos/commit/d6930633bfe02ec11cc47544f9d1ac9632d96b8c) fix: clean up and overhaul mount ops
* [`a0219404d`](https://github.com/siderolabs/talos/commit/a0219404d0f95af50c35804ee0375e66cdf44979) fix: cgroups cleanup
* [`58d8b71c4`](https://github.com/siderolabs/talos/commit/58d8b71c420f97aeb259aa04e2c776fcbbeaf261) fix: stop the log persistence and close all files on shutdown
* [`4b32ebc17`](https://github.com/siderolabs/talos/commit/4b32ebc17e4e37f8bd0547a496bc7a86f7b60bad) refactor: simplify trustd/apid rootfs setup
* [`dc98e3553`](https://github.com/siderolabs/talos/commit/dc98e3553d513aadb1588a0df605fb28de4e3f8e) feat: implement filesystem trim support
* [`897bef633`](https://github.com/siderolabs/talos/commit/897bef633ec9a1e0e9e292ab68a9dbf5704f9ac9) feat: introduce KubeProxyConfig multi-doc
* [`ebde543cf`](https://github.com/siderolabs/talos/commit/ebde543cf32a46c6a5f520aea52f7bdea8deb1d1) feat: introduce BootID resource
* [`cd178b9f3`](https://github.com/siderolabs/talos/commit/cd178b9f34bdcdf6d2a0a6d59de1c97e8523c9a5) fix: ensure consistent manifest apply order
* [`19fac6151`](https://github.com/siderolabs/talos/commit/19fac61511744a2b50342c512c06a6917787ebfc) feat: remote provisioner
* [`b6412e031`](https://github.com/siderolabs/talos/commit/b6412e031b6d8ac7d88addf79d1967fceb35c45b) fix: drop one more reference to removed 'nodes'
* [`be7f7a7db`](https://github.com/siderolabs/talos/commit/be7f7a7db50caf095fecebd62e881f6b96a4a9de) feat: add human-readable size fields to LVM resources
* [`d4e0ca1ba`](https://github.com/siderolabs/talos/commit/d4e0ca1ba94c48b7d6bd9de686c78a7fed7c6d39) fix: make LVM reconciliation robust and idempotent
* [`0dbc1e529`](https://github.com/siderolabs/talos/commit/0dbc1e529536f90727663184d8870a7e37b283e7) chore(ci): fix flaky test
* [`b687a47ab`](https://github.com/siderolabs/talos/commit/b687a47ab9ab58e43b4e045c0851b19e04f92680) feat: implement an option to allow discards on encrypted volumes
* [`3fc981c57`](https://github.com/siderolabs/talos/commit/3fc981c570205e3a71df32ad1a607c9730beaf1c) fix: improve security of scheduler/controller-manager
* [`5d4af9f33`](https://github.com/siderolabs/talos/commit/5d4af9f337a241a012dc324884ce6b9acdf461da) fix: gracefully stop node containers before removal
* [`c1593d8a3`](https://github.com/siderolabs/talos/commit/c1593d8a303ea95ab9733d5dc75779b3a18c7959) fix: honor FailurePauseTimeout when pausing before reboot
* [`506dc1323`](https://github.com/siderolabs/talos/commit/506dc132344d2d966ac2f6e31ea4b9e55c4d9107) feat: add imager flag to set the SecureBoot key enrollment mode
* [`5d4ba702e`](https://github.com/siderolabs/talos/commit/5d4ba702e8acd2ad8ec4948de899047d2a193d5b) refactor: generate pod definitions in k8stemplates
* [`995bc30d5`](https://github.com/siderolabs/talos/commit/995bc30d5df5beadd28e37a04959f315942d2182) feat: drop apply config method reboot
* [`18f6cb4d0`](https://github.com/siderolabs/talos/commit/18f6cb4d008b6a39493611e819d1915fde369b09) fix: increment time epoch on wall-clock jump when time sync is disabled
* [`755a8c8eb`](https://github.com/siderolabs/talos/commit/755a8c8eb5b8bd1eb935339288af1c4d5e621077) feat: update etcd to 3.7.0-rc.0
* [`a0c76fad1`](https://github.com/siderolabs/talos/commit/a0c76fad133234243956c3411ee386cea2416d07) feat(talosctl): implement cluster logs
* [`db052165c`](https://github.com/siderolabs/talos/commit/db052165c4cee767036668c490108d11e0eb1558) feat(talosctl): support rebooting cluster nodes
* [`0a04f463a`](https://github.com/siderolabs/talos/commit/0a04f463a1d3f19729e2ba0b1884017411df3783) feat(talosctl): use gateway dns for cluster
* [`cf3eb1cad`](https://github.com/siderolabs/talos/commit/cf3eb1cad1eeae4b90e010bca5bfa87c6fe334ed) chore(talosctl): disable kexec for cluster create on arm64
* [`180182b0f`](https://github.com/siderolabs/talos/commit/180182b0f50bcdb8b50d36112c308693b146bc70) fix: correct the link alias condition
* [`ac9014f05`](https://github.com/siderolabs/talos/commit/ac9014f051716adda826ee5a319ba83415fa87b3) fix: introduce pull attempt stall detection for image pull
* [`f2286d616`](https://github.com/siderolabs/talos/commit/f2286d616e86ce19c59a29a32e8334914cee9d3a) fix: move Flannel netpol patch to the controlplane
* [`9986c0b16`](https://github.com/siderolabs/talos/commit/9986c0b16a984974bdcd421d95bd533cc3c63fad) feat: bump kernel to 6.18.35
* [`e8845fba6`](https://github.com/siderolabs/talos/commit/e8845fba6b7e4a6342def670387ad498eee97810) fix: route ProxyURL test via reachable endpoint
* [`44acedf30`](https://github.com/siderolabs/talos/commit/44acedf3093a3e19b92f561a699ecc65f5cd5c96) feat: add declarative LVM logical volume provisioning
* [`f6058a11b`](https://github.com/siderolabs/talos/commit/f6058a11bc1b5fb359d512cc1601ff7f43feec96) feat: grab support bundle via client factory
* [`cdd719773`](https://github.com/siderolabs/talos/commit/cdd719773f5bd225689c6fa3e13eecfd82261325) feat: add CPUCores resource
* [`8e41eb1bd`](https://github.com/siderolabs/talos/commit/8e41eb1bd6f6970c9f7dc41bead1f7cd2d554a39) feat: verify go.mod tidiness in generate target
* [`b19e2ea42`](https://github.com/siderolabs/talos/commit/b19e2ea42dffdcf9a32b44fbd6d5369ed6a123ae) feat: add kube-apiserver probes
* [`a321a1dcc`](https://github.com/siderolabs/talos/commit/a321a1dccc449c586a6787f2c09cc79c3b26d3b9) feat: support proxy-url in talosconfig context
* [`bb2ac7546`](https://github.com/siderolabs/talos/commit/bb2ac7546cd5e93e88c305b2a0c4237227f2250f) feat: parse schematic info out of extension status
* [`0c02a5a07`](https://github.com/siderolabs/talos/commit/0c02a5a074115acd0d1899af5c973a4a5b743436) fix: align flannel MTU with kubespan to avoid permanent fragmentation
* [`3d5fd822c`](https://github.com/siderolabs/talos/commit/3d5fd822c57e6e87d574c59a1f0198c56a3cfe3c) feat: expose disk firmware and BIOS version
* [`30115981c`](https://github.com/siderolabs/talos/commit/30115981c4ce815437feeff6c43a18f6c308cca1) fix: relax LUKS header validation
* [`5923199fb`](https://github.com/siderolabs/talos/commit/5923199fba6107113541700c5485117aef6976d7) refactor: use ClientFactory for the action tracker
* [`72c0ced3c`](https://github.com/siderolabs/talos/commit/72c0ced3ca15ac1a5ca98d4af7560c84191af191) refactor: deprecate sysfs and sysctl in machineconfig
* [`ee74a41fb`](https://github.com/siderolabs/talos/commit/ee74a41fbb6d31b1fb44b0293943a8d1082ef610) fix: handle cluster-scoped resources with a namespace correctly
* [`9df5a647a`](https://github.com/siderolabs/talos/commit/9df5a647af72be48dd431eb5d506dfa0dc70b2fc) feat: allow to disable access time for EPHEMERAL partition
* [`9b667dbde`](https://github.com/siderolabs/talos/commit/9b667dbdeeddac1318166c4f59fa60da427d1847) chore: fix lint error in test
* [`311378386`](https://github.com/siderolabs/talos/commit/31137838691f03aad157339734d10d8275e75c09) test: increase resource inmem buffer to stabilize the tests
* [`6f85ce3d2`](https://github.com/siderolabs/talos/commit/6f85ce3d2bb0cd8884e84da83d464d578298524f) docs: hack/release.toml explains kernelmodulestatus
* [`9bb0a5d01`](https://github.com/siderolabs/talos/commit/9bb0a5d01a02460689e92f8a27dac511fbb55104) fix(talosctl): add scrolling to dashboard footer node list
* [`4c029c2d6`](https://github.com/siderolabs/talos/commit/4c029c2d64c6523625a0533418e5e446d33b1ec6) fix: machine configuration schemas
* [`c3052e845`](https://github.com/siderolabs/talos/commit/c3052e845e0e1c001386dc89511a0d2ed99fafa7) feat: move CNI config out of v1alpha1 config
* [`1d2f1208c`](https://github.com/siderolabs/talos/commit/1d2f1208c02f879ee13e1674305cca7e823f088b) feat: add declarative LVM volume group provisioning
* [`85f1d428f`](https://github.com/siderolabs/talos/commit/85f1d428f19e362b3bb2b10d1dfec7b4d1685695) chore: refactor tests to use debug api
* [`c901d47a5`](https://github.com/siderolabs/talos/commit/c901d47a57b742abd6520308f0605b9c4be3cd9e) refactor: talosctl streaming commands and more fixes
* [`166854959`](https://github.com/siderolabs/talos/commit/1668549593350110884c5c0a40219ea654359edc) fix: mark more resources as sensitive
* [`58adf2e00`](https://github.com/siderolabs/talos/commit/58adf2e00de7c55dd098c318d487ab5fec52d278) fix: classify installer and imager exits
* [`9549930ff`](https://github.com/siderolabs/talos/commit/9549930ff7cadb58215721c160b9fa3993ae2f1f) feat: update Flannel to v0.28.5
* [`27362d18e`](https://github.com/siderolabs/talos/commit/27362d18ee6637484e0d69458257b6e76fd7470c) refactor: replace the callback strategy for most commands
* [`cb42d9d9a`](https://github.com/siderolabs/talos/commit/cb42d9d9a8856a3607b4968f2d8d171dc244a9e6) feat: implement support bundle encryption
* [`9ae260b55`](https://github.com/siderolabs/talos/commit/9ae260b555074a79b06984a7e63d723733b98407) feat: enable NRI by default
* [`d1d5847b0`](https://github.com/siderolabs/talos/commit/d1d5847b0eb59e5fbf48a575db291aab9a4ca8ad) fix: flaky test
* [`0f2331586`](https://github.com/siderolabs/talos/commit/0f23315866b45d7506f6084e8d970883720480a7) feat: support external secureboot and pcr signers
* [`b349d919d`](https://github.com/siderolabs/talos/commit/b349d919db6b7b5aa327f60b7fc7cf5ede64d578) feat: enforce strict QoS ordering in OOM victim selection
* [`76d9b49bd`](https://github.com/siderolabs/talos/commit/76d9b49bd16e57942dbd5a156839a6543da41778) fix(ci): aws nvidia tests
* [`3131826cd`](https://github.com/siderolabs/talos/commit/3131826cded77002b18fd30f59b081b40ec8b55f) fix: provide NTS sync with bad initial clock state
* [`89e307e58`](https://github.com/siderolabs/talos/commit/89e307e5847befce2894a295c9479d8101078b6c) fix: etcd client leak in the (legacy) Upgrade API
* [`476c4d050`](https://github.com/siderolabs/talos/commit/476c4d0500d3c2f7753cb1f583f369256c3148e2) fix: recreate dns server and listeners on host DNS runner restart
* [`9a283d9b1`](https://github.com/siderolabs/talos/commit/9a283d9b190f9a1042246bc4830e75dae5826ce8) feat: bump Go to 1.26.4
* [`4759dc246`](https://github.com/siderolabs/talos/commit/4759dc24699dc03690511f7648412d0c9e70877f) chore: bump dependencies
* [`26a25a073`](https://github.com/siderolabs/talos/commit/26a25a0736c3575b0acd1661d5f67bf3ee72c0f2) chore(ci): drop homebrew workflow
* [`fa8a55192`](https://github.com/siderolabs/talos/commit/fa8a551928783ceb9efbb771837004474f64fa68) feat: update etcd to v3.6.12
* [`41fcab476`](https://github.com/siderolabs/talos/commit/41fcab476b06ce1d150c14d8a41a05727b064a25) feat: update kernel to 6.18.34
* [`8ba00612b`](https://github.com/siderolabs/talos/commit/8ba00612bea4303d28cd36ef97715a840238d9d5) feat: update dependencies
* [`6e2dec1ea`](https://github.com/siderolabs/talos/commit/6e2dec1eaf5ed62bea7e1563295bb590754e6546) refactor: update talosctl commands to stop using WithNodes
* [`f9ad63a35`](https://github.com/siderolabs/talos/commit/f9ad63a35a40cfb0bd3bc15a0368b50334fa5f11) feat: add custom logging convention linter
* [`30dbce03f`](https://github.com/siderolabs/talos/commit/30dbce03f359cdf3718659eb335b4b23eb95b450) chore: make oci images reproducible
* [`38244fd5b`](https://github.com/siderolabs/talos/commit/38244fd5b5aadd3e474f78b6a1c2c371b6ada766) feat: add sbom builder
* [`5177c50e2`](https://github.com/siderolabs/talos/commit/5177c50e2ea586d09c328190a65802571210b987) refactor: deprecate loadedkernelmodule
* [`c2eef3645`](https://github.com/siderolabs/talos/commit/c2eef3645a91bd76c69379b61a34470c927c37cc) fix: health request server-side
* [`d6eff8eff`](https://github.com/siderolabs/talos/commit/d6eff8eff42f2c324c9e08ef9da3edf8db17a515) refactor: drop multi-nodes proxying for the dashboard
* [`2e547a964`](https://github.com/siderolabs/talos/commit/2e547a964b20f91bf06fce4a5a35b8c5ef56a6e3) refactor: deprecate multi-node proxying
* [`ddcc519e1`](https://github.com/siderolabs/talos/commit/ddcc519e12e487123e19594c6ed0c7a406c0539c) fix: add --fail to image-signer curl download
* [`e5b0b1dde`](https://github.com/siderolabs/talos/commit/e5b0b1dded16b01d56e44182163652d2a66d1cee) fix: normalize log fields
* [`d8e95c396`](https://github.com/siderolabs/talos/commit/d8e95c3965b8690fd4e3d575d77dedec550328fd) fix: drop installer from bundle
* [`7aad9ec81`](https://github.com/siderolabs/talos/commit/7aad9ec81bc1a2f8ae6e4917173637b465613b30) feat: update pkgs, tools, Go dependencies
* [`b50ee396f`](https://github.com/siderolabs/talos/commit/b50ee396f6c6acdd97b79d92f149a25efa38aa22) fix: fix trace fix to also lookup release branches
* [`027c93d25`](https://github.com/siderolabs/talos/commit/027c93d254debaa8d7f62936cf05ff3c4ed9872c) release(v1.14.0-alpha.1): prepare release
* [`4eb862d09`](https://github.com/siderolabs/talos/commit/4eb862d09082e5bf187ccc3df10f33e02644cf48) feat: add LVMService for VG/LV/PV removal
* [`b88f16a52`](https://github.com/siderolabs/talos/commit/b88f16a529ff69a1b77ac3f6c14d90922399912e) fix: use POSIX shell idioms for error propagation
* [`5290eb374`](https://github.com/siderolabs/talos/commit/5290eb3742f1258669b5c608e3e43dcddd85c7c5) fix: suppress ICMP redirects by default
* [`7b4aba2e5`](https://github.com/siderolabs/talos/commit/7b4aba2e578f66e5282a1e20fab3dcf4043ba316) fix: marshal kube-scheduler config correctly with int types
* [`894be9bf5`](https://github.com/siderolabs/talos/commit/894be9bf5f9f3e4f365d43956e92c155355815ef) fix: touch rootfs files with SOURCE_DATE_EPOCH
* [`cde82224e`](https://github.com/siderolabs/talos/commit/cde82224e06af4fcc763a182a263fb68ba801536) fix: ignore cgroups with zero rank in OOM handler
* [`bc0372411`](https://github.com/siderolabs/talos/commit/bc03724111dff9e8c22edee4cf4e3ba1b7640317) fix: bring in a change to BCM2712_MIP
* [`f572c33f1`](https://github.com/siderolabs/talos/commit/f572c33f1c8b56f8b0b55fc74477689b8e354794) chore: fail on makefile error
* [`e317d4b47`](https://github.com/siderolabs/talos/commit/e317d4b47239695f87a67e9b9484bb9aa4812fe7) fix: drop modprobe path and enforce usermode helper
* [`89e53f610`](https://github.com/siderolabs/talos/commit/89e53f6102ea74f37b144be80ac6be185e5cff62) fix(machined): make built-in mod state always 'permanent'
* [`cfbec9bd5`](https://github.com/siderolabs/talos/commit/cfbec9bd584da34d33201c1404e7291166b3fc2b) test: skip UEFI vars wipe if TPM is enabled
* [`1e31deda3`](https://github.com/siderolabs/talos/commit/1e31deda39dfc38118b176bed276788103268895) fix: create parent directories when extracting tar archives
* [`14dc188bd`](https://github.com/siderolabs/talos/commit/14dc188bd68a54761bb44c58af87d1ba48707037) chore: verify go-containerregistry preserves symlinks
* [`951922dfb`](https://github.com/siderolabs/talos/commit/951922dfbc8fb72e87f26d09c95d6bf958efef07) fix: guard apply config API call
* [`3e173adf4`](https://github.com/siderolabs/talos/commit/3e173adf418950166a720501b32a37f8a3c92ac2) feat: move kube-controller-manager config to multi-doc
* [`b5cda3438`](https://github.com/siderolabs/talos/commit/b5cda3438c8ab2d9b2506087287a1fb2df87f8fb) fix: reset QEMU UEFI variable store when disk is wiped
* [`4a17ac6ac`](https://github.com/siderolabs/talos/commit/4a17ac6ac09373a92cb8690c6fe2364c4852fe1c) chore: script for tracking fixes made in upstream toolchain/tools/pkgs
* [`d71edeead`](https://github.com/siderolabs/talos/commit/d71edeead9d9c38018e69524cf5e2d764109a57e) feat: add LVM status resource definitions
* [`4aeba1cde`](https://github.com/siderolabs/talos/commit/4aeba1cde061e45ef523033ecae575f2475b0fce) fix: perform backwards-compatible kernel args cleanup
* [`9b7b2bf36`](https://github.com/siderolabs/talos/commit/9b7b2bf36a9790da34a7b8121c69d4df698c78cb) feat: implement support for btrfs user volumes
* [`03ee8ee3a`](https://github.com/siderolabs/talos/commit/03ee8ee3a3d7c40a0c060770ca8e20f9de3dd2a4) feat(machined): support instance tags on Akamai
* [`d19f9ade0`](https://github.com/siderolabs/talos/commit/d19f9ade01226d6092815a4545987a1e147f8ae4) fix: memorymodules resource reporting
* [`a6edcf6f3`](https://github.com/siderolabs/talos/commit/a6edcf6f33a44c8ce443805fd97477f3b7e41344) chore: move out adv library
* [`40e66eac7`](https://github.com/siderolabs/talos/commit/40e66eac7d6ef925f937de024b8d4fa9efdda231) fix: bump Go golang.org/x modules
* [`e23ca4a0a`](https://github.com/siderolabs/talos/commit/e23ca4a0abfa2aaf97e1ad3e89488e6f399a45e1) chore(ci): add upgrade tests for trustedboot
* [`e3003c0ec`](https://github.com/siderolabs/talos/commit/e3003c0ec7ec4961ee16cc053afc236314cecbad) chore: bump tpm nonce size to match the algorithm used
* [`8fd04da1f`](https://github.com/siderolabs/talos/commit/8fd04da1f75a61ef04ced0131a2effc1ae747211) feat: add bnxt_re module to the rootfs
* [`1cfab00f1`](https://github.com/siderolabs/talos/commit/1cfab00f14ad29b957f61b266f0d82c8d8d72114) fix: update etcd experimental args
* [`ad96fc6ae`](https://github.com/siderolabs/talos/commit/ad96fc6ae7a1be3aaff9ce89d7ee4e7ae8329085) fix: relax hostname config validation
* [`efd735334`](https://github.com/siderolabs/talos/commit/efd73533490e678a40fa7a9b89e0510e2b829653) chore(ci): add missing labels, move release metadata check to job
* [`9ec045059`](https://github.com/siderolabs/talos/commit/9ec04505954cde51be41d6fff8f0ceace0d84a39) feat: update containerd to 2.3.1
* [`42f4144a1`](https://github.com/siderolabs/talos/commit/42f4144a175c5bc1bccdbbcf82696e15af378d55) feat: introduce new KubeSchedulerConfig
* [`f2b7f39db`](https://github.com/siderolabs/talos/commit/f2b7f39dbfa816599d330b256de4dc5c20a97853) refactor: move Args type out of config/v1alpha1
* [`b959dcb3e`](https://github.com/siderolabs/talos/commit/b959dcb3eb1b7243e7fb4618e3d9bef19657dd0e) fix: bump Kubernetes to 1.36.1 in one more place
* [`8ecc77f1a`](https://github.com/siderolabs/talos/commit/8ecc77f1a8ad9fdd619f7c427277e876146f3b80) feat: update default Kubernetes version to 1.36.1
* [`cbd9c3745`](https://github.com/siderolabs/talos/commit/cbd9c374591f1967c98bb5a2194832bcb3f98ccb) chore: rekres to secure slack workflows
* [`6a92fc653`](https://github.com/siderolabs/talos/commit/6a92fc6535f6dfb3df00d00556c3108817b27368) test: update Canal version used in the tests
* [`be12d3d08`](https://github.com/siderolabs/talos/commit/be12d3d081d7f10aff8b305b2f7685b4aad8451e) feat: support 4k sector size disk images
* [`a7e8f4c28`](https://github.com/siderolabs/talos/commit/a7e8f4c282d461840d04d6a045465aab7da3765c) chore(ci): fix cloud image upload job name
* [`4319399f6`](https://github.com/siderolabs/talos/commit/4319399f62c144aa9b66d9defe803d9d81771dcd) feat: introduce more modular Linux kernel
* [`ed5df89f6`](https://github.com/siderolabs/talos/commit/ed5df89f6ec161b80e0a9a522db3fce6d5b94a04) feat(ci): rotate credentials
* [`a6a984ff7`](https://github.com/siderolabs/talos/commit/a6a984ff733dbd5b91568aed62a5c67348b6e047) chore(ci): fix the job conditions
* [`ecb7d4588`](https://github.com/siderolabs/talos/commit/ecb7d45883bc1de46c4adbe37dcf556c59e78a25) feat: enable Flannel nftables mode
* [`9919ff781`](https://github.com/siderolabs/talos/commit/9919ff781326a01dd1776d4cfaa2834f701c3e05) feat: update Linux to 6.18.32
* [`1a7d136e4`](https://github.com/siderolabs/talos/commit/1a7d136e41a327957c7e6e75010ff7d191acbca6) feat: add Azure Secure Boot imager profile
* [`df68e7391`](https://github.com/siderolabs/talos/commit/df68e73912da0d57ecebe340d1bf70ecaa3e47c9) feat: implement kernel module status resource
* [`e98ee99d4`](https://github.com/siderolabs/talos/commit/e98ee99d42491bca8ac15eafd7b03fc2bb2c7a6b) fix: streamline config validation flow
* [`d7f0a2fd4`](https://github.com/siderolabs/talos/commit/d7f0a2fd49efae514f4e7315201cd57be54ac8ae) feat: update Linux to 6.18.31
* [`2b66e25a5`](https://github.com/siderolabs/talos/commit/2b66e25a50de71bcd3bb5c4d8b059de0e08ffb9b) chore: update image signer
* [`5aa1795f9`](https://github.com/siderolabs/talos/commit/5aa1795f973daed1c87382d6a08e048345b504f4) chore: drop e2e step dependencies
* [`d42b3b396`](https://github.com/siderolabs/talos/commit/d42b3b396fb14036720cda44f9b2044e98c62f06) feat: update Linux to 6.18.30
* [`c3f6f3507`](https://github.com/siderolabs/talos/commit/c3f6f3507816c8dd81daa8baa2e3900ec78140c9) feat: implement static host resolving via host DNS
* [`2f06a68ef`](https://github.com/siderolabs/talos/commit/2f06a68efa525b917c9626be2e7adeb5186aff60) refactor: split host DNS handler
* [`e99c5be5a`](https://github.com/siderolabs/talos/commit/e99c5be5adf24c56a5fb482356d679e914a467f7) feat: implement DNS over HTTP(S)
* [`cf6065238`](https://github.com/siderolabs/talos/commit/cf606523846a63066d193be7e6669c0fb2d3ce23) chore: stop publishing installer to ghcr
* [`0edabd29c`](https://github.com/siderolabs/talos/commit/0edabd29c42cf01025f5c0e49e32440bb7a768ef) fix: restore some shared (and some lower tier slave) mount propagation
* [`f1578dc63`](https://github.com/siderolabs/talos/commit/f1578dc634f23d797e7d0c491bc516eada9f8171) fix: image verification issue with registry.k8s.io
* [`46b1f8a24`](https://github.com/siderolabs/talos/commit/46b1f8a24f92b449084b6897d2f299c762e9e9dc) fix: rework how scheduler config is marshaled
* [`820a9fa59`](https://github.com/siderolabs/talos/commit/820a9fa59f7a8096fce77f7f7ac1c9dc0215331f) chore: fix typos in comments
* [`649a384a9`](https://github.com/siderolabs/talos/commit/649a384a961c64c036dcc0b82ecbc102b951d2e2) feat: move more kernel stuff to modules
* [`4f3ab2012`](https://github.com/siderolabs/talos/commit/4f3ab2012bdefd9191c2f1772cdc8afd607381db) chore(ci): try fixing homebrew action
* [`600c0ab5d`](https://github.com/siderolabs/talos/commit/600c0ab5d8fefeb197c35c2ac1ad6f01e7e7ef67) feat(ci): validate that extensions PKGS and TOOLS sync with talos
* [`76080416b`](https://github.com/siderolabs/talos/commit/76080416beee6460375c0d0db3cab76cc493ad3e) feat: redact more machine config secrets and audit redactors
* [`aabf63957`](https://github.com/siderolabs/talos/commit/aabf63957d70b2a3a53def32b4bcf6ef35885ea8) docs: drop controlplane endpoint examples
* [`b48a2bef4`](https://github.com/siderolabs/talos/commit/b48a2bef4d4df861ffd035a1cd80f2a40e0a4486) test: relax kernel-default routing rule assertion
* [`d2208b034`](https://github.com/siderolabs/talos/commit/d2208b0348ca9a9e0054a61807fdb0551d6ccd56) refactor(talosctl): propagate command context throughout, handle interrupts
* [`0760b5c28`](https://github.com/siderolabs/talos/commit/0760b5c28b998b0d939f6ff3de1158dbb55cbbcb) fix: normalize source name for syft consistency
* [`c49ac0ec2`](https://github.com/siderolabs/talos/commit/c49ac0ec2f2774b601c79978787d90a2effa0adb) docs: document release policy
* [`ec7e6ef9f`](https://github.com/siderolabs/talos/commit/ec7e6ef9f461b3fde20a8534e10123d515558658) feat: bump in-toto indirect dependency
* [`21858a674`](https://github.com/siderolabs/talos/commit/21858a6745adfa356c81d6e78d92c20396c12d0b) feat: update kernel to 6.18.29
* [`5a49dc61d`](https://github.com/siderolabs/talos/commit/5a49dc61dbebbe91f3bb00bb043351ffabf3c6c2) feat: migrate Image Cache config to multi-doc
* [`574298ec1`](https://github.com/siderolabs/talos/commit/574298ec11a9aaeb3375733cc539c308bbdc4984) fix: handle empty GCP operation errors
* [`366b10b79`](https://github.com/siderolabs/talos/commit/366b10b7990bf01a4bbec7bf543f198842fa4df6) feat: dockerfile improvements
* [`9a1d9d0af`](https://github.com/siderolabs/talos/commit/9a1d9d0afd8b726f952f7d96af229381bd407b6b) feat: bump go 1.26.3
* [`6eec1c229`](https://github.com/siderolabs/talos/commit/6eec1c2293d7ca86c73d2fcac05e9e73befadc26) feat: support DNS over TLS for upstream resolvers
* [`dee139aef`](https://github.com/siderolabs/talos/commit/dee139aef05edfb1b4f23c24336102c1ee2a76e8) feat: revert update CoreDNS to 1.14.3
* [`087bc4c18`](https://github.com/siderolabs/talos/commit/087bc4c188e2096e1536202ffff07771235dacce) chore: lint packages under tools
* [`9e7516fae`](https://github.com/siderolabs/talos/commit/9e7516faee8c08195ffb85fa0874dc11992c9d82) fix: clarify documentation for image verification pattern
* [`41c8e9dc4`](https://github.com/siderolabs/talos/commit/41c8e9dc496bde6ee67c39368858844dc35c5bb7) feat: bump dependencies
* [`2b6c06ef5`](https://github.com/siderolabs/talos/commit/2b6c06ef51fdd3c7ffcbf621eacbb37d798e5228) feat: update CoreDNS to 1.14.3
* [`6b6f7978b`](https://github.com/siderolabs/talos/commit/6b6f7978be5b555a7124903f26bcb5b669b02e90) feat: update containerd to 2.3.0
* [`f9c4f90da`](https://github.com/siderolabs/talos/commit/f9c4f90da7cde25f84fff99b94844cb82fc23e95) feat(ci): longhorn v2 ublk tests
* [`84d169c62`](https://github.com/siderolabs/talos/commit/84d169c62a69c9a35f9dc7a4dd0851e503b7d539) fix: make dnsd retry listening
* [`689974bd5`](https://github.com/siderolabs/talos/commit/689974bd55ca28ccdd519f11dee0c3b8bc727601) fix: volume mount permissions
* [`ff0f66bdf`](https://github.com/siderolabs/talos/commit/ff0f66bdfa57a583398808b3dda23da4f900098b) fix: skip reserved routing rule priorities
* [`850e2c754`](https://github.com/siderolabs/talos/commit/850e2c754ce3c81549a8e390031a8c3ef6da61d7) feat: drop fakeroot, use go helper
* [`0c1bd701a`](https://github.com/siderolabs/talos/commit/0c1bd701a028e7e858224e0fa91492dda699f985) feat: add golangci-lint fmt target
* [`53bd66956`](https://github.com/siderolabs/talos/commit/53bd669562401409a03365c591586179f5b4c442) feat: support conditional start of IPv6 dns servers
* [`b31d93e0d`](https://github.com/siderolabs/talos/commit/b31d93e0d0e61d93fe18a77d80658d1a0aba42c3) feat: auto-enroll SecureBoot keys for disk images
* [`849a68006`](https://github.com/siderolabs/talos/commit/849a68006399dd677e4fb5152ef13c87578ba652) test: update pkgs to test new extensions
* [`c30a6dfcb`](https://github.com/siderolabs/talos/commit/c30a6dfcb1780860bed3ebed0cf9ee1d3af0d8ab) fix: preserve DHCP DNS servers
* [`5b81b20d3`](https://github.com/siderolabs/talos/commit/5b81b20d3cf2751b52ec0621aa74b627d6e11c49) feat: apply DHCP search domains
* [`4e5ff8fa2`](https://github.com/siderolabs/talos/commit/4e5ff8fa2191af92d5e8f3833dd279a01ed78d30) fix(ci): zfs test
* [`14abe5140`](https://github.com/siderolabs/talos/commit/14abe514006e4a30eb59a92742c92ca685962826) fix: handle gateways which are not on-link routes in dhcp4
* [`e1f759af8`](https://github.com/siderolabs/talos/commit/e1f759af80c7a51f4649cb64700286dc822f2d7a) chore: fix lint issues automatically
* [`664c5f643`](https://github.com/siderolabs/talos/commit/664c5f6432168e54e0f7a4723502df3705e1ce0a) chore: update tools
* [`c64df2b61`](https://github.com/siderolabs/talos/commit/c64df2b61971304b21f8333a993098ea0c3d15b7) fix: add missing kernel modules in rootfs
* [`f73c24594`](https://github.com/siderolabs/talos/commit/f73c24594132364fb39c68b901b1f70b57e72674) feat: run depmod with verification on rootfs build
* [`1371596d7`](https://github.com/siderolabs/talos/commit/1371596d750b431776cdcb4fd68916d4e9b710bb) fix: provide proper AWS platform metadata
* [`4f11f021d`](https://github.com/siderolabs/talos/commit/4f11f021de71015d5a7e1c39e38d1da21456a448) feat: implement etcd encryption config (kube-apiserver)
* [`876f83643`](https://github.com/siderolabs/talos/commit/876f8364302f7ab5fa9638fd297cd3fbfb3863ab) feat: add support for HTTP Probes
* [`9b776d598`](https://github.com/siderolabs/talos/commit/9b776d59819c98539fcc7649de464e70af27f606) feat: update etcd to 3.6.11
* [`631a1bc5e`](https://github.com/siderolabs/talos/commit/631a1bc5e1230b14ddc018e7f5116b2d0ebe8821) fix: bring in hardened kernel
* [`a349dac03`](https://github.com/siderolabs/talos/commit/a349dac03688971cc2322d2e64647efcf949a191) fix: stale discovered volume children
* [`13ce01879`](https://github.com/siderolabs/talos/commit/13ce0187959297bfaa7ba80080c5e472fe1bd0c3) fix: re-enable kexec on arm64
* [`32539d4ac`](https://github.com/siderolabs/talos/commit/32539d4ac4b4a62803dd840f84183da9597653ad) fix: deadlock in the makefs ext4 with populated source
* [`0f3e1966a`](https://github.com/siderolabs/talos/commit/0f3e1966af51cf815cf05f7a8c9b9e0e175b25fe) fix: panic in Kubernetes manifest sync
* [`3bae01ac1`](https://github.com/siderolabs/talos/commit/3bae01ac11cd64265f0aaaa9e2e7f83e39bd7d73) fix: do not pick up a system disk from a loop device
* [`dedb7a96c`](https://github.com/siderolabs/talos/commit/dedb7a96c16b8d90dd4dfa7e3b5a622952da246d) fix(talosctl): protect k8sNames map writes with mutex
* [`cc2be213a`](https://github.com/siderolabs/talos/commit/cc2be213a81a63f8f01534263924f152f5f083d5) fix: drop explicit platform matcher
* [`1dffebaf2`](https://github.com/siderolabs/talos/commit/1dffebaf2abbb9628cc09b29ec0881738c3756cd) fix: mount throws EPERM on virtiofs with SELinux
* [`48a481c29`](https://github.com/siderolabs/talos/commit/48a481c29fdabb82727acc6de9a4b7cb58156982) fix: replace Canal manifest with a more recent one
* [`6a445406e`](https://github.com/siderolabs/talos/commit/6a445406e0df4e33044cc429fabd92b658d05743) fix: make lacp active nilable
* [`0d1d95c7d`](https://github.com/siderolabs/talos/commit/0d1d95c7dac0b60f65b213ecaa5ab7c662147687) fix: bump go-kmsg to fix the timestamp drift
* [`bd344fd53`](https://github.com/siderolabs/talos/commit/bd344fd53ff5b74c7a9379bec00b33e5eb62879b) fix: reset the ticker when the KubeSpan is disabled/enabled
* [`462015bcd`](https://github.com/siderolabs/talos/commit/462015bcd9c196e458d318f0c1de3202aceed467) release(v1.14.0-alpha.0): prepare release
* [`8a037a56e`](https://github.com/siderolabs/talos/commit/8a037a56ed501b99757ca29f718c6ad7dfa2f223) test: fix flaky tests
* [`08c81d838`](https://github.com/siderolabs/talos/commit/08c81d8380b80090183df51f3a8b02ed5339adb4) feat: bump kernel to 6.18.25
* [`fe40b6e58`](https://github.com/siderolabs/talos/commit/fe40b6e588c38628e5cd9298dcaa56d2f2590827) fix(ci): fetch empty pr labels
* [`837a9ed07`](https://github.com/siderolabs/talos/commit/837a9ed077156ad00a1d31e731cf396c466bf6f6) feat: move host DNS config into ResolverConfig
* [`96a8ecd1e`](https://github.com/siderolabs/talos/commit/96a8ecd1eed06f3d04fea853a8673699130dded8) feat: default to factory installer image
* [`f19eef78b`](https://github.com/siderolabs/talos/commit/f19eef78b9cc01c107f86a6eddf24da0d288d124) fix: revert add extraArgs from service-account-issuer
* [`6821225b6`](https://github.com/siderolabs/talos/commit/6821225b64ddd48e5cc0d16ab80204d539110f78) fix: revert use append instead of prepend in service-account-issuer
* [`b43c3a124`](https://github.com/siderolabs/talos/commit/b43c3a124f6c6d1523c1feaddc9c4a23454eeb56) feat: add quirk for talosctl factory downloads
* [`df0b9a8da`](https://github.com/siderolabs/talos/commit/df0b9a8da1423842d830261e5ddc5dc8f5a234c1) refactor: make all controller unit-test follow modern patterns
* [`c2948cef2`](https://github.com/siderolabs/talos/commit/c2948cef232f6a175312636369b444124cb995db) feat: support auth for Image Factory in cluster create
* [`560bcf0ca`](https://github.com/siderolabs/talos/commit/560bcf0cae764015520b1d1efbef2a0bb4fe88b7) feat: enforce TLS 1.3 minmum version for Kubernetes components
* [`3db14309e`](https://github.com/siderolabs/talos/commit/3db14309e058cacc2ab8664944fc18f80a3bb747) fix(talosctl): ensure uncordon runs after reboot/upgrade errors
* [`ecf2fa855`](https://github.com/siderolabs/talos/commit/ecf2fa855b8eb19731b228990a3acbe1430ccad4) feat: update Kubernetes to v1.36.0
* [`71557eadd`](https://github.com/siderolabs/talos/commit/71557eadda51ba62fcc10d4ed859c390a93c565d) fix(ci): skip misc jobs not on pull request
* [`026313b7c`](https://github.com/siderolabs/talos/commit/026313b7cc103a2dc7efdee1dfbad32c8050daf6) docs: rename security-insights.yml to lowercase for LFX detection
* [`dc4ffd490`](https://github.com/siderolabs/talos/commit/dc4ffd490d878621b929af1ba1aca1d32e2530de) fix(ci): fix jobs not interpolating matrix due to condition
* [`25e2f37e2`](https://github.com/siderolabs/talos/commit/25e2f37e2b1c3b6bdc5ee04ffa86e6fe34cf582a) chore: generate comments for fields in resource proto
* [`149592fa5`](https://github.com/siderolabs/talos/commit/149592fa59d20c5aa29e4c0af9a3760585f378ce) fix: watch kubelet's kubeconfig and time out for cache sync
* [`1f315e6e9`](https://github.com/siderolabs/talos/commit/1f315e6e903ec81e2989eb02404522a8b3c2dab7) feat: update Linux to 6.18.23
* [`0198eedc2`](https://github.com/siderolabs/talos/commit/0198eedc2b39477a62a2d6e6450934ff29bce8b3) feat: add NTS (Network Time Security) support for NTP time sync
* [`6830a8b97`](https://github.com/siderolabs/talos/commit/6830a8b97df4a08f27516869363e13a53121b2e4) fix(ci): matrix jobs cleanups
* [`71aeb347f`](https://github.com/siderolabs/talos/commit/71aeb347f90969cb6057651666bfda205269d917) test: fix OOM test flake
* [`9b9542cc5`](https://github.com/siderolabs/talos/commit/9b9542cc55ee6d08f3490d270c1b497c7b9d3049) test: fix a flake in the manifest sync test
* [`863d882b6`](https://github.com/siderolabs/talos/commit/863d882b6cbd50abcc4fc8717e5921c92a1f0f0b) test: add image verification for factory.talos.dev
* [`bba0b4aee`](https://github.com/siderolabs/talos/commit/bba0b4aeefd7ec0daf7cc048e48c66d8b614f576) chore(ci): nvidia update helm values
* [`3399ff4de`](https://github.com/siderolabs/talos/commit/3399ff4de05b4fafb8511d6399e919436f1178da) fix: propagate route table down to the resource
* [`c684ec60e`](https://github.com/siderolabs/talos/commit/c684ec60ea5035e84517dac05a16eabf04f06a33) chore: prepare for Talos 1.14 release
* [`ed9545d0d`](https://github.com/siderolabs/talos/commit/ed9545d0db55cdff8ad7f7755398913780a7540e) chore(ci): bump gpu operator version
* [`4de3e4393`](https://github.com/siderolabs/talos/commit/4de3e4393e6ee968a7ef315c1a0f9fe4d86f449c) fix(ci): cron triggered workflows
* [`212182e6f`](https://github.com/siderolabs/talos/commit/212182e6f655f61e8917059868fc381728e4a959) chore: bump container registry library
* [`c028db0b8`](https://github.com/siderolabs/talos/commit/c028db0b8d25e85a4b580e10252d964785320291) fix: do not flip machine stage to rebooting during shutdown
* [`6ce62d9e8`](https://github.com/siderolabs/talos/commit/6ce62d9e8eea41a37e90fec5551ac06d26ef8b28) fix(ci): workflow runs with `workflow_run`
* [`509cd9733`](https://github.com/siderolabs/talos/commit/509cd9733926a6994843fb58ccdf38e5cd63a382) fix: boot entry detection
* [`5e3f30188`](https://github.com/siderolabs/talos/commit/5e3f301887546bfc83b9819bbc3ae05fe92f3471) feat(ci): rework to schedule daily runs after a cron
* [`7fa4d3919`](https://github.com/siderolabs/talos/commit/7fa4d39197e1a9e54ba8a259c111f2cb8047ef9c) fix: zfs extensions test
* [`1ef8e630a`](https://github.com/siderolabs/talos/commit/1ef8e630ab77b3c849e7da6d1ff83e7c6795f070) test: allow more tests to run in FIPS strict mode
* [`bdcc9321b`](https://github.com/siderolabs/talos/commit/bdcc9321b637da77f1007a571193c2e03c984b8b) fix: reduce memory dashboard usage
* [`2d177af82`](https://github.com/siderolabs/talos/commit/2d177af82b96cefdc7aebb62d593d0ffcba1a418) chore: update Syft to v1.42.4+patches
* [`0d8362119`](https://github.com/siderolabs/talos/commit/0d8362119e4415182caa9349e0ddfb27ea290d90) fix: return failed precondition on upgrade when not installed
* [`be58eafab`](https://github.com/siderolabs/talos/commit/be58eafaba98bb7b1bcd20ac1ed8f8b03734c7e0) fix: wrong slot of encryption key was logged
* [`015081c76`](https://github.com/siderolabs/talos/commit/015081c768ec85c3fb3b74ea22dd0b981db7c96a) feat: update dependencies
* [`9fbb7c95d`](https://github.com/siderolabs/talos/commit/9fbb7c95df2b1dcd68fafa23865412bbd8300f4b) fix: audit trustd code for security
* [`986e97fc7`](https://github.com/siderolabs/talos/commit/986e97fc757824bc998d81933e60108250316e5e) feat: update Flannel to 0.28.4
* [`f3817d1d1`](https://github.com/siderolabs/talos/commit/f3817d1d1c90bb2f2c19c209af154dc1a93eb507) chore: update sign images to support image name suffix
* [`e776721f3`](https://github.com/siderolabs/talos/commit/e776721f33b1fedff1dff310298035b3d603e676) feat: update Kubernetes 1.36.0-rc.1
* [`f6e7346fa`](https://github.com/siderolabs/talos/commit/f6e7346fa725a703ac4281854150d7a3be12c8d1) fix: encode extra args fields in resources with new id
* [`3c7bb80ba`](https://github.com/siderolabs/talos/commit/3c7bb80bab0323d72a1727256ccf339d2c79804c) chore: bump tools
* [`3ba35c9b9`](https://github.com/siderolabs/talos/commit/3ba35c9b9fca9c54e596d5c6df61d515a4a39555) chore(ci): nvidia try UKI boot
* [`e3e8f01ca`](https://github.com/siderolabs/talos/commit/e3e8f01ca66ee74898ebba5dadf4f199775d278e) chore: bump tools
* [`181584a5f`](https://github.com/siderolabs/talos/commit/181584a5f1850f2bfb2a837c0d05bd9e30ee48b5) fix: handle boot failure
* [`c464c7e88`](https://github.com/siderolabs/talos/commit/c464c7e88a3f058cb2bbc36af1910d69d903cd07) fix: upgrade API in maintenance mode (legacy)
* [`b7512d912`](https://github.com/siderolabs/talos/commit/b7512d9125b623d2bb92e3a8b5839e85e1309a39) feat: update Kubernetes to 1.36.0-rc.0
* [`4ba11156f`](https://github.com/siderolabs/talos/commit/4ba11156fd164a0d94538508f5c028f249deed50) refactor: allow overriding out image name suffix
* [`c81aa125c`](https://github.com/siderolabs/talos/commit/c81aa125c85d3886c5b9bb4d7f77ec2def104f21) fix: panic in reading PCR values
* [`6a3ab87c5`](https://github.com/siderolabs/talos/commit/6a3ab87c54f83f70869a2e298e6ed7722cf4afad) feat(ci): add nvidia arm64 matrix
* [`21f459aab`](https://github.com/siderolabs/talos/commit/21f459aab5d8ac2841aa69a9237ca3faa06da7df) fix(talosctl): always use default GRPC dial options
* [`ca208e514`](https://github.com/siderolabs/talos/commit/ca208e51492c4584f9a4cea4d0762c2199f703e7) fix: validate hostDNS forwarding requires hostDNS to be enabled
* [`9fcb9e05b`](https://github.com/siderolabs/talos/commit/9fcb9e05b668ba2fbc7df776ab32e57b1c15e221) feat: bump go to 1.26.2
* [`0bfdf7f70`](https://github.com/siderolabs/talos/commit/0bfdf7f7035fefe804ec4b568709cd6a09195293) fix: create correct blackhole routes for IPv4
* [`52b920032`](https://github.com/siderolabs/talos/commit/52b920032e97e1b241c1e0bd89c6e41cbc1c9a47) feat: add client-side Kubernetes node drain to reboot and upgrade commands
* [`968ec1e0c`](https://github.com/siderolabs/talos/commit/968ec1e0ca26eb1f0de0836e0a55df09dea7dafe) refactor: propagate NAME properly, allow to set on build
* [`acc69c346`](https://github.com/siderolabs/talos/commit/acc69c346f8816324b632fd33a5d0cb3f4b73509) fix: set the minimum TLS version to 1.3
* [`0cfa6e302`](https://github.com/siderolabs/talos/commit/0cfa6e3024100e34692a0b10e9dacb762c16a626) chore: bump some tool dependencies
* [`4229bb9d2`](https://github.com/siderolabs/talos/commit/4229bb9d2ed263c309d0b0082f6e21d2f002c925) feat: add dis-vulncheck tool
* [`d697f5538`](https://github.com/siderolabs/talos/commit/d697f5538a7a624a1ac7bafdfebc67dd9418c434) fix: don't set xattrs while decompressing extensions
* [`34fb2cbe5`](https://github.com/siderolabs/talos/commit/34fb2cbe5148a9f60fd888551ba6eceb84b550cf) refactor: remove manual shell completion and replace with cobra completion
* [`79fa2e300`](https://github.com/siderolabs/talos/commit/79fa2e3001082cf21be92c52b3da4e844313184d) feat: allow more nvidia and nvme files from extensions
* [`414f78a29`](https://github.com/siderolabs/talos/commit/414f78a298fc1a196fe310b17b89d3aadc15e1b4) feat: allow glibc ld files in etc
* [`1bbba4301`](https://github.com/siderolabs/talos/commit/1bbba4301495e256f2686a6b0d44663d3fdad2c4) feat: update Flannel to v0.28.2
* [`55815e0fa`](https://github.com/siderolabs/talos/commit/55815e0fa545de42997b89beaa7bf15ef9aa36f3) fix: handle ISOs with zeroes in volume labels
* [`7b6ab0c1c`](https://github.com/siderolabs/talos/commit/7b6ab0c1c3cec7b6260e27dd5b6e72faa1975ab0) feat: add flag to force fallback to legacy upgrade
* [`5e24d5265`](https://github.com/siderolabs/talos/commit/5e24d5265bde9adee92c02e675140de87ee126bf) feat: add resource view to talosctl dashboard
* [`649ab7fe4`](https://github.com/siderolabs/talos/commit/649ab7fe4234de1a947071926603377e00910cb9) fix: add os:meta:writer role to the dashboard
* [`10cdfa909`](https://github.com/siderolabs/talos/commit/10cdfa9099a3e40ca8182ecb69d836c06ca621e3) fix: drop talosctl install
* [`087ced85f`](https://github.com/siderolabs/talos/commit/087ced85f5130656cbc647c2e4d838cab3ff1737) fix: unseal with "slow" TPM
* [`11ab0a8c5`](https://github.com/siderolabs/talos/commit/11ab0a8c5aec1537542bddb851a9f71e92888e3b) fix: drop unused type from ExternalVolume schema
* [`e2df0f6ce`](https://github.com/siderolabs/talos/commit/e2df0f6ce8c47b0dc3e93bf257afb8a1ae9243fb) fix: always grow disks
* [`919d8c365`](https://github.com/siderolabs/talos/commit/919d8c36552a46ed326c9cb01bb474cee21e8d0a) chore: drop debug shell
* [`783a35851`](https://github.com/siderolabs/talos/commit/783a35851ed1bac4ddd0f1fed583fc1b6477614d) fix: add metal-agent mode to runtime capabilities
* [`37b2221cc`](https://github.com/siderolabs/talos/commit/37b2221ccfff64f37461397712c8b08ea3736dc0) docs: add SECURITY-INSIGHTS.yml for OSPS Baseline QA-04.01
* [`bed2bd414`](https://github.com/siderolabs/talos/commit/bed2bd414ea57866b5b31cb09f562fc7161ca74a) feat: add graceful power off support to QEMU VM launcher
* [`3400059cc`](https://github.com/siderolabs/talos/commit/3400059ccf4811140a4326397d972f68693c708c) fix: incorrect route source for on-link routes
* [`b3dfbf743`](https://github.com/siderolabs/talos/commit/b3dfbf743e6c2fd44020911ee1e0eea3a7676579) feat: bump musl to 1.2.6
* [`4227921b3`](https://github.com/siderolabs/talos/commit/4227921b3979d3a8542946fed4ceb622747adb00) test: fix the PKI mismatch test flake
* [`f2bc2dcc6`](https://github.com/siderolabs/talos/commit/f2bc2dcc6e0391dbd4aa19e8366d657b2056790f) feat: update NVIDIA production drivers to 595.58.03
* [`aa5946dd3`](https://github.com/siderolabs/talos/commit/aa5946dd385a2b99d572f9318e4eeeeee441b51b) test: fix cron failures for provision-1 & provision-2
* [`1dd701efa`](https://github.com/siderolabs/talos/commit/1dd701efa8119b6515a62ff68c430c99a96f2b68) fix: allow blockdevice wipe in maintenance mode
* [`786bf00ab`](https://github.com/siderolabs/talos/commit/786bf00abb309955616e440cd06fd0718b1b77ab) feat: add --platform=all support to image cache-create
* [`e1f645e3c`](https://github.com/siderolabs/talos/commit/e1f645e3cbeee5306dc0075deb8942793eb80a81) feat: validate luks headers for tampering
* [`ad72c7300`](https://github.com/siderolabs/talos/commit/ad72c73006abc3b51e5371496c61d8637b2222f0) test: improve maintenance API provision tests
* [`70cefab6a`](https://github.com/siderolabs/talos/commit/70cefab6af3dacdc80921b55ca8dbf5644501c6c) test: fix the flakes in tests with trusted roots
* [`aacff17f4`](https://github.com/siderolabs/talos/commit/aacff17f4c8890d6cada8efc6e715f69750f79cd) test: bump memory for Flannel netpolicy tests
* [`9c3459114`](https://github.com/siderolabs/talos/commit/9c34591144f1e2fc759fdc6d56694541eb9f241a) feat: update Linux to 6.18.19, CNI to 1.9.1
* [`038cb8735`](https://github.com/siderolabs/talos/commit/038cb87354eea1c1ff4612bdd13d1e77e595955a) feat: enforce PID check on connections to services over file sockets
* [`e2b2dd3ea`](https://github.com/siderolabs/talos/commit/e2b2dd3ea7eed8bc139cd0bd812253baee0dd95c) chore: update go-kubernetes library
* [`9597714f6`](https://github.com/siderolabs/talos/commit/9597714f625ac07bf74de32a24c3e6dad5abdc91) fix: add symlinks nvidia-ctk and nvidia-cdi-hook in /usr/bin
* [`8ac47d677`](https://github.com/siderolabs/talos/commit/8ac47d677703624ec6568294d94dcad7e533e6c4) fix: unset rlimits for extension services
* [`b1a02f368`](https://github.com/siderolabs/talos/commit/b1a02f3681c7e361ee6a3ef3d230b47480b48408) feat: update Kubernetes to 1.36.0-beta.0
* [`362fdc9ec`](https://github.com/siderolabs/talos/commit/362fdc9ece81e805a5a6a4e0303bdf78a6b2c35d) feat: update etcd to 3.6.9
* [`0a47f40b3`](https://github.com/siderolabs/talos/commit/0a47f40b3cdf304a079c6b3fa964e9f82e91ec63) fix(machined): clear stale bond ARP/NS targets on decode
* [`86344639f`](https://github.com/siderolabs/talos/commit/86344639fcb76d9430ac1e975c98db4488701e43) fix: update diff library to v1.0.1
* [`eff89d1ed`](https://github.com/siderolabs/talos/commit/eff89d1ed46e5f3c709305a8cb134dabae925420) fix: panics in diff algorithms
* [`8e1c8a7a9`](https://github.com/siderolabs/talos/commit/8e1c8a7a90fb039fd8a639a1218c169bc683d141) test: fix the apid test against AWS/GCP
</p>
</details>

### Dependency Changes

* **github.com/bougou/go-ipmi**                  v0.8.3 -> v0.9.1
* **github.com/cosi-project/runtime**            v1.16.2 -> v1.16.3
* **github.com/planetscale/vtprotobuf**          ba97887b0a25 -> 8ae5a48058df
* **github.com/siderolabs/talos**                v1.13.6 -> v1.14.1
* **github.com/siderolabs/talos/pkg/machinery**  v1.13.6 -> v1.14.1
* **github.com/stretchr/testify**                v1.11.1 -> v1.12.1
* **golang.org/x/sync**                          v0.22.0 -> v0.23.0
* **google.golang.org/grpc**                     v1.82.1 -> v1.84.0
* **google.golang.org/protobuf**                 f2248ac996af -> v1.36.12

Previous release can be found at [v0.1.6](https://github.com/siderolabs/talos-metal-agent/releases/tag/v0.1.6)

## [talos-metal-agent 0.1.6](https://github.com/siderolabs/talos-metal-agent/releases/tag/v0.1.6) (2026-07-21)

Welcome to the v0.1.6 release of talos-metal-agent!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/talos-metal-agent/issues.

### Contributors

* Andrey Smirnov
* Noel Georgi
* Mateusz Urbanek
* Erwan Leboucher
* Maja Bojarska
* Utku Ozdemir
* Aleksei Sviridkin
* Jonny
* Orzelius

### Changes
<details><summary>1 commit</summary>
<p>

* [`54525a6`](https://github.com/siderolabs/talos-metal-agent/commit/54525a6d6dd9f2f2f48ccf807d1e81b411b73e5c) chore: bump deps, rekres, and boot assets Talos version
</p>
</details>

### Changes from siderolabs/talos
<details><summary>88 commits</summary>
<p>

* [`04318854e`](https://github.com/siderolabs/talos/commit/04318854eb64c90e99308b844b45b26b0077489e) release(v1.13.6): prepare release
* [`9d8e47dd3`](https://github.com/siderolabs/talos/commit/9d8e47dd3cbb6cadf697dc05e29d4343f29a8d7a) chore: update pkgs and tools
* [`31552f400`](https://github.com/siderolabs/talos/commit/31552f40099cdcf181b8bee820055847b86ecda4) fix: shutdown/reboot via usermode helpers
* [`bc0c3f3d3`](https://github.com/siderolabs/talos/commit/bc0c3f3d39327b04dcaa0c15bea8b06a8c64b8a1) fix: flaky serviceaccount suite test
* [`3e7559258`](https://github.com/siderolabs/talos/commit/3e7559258e8da86b94114e2fc6b5b293d465dc09) fix: flaky tests
* [`fbe4d900d`](https://github.com/siderolabs/talos/commit/fbe4d900d983ab81fe5b5faacdbe5d5d124a613e) fix: data race in manifest sync
* [`6df3a452b`](https://github.com/siderolabs/talos/commit/6df3a452b0aaea5dab75e3221c2a600f5ab1b94e) fix: provide cooldown period for the QoS trigger
* [`85f8dd63e`](https://github.com/siderolabs/talos/commit/85f8dd63ed7dac59ab939ee0be1acb7e66659a08) fix: decode extraArgs list values correctly
* [`c2a56d592`](https://github.com/siderolabs/talos/commit/c2a56d592a9c7f1246052148e341e9bd5d78fb29) fix: kubelet stuck restarting
* [`871440858`](https://github.com/siderolabs/talos/commit/8714408587bea7d9f70be41e29ce69e61978abdc) chore: bump rekor for GHSA-47q9-m4ww-924m
* [`3e37ef8cd`](https://github.com/siderolabs/talos/commit/3e37ef8cddffed23e67adedc9bef5398353b70db) fix: handle image cache being disabled
* [`466bcd804`](https://github.com/siderolabs/talos/commit/466bcd80466c2897c7a254a21f0ef85b80da4cd0) fix: align documented image cache partition label
* [`d3cf09bcb`](https://github.com/siderolabs/talos/commit/d3cf09bcb8d58b6d548871fff6a57594744fc034) fix: image verification with referrers
* [`e9609b992`](https://github.com/siderolabs/talos/commit/e9609b992fde91094acc1eb1338b1f60d2d5252b) feat: add AMD XGBE driver to initramfs
* [`f18efcc4d`](https://github.com/siderolabs/talos/commit/f18efcc4d6f603dde1967a2045bea161768fceaf) chore: update deps
* [`51b0d8ed0`](https://github.com/siderolabs/talos/commit/51b0d8ed0097b1137b748a90ec7dcc3e902f1949) release(v1.13.5): prepare release
* [`c5089c655`](https://github.com/siderolabs/talos/commit/c5089c6552c9f98a1b65a938f845de26224094d8) fix: bump number of open files for etcd
* [`e0b4d9d75`](https://github.com/siderolabs/talos/commit/e0b4d9d75133322e0229509358442492d883a665) fix: stop the log persistence and close all files on shutdown
* [`23a080dcf`](https://github.com/siderolabs/talos/commit/23a080dcffa33a5003260154309b50c1a4ac46e3) fix: honor FailurePauseTimeout when pausing before reboot
* [`9adc63a32`](https://github.com/siderolabs/talos/commit/9adc63a32bf110324f6f603aaa74a4ec3880ea6b) fix: correct the link alias condition
* [`b902f9de9`](https://github.com/siderolabs/talos/commit/b902f9de90184c60cc1f1137f85ab4377a2de988) feat: verify go.mod tidiness in generate target
* [`765f0a1dc`](https://github.com/siderolabs/talos/commit/765f0a1dc9cb444c459f7be14f1d824032b9bb2e) fix: relax LUKS header validation
* [`d63aba4c7`](https://github.com/siderolabs/talos/commit/d63aba4c7302d24e12dce130e9018789e9198f58) feat: update pkgs and Kubernetes
* [`f0a5842ab`](https://github.com/siderolabs/talos/commit/f0a5842abeb961822be4c01785bc6d8abc7510fa) fix: update go.mod and rekres
* [`707dbd89b`](https://github.com/siderolabs/talos/commit/707dbd89b8716a930d83a31857c19b796e2ca290) release(v1.13.4): prepare release
* [`27d7a1985`](https://github.com/siderolabs/talos/commit/27d7a1985c4fff351846432c314373cf54d60340) fix: handle cluster-scoped resources with a namespace correctly
* [`fe74e00fb`](https://github.com/siderolabs/talos/commit/fe74e00fb46d96ac8a9348e09491d3aeb4ea8e16) chore: update deps
* [`f44cafbcd`](https://github.com/siderolabs/talos/commit/f44cafbcd2ca614373c7051874f1167f190a0e3d) fix: recreate dns server and listeners on host DNS runner restart
* [`5ed296b76`](https://github.com/siderolabs/talos/commit/5ed296b76eb035428db70f20bfa605bfbe068bd5) fix: marshal kube-scheduler config correctly with int types
* [`5992015b0`](https://github.com/siderolabs/talos/commit/5992015b0e119743696724a4d112da3396cee195) fix: machine configuration schemas
* [`b8dfda7ee`](https://github.com/siderolabs/talos/commit/b8dfda7ee4d88be4e552a55ca180dd0d55902185) fix: mark more resources as sensitive
* [`aad841b7f`](https://github.com/siderolabs/talos/commit/aad841b7f3041404b0984c71235a78e093ea8591) feat: update Flannel to v0.28.5
* [`7c0900b85`](https://github.com/siderolabs/talos/commit/7c0900b856b3d6276c3cb90443fd54d9e4cff1f3) fix(ci): aws nvidia tests
* [`9f5122db7`](https://github.com/siderolabs/talos/commit/9f5122db76c7d1528fcf5186b56d60eb93e8fa6d) fix: flaky test
* [`cf62af3e2`](https://github.com/siderolabs/talos/commit/cf62af3e27b68d18a031339cb49055d8ecdc1139) fix: etcd client leak in the (legacy) Upgrade API
* [`d5c3136e0`](https://github.com/siderolabs/talos/commit/d5c3136e07994497f28a94f657501b89ccf52911) feat: enforce strict QoS ordering in OOM victim selection
* [`b5ad39e65`](https://github.com/siderolabs/talos/commit/b5ad39e657777e3147d7018614e903bf8d7e99ad) feat: update etcd to v3.6.12
* [`c83dad3c5`](https://github.com/siderolabs/talos/commit/c83dad3c55502ba5268086be947e92164842b771) fix: health request server-side
* [`577cc6f6c`](https://github.com/siderolabs/talos/commit/577cc6f6c305b9328060d0e8e01f03723e4c491e) fix: bring in a change to BCM2712_MIP
* [`29da68ae2`](https://github.com/siderolabs/talos/commit/29da68ae288920360ef2547370bdb5390c9f26cb) fix: touch rootfs files with SOURCE_DATE_EPOCH
* [`b19a03bc2`](https://github.com/siderolabs/talos/commit/b19a03bc2c8faabda6c92d5b70d703a4f7333d31) fix: ignore cgroups with zero rank in OOM handler
* [`befeda7cb`](https://github.com/siderolabs/talos/commit/befeda7cb3b8a771a965e0f944b13d80b8151a95) release(v1.13.3): prepare release
* [`f4d451054`](https://github.com/siderolabs/talos/commit/f4d45105468c9b8fe7f624f0e3587f813ea2fd5b) feat(ci): rotate credentials
* [`01b434870`](https://github.com/siderolabs/talos/commit/01b4348701893f70c7e9764bdd465f72a0c6ae8b) fix: guard apply config API call
* [`a42c37f24`](https://github.com/siderolabs/talos/commit/a42c37f24eb32130bb5207b248a7408d1e047a42) feat(machined): support instance tags on Akamai
* [`d62d54ca7`](https://github.com/siderolabs/talos/commit/d62d54ca74b604e7ba6058e6d1b796d8217b3ee8) fix: memorymodules resource reporting
* [`b673b4be7`](https://github.com/siderolabs/talos/commit/b673b4be77f7e8f8563d5dd6b87796ce902cae4a) fix: bump Go golang.org/x modules
* [`19755ad14`](https://github.com/siderolabs/talos/commit/19755ad14ed454cea97881e96ec3ada01636eed4) feat: add bnxt_re module to the rootfs
* [`532bc6baa`](https://github.com/siderolabs/talos/commit/532bc6baa9029ce1bda696634247c8bace33ef74) fix: relax hostname config validation
* [`3bbd3ed35`](https://github.com/siderolabs/talos/commit/3bbd3ed35172d640ccd58dba9927db56525b6a45) fix: bump Kubernetes to 1.36.1 in one more place
* [`472b9d991`](https://github.com/siderolabs/talos/commit/472b9d99102ae33b4f561a37bc79ce46b601124c) feat: update default Kubernetes version to 1.36.1
* [`6d53ce0d5`](https://github.com/siderolabs/talos/commit/6d53ce0d5846f4043a84882f40b16766c303c759) chore(ci): fix cloud image upload job name
* [`5633c7791`](https://github.com/siderolabs/talos/commit/5633c779121ca8b3420b4a4573e242b1beba751a) fix: rework how scheduler config is marshaled
* [`52f056084`](https://github.com/siderolabs/talos/commit/52f0560845d5050c5e4dc221fbdc26b847c49cda) fix: restore some shared (and some lower tier slave) mount propagation
* [`9de3c12d9`](https://github.com/siderolabs/talos/commit/9de3c12d960f523ccb7a84a52115a18a13e78e41) fix: image verification issue with registry.k8s.io
* [`7dc716d85`](https://github.com/siderolabs/talos/commit/7dc716d850e43373ba8702c702105f4edd68fceb) feat: redact more machine config secrets and audit redactors
* [`d5448c60d`](https://github.com/siderolabs/talos/commit/d5448c60d50ec2d8548fd307c28b1725bd20b77c) chore(ci): try fixing homebrew action
* [`ef9f0bf02`](https://github.com/siderolabs/talos/commit/ef9f0bf021adb6e13e1ffd825aa4e94a6192ac1e) docs: drop controlplane endpoint examples
* [`7ee3e787b`](https://github.com/siderolabs/talos/commit/7ee3e787b8f399a8dd8a8eeb398d83068046256f) feat: update Linux to 6.18.33
* [`e99744bad`](https://github.com/siderolabs/talos/commit/e99744badec5ac74eb7fe44b90a9056993c537c6) fix: update containerd to 2.2.4
* [`c5d7c6536`](https://github.com/siderolabs/talos/commit/c5d7c65366e9bd767175faa1c8644eef2dd30697) release(v1.13.2): prepare release
* [`7df617aa7`](https://github.com/siderolabs/talos/commit/7df617aa74a44aa353aaede8bdb60be4f3b46f50) release(v1.13.1): prepare release
* [`09ead22a3`](https://github.com/siderolabs/talos/commit/09ead22a3cb86d977bfa17919b6f7edcb3e7101e) test: relax kernel-default routing rule assertion
* [`817609677`](https://github.com/siderolabs/talos/commit/817609677f1e3c30abd2be56d0638fbb4fd69ef0) feat: update Go to 1.26.3
* [`a5f32abda`](https://github.com/siderolabs/talos/commit/a5f32abda13832586ccbc1ef49584cee529f7e18) fix: normalize source name for syft consistency
* [`f8298948a`](https://github.com/siderolabs/talos/commit/f8298948a8316c2b01147451340114925e98695e) feat: bump in-toto indirect dependency
* [`ded9a2d78`](https://github.com/siderolabs/talos/commit/ded9a2d78340d95a9cdf8e33bd5e7c7ba876758d) feat: update kernel to 6.18.29
* [`755628239`](https://github.com/siderolabs/talos/commit/75562823938217b092c45fecffff90c7e8200e1e) fix: handle empty GCP operation errors
* [`e7645ba1c`](https://github.com/siderolabs/talos/commit/e7645ba1ccae5b0c22b5beab3936a4b12bd91df1) fix: clarify documentation for image verification pattern
* [`e85d01a07`](https://github.com/siderolabs/talos/commit/e85d01a07ba91ef971735612fc33a369bb132c0d) fix: skip reserved routing rule priorities
* [`c5a81f2cc`](https://github.com/siderolabs/talos/commit/c5a81f2cc88f6c0bf9759008fded3bb9a87c3c9c) feat: update etcd to 3.6.11
* [`38ca2bca6`](https://github.com/siderolabs/talos/commit/38ca2bca6d20da0a34d48908d67e8e9196e4b091) fix: add missing kernel modules in rootfs
* [`dc30ad327`](https://github.com/siderolabs/talos/commit/dc30ad327568f47767966873b37c3f26bf7fca4c) fix: preserve DHCP DNS servers
* [`d8e32fa73`](https://github.com/siderolabs/talos/commit/d8e32fa73d1a330652b53f524cc426d7e83d22d7) fix: stale discovered volume children
* [`80c110c87`](https://github.com/siderolabs/talos/commit/80c110c87c347d67f57a6941ada7b7fb03465e88) fix: re-enable kexec on arm64
* [`bd9ac044e`](https://github.com/siderolabs/talos/commit/bd9ac044e2851864c3cebb9d93ae665a0e5786a0) fix: provide proper AWS platform metadata
* [`549f3c0b4`](https://github.com/siderolabs/talos/commit/549f3c0b4c32e79246352f491ffde44e51eec0f1) fix: panic in Kubernetes manifest sync
* [`29eb6651d`](https://github.com/siderolabs/talos/commit/29eb6651d67ebbe671c9b6e4522f7e32c708ad3d) fix(ci): zfs test
* [`4b36fc9c2`](https://github.com/siderolabs/talos/commit/4b36fc9c260133d915ba11c2eef8608c57fb994c) fix: deadlock in the makefs ext4 with populated source
* [`fdf4f9f6c`](https://github.com/siderolabs/talos/commit/fdf4f9f6c77f9a7fe5a691bba3419ff81d6153af) fix: do not pick up a system disk from a loop device
* [`4ff29cc9f`](https://github.com/siderolabs/talos/commit/4ff29cc9fbd6ccdd9b25323d9a379cadd873377d) fix(talosctl): protect k8sNames map writes with mutex
* [`ff53434c9`](https://github.com/siderolabs/talos/commit/ff53434c96a50888444edccb4ead62f219f67389) fix: mount throws EPERM on virtiofs with SELinux
* [`16cc0a99c`](https://github.com/siderolabs/talos/commit/16cc0a99cd421fe070147775e36faad2e432c933) fix: drop explicit platform matcher
* [`ddb631aba`](https://github.com/siderolabs/talos/commit/ddb631aba8b7deee9254d63de7594f6617c5c5c5) fix: bump go-kmsg to fix the timestamp drift
* [`595470849`](https://github.com/siderolabs/talos/commit/595470849dae91026bb9404393ff4c276c8832ab) fix: make lacp active nilable
* [`879e31a65`](https://github.com/siderolabs/talos/commit/879e31a65243c312edc2fb5631a2acce45b8a639) test: fix flaky tests
* [`ef1d9ffc3`](https://github.com/siderolabs/talos/commit/ef1d9ffc36b551549e88949abe91788e18345075) fix: reset the ticker when the KubeSpan is disabled/enabled
* [`ce89d6727`](https://github.com/siderolabs/talos/commit/ce89d672708e860a47c8023c85ef751b0a38b655) fix: replace Canal manifest with a more recent one
</p>
</details>

### Dependency Changes

* **github.com/bougou/go-ipmi**                  v0.8.1 -> v0.8.3
* **github.com/cosi-project/runtime**            v1.14.1 -> v1.16.2
* **github.com/siderolabs/talos**                v1.13.0 -> v1.13.6
* **github.com/siderolabs/talos/pkg/machinery**  v1.13.0 -> v1.13.6
* **go.uber.org/zap**                            v1.27.1 -> v1.28.0
* **golang.org/x/sync**                          v0.20.0 -> v0.22.0
* **google.golang.org/grpc**                     v1.80.0 -> v1.82.1

Previous release can be found at [v0.1.5](https://github.com/siderolabs/talos-metal-agent/releases/tag/v0.1.5)

## [talos-metal-agent 0.1.4](https://github.com/siderolabs/talos-metal-agent/releases/tag/v0.1.4) (2026-02-02)

Welcome to the v0.1.4 release of talos-metal-agent!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/talos-metal-agent/issues.

### Contributors

* Andrey Smirnov
* Noel Georgi
* Mateusz Urbanek
* Orzelius
* Dmitrii Sharshakov
* Amarachi Iheanacho
* Justin Garrison
* Laura Brehm
* Oguz Kilcan
* Orzelius
* Utku Ozdemir
* Steve Francis
* Till Hoffmann
* Andrew Longwill
* Spencer Smith
* Alexandre GV
* Artem Chernyshev
* Bryan Lee
* George Gaál
* 459below
* Adrian L Lange
* Alp Celik
* Alvaro "Chamo" Linares Cabre
* Brian Brookman
* Bryan Mora
* Chris Sanders
* Christopher Puschmann
* Clément Nussbaumer
* Damien
* David R
* Dmitry
* Edward Sammut Alessi
* Febrian
* Florian Grignon
* Giau. Tran Minh
* Grzegorz Rozniecki
* Joakim Nohlgård
* Jonas Lammler
* Jorik Jonker
* Justin Seely
* Lennard Klein
* Luke Cousins
* Marat Bakeev
* Markus Freitag
* Markus Reiter
* Martyn Ranyard
* Max Makarov
* Michael Moerz
* Michael Robbins
* Michael Smith
* Mike
* Mike Beaumont
* Misha Aksenov
* MrMrRubic
* Olav Thoresen
* Olivier Doucet
* Pranav
* Serge Logvinov
* Skye Soss
* Skyler Mäntysaari
* SuitDeer
* Tan Siewert
* Thibault VINCENT
* Tom
* Tom Keur
* aurh1l
* frozenprocess
* frozensprocess
* jvanthienen-gluo
* kassad
* killcity
* leppeK
* samoreno
* theschles
* winnie
* yashutanu

### Changes
<details><summary>3 commits</summary>
<p>

* [`f2f51f9`](https://github.com/siderolabs/talos-metal-agent/commit/f2f51f98b903202a98236ecaaeb3337ef7a57f0f) fix: default to IPMI port 623 when it is unsupported
* [`b475ccc`](https://github.com/siderolabs/talos-metal-agent/commit/b475ccc8d14ce28e03323f9b5ab8eafb581c3207) chore: bump deps, rekres
* [`8e92d6e`](https://github.com/siderolabs/talos-metal-agent/commit/8e92d6eeedd1cefb8e0473f1051d274807df2292) chore: bump extensions ref in boot assets image
</p>
</details>

### Changes from siderolabs/talos
<details><summary>609 commits</summary>
<p>

* [`54e5b438d`](https://github.com/siderolabs/talos/commit/54e5b438d8dcf6395e6424808d1155d02abf3bc0) release(v1.12.2): prepare release
* [`30da0bc19`](https://github.com/siderolabs/talos/commit/30da0bc19eb699dabf966cce38ef4477add193d4) fix: oracle platform file format
* [`7ddb37b1f`](https://github.com/siderolabs/talos/commit/7ddb37b1f3e2abf6c3406d35be92093fe4512eff) fix: make OOM expression a bit less sensitive
* [`e438ec23e`](https://github.com/siderolabs/talos/commit/e438ec23eefef97bbaa160dd6bb133b48a267ac7) fix: marshal of FailOverMac property
* [`717ed7265`](https://github.com/siderolabs/talos/commit/717ed726569d1270e2fb48df60e5fd7f43d1885b) fix: check if the device is not mounted when wiping
* [`c95c9fd06`](https://github.com/siderolabs/talos/commit/c95c9fd06508f02a770100f87da754a6fd3b9fa8) fix: wipe the first/last 1MiB in addition to wiping by signatures
* [`52bed358d`](https://github.com/siderolabs/talos/commit/52bed358d3606d04e6b4acded5dfe26cdb5f0ec9) fix: add talos version to Hetzner Cloud client user agent
* [`0e447a431`](https://github.com/siderolabs/talos/commit/0e447a4318ff2b7a398a719144690b22dce1e3f7) fix: make OOM controller more precise by considering separate cgroup PSI
* [`3b974b99e`](https://github.com/siderolabs/talos/commit/3b974b99e583c3a5bdd80e239517ef1ebc19de9c) fix: sort mirrors and tls configs when generating the machine config
* [`8b16fe50b`](https://github.com/siderolabs/talos/commit/8b16fe50bb44c7cb4bd3f50580a3ea18cdc3a727) feat: add VLAN support to OpenStack platform
* [`eb8480c4c`](https://github.com/siderolabs/talos/commit/eb8480c4ce088bd9fe705302c7e588aa01da207b) fix: panic in configpatcher when the whole section is missing
* [`4d44306dd`](https://github.com/siderolabs/talos/commit/4d44306dd148c872803578dc3880bbab307612b9) fix: wipe disk by signatures
* [`cca4cd269`](https://github.com/siderolabs/talos/commit/cca4cd269b0a4ac24627d195fad4bd9fa00c3f85) feat: add it87 hwmon module
* [`d9480eef2`](https://github.com/siderolabs/talos/commit/d9480eef2ed45b35d5f1782b651c1499451536c5) fix: resolve SideroLink Wireguard endpoint on reconnect
* [`e16c2d5bb`](https://github.com/siderolabs/talos/commit/e16c2d5bba1b6dce241905dc9e4846d45a774f78) fix: handle correctly incomplete RegistryTLSConfig
* [`dedd273df`](https://github.com/siderolabs/talos/commit/dedd273dfcd5d721e63cbe0124623ce2b5e50df4) fix: bond config via platform
* [`f527cff23`](https://github.com/siderolabs/talos/commit/f527cff239cf246891ef6e053d0aec5ce8900e22) fix: allow HostnameConfig to be used with incomplete machine config
* [`10918136c`](https://github.com/siderolabs/talos/commit/10918136c6338506d08dd86b57d82b880ea50348) fix: lock down etcd listen address to IPv4 localhost
* [`9f8d938db`](https://github.com/siderolabs/talos/commit/9f8d938db68f4c872ccf65573339e4761b4a09d4) fix: print talosctl images to release notes
* [`95433c167`](https://github.com/siderolabs/talos/commit/95433c167493a7650513379866e544bdb0adbc2e) fix: update VIP config example
* [`919394fee`](https://github.com/siderolabs/talos/commit/919394fee8122bd583ac1f0cfc55d8a0d3e3d3cb) feat: update Go to 1.25.6
* [`7ea2ef7cf`](https://github.com/siderolabs/talos/commit/7ea2ef7cf4d0d48ac9b30eca9b7ec17aa83fde50) release(v1.12.1): prepare release
* [`78a785604`](https://github.com/siderolabs/talos/commit/78a785604ad40eb9f1634c9db5477bd6ce99428c) chore: run rekres and update dependencies
* [`c31067173`](https://github.com/siderolabs/talos/commit/c3106717392a34fcca959b414f5064d6c799eaa3) fix: disable swap for system services
* [`a7e8426cf`](https://github.com/siderolabs/talos/commit/a7e8426cfb46f4c46476243032e2f4ade1fe9dfc) test: skip the source bundle on exact tag
* [`943984167`](https://github.com/siderolabs/talos/commit/943984167c22af0853d2c956677a241acece807f) fix: probe small images correctly
* [`42df71637`](https://github.com/siderolabs/talos/commit/42df71637763b1bf10bdf0fe89f650c367605b8c) fix: invalid versions check in talos-bundle
* [`a3e90e445`](https://github.com/siderolabs/talos/commit/a3e90e445f0f99b050eb98fcd9565b2b5e3397bf) fix: make upgrade work with SELinux enforcing=1
* [`ac91ade2c`](https://github.com/siderolabs/talos/commit/ac91ade2c7e435e63ed2546244d428a81abd22ad) release(v1.12.0): prepare release
* [`82553b2a1`](https://github.com/siderolabs/talos/commit/82553b2a1a713836f496b822e86e5e6788c5ebd1) fix: mount volume mount/unmount race
* [`33f6e22ec`](https://github.com/siderolabs/talos/commit/33f6e22ecb3b393d1488730c67d6f973a46b0b39) fix: bond setting change detection
* [`d5be50ac5`](https://github.com/siderolabs/talos/commit/d5be50ac55cac1c1c1deff4971fd991f364696a1) docs: split talosctl commands into groups
* [`70d3ab9ac`](https://github.com/siderolabs/talos/commit/70d3ab9ac090095c2fc8cbbfaa9c5c472d76c794) feat: update Kubernetes to 1.35.0
* [`101814d88`](https://github.com/siderolabs/talos/commit/101814d889924afe7c049106c638a32ae107a139) feat: update etcd 3.6.7, CoreDNS 1.13.2
* [`ce286825a`](https://github.com/siderolabs/talos/commit/ce286825a7f969f847ea7ad17bd2a31fa85d301c) fix: drop the Omni API URL check on IP address
* [`96f724adc`](https://github.com/siderolabs/talos/commit/96f724adccbc6fac844f9a341e36eede331b3947) feat: enable Powercap and Intel RAPL
* [`e195427c1`](https://github.com/siderolabs/talos/commit/e195427c17a004b5bcaa6f1870ce6c855ae61f1d) docs: fix the talosctl cluster create help output
* [`e025355b7`](https://github.com/siderolabs/talos/commit/e025355b759bb110925631f5f84230e99b9069df) feat(talosctl): support running qemu x86 on Mac
* [`21a914a1d`](https://github.com/siderolabs/talos/commit/21a914a1d1ca48d6bb4d47ddc8be0d0fdf74800d) fix: exclude new Virtual IPs configured with new config
* [`ca645777d`](https://github.com/siderolabs/talos/commit/ca645777dae5ad07501501dafc4717e7383045b0) fix: provide json support in `nft` binary
* [`6dd0558a3`](https://github.com/siderolabs/talos/commit/6dd0558a314af9a0dfda77b4f58a7115ef86b6fc) feat: sync pkgs
* [`c931847cc`](https://github.com/siderolabs/talos/commit/c931847ccaadf84f84e5f2befadaffb55740b592) feat: update containerd to v2.1.6
* [`a2a77004d`](https://github.com/siderolabs/talos/commit/a2a77004deac3efe6ac14f906a8bd0a3b0f926ca) release(v1.12.0-rc.1): prepare release
* [`47198780b`](https://github.com/siderolabs/talos/commit/47198780bfc084347b9ae675aaeb27a1c1d58d38) fix: bond configuration with new settings
* [`03a424bdf`](https://github.com/siderolabs/talos/commit/03a424bdf1b8a270dd694fc2738d81a3261d80cf) fix: disable kexec on arm64
* [`688fb789b`](https://github.com/siderolabs/talos/commit/688fb789beb979544e16447e512419629ea61b21) feat: add Secure Boot to CloudStack platform config
* [`66e67fd13`](https://github.com/siderolabs/talos/commit/66e67fd1394946b3425543a1aac52d4a8338e375) fix: discard better klog message from Kubernetes client
* [`d8403498c`](https://github.com/siderolabs/talos/commit/d8403498c92e0f9c37b04ad6786b2c84df5e7c95) fix: disable kexec in talosctl cluster create on arm64
* [`5ced4258c`](https://github.com/siderolabs/talos/commit/5ced4258c18f5649590a50c2927ab8e16db298ec) fix: do not override DNS on MacOS
* [`fabf3f0e7`](https://github.com/siderolabs/talos/commit/fabf3f0e73918b650b33ef0f009cacb9a15ecbc0) fix: selection of boot entry
* [`93cec4b9d`](https://github.com/siderolabs/talos/commit/93cec4b9dfdef0566152ef80c28439a7dbb0c320) fix: update CNI plugins to 1.9.0
* [`964098d96`](https://github.com/siderolabs/talos/commit/964098d9696a804de5d27284cd79dccffa7c81b9) fix: update KubeSpan MSS clamping
* [`bce04084d`](https://github.com/siderolabs/talos/commit/bce04084d6f5a9c703c7d63d1558d7d43c54dfbf) feat: rename image default and source bundle
* [`d1abc0f84`](https://github.com/siderolabs/talos/commit/d1abc0f8473c1a562e37a712624f803ce0f60fec) chore: update pkgs
* [`061307687`](https://github.com/siderolabs/talos/commit/0613076873bbd2d763da30ae2e9e1903486f7cb8) release(v1.12.0-rc.0): prepare release
* [`bc4de5b79`](https://github.com/siderolabs/talos/commit/bc4de5b7926a9a2e7a7af9da4763effb5c33693e) fix: constants file
* [`4a15763a9`](https://github.com/siderolabs/talos/commit/4a15763a962cad0c020e01f66948ba1f326c9201) docs: update release notes
* [`297336549`](https://github.com/siderolabs/talos/commit/29733654902be5cb72b71a9a64ea0ed3c0a0f011) fix: correct condition to use UKI cmdline in GRUB
* [`0ac58929d`](https://github.com/siderolabs/talos/commit/0ac58929db6960ef91c1bcfbc891264e18e1e930) docs: drop machine.network example
* [`184a45c40`](https://github.com/siderolabs/talos/commit/184a45c405530c73c31d5b6c642cda4ddd1772ca) test: bird2 extension
* [`8eac9f37d`](https://github.com/siderolabs/talos/commit/8eac9f37d9dddc507c988cfb187b939a5624f563) docs: add omni join token example to create qemu command
* [`e79a94d57`](https://github.com/siderolabs/talos/commit/e79a94d57781d6ede61e6205f6f5d0f0708a8ddb) fix: adapt SELinuxSuite.TestNoPtrace to new strace version
* [`7a1bb4c26`](https://github.com/siderolabs/talos/commit/7a1bb4c26a99c7f4e37196b40aced6334eeda731) fix: mark secureboot as supported for metal
* [`5c6ee6ace`](https://github.com/siderolabs/talos/commit/5c6ee6aceeb87785c08a05f2ddc6b7cbcad0bc9a) fix: clear provisioning data on SideroLink config change
* [`2e6fe4684`](https://github.com/siderolabs/talos/commit/2e6fe4684b98ca4432284b7b51dfcd1a8b91a03c) feat: update Linux to 6.18.0
* [`473bc17c1`](https://github.com/siderolabs/talos/commit/473bc17c199165dd0f925981753dec431cc5613b) feat: update Kubernetes to 1.35.0-rc.0
* [`6dc8e82b3`](https://github.com/siderolabs/talos/commit/6dc8e82b31d095a357b9f6d99420bb860e51261c) fix: add a timeout for DNS resolving for NTP
* [`a7dbbbd4d`](https://github.com/siderolabs/talos/commit/a7dbbbd4d87feeace427e4c63f67880c72f7cd22) fix: don't disable LACP by default
* [`3ca342c09`](https://github.com/siderolabs/talos/commit/3ca342c0979ffcfe7bee95a4e56c98ddece8abb5) chore: fix longhorn test
* [`364ebb6ba`](https://github.com/siderolabs/talos/commit/364ebb6baf3c77a1e2dd28d83b6af7cfe821e1e8) fix: selection of boot entry
* [`aa286d3f6`](https://github.com/siderolabs/talos/commit/aa286d3f6eb28a813c982a9cc1230c138e56b33a) feat: update Kubernetes to 1.35.0-beta.0
* [`f4891eebb`](https://github.com/siderolabs/talos/commit/f4891eebb192d2895f27f85502fd223290217d90) feat: implement logs persistence
* [`c9a4f95b4`](https://github.com/siderolabs/talos/commit/c9a4f95b42c3347266f60215558f6bde77d4f8a5) release(v1.12.0-beta.1): prepare release
* [`d321d7da0`](https://github.com/siderolabs/talos/commit/d321d7da04fa87e0622f6ec7b5311d5578c534ba) chore: correct condition for running k8s integration tests
* [`736f32a80`](https://github.com/siderolabs/talos/commit/736f32a8077aea0f4a72f3545571882b9d79207c) chore: disable k8s integration tests for 1GiB worker nodes
* [`d9de616c4`](https://github.com/siderolabs/talos/commit/d9de616c48056fc079e693439d4c91a85e154222) chore(ci): skip multipath extension tests
* [`57d6683cd`](https://github.com/siderolabs/talos/commit/57d6683cde0195194acf6880ee85c406216fecc1) chore: update pkgs and tools version
* [`949323ab5`](https://github.com/siderolabs/talos/commit/949323ab51bf5cb95922af7169b698d333c5c9ab) feat: present kernel log as `talosctl logs kernel`
* [`7531fcbc7`](https://github.com/siderolabs/talos/commit/7531fcbc76f3e59e2e8af823d72ffad2cfcaa40a) test: fix flaky LinkSpec/Wireguard test
* [`1dbc64d69`](https://github.com/siderolabs/talos/commit/1dbc64d698f6654e8f8ca5baa13ae9d56745fe6a) fix: simplify OOM expression
* [`0ffb1d857`](https://github.com/siderolabs/talos/commit/0ffb1d8577c9b4da0850a36e80708122b93de303) fix: trim trailing dots from certificate SANs
* [`9a2f6d9c9`](https://github.com/siderolabs/talos/commit/9a2f6d9c9ec5670a12fb033935661f70a80da503) fix: support specifying patch file without '@' symbol
* [`582b0feab`](https://github.com/siderolabs/talos/commit/582b0feab2845d3265cdc852adac78a723953408) fix: assign value of multicast setting properly
* [`16aa6ac47`](https://github.com/siderolabs/talos/commit/16aa6ac471d98b5cdea11d7a4d22ea1048cbd2ce) feat: update etcd to 3.6.6
* [`4396f09c8`](https://github.com/siderolabs/talos/commit/4396f09c8c82ca15b7c09dde8ff1c69a1fe32b08) docs: add API Server Cipher Suites changelog
* [`fdf6fe8e6`](https://github.com/siderolabs/talos/commit/fdf6fe8e6299d620abb3f5c23dcab3cb38fb9367) feat: update TLS cipher suites for API server
* [`139cce3b4`](https://github.com/siderolabs/talos/commit/139cce3b45a7643144aac3042d2bf291e097199d) fix: add CA subject to generated certificate
* [`9b294af22`](https://github.com/siderolabs/talos/commit/9b294af225677a87524491ebd2f21106931dead1) feat: generate mirrors patch
* [`15465f0c5`](https://github.com/siderolabs/talos/commit/15465f0c513ed46886c9f4179c996368843a2daf) fix: add more resilient move
* [`b4147e3a1`](https://github.com/siderolabs/talos/commit/b4147e3a17eebc775cc8ae6087ded6fced11a261) feat: extend flags of cache-cert-gen
* [`72d3d1c9f`](https://github.com/siderolabs/talos/commit/72d3d1c9f53e9b62c189a6369a3060aee4c98d9c) chore: remove spammy 'clean up unused volumes' logs
* [`d6c78de84`](https://github.com/siderolabs/talos/commit/d6c78de84745f27f3051c971451339e760c71397) feat: support TALOS_HOME env var
* [`4040e0814`](https://github.com/siderolabs/talos/commit/4040e0814fc186b2f4e1a2c25520ac08c4d07633) feat: implement multicast setting
* [`eb636dc1f`](https://github.com/siderolabs/talos/commit/eb636dc1f96d1739f1858c4bf825cedc3e0d11e2) feat: add multicast to linkconfig
* [`e34e458c4`](https://github.com/siderolabs/talos/commit/e34e458c4b141ace9604a49b890b2714a59a614e) feat: update dependencies
* [`36152d278`](https://github.com/siderolabs/talos/commit/36152d2787f0cbf3b2efda9c30596f991a811022) fix: add riscv64 talosctl to release artifacts
* [`aebbbaf27`](https://github.com/siderolabs/talos/commit/aebbbaf2746956dc5f88cce6a95061ba447bb36a) feat: support relative voume size
* [`3d997d742`](https://github.com/siderolabs/talos/commit/3d997d7421f3d1b3fda55c92d0e11d75d16daf26) release(v1.12.0-beta.0): prepare release
* [`e62384ba3`](https://github.com/siderolabs/talos/commit/e62384ba34031d43fadebdc84a7d31dd41bf0678) fix: re-creating STATE after partition drop
* [`6919d232a`](https://github.com/siderolabs/talos/commit/6919d232abbaaf44120b9c882e2bc27e4b95deee) docs: update kernel args size
* [`887b296dc`](https://github.com/siderolabs/talos/commit/887b296dc5b111cf54961c1346c4dca4744ccdf9) test: randomize MAC addresses used in the unit-tests
* [`6063fbf91`](https://github.com/siderolabs/talos/commit/6063fbf9124d1953d3bd933bed7f70d42ede2afb) feat: update dependencies
* [`542a67a06`](https://github.com/siderolabs/talos/commit/542a67a066a842a5673755323a3936894b0825ef) feat: add riscv64 build of talosctl
* [`68560b53a`](https://github.com/siderolabs/talos/commit/68560b53ab81335057c0c5524af6f6d2b6882bcf) fix: split volume/disk locators
* [`2c3d30e94`](https://github.com/siderolabs/talos/commit/2c3d30e94f426f2567e9cb97cc3ca9499f53cc7f) docs: fix image-cache-path flag description
* [`93f2e87c2`](https://github.com/siderolabs/talos/commit/93f2e87c2d00c69aacc5f4422182db01b9e617fd) feat: shorthand for generating secrets to stdout
* [`5e1de0035`](https://github.com/siderolabs/talos/commit/5e1de003596837ffe4cf9dd90df4ea121fa2eacc) feat: implement time and resolvers multi-doc configuration
* [`399240be3`](https://github.com/siderolabs/talos/commit/399240be3a51c7053afb9ac60b9e19bd05857615) feat: drop partitions on reset with system partitions wipe
* [`5cca96655`](https://github.com/siderolabs/talos/commit/5cca966557651bb3018ba15d01e0b87146e508fe) feat: add new rockchip sbcs
* [`00fe50d86`](https://github.com/siderolabs/talos/commit/00fe50d868b0463fa32f56ec154bd92bae732f11) fix: uefi bootorder setting
* [`3a881184b`](https://github.com/siderolabs/talos/commit/3a881184bf149410b93657e885796ecf5005b547) chore: improve error handling for system disk reset
* [`859194e67`](https://github.com/siderolabs/talos/commit/859194e6780018ec8e637e87884aa16d3a14cfa6) chore: extract system+user volume config transformers, test
* [`308c6bc41`](https://github.com/siderolabs/talos/commit/308c6bc414d5c6c207bc021ca2949df602725e52) feat: add full disk volumes
* [`82ac1119e`](https://github.com/siderolabs/talos/commit/82ac1119ec102cc591935bbf0afb73431832b775) feat: implement new registry configuration
* [`106f45799`](https://github.com/siderolabs/talos/commit/106f45799d29c7436592b9f1194f6beeed5e394a) feat: update Linux kernel with userfaultfd/VDPA
* [`721a1e0d7`](https://github.com/siderolabs/talos/commit/721a1e0d7cc0cb3eb4d957510accff7762ff366c) chore: rename+improve `client.ErrEventNotSupported`
* [`43f4e317f`](https://github.com/siderolabs/talos/commit/43f4e317f1976762f2999e71ccd6761248a85f12) fix: race between VolumeConfigController and UserVolumeConfigController
* [`66c01a706`](https://github.com/siderolabs/talos/commit/66c01a706f0b1dba88e30dbc1781d7fb7ef57756) chore: deprecate interactive installer mode
* [`957770f65`](https://github.com/siderolabs/talos/commit/957770f65af0d50670b7bbe3758246ced37e9a3e) feat(machined): add panic/force mode reboot
* [`60be0daf8`](https://github.com/siderolabs/talos/commit/60be0daf8414a69b1a60970b14aceb872b31e415) feat: implement multi-doc Wireguard config
* [`cf014cb5d`](https://github.com/siderolabs/talos/commit/cf014cb5d3294ecdcf769315f4795fb8f82a239f) fix: only set default bootloader if none is set
* [`e9b016f80`](https://github.com/siderolabs/talos/commit/e9b016f809d83da33e57492df4a96d68a270ed8c) fix: use strict platform match when pulling images
* [`fafab391b`](https://github.com/siderolabs/talos/commit/fafab391b4d3947daad014438a833ae67b8995fe) feat: update Kubernetes to 1.35.0-alpha.3
* [`7bf3aaca9`](https://github.com/siderolabs/talos/commit/7bf3aaca9129ad40d49f9eadf7ad9be23cf99b32) feat: allow glibc aarch64 so files in extensions
* [`c8561ee2d`](https://github.com/siderolabs/talos/commit/c8561ee2d04c7f9f06c9ec1b3be34ef2a7057efc) feat: implement bridge multi-document config
* [`f4ad3077b`](https://github.com/siderolabs/talos/commit/f4ad3077b0c56b200a37e97abd1a51c63a04c648) feat: implement bond multi-doc configuration
* [`75fe47582`](https://github.com/siderolabs/talos/commit/75fe475828580d9b9a18a2fde0e59f7a9f047ca3) fix: stop attaching to tearing down mount parents
* [`c93a9c6b4`](https://github.com/siderolabs/talos/commit/c93a9c6b41396fe8f8f3f49f475d622e4a45b689) fix: improve OOM controller stability and make test strict on false positives
* [`021bbfefb`](https://github.com/siderolabs/talos/commit/021bbfefbecc688fc4c61876c264416f72c7a7a2) feat: update Go 1.25.4, containerd 2.1.5
* [`e25db484f`](https://github.com/siderolabs/talos/commit/e25db484f54414dcd7b8f08c1a741b58435e52f5) test: disable parallelism in Longhorn tests
* [`54b93aff0`](https://github.com/siderolabs/talos/commit/54b93aff0c372761dfe9621a782a347b6877c2e9) feat: update Linux 6.17.7, runc 1.3.3
* [`2af69ff35`](https://github.com/siderolabs/talos/commit/2af69ff35712ac843c66e30fdf6a380aae2ed499) fix: provide minimal platform metadata always
* [`92eeaa482`](https://github.com/siderolabs/talos/commit/92eeaa4826cf71a5962da8ea055a11732fbc851e) fix: update YAML library
* [`aa24da9aa`](https://github.com/siderolabs/talos/commit/aa24da9aab9c5dc2f51401ae8ba0161e63c09924) fix: bump kubelet credendial provider config to v1
* [`335f91761`](https://github.com/siderolabs/talos/commit/335f9176151f7d45c0f847abecb20184483a6cd3) feat: add short -c flag for --cluster
* [`4c095281b`](https://github.com/siderolabs/talos/commit/4c095281be93cb11290eb43f60b4cc1a168bef17) fix: set a timeout for SideroLink provision API call
* [`75e4c4a59`](https://github.com/siderolabs/talos/commit/75e4c4a598181a18638aadcb77c89fbe762c6b9f) fix: log duplication on log senders
* [`e3cbc92c0`](https://github.com/siderolabs/talos/commit/e3cbc92c0579beb0262d2d2d6a0d00d56bbbdc17) fix: add video kernel module to arm
* [`d69305a67`](https://github.com/siderolabs/talos/commit/d69305a670ac982ba7dd00cfc8e7cf736cbfb385) fix: userspace wireguard handling
* [`ee5fee7c8`](https://github.com/siderolabs/talos/commit/ee5fee7c8a0f482894534bd2f8e5b0c2b2076854) fix: image-signer commands
* [`be028b67a`](https://github.com/siderolabs/talos/commit/be028b67a068c0d0d4465725c96b28ad9b276e8a) feat: add support for multi-doc VLAN config
* [`f3df0f80b`](https://github.com/siderolabs/talos/commit/f3df0f80b9d64e282bf163ba04ed9363e40865a3) feat: add directory backed UserVolumes
* [`0327e7790`](https://github.com/siderolabs/talos/commit/0327e77902a05978c79a9efb92bc50a792e4e0be) feat: add support for dashboard custom console parameter
* [`fed948b8a`](https://github.com/siderolabs/talos/commit/fed948b8ae416db886df6ed783bde60aae2a25c8) release(v1.12.0-alpha.2): prepare release
* [`fb4bfe851`](https://github.com/siderolabs/talos/commit/fb4bfe851c7c308eeaf4a11e0ac5c944f66dc0c4) chore: fix LVM test
* [`f4ee0d112`](https://github.com/siderolabs/talos/commit/f4ee0d1128ba2f35d54ec3d35a83fc62fd222f2e) chore: disable VIP operator test
* [`288f63872`](https://github.com/siderolabs/talos/commit/288f6387260843570d53d28a4d77e564b3182979) feat: bump deps
* [`b66482c52`](https://github.com/siderolabs/talos/commit/b66482c529beda8b1abf9ed6b71ece354c1540be) feat: allow disabling injection of extra cmdline in cluster create
* [`704b5f99e`](https://github.com/siderolabs/talos/commit/704b5f99e6bef4410629427ac65fd2742ddb335d) feat: update Kubernetes to 1.35.0-alpha.2
* [`1dffa5d99`](https://github.com/siderolabs/talos/commit/1dffa5d9965a6c7d872f052bfb1750ea550671c2) feat: implement virtual IP operator config
* [`43b1d7537`](https://github.com/siderolabs/talos/commit/43b1d7537507a916629cc2d6db7440a99ffcb748) fix: validate provisioner when destroying local clusters
* [`b494c54c8`](https://github.com/siderolabs/talos/commit/b494c54c81e6ca81cef8ce26da772c1fc336ea8d) fix: talos import on non-linux
* [`61e95cb4b`](https://github.com/siderolabs/talos/commit/61e95cb4b7b354d175d1dfce3d0fa43deefad187) feat: support bootloader option for ISO
* [`d11072726`](https://github.com/siderolabs/talos/commit/d110727263c57c02392f201938d2b71976b8c4d6) fix: provide offset for partitions in discovered volumes
* [`39eeae963`](https://github.com/siderolabs/talos/commit/39eeae96311be2b8e2d3660d878f852ba92ca064) feat: update dependencies
* [`9890a9a31`](https://github.com/siderolabs/talos/commit/9890a9a31deb11ab170b94c667143314db08f76f) test: fix OOM test
* [`c0772b8ed`](https://github.com/siderolabs/talos/commit/c0772b8eda429675a06899b9c4a4d1dd7d5f6a5f) feat: add airgapped mode to QEMU backed talos
* [`ac60a9e27`](https://github.com/siderolabs/talos/commit/ac60a9e27deed63db0e4e61ffa30d46f4cab590a) fix: update test for PCI driver rebind/IOMMU
* [`6c98f4cdb`](https://github.com/siderolabs/talos/commit/6c98f4cdb049c58ef4f6e8193ef66c2338a2877d) feat: implement new DHCP network configuration
* [`da92a756d`](https://github.com/siderolabs/talos/commit/da92a756d9668fa043b4794db45d5c985d8ea4a6) fix: drop 'ro' falg from defaults
* [`28fd2390c`](https://github.com/siderolabs/talos/commit/28fd2390cb6e02f400bb237dd674c7d0d40f8ed3) fix: imager build on arm64
* [`4e12df8c5`](https://github.com/siderolabs/talos/commit/4e12df8c5c27ae115c4eac70a7e2fceb03dac5f5) test: integration test for OOM controller
* [`7e498faba`](https://github.com/siderolabs/talos/commit/7e498faba93f972ba82edf41550d3b94256e83e9) feat: use image signer
* [`eccb21dd3`](https://github.com/siderolabs/talos/commit/eccb21dd3ba03eb4ab03c4da87a51a4e3d8da49a) feat: add presets to the 'cluster create qemu' command
* [`ec0a813fa`](https://github.com/siderolabs/talos/commit/ec0a813facf5be5ca3e9ba65924ae18b2b05a7d9) feat: unify cmdline handling GRUB/systemd-boot
* [`37e4c40c6`](https://github.com/siderolabs/talos/commit/37e4c40c6a2477e45bbf067effc4389d4639c905) fix: skip module signature tests on docker provisioner only
* [`8124efb42`](https://github.com/siderolabs/talos/commit/8124efb42fd5a3eb81f41e84974e4242246ca7c4) fix: cache e2e
* [`4adcda0f5`](https://github.com/siderolabs/talos/commit/4adcda0f5427e1bae49f6dda58318324a3b24ac5) fix: reserve the apid and trustd ports from the ephemeral port range
* [`ced57b047`](https://github.com/siderolabs/talos/commit/ced57b047a389e26f7e5bfa3efab5b64f3fced87) feat: support optionally disabling module sig verification
* [`1e5c4ed64`](https://github.com/siderolabs/talos/commit/1e5c4ed644cbc60d8518fe4298e63a5cf5dc8cf5) fix: build talosctl image cache-serve non-linux
* [`dbdd2b237`](https://github.com/siderolabs/talos/commit/dbdd2b237e0aefbba439b90472abf9ec7eea6aa6) feat: add static registry to talosctl
* [`77d8cc7c5`](https://github.com/siderolabs/talos/commit/77d8cc7c589a190c8cb86e6e1684233129b648a1) chore: push `latest` tag only on main
* [`59d9b1c75`](https://github.com/siderolabs/talos/commit/59d9b1c75dbff09e405906ebcfb3ad1a69cb8f4b) feat: update dependencies
* [`bf6ad5171`](https://github.com/siderolabs/talos/commit/bf6ad51710c367764e582ccc1fb77b4d989c874d) feat: add back install script
* [`da451c5ba`](https://github.com/siderolabs/talos/commit/da451c5ba4ee97e7ef108bb6d73d5aa8bc7c72fd) chore: drop documentation except for fresh reference
* [`2f23fedeb`](https://github.com/siderolabs/talos/commit/2f23fedeb725a5786b6ffac2aef8125eecd6cb6e) fix: file leak in reading cgroups
* [`b412ffdbc`](https://github.com/siderolabs/talos/commit/b412ffdbc29d77a81aed88be62f21bc2999afcde) docs: update README.md for docs link
* [`8dc51bae7`](https://github.com/siderolabs/talos/commit/8dc51bae79a37b56c058d40787dbda6e828fd0d3) feat: add drm_gpuvm and drm_gpusvm_helper modules
* [`4ca58aeb8`](https://github.com/siderolabs/talos/commit/4ca58aeb81145cb7ebef071865b3d853a4712729) fix: make Akamai platform usable
* [`061f8e76f`](https://github.com/siderolabs/talos/commit/061f8e76fd58906ff823a0e467d6efcf5161ed9f) feat: bump pkgs
* [`a9fa852da`](https://github.com/siderolabs/talos/commit/a9fa852dadd75740d73588fd2156f6f1ad782fdd) feat: update uefi image to talos linux logo
* [`04753ba69`](https://github.com/siderolabs/talos/commit/04753ba6983b6ff2754cf62b8d60cc6065921dbd) feat: update go to 1.25.2
* [`9a42b05bd`](https://github.com/siderolabs/talos/commit/9a42b05bdac2bf0cbbc97d040be7860f48c69386) feat: implement link aliasing
* [`d732bd0be`](https://github.com/siderolabs/talos/commit/d732bd0be73c3d17d140c00be0e9d27ea621909b) chore(ci): run only nvidia tests for NVIDIA workflows
* [`8d1468209`](https://github.com/siderolabs/talos/commit/8d1468209aa28f59df9dc52466c506defa8c3cc3) fix: stop populating apiserver cert SANs
* [`02473244c`](https://github.com/siderolabs/talos/commit/02473244c17ef0149515f300bcd201f9347acabc) fix: wait for mount status to be proper mode
* [`825622d90`](https://github.com/siderolabs/talos/commit/825622d90a7716f7b6027651a5b9389173432393) fix: resource proto definitions
* [`2c6003e79`](https://github.com/siderolabs/talos/commit/2c6003e790003f6ef1a03b8d2af8030fb57c5d02) docs: add Project Calico installation in two mode
* [`4fb4c8678`](https://github.com/siderolabs/talos/commit/4fb4c86780def54eed4d999b1f0ce93042269076) feat: add disk.EnableUUID to generated ova
* [`33fb48f8f`](https://github.com/siderolabs/talos/commit/33fb48f8f90ccf44e95c93ac7ec1adcd1b4e0373) fix: add dashboard spinner
* [`053fd0bd4`](https://github.com/siderolabs/talos/commit/053fd0bd4d324bc21e076b3a30466ed61c7684e1) feat: update Linux to 6.17
* [`34e107e1b`](https://github.com/siderolabs/talos/commit/34e107e1bd14b0a56ebfa0c65e0c7da715976d99) docs: fix broken link
* [`dfbece56b`](https://github.com/siderolabs/talos/commit/dfbece56bd45e95c9ec477af4b53ffcefdfec66c) docs: update the kubespan docs
* [`8b041a72c`](https://github.com/siderolabs/talos/commit/8b041a72ca9c07985c024c1136c85c85df92beda) docs: update scaleway.md
* [`435dcbf82`](https://github.com/siderolabs/talos/commit/435dcbf820cd9f8cc9fecc0f7d42819acef36106) fix: provide nocloud metadata with missing network config
* [`ec3bd878f`](https://github.com/siderolabs/talos/commit/ec3bd878f9770ceb932b654aabad1711880da829) refactor: remove the go-blockdevice v1 completely
* [`33544bde9`](https://github.com/siderolabs/talos/commit/33544bde9c15745f4ae692c7647d661b32d4bed4) fix: minor improvements to fs
* [`fd2eebf7f`](https://github.com/siderolabs/talos/commit/fd2eebf7fa4831d33383a53d6d058c74789553e4) feat: create merge patch from diff of two machine configs
* [`eadbdda94`](https://github.com/siderolabs/talos/commit/eadbdda9471289fae5159c8cc024a735a1547807) fix: uefi boot order setting
* [`cd9fb2743`](https://github.com/siderolabs/talos/commit/cd9fb274342c5a973b3d087b991a7eea5df4142a) fix: support secure HTTP proxy with gRPC dial
* [`adf87b4b9`](https://github.com/siderolabs/talos/commit/adf87b4b931ded1edeb64217b0e9d5edfd046004) feat: update Flannel to v0.27.4
* [`5dfb7e1fe`](https://github.com/siderolabs/talos/commit/5dfb7e1fe7d9cc6db3e4c2b6f587e641b4a0842b) feat: serve etcd image from registry.k8s.io
* [`5ca841804`](https://github.com/siderolabs/talos/commit/5ca8418049e3b878585014a3764021f2d30a0df7) fix: nftables flaky test
* [`a940e45a7`](https://github.com/siderolabs/talos/commit/a940e45a7fe041b17437f774eb52b9f3a42e3633) feat: generate list of images required to build talos
* [`3472d6e79`](https://github.com/siderolabs/talos/commit/3472d6e79caa13fd42df7774101397b0a30f62f5) fix: revert "chore: use new mount/v3 package in efivarfs"
* [`42c0bdbf3`](https://github.com/siderolabs/talos/commit/42c0bdbf320bf24311b2d56b2e0f7155e86b3713) feat: add provisioner flag to images default command
* [`6bc0b1bcf`](https://github.com/siderolabs/talos/commit/6bc0b1bcf7d9dc9f2417a7db63d1e76e7ddc6aa3) feat: drop and lock deprecated features
* [`362a8e63b`](https://github.com/siderolabs/talos/commit/362a8e63b798c4a4fc31fe5e728d2429fc953166) fix: change the compression format
* [`6e58f58aa`](https://github.com/siderolabs/talos/commit/6e58f58aaeb6e16883d8dc8757ad92b6b6da7e84) fix: mkdir artifacts path
* [`3165a2b84`](https://github.com/siderolabs/talos/commit/3165a2b84cb80dd5fd09bf496fdccaf1628593d0) release(v1.12.0-alpha.1): prepare release
* [`e455c7ea9`](https://github.com/siderolabs/talos/commit/e455c7ea9c919a2f70ddecceaa8f3b4e25566048) chore: use testing/synctest in tests
* [`7f048e962`](https://github.com/siderolabs/talos/commit/7f048e962e217687ab67ed7027c5228e8ccb7d16) feat: update dependencies
* [`fe36b3d32`](https://github.com/siderolabs/talos/commit/fe36b3d3200db57f3e21017ff7a4808b330a1d55) fix: stop returning EINVAL on remount of detached mounts
* [`c6279e04c`](https://github.com/siderolabs/talos/commit/c6279e04c45504af243c0aef9f255317426b4ca0) chore: use new mount/v3 package in efivarfs
* [`d5197effb`](https://github.com/siderolabs/talos/commit/d5197effb0b48290d613140b68796cb8f30b9a70) feat: update etcd 3.6.5, CoreDNS 1.12.4
* [`33714b715`](https://github.com/siderolabs/talos/commit/33714b7158a0d569be1d0b1d7b012280856db484) feat: release cloud image using factory
* [`d10a2747e`](https://github.com/siderolabs/talos/commit/d10a2747e0e835876aff158e6b6f7882cef9fa44) docs: deprecate JSON6902 patches and interactive installer
* [`1e604cbf5`](https://github.com/siderolabs/talos/commit/1e604cbf514bece1e112d8afd5d1cd6ccb1045c3) fix: don't set broadcast for /31 and /32 addresses
* [`65a66097a`](https://github.com/siderolabs/talos/commit/65a66097a05e5c0e2334d5eff494a0e71534716f) refactor: split cluster create logic into smaller parts
* [`ab847310e`](https://github.com/siderolabs/talos/commit/ab847310efde540b5bfe17570b99af1bb705832b) fix: provide refreshing CA pool (resolvers)
* [`d63c3ed7d`](https://github.com/siderolabs/talos/commit/d63c3ed7db2b22f7e394fc45d101d03cba463177) docs: update secureboot docs
* [`493f7ed9d`](https://github.com/siderolabs/talos/commit/493f7ed9d2710eb240eab6b6ab532f41abc818c1) feat: support embedded config
* [`251df70f6`](https://github.com/siderolabs/talos/commit/251df70f6d33f1d5a3b1b9e4c0c249d8bc85c4b3) feat: add a userspace OOM controller
* [`7bae5b40b`](https://github.com/siderolabs/talos/commit/7bae5b40b4f22f0f07a586ebd9cda9436086a5f8) feat: implement link configuration
* [`724857dec`](https://github.com/siderolabs/talos/commit/724857decb95ddeebb2ac5d33c38a71bf7512805) fix(ci): skip netbird extension for tests
* [`e06a08698`](https://github.com/siderolabs/talos/commit/e06a086989331f28406e8d4234e02d9a6b83f87d) fix: default gateway as string
* [`7ed07412e`](https://github.com/siderolabs/talos/commit/7ed07412e963e6ee91615adbea095944aa6a56e5) fix: uefi boot entry handling logic
* [`ea4ed165a`](https://github.com/siderolabs/talos/commit/ea4ed165ad860a5beea17ca2d404bdaa6e5ad933) refactor: efivarfs mock and tests
* [`1fca111e2`](https://github.com/siderolabs/talos/commit/1fca111e24bcae81b78f007e67b71c9155c0169f) feat: support setting wake-on-lan for Ethernet
* [`94f78dbe7`](https://github.com/siderolabs/talos/commit/94f78dbe798cb227a0c38b70a1d6840803989290) docs: add a documentation for running Talos in KVM
* [`46902f8fd`](https://github.com/siderolabs/talos/commit/46902f8fdee257a09be4bc1753c6b3f845ef8089) docs: add TrueFullstaq to adopters
* [`a28e5cbd5`](https://github.com/siderolabs/talos/commit/a28e5cbd50d11aa6c253a6a9ce1999b9d45effad) chore: update pkgs and tools
* [`7cf403db8`](https://github.com/siderolabs/talos/commit/7cf403db8ca0e1719195001895cfbc12835b0fdd) docs: step-by-step scaleway documentation to get an image
* [`687285fa2`](https://github.com/siderolabs/talos/commit/687285fa26ec42dadbfb72580099f6e20bbaf85e) docs: remove 'curl' in wget command
* [`9db6dc06c`](https://github.com/siderolabs/talos/commit/9db6dc06c3010cd89ce4cb0ec0bde178db0447a4) feat: stop mounting state partition
* [`53ce93aae`](https://github.com/siderolabs/talos/commit/53ce93aaed3bd5bfcbe926fa69ca3b4b8b45c74f) test: try to clear connection refused more aggressively
* [`51db5279c`](https://github.com/siderolabs/talos/commit/51db5279c423e4b8637a05e52b26dfc5aa719cbc) fix: bump trustd memory limit
* [`25204dc8a`](https://github.com/siderolabs/talos/commit/25204dc8a8df79bc876a0bec2492e1147a81d954) fix(machined): change `constants.MinimumGOAMD64Level` using build tag
* [`9cd2d794d`](https://github.com/siderolabs/talos/commit/9cd2d794d060b637dbac5263ae417a4e83d54efe) feat: ship nft binary with Talos rootfs
* [`b1416c9fe`](https://github.com/siderolabs/talos/commit/b1416c9fe1d5ea9cd68f9b6b766a288a267cee61) feat: record last log the failed service
* [`0b129f9ef`](https://github.com/siderolabs/talos/commit/0b129f9efdf57dd9692f7cece6b97719a7ccf80e) feat: enforce more KSPP and hardening sysctls
* [`11872643c`](https://github.com/siderolabs/talos/commit/11872643c310212c52b4fd7e13b6cc7d6ec7e4fc) chore: drop docs folder
* [`d30fdcd88`](https://github.com/siderolabs/talos/commit/d30fdcd88f421824cf17b9ecec25be7c8044e857) chore: pass in github token to imager
* [`b88f27d80`](https://github.com/siderolabs/talos/commit/b88f27d804d60a706f598b50676dad5dd2a9726a) chore: make reset test code a bit better
* [`1cde53d01`](https://github.com/siderolabs/talos/commit/1cde53d0173fd1ae637855e15fe34bb74bb027a0) test: fix several issues with tests
* [`16cd127a0`](https://github.com/siderolabs/talos/commit/16cd127a04bb5fc907b7ca04f1c81d4c7150eab2) docs: add docs on updating image cache
* [`c3ae92b14`](https://github.com/siderolabs/talos/commit/c3ae92b1424d4a2c9bc18cfa394b10eda6c9a20f) fix: build kernel checks only on linux
* [`2120904ec`](https://github.com/siderolabs/talos/commit/2120904ec534a91f66dcea419b5a29e36a16f6e4) feat: create detached tmpfs
* [`6bbee6de5`](https://github.com/siderolabs/talos/commit/6bbee6de5b18b25deb4e6f515251187e259aa424) docs: remove 'ceph-data' from volume examples/docs
* [`07acb3bd2`](https://github.com/siderolabs/talos/commit/07acb3bd2d4f92e80706d1835130bbe6e944d096) fix: use correct order to determine SideroV1 keys directory path
* [`2d57fa002`](https://github.com/siderolabs/talos/commit/2d57fa00281f8090b85097c66df634101b0cde79) fix: trim zero bytes in the DHCP host & domain response
* [`451cb5f78`](https://github.com/siderolabs/talos/commit/451cb5f78fac3b2ddfec7d545629fe8c88ea2367) docs: clarify disk partition confusion
* [`a2122ee5c`](https://github.com/siderolabs/talos/commit/a2122ee5cb9c84f33e0c4b30e9223bb239621d55) feat: implement HostConfig multi-doc
* [`69ab076b4`](https://github.com/siderolabs/talos/commit/69ab076b4d6e52484677ee7f68a853dc4edfe2bc) fix: re-create cgroups when restarting runners
* [`297b5cc28`](https://github.com/siderolabs/talos/commit/297b5cc2856710b74b4e0e46b00ae33aea4c1bf7) docs: add docs on node labels
* [`e168512dd`](https://github.com/siderolabs/talos/commit/e168512dd020da9eac654dae2ba891cf33415c44) fix: apply 'ro' flag to iso9660 filesystems
* [`7f7acfbb9`](https://github.com/siderolabs/talos/commit/7f7acfbb9f10c243d0b132c1ef079cb77d2727e0) docs: fix typo in doc
* [`d57882b18`](https://github.com/siderolabs/talos/commit/d57882b1830504fe4bfd5344edae613168db7f0e) feat: update Kubernetes to 1.34.1
* [`f85f82f32`](https://github.com/siderolabs/talos/commit/f85f82f32f098f97588f404550f72d64786fe329) test: fix flakiness in RawVolumes test
* [`82569e319`](https://github.com/siderolabs/talos/commit/82569e319eb57b1199db6bfd3e612fb771c8c7cd) feat: update Linux 6.16.6
* [`2fd2ab4e4`](https://github.com/siderolabs/talos/commit/2fd2ab4e43e06910154705d6ef1d0576a7c04a2b) fix: remove CoreDNS cpu limit
* [`ce9bc32a0`](https://github.com/siderolabs/talos/commit/ce9bc32a08695873d9054afe2608a76cf7c6088a) chore(ci): rekres to use new runner groups
* [`8b64f68f6`](https://github.com/siderolabs/talos/commit/8b64f68f6946c2979f6fe2bf617f31639a927bf8) test: improve test stability
* [`272cb860d`](https://github.com/siderolabs/talos/commit/272cb860d4cfb8464b29ff31567e25fe6c275849) chore: drop the --input-dir flag from the cluster create command
* [`1b6533675`](https://github.com/siderolabs/talos/commit/1b65336752933acdcbf681767785157714866f88) docs: add note about ca-signed certs for secureboot
* [`d3f88f50c`](https://github.com/siderolabs/talos/commit/d3f88f50c5394536ee80d19464359408a37d81ff) docs: document talos vip failover behavior
* [`005fc8bd5`](https://github.com/siderolabs/talos/commit/005fc8bd50fbc4b15b26032b43d1d32c1da22f11) docs: add docs on syncing configs after a kube upgrade
* [`4d876d9af`](https://github.com/siderolabs/talos/commit/4d876d9af9fcc9828f09d05db124fbdce9c17785) feat: update Go to 1.25.1
* [`2b556cd22`](https://github.com/siderolabs/talos/commit/2b556cd22a3563f1d86a648ea6c69a4d45edad76) feat: implement multi-doc StaticHostConfig
* [`a7b776842`](https://github.com/siderolabs/talos/commit/a7b7768420566b6840fc52bb2152e9bf165f8cd3) docs: replace Raspberry Pi 5 links with Talos builder
* [`a349b20ed`](https://github.com/siderolabs/talos/commit/a349b20ed4b3c05dcd0175541b795331f0f7c64d) docs: clarify that talos does not support intermediate ca
* [`895133de9`](https://github.com/siderolabs/talos/commit/895133de99158ce3f50b557b77c81d4f0f9d6b40) feat: support configuring PCR states to bind disk encryption
* [`c1360103b`](https://github.com/siderolabs/talos/commit/c1360103b5e037cf713b7d787436f01e7182821c) docs: fix command for uploading image on Hetzner
* [`43b5b9d89`](https://github.com/siderolabs/talos/commit/43b5b9d8992ad6df37619b3719b57948e4bd9671) fix: correctly handle status-code 204
* [`feeb0d312`](https://github.com/siderolabs/talos/commit/feeb0d312ecacb451e5313390939c7c9349d2ba6) feat: update runc to 1.3.1
* [`421634a14`](https://github.com/siderolabs/talos/commit/421634a1417f529551a75d0bb9be08b73f1120b1) docs: add docs on multihoming
* [`41af2d230`](https://github.com/siderolabs/talos/commit/41af2d230c2dd5dce5bc931f76a2eb69405dc554) refactor: clean up internal cluster creation code
* [`3000d9e43`](https://github.com/siderolabs/talos/commit/3000d9e431deaf952d08da724da40789cd743f2c) fix: don't bootstrap talos cluster if there's no config present
* [`79cb871d0`](https://github.com/siderolabs/talos/commit/79cb871d088e5b1c3a3488610ded14e7a28cec29) feat: use the id of the volume in the mapped luks2 name
* [`6c322710d`](https://github.com/siderolabs/talos/commit/6c322710d64786f19e2e0e39d65596c8dce71952) chore: refactor mount package
* [`ced7186e2`](https://github.com/siderolabs/talos/commit/ced7186e2a5f0634d9441b12a5340f5ca4c451ff) refactor: update COSI to 1.11.0
* [`de2e24fcd`](https://github.com/siderolabs/talos/commit/de2e24fcda590a1ef3f80a5372bb70865a2f47c3) docs: clarify that install-cni image is deprecated
* [`bef8ef509`](https://github.com/siderolabs/talos/commit/bef8ef509380aba259efcc2f5d1f6632e034160b) docs: add docs on cilium's compatibility with kubespan
* [`e5acb10fc`](https://github.com/siderolabs/talos/commit/e5acb10fcceba69060507a35caea21281bdc71cc) feat: update pkgs
* [`c4c1daf0e`](https://github.com/siderolabs/talos/commit/c4c1daf0e2e6675626b974b0c008e101d919c8b5) docs: add info about br_netfilter
* [`5c52ecac3`](https://github.com/siderolabs/talos/commit/5c52ecac364f917e5f45859f680494a08f85cb90) docs: clarify interactive dashboard resolution control
* [`15ecb02a4`](https://github.com/siderolabs/talos/commit/15ecb02a4545639ffb8ba5c6e5a413e53129b619) feat: update Linux kernel (memcg_v1, ublk)
* [`53f18c2f6`](https://github.com/siderolabs/talos/commit/53f18c2f60c84c4b0f944cc343ae1f538e8d1236) fix: enable support for VMWare arm64
* [`3bbe1c0da`](https://github.com/siderolabs/talos/commit/3bbe1c0da5485b6cd3e7fadd8f020e0d0aca406a) docs: add docs on grow flag
* [`b9fb09dcd`](https://github.com/siderolabs/talos/commit/b9fb09dcdbcca60f695ac317c45e18fa092541a8) release(v1.12.0-alpha.0): prepare release
* [`6a389cad3`](https://github.com/siderolabs/talos/commit/6a389cad35f80b27fe9c43db9e701ee9f6f6142a) chore: update dependencies
* [`9d98c2e89`](https://github.com/siderolabs/talos/commit/9d98c2e891258dcf2ef90519d38d0aefb77cd0db) feat: add a cgroup preset for PSI and --skip-cri-resolve
* [`072f77b16`](https://github.com/siderolabs/talos/commit/072f77b1623cdc838093465b7266b26e20a248ea) chore: prepare for future Talos 1.12-alpha.0 release
* [`96f41ce88`](https://github.com/siderolabs/talos/commit/96f41ce8840783f783fcc8e0fd6b43302b9bfe43) docs: update qemu and docker docs
* [`a751cd6b7`](https://github.com/siderolabs/talos/commit/a751cd6b7474a4dc20137e917dbb2229fe9cc8bd) docs: activate Talos v1.11 docs by default
* [`e8f1ec1c5`](https://github.com/siderolabs/talos/commit/e8f1ec1c5bbd8a6cfb68886e6283e7caaf5fb063) docs: fix broken create qemu command v1.11 docs
* [`639f0dfdd`](https://github.com/siderolabs/talos/commit/639f0dfdd88c5596439601f3f9600b3aafb24227) feat: update Linux to 6.16.4
* [`8aa7b3933`](https://github.com/siderolabs/talos/commit/8aa7b3933d07ea45a96844b9c91347a08950e243) fix: bring back linux/armv7 build and update xz
* [`9cae7ba6b`](https://github.com/siderolabs/talos/commit/9cae7ba6b97a67a5d282c6f667ccb4c3e2111447) feat: update CoreDNS to 1.12.3
* [`cfef3ad45`](https://github.com/siderolabs/talos/commit/cfef3ad4544498a47de17f6b05fb8374c35e3dd8) fix: drop linux/armv7 build
* [`42ea2ac50`](https://github.com/siderolabs/talos/commit/42ea2ac5058457dafe666f8d79f08d3c8ee60cfb) fix: update xz module (security)
* [`4fcfd35b9`](https://github.com/siderolabs/talos/commit/4fcfd35b9510f45d0ef7ae3657eb0916d549d2dd) docs: fix module name example
* [`50824599a`](https://github.com/siderolabs/talos/commit/50824599a4fa7b72d563a35a4746ca063becf672) chore: update some tools
* [`bcd297490`](https://github.com/siderolabs/talos/commit/bcd297490c608f593b6dd274945aa2b73c3fd3ee) feat: allow Ed25119 in FIPS mode
* [`5992138bb`](https://github.com/siderolabs/talos/commit/5992138bb981e84dae917f0f0fdafee4049bc5ec) test: ignore one leaking goroutine
* [`d155326c1`](https://github.com/siderolabs/talos/commit/d155326c1206979f30a5355f7bdb23cb051e9b78) docs: add sbc unofficial ports docs
* [`285fa7d22`](https://github.com/siderolabs/talos/commit/285fa7d222be1f5e63c0bb725b206966e2722a3b) docs: add the deploy application docs
* [`527791f09`](https://github.com/siderolabs/talos/commit/527791f0974afe9c8558b82fa19f4354487693ed) feat: update Kubernetes to 1.34.0
* [`a1c0e237d`](https://github.com/siderolabs/talos/commit/a1c0e237d6e047bb59c4fbd48e2c2b9e36dd4808) feat: update Linux to 6.15.11, Go to 1.25
* [`4d7fc25f8`](https://github.com/siderolabs/talos/commit/4d7fc25f8bf20d4489080795a3d0ce0dfb1bc6b8) docs: switch order of wipe disk command
* [`7368a994d`](https://github.com/siderolabs/talos/commit/7368a994df07cc4e50e3709ac766d8062db070a0) feat: add SOCKS5 proxy support to dynamic proxy dialer
* [`d63591069`](https://github.com/siderolabs/talos/commit/d635910697b221aee3e9afa6d9e5b398236b6a21) chore: silence linter warnings
* [`07eb4d7ec`](https://github.com/siderolabs/talos/commit/07eb4d7ec148a7e3c4c6dde080469c1a2fb410fb) fix: set default ram unit to MiB instead of MB
* [`6b732adc4`](https://github.com/siderolabs/talos/commit/6b732adc43684facfd329f424a34a7e4df36d77b) feat: update Linux to 6.12.43
* [`b6410914f`](https://github.com/siderolabs/talos/commit/b6410914f74ce01672fdef7e912e37970909281c) feat: add human readable byte size cli flags
* [`ec70cef99`](https://github.com/siderolabs/talos/commit/ec70cef99005fd7e383fea63b5c23774882fcf28) feat: update NVIDIA drivers and kernel
* [`0879efa69`](https://github.com/siderolabs/talos/commit/0879efa690ad657e4aed251fcbeba8f5645d73ce) feat: update Kubernetes default to v1.34.0-rc.2
* [`f504639df`](https://github.com/siderolabs/talos/commit/f504639df4388619f731196ed8e79a6818b6ed5f) feat: add a user-facing create qemu command
* [`558e0b09a`](https://github.com/siderolabs/talos/commit/558e0b09ab65b353e83b98c9ddf6cb2b67fd060e) test: fix the Image Factory PXE boot test
* [`d73f0a2e5`](https://github.com/siderolabs/talos/commit/d73f0a2e5b788c7b69c2fb827f7111d5f9c8e706) docs: make readme badges consistent
* [`f1369af98`](https://github.com/siderolabs/talos/commit/f1369af98e1f6d48fed137e31237956abbd28b0f) chore: use new filesystem api on STATE partition
* [`366cedbe7`](https://github.com/siderolabs/talos/commit/366cedbe7495ce15bcd0e6c6f7f0add65a41a861) docs: link to kubernetes linux swap tuning
* [`2f5a16f5e`](https://github.com/siderolabs/talos/commit/2f5a16f5e4ae186a309aef5e3d285897d0fe2df1) fix: make --with-uuid-hostnames functionality available to qemu provider
* [`70612c1f9`](https://github.com/siderolabs/talos/commit/70612c1f9fc9056e8a3669ff10a385c4e8e03350) refactor: split the PlatformConfigController
* [`511748339`](https://github.com/siderolabs/talos/commit/51174833997fd9a0a599ab1dde947834b682ab14) docs: add system extension tier documentation
* [`009fb1540`](https://github.com/siderolabs/talos/commit/009fb1540e0b9f5daac6302f42e8813e596fc87c) test: don't run nvidia tests on integration/aws
* [`99674ef20`](https://github.com/siderolabs/talos/commit/99674ef20d34166d60563d4bf46fbbfc57399509) docs: apply fixes for what is new
* [`92db677b5`](https://github.com/siderolabs/talos/commit/92db677b5d32de32ec7e785531b32202e03283b4) fix: image cache lockup on a missing volume
* [`9c97ed886`](https://github.com/siderolabs/talos/commit/9c97ed886b89b2fb84f47866abdf1000839143c4) fix: version contract parsing in encryption keys handling
* [`1fc670a08`](https://github.com/siderolabs/talos/commit/1fc670a08dc7af8eaeabdc7134eb77a5c939df40) fix: dial with proxy
* [`18447d0af`](https://github.com/siderolabs/talos/commit/18447d0afdbcc8fa7db6ae008e4bc4d5b0a0b00a) feat: update Linux to 6.12.41
* [`f65f39b78`](https://github.com/siderolabs/talos/commit/f65f39b78b0c7881e5f51c66ad022c17c2cd4960) fix: provide mitigation CVE-1999-0524
* [`8817cc60c`](https://github.com/siderolabs/talos/commit/8817cc60cfaf4b50f11c38d3b25df7df48382033) fix: actually use SIDEROV1_KEYS_DIR env var if it's provided
* [`b08b20a10`](https://github.com/siderolabs/talos/commit/b08b20a1005256a9e3fc7cae8bcf8eea87f6ac09) feat: use key provider with fallback option for auth type SideroV1
* [`7a52d7489`](https://github.com/siderolabs/talos/commit/7a52d7489c9709708d55f8f001d70700addc7e1e) fix: kubernetes upgrade options for kubelet
* [`ea8289f55`](https://github.com/siderolabs/talos/commit/ea8289f550787593b1cd35f2d8da59aa5311880e) feat: add a user facing docker command
* [`54ad64765`](https://github.com/siderolabs/talos/commit/54ad64765090d90013e4917d1bf494592069beec) chore: re-enable vulncheck
* [`26bbddea9`](https://github.com/siderolabs/talos/commit/26bbddea95669278363c604316ed85986f312d71) fix: darwin build
* [`b5d5ef79e`](https://github.com/siderolabs/talos/commit/b5d5ef79e7a2d76e29a7c872c1c418fffc63b0df) fix: set secs field in DHCPv4 packets
* [`c07911933`](https://github.com/siderolabs/talos/commit/c0791193373e36c35f29c70318432331b4c6ab2a) chore: refactor how tools are being installed
* [`34f25815c`](https://github.com/siderolabs/talos/commit/34f25815c036d2c91bdfddc9c7d40ca2edf677bd) docs: fork docs for v1.12
* [`b66b995d3`](https://github.com/siderolabs/talos/commit/b66b995d34306192cbaa4ef68fe39f821b37d1f0) feat: update default Kubernetes to v1.34.0-rc.1
* [`b967c587d`](https://github.com/siderolabs/talos/commit/b967c587d9f217f25798e0bee0c90393e55dc085) docs: fix clone URL to include `.git`
* [`b72c68398`](https://github.com/siderolabs/talos/commit/b72c6839806103ac0a76acd46f30eabea0375790) docs: edit the insecure, etcd-metrics, inline and extramanifests
* [`e5b9c1fff`](https://github.com/siderolabs/talos/commit/e5b9c1ffffec9fd49ffb84a36c918e75eaa8f1ef) docs: remov RAS Syndrome
* [`701fe774b`](https://github.com/siderolabs/talos/commit/701fe774bd19de7c9f21e043e1520161a8c5fff7) docs: fix cilium links and bump to 1.18.0
* [`d306713a1`](https://github.com/siderolabs/talos/commit/d306713a13a18d7af6caffd5890d54d91d22cad7) feat: update Go to 1.24.6
* [`721595a00`](https://github.com/siderolabs/talos/commit/721595a0009f78a2722802ab665957fd767c4d1e) chore: add deadcode elimination linter
* [`dc4865915`](https://github.com/siderolabs/talos/commit/dc4865915d567942adea3efa66f8ad360f9c4cce) refactor: stop using `text/template` in `machined` code paths
* [`545be55ed`](https://github.com/siderolabs/talos/commit/545be55edc863245638d4387cb9ee7e7b068f2ba) feat: add a pause function to dashboard
* [`06a6c0fe3`](https://github.com/siderolabs/talos/commit/06a6c0fe332940b7a70ea2652bc2a5e7bc51bbf3) refactor: fix deadcode elimination with godbus
* [`2dce8f8d4`](https://github.com/siderolabs/talos/commit/2dce8f8d4693a85d2f3bf46169af8cf502d49f9d) refactor: replace containerd/containerd/v2 module for proper DCE
* [`9b11d8608`](https://github.com/siderolabs/talos/commit/9b11d86081df8cf77860d2d27eed5d8001ff721e) chore: rekres to configure slack notify workflow for CI failures
* [`5ce6a660f`](https://github.com/siderolabs/talos/commit/5ce6a660f67f4e2776550a1e621179beb8a6788c) docs: augment the pod security docs
* [`ada51ff69`](https://github.com/siderolabs/talos/commit/ada51ff696011e15dcd9c661da1d839bdc341745) fix: unmarshal encryption STATE from META
* [`b9e9b2e07`](https://github.com/siderolabs/talos/commit/b9e9b2e07a645f53ca23355810d485a2622870c9) docs: add what is new notes for 1.11
* [`53055bdf4`](https://github.com/siderolabs/talos/commit/53055bdf49ce4c81f63c159cdbaa8ea85d9ca2b8) docs: fix typo in kubevirt page
* [`8d12db480`](https://github.com/siderolabs/talos/commit/8d12db480c38ec37aee5ae7721b2e5ca55ad733e) fix: one more attempt to fix volume mount race on restart
* [`34d37a268`](https://github.com/siderolabs/talos/commit/34d37a268a9e0098179369af128261dbfc956d1d) chore: rekres to use correct slack channel for slack-notify
* [`326a00538`](https://github.com/siderolabs/talos/commit/326a00538210bf98b01795d314c1e154a74d2d58) feat: implement `talos.config.early` command line arg
* [`a5f3000f2`](https://github.com/siderolabs/talos/commit/a5f3000f2e8a79d4e9a5be95fbcac91a2d78675b) feat: implement encryption locking to STATE
* [`c1e65a342`](https://github.com/siderolabs/talos/commit/c1e65a34256944743e768613b119c0caa517b54d) docs: remove talos API flags from mgmt commands
* [`181d0bbf5`](https://github.com/siderolabs/talos/commit/181d0bbf5381343d35a01190da45e3442320d7c5) feat: bootedentry resource
* [`7ad439ac3`](https://github.com/siderolabs/talos/commit/7ad439ac35859695074d3a3efdcdb5c0cab1a5c6) fix: enforce minimum size on user volumes if not set explicitly
* [`50e37aefd`](https://github.com/siderolabs/talos/commit/50e37aefdbde973bcc8aa352639946490fbe7d94) fix: live reload of TLS client config for discovery client
* [`87efd75ef`](https://github.com/siderolabs/talos/commit/87efd75efb3e62b88b4f65a221f9fbdd4b4d6ef9) feat: update containerd to 2.1.4
* [`724b9de6d`](https://github.com/siderolabs/talos/commit/724b9de6d5195bcccc5f484c696429b2f09ab16e) feat: add F71808E watchdog driver
* [`8af96f7af`](https://github.com/siderolabs/talos/commit/8af96f7afdac1c4d5e2697b897b81e2bddd15f66) docs: add ETCD downgrade documentation
* [`44edd205d`](https://github.com/siderolabs/talos/commit/44edd205d5fdffab39b65ee62695a40e22ef188c) docs: add remark about 'exclude-from-external-load-balancers' label
* [`727101926`](https://github.com/siderolabs/talos/commit/7271019263b0dc5b28d2764d19fe531e473222fc) fix(ci): use a random suffix for ami names
* [`d621ce372`](https://github.com/siderolabs/talos/commit/d621ce3726f20ee568ea3b6ac57d9e8dfa0580cc) fix: grype scan
* [`d62e255c2`](https://github.com/siderolabs/talos/commit/d62e255c260810a5f0f2959e32592a3331df28d3) fix: issues with reading GPT
* [`5d0883e14`](https://github.com/siderolabs/talos/commit/5d0883e147163c12a77cd926db799ffed854aedf) feat: update PCI DB module to v0.3.2
* [`3751c8ccf`](https://github.com/siderolabs/talos/commit/3751c8ccfa1bab9fcd435290f36e9012a5626e40) test: wait for service account test job longer
* [`a592eb9f9`](https://github.com/siderolabs/talos/commit/a592eb9f98788883a7ec6d17772e10707230a0d8) feat: update Linux to 6.12.40
* [`4c40e6d3f`](https://github.com/siderolabs/talos/commit/4c40e6d3fb4c2f451a8d7a671df5f6254161bd5d) feat: update etcd to 3.6.4
* [`2bc37bd2c`](https://github.com/siderolabs/talos/commit/2bc37bd2c9679c8055fd7b52eb310f23a329af4e) docs: fix error in kernel module guide
* [`bfc57fb86`](https://github.com/siderolabs/talos/commit/bfc57fb863224f7626f49e5b26be06f77bea2e40) chore: tag aws snapshots created via ci with the image name
* [`06ef7108a`](https://github.com/siderolabs/talos/commit/06ef7108a6050b3a8fd7535f01a469f09042bf56) fix: issue with volume remount on service restart
* [`03efbff18`](https://github.com/siderolabs/talos/commit/03efbff18e420c4fe960f490f91dd9f4751ece04) docs: add SBOM documentation
* [`af8a2869d`](https://github.com/siderolabs/talos/commit/af8a2869dbbec073ffaf72a1378682e109b053ec) fix: do not download artifacts for cron Grype scan
* [`5f442159b`](https://github.com/siderolabs/talos/commit/5f442159b224c96c90badc7176fed17bfb561709) feat: unify disk encryption configuration
* [`38e176e59`](https://github.com/siderolabs/talos/commit/38e176e594edb3d271d98f78417b9fd5ba0c5288) chore(ci): fix datasource versioning
* [`85d6b9198`](https://github.com/siderolabs/talos/commit/85d6b919890a1aa9c4f94d5b18861cc617134ff9) feat: update etcd to v3.5.22
* [`dd7bd2dab`](https://github.com/siderolabs/talos/commit/dd7bd2dab8cf09334e3e353d6a477509bbaa303e) docs: rewrite the getting started and prod docs for v1.10 and v1.11
* [`136a899aa`](https://github.com/siderolabs/talos/commit/136a899aa25b3fdcdd771594668278d563f09192) chore: regenerate release step with signing fixes
* [`450b30d5a`](https://github.com/siderolabs/talos/commit/450b30d5a986563869efdbaa074e82d612f6f2ef) chore(ci): add more nvidia test matrix
* [`451c2c4c3`](https://github.com/siderolabs/talos/commit/451c2c4c39e70c20df58fc31459cd5c789a0e46f) test: add talosctl:latest to the image cache
* [`3039162dc`](https://github.com/siderolabs/talos/commit/3039162dc44d7176c9fccac8203c5463286455a2) feat: update Flannel to v0.27.2
* [`7e6052e63`](https://github.com/siderolabs/talos/commit/7e6052e63acb7e6a446b8dac2da831b48676fa8b) feat: increase boot partition to 2 GiB
* [`cb7ca17bb`](https://github.com/siderolabs/talos/commit/cb7ca17bba209c93c9683b37b28c2da7552e1148) feat: implement ExistingVolumeConfig
* [`a857c696f`](https://github.com/siderolabs/talos/commit/a857c696faf6432b3e9d7ef338ae2e1cfc637301) chore(machined): remove deprecated Endpoints
* [`a60101c55`](https://github.com/siderolabs/talos/commit/a60101c5515b767b58e26ee2f9dc5f0210fa5cb8) fix: fill serial using helpers
* [`5420e9979`](https://github.com/siderolabs/talos/commit/5420e9979b47b65d0abb3284240a53990440a521) refactor: output default selection for profiles
* [`023a24cd4`](https://github.com/siderolabs/talos/commit/023a24cd4de87ffdb8dc5b18d63ee0f3b7487525) test: use Grype to scan SBOM for vulnerabilities
* [`96896fddb`](https://github.com/siderolabs/talos/commit/96896fddb926339597438af931a9235796271f4d) chore: build less images by default
* [`75b5dec06`](https://github.com/siderolabs/talos/commit/75b5dec061bb458fce2799f99f613a98318c055c) fix: sd-boot kexec with disk images
* [`10546d6f8`](https://github.com/siderolabs/talos/commit/10546d6f8f206f1990e474711c88fc56cb0d1108) feat: update Kuberentes 1.34.0-beta.0
* [`3f35b83ae`](https://github.com/siderolabs/talos/commit/3f35b83ae35a85bb4d31eb93e6848de1b9d92aeb) fix: ignore absent extensions SBOM directory
* [`9920da3e1`](https://github.com/siderolabs/talos/commit/9920da3e1ad34599d4ee50794ef4bfb8d5f5b1a6) feat: add etcd downgrade API
* [`c38682279`](https://github.com/siderolabs/talos/commit/c38682279bfa77e827bf5975e3a878d204bed733) feat: bump pkgs and tools, read extensions' SBOMs, rekres
* [`9c0d2706c`](https://github.com/siderolabs/talos/commit/9c0d2706ce1abdcfb038f717770891a3bce43cd2) docs: add release notes about v3.6.x bug
* [`d21994210`](https://github.com/siderolabs/talos/commit/d2199421009cb52b90d101715191dbbadf6ebd5f) test: refactor various merge controller tests
* [`da5a4449f`](https://github.com/siderolabs/talos/commit/da5a4449f1a9062c234af87b33d507489b946d45) feat: implement raw volume support
* [`41adda1cf`](https://github.com/siderolabs/talos/commit/41adda1cffef3423122876b85de5d1f12f1ff38d) docs: add secure boot setup mode note for Xen
* [`993b4ade8`](https://github.com/siderolabs/talos/commit/993b4ade8407f6d0f3585d5cc40a4495b24c21e3) docs: fix typo in hugo config: pre-releaase
* [`130b7fd6e`](https://github.com/siderolabs/talos/commit/130b7fd6e6e253acd0f37a87b0f650db7acd5314) test: fix flaky TestDNS
* [`35b45ae6e`](https://github.com/siderolabs/talos/commit/35b45ae6e7066215c69cad8f264c495950177fcc) feat(talosctl): support tpm operation on mac
* [`24628db20`](https://github.com/siderolabs/talos/commit/24628db207dfa921ee8724497aea94c069f43136) feat: update Kubernetes to v1.34.0-alpha.3
* [`ff68286d1`](https://github.com/siderolabs/talos/commit/ff68286d1739d92a7f6bdf4663f1eed154931643) feat: include hwrandom modules
* [`a5b07c9a5`](https://github.com/siderolabs/talos/commit/a5b07c9a50d6c48a8a7030def898976af78d2590) test: split tests and lint from the default pipeline
* [`a957ef416`](https://github.com/siderolabs/talos/commit/a957ef4163db9a87d9c8f8f5187ba827bbf15539) feat: add SBOMs to the imager container
* [`506212a71`](https://github.com/siderolabs/talos/commit/506212a712809a3f4e217d57258e969e8a584e5a) feat: include AMD encrypted mem modules into base
* [`a966321cc`](https://github.com/siderolabs/talos/commit/a966321cce1aa78faaa9b6e97c59eaaea6a889bd) fix: add more bootloader probe logs on upgrade
* [`b38fa568a`](https://github.com/siderolabs/talos/commit/b38fa568ac48deb8e72062600dd3fa62404b34de) feat: add validation for secrets bundle
* [`2d89bcc71`](https://github.com/siderolabs/talos/commit/2d89bcc71fe0e4c395455b13a701a5512bf11b41) feat: bump Linux, Go and other packages
* [`0b8c180b8`](https://github.com/siderolabs/talos/commit/0b8c180b82aa8bed314c9e8e49e8cfe453c5b776) fix: rename instances to referenceCount
* [`378fe4f2f`](https://github.com/siderolabs/talos/commit/378fe4f2f41f525006421d659d8c3c2717187b2d) feat: support writing EFI boot order
* [`9f0792632`](https://github.com/siderolabs/talos/commit/9f07926325047e6de53c9e5d73dbac52b7a56a2b) fix: improve volume provisioning errors
* [`b8fcf3c71`](https://github.com/siderolabs/talos/commit/b8fcf3c71230207c863c639c47d394478a651112) fix: change module instance evaluation
* [`d680e560d`](https://github.com/siderolabs/talos/commit/d680e560dafdc5ea12bb64790cfaa9cd7f55309d) docs: create FUNDING.yml
* [`641505584`](https://github.com/siderolabs/talos/commit/6415055847a17235fff0dfd3de72269166869f77) feat: support project quota support for user volumes
* [`52656cc3c`](https://github.com/siderolabs/talos/commit/52656cc3c11e2d326aa19cb5d70a5b2ac0ebb217) feat: allow taloscl disk wipe in maintenance mode
* [`850579448`](https://github.com/siderolabs/talos/commit/850579448e9df7e92d427d541800464026c414b9) feat: export SBOM as resources
* [`4f3a2ffab`](https://github.com/siderolabs/talos/commit/4f3a2ffabe5b3d9215aed50beccb3f98a866c87f) test: update unit-test runner
* [`d531b682c`](https://github.com/siderolabs/talos/commit/d531b682cb4c0ccc554f7c4a0dab9d1531d40b46) fix: provide FIPS 140-3 compliance
* [`3e3129d36`](https://github.com/siderolabs/talos/commit/3e3129d3639847ee14cac76544767bc135379d4b) feat: include packages into SBOM
* [`54bd50be3`](https://github.com/siderolabs/talos/commit/54bd50be3677df67281d2dacf79102b4db491c0d) fix: talos endpoint might not be created in Kubernetes
* [`8789a02c3`](https://github.com/siderolabs/talos/commit/8789a02c3ecb1883984941d6db0a5b5596f69535) feat: present loaded kernel modules
* [`33ecbaec6`](https://github.com/siderolabs/talos/commit/33ecbaec6d724e57c4365dc847b27f2ea3e6a45a) test: update apply config tests
* [`7d2fd390c`](https://github.com/siderolabs/talos/commit/7d2fd390c55ea2fa7f239eaae313f34121563e06) chore: bump Talos version in the Image Factory CI pipeline
* [`de77f2142`](https://github.com/siderolabs/talos/commit/de77f2142acbe781fb708fa855b2281536196fee) docs: add example for fluentbit config
* [`1f1f78106`](https://github.com/siderolabs/talos/commit/1f1f7810605f15019c06f5a6ead2d9616c356e98) fix: add limited retries for not found images
* [`3d6a2c14e`](https://github.com/siderolabs/talos/commit/3d6a2c14e484040fc6808b5df87070633bcc5919) chore: generate and upload signatures on release
* [`380141330`](https://github.com/siderolabs/talos/commit/38014133090d9e2bffdb1dff05e8f8bad3d2e104) feat: expose kernel cmdline as a resource
* [`4c6b3b14d`](https://github.com/siderolabs/talos/commit/4c6b3b14d9f4221ad2ee5359fcb9112997787313) docs: document disabling SELinux
* [`3a6e5a71e`](https://github.com/siderolabs/talos/commit/3a6e5a71ed8c007965a554a5d8b104ea79fde476) feat: add talosctl mulitarch bundle image
* [`be671ee6d`](https://github.com/siderolabs/talos/commit/be671ee6d1636712f1d0f25aea1e9b8b6963477e) chore: add sbom step to the release pipeline
* [`7fd0e8fc7`](https://github.com/siderolabs/talos/commit/7fd0e8fc78d86888033400c4db27a8a740311289) release(v1.11.0-alpha.3): prepare release
* [`777335f23`](https://github.com/siderolabs/talos/commit/777335f2342abf1c04a738456678980fcc375e1b) chore: improve cloud image uploader resilience
* [`14e5eee7d`](https://github.com/siderolabs/talos/commit/14e5eee7d14bdb95e7e632c54705d8753627ab2a) release(v1.11.0-alpha.2): prepare release
* [`1e5a008f5`](https://github.com/siderolabs/talos/commit/1e5a008f5740af9dd9297ec5616bde9fd102f21f) fix: hold user volume mount point across kubelet restarts
* [`cdad50590`](https://github.com/siderolabs/talos/commit/cdad50590d4436eb12b959f2ff04457d5632f941) docs: user volumes and kubernetes upgrade updates
* [`c880835c8`](https://github.com/siderolabs/talos/commit/c880835c809c2a02f0bb6d0450d15df042a50781) feat: implement zswap support
* [`7f0300f10`](https://github.com/siderolabs/talos/commit/7f0300f108e7f2e9192214f87a13c8ff2ea25866) feat: update dependencies, Kubernetes 1.34.0-alpha.2
* [`61afbe3d2`](https://github.com/siderolabs/talos/commit/61afbe3d216862a9b9a5c8f521475a0f39cd710e) docs: add vc4 documentation
* [`b9dbdc8e7`](https://github.com/siderolabs/talos/commit/b9dbdc8e7213c305e4de71516b990641e0fed706) fix: etcd recover with multiple advertised addresses
* [`19d94c357`](https://github.com/siderolabs/talos/commit/19d94c3574b7b3ee3fbe21fdb56cff5a18e7b91e) feat: update Linux to 6.12.35, containerd to 2.1.3
* [`44a1fc3b7`](https://github.com/siderolabs/talos/commit/44a1fc3b78589540f5a0d9b8ea4d898474da3a80) fix: treat context canceled as expected error on image pull
* [`4da2dd537`](https://github.com/siderolabs/talos/commit/4da2dd537d5dae884f47bd3f04ddcd05ac6cd222) feat: enforce Kubernetes version compatibility
* [`6c7f8201a`](https://github.com/siderolabs/talos/commit/6c7f8201a9ceeec6ecfd0a35b308805ec149f3de) fix: set default MTU on Azure to 1400
* [`091cd6989`](https://github.com/siderolabs/talos/commit/091cd6989ce8c09885b3ae3e8c594c4770bd0748) docs: small yaml typo fix
* [`66ecbd48f`](https://github.com/siderolabs/talos/commit/66ecbd48fdaf509bbb2b37327eb0e0891dd81910) docs: update support matrix with omni version
* [`c948d7617`](https://github.com/siderolabs/talos/commit/c948d7617d1579c462a809b37956fc98270fcce4) docs: minor fixes for creating kernel modules
* [`cc14c4a25`](https://github.com/siderolabs/talos/commit/cc14c4a25d355910a00e60c69ed641abbb7b40f6) docs: add docs for creating kernel modules
* [`93bcd3b56`](https://github.com/siderolabs/talos/commit/93bcd3b5623d900a0f731c0f60d3ce0d69c9c32c) docs: create SBOM for Go dependencies
* [`38c4ce415`](https://github.com/siderolabs/talos/commit/38c4ce415dc8535b4a7403f7a35c5440f2f4aeb6) feat: add user-space InfiniBand modules
* [`251dc934f`](https://github.com/siderolabs/talos/commit/251dc934f3f4d9d81d6d11fd66cf4e52517d9878) feat: arm64 support for platform vmware
* [`09b3ad577`](https://github.com/siderolabs/talos/commit/09b3ad5771b4ee813dcb4d53ad8d291b74b8d8fa) feat: update containerd to 2.1.2
* [`0767dd07b`](https://github.com/siderolabs/talos/commit/0767dd07b9067aeb3470d463ff32874c69082853) chore: enable --with-siderolink-agent on Darwin
* [`9642198d7`](https://github.com/siderolabs/talos/commit/9642198d76963bd9f6bdda03fb31c165f31f8087) fix: userspace wireguard library overrides
* [`208f0763e`](https://github.com/siderolabs/talos/commit/208f0763ef2db94a913606051b5d223d1de61f24) chore: fix talosctl build on non-Linux hosts
* [`87421af87`](https://github.com/siderolabs/talos/commit/87421af87a88851b78e576b2f9b4af9a48f0acb8) docs: expand documentation description
* [`d32ccfa59`](https://github.com/siderolabs/talos/commit/d32ccfa598284450477af166734595dc952021fa) feat: implement swap support
* [`8f5cf81db`](https://github.com/siderolabs/talos/commit/8f5cf81dba80015f66037ee181f17eb2294bb8a2) docs: update kvm documentation
* [`8e84c8b0f`](https://github.com/siderolabs/talos/commit/8e84c8b0f8405be519a9f0530e34a612ff054373) fix: nil pointer deref in quirk
* [`6e74a3676`](https://github.com/siderolabs/talos/commit/6e74a367636dc21e2bf017d6284bbf998a4bad7d) docs: aad ery basic details on how to run on scaleway
* [`260d1bc9a`](https://github.com/siderolabs/talos/commit/260d1bc9a93f5f6added5e6998f3d2f08fedb770) fix: correctl close encrypted volumes
* [`034ef42af`](https://github.com/siderolabs/talos/commit/034ef42af25ee3dacf5dd0391385ea881b6d5d32) fix: update siderolink library for wgtunnel panic fix
* [`3035744a8`](https://github.com/siderolabs/talos/commit/3035744a8096270691f6bdccfabe34ad53da489c) fix: correctly predict interface name on darwin
* [`cfcfad3c4`](https://github.com/siderolabs/talos/commit/cfcfad3c45376b8ebb989b865f3c13729c87d388) chore: move `checkUnknownKeys` function to `github.com/siderolabs/gen`
* [`5ecc53c69`](https://github.com/siderolabs/talos/commit/5ecc53c695ec578dbc32f00fa7df65b31a5e77aa) docs: add macos section to developing-talos.md
* [`b5b35307f`](https://github.com/siderolabs/talos/commit/b5b35307fe950d0de9ee2ff1d5686af858db13b4) chore: update Go to 1.24.4
* [`fde772d8d`](https://github.com/siderolabs/talos/commit/fde772d8d82e9d6bc7e63b49c965b8d924e308ab) feat: update Flannel to 0.27.0
* [`81ca27949`](https://github.com/siderolabs/talos/commit/81ca27949427c546f43b0409b56f733becabc2f6) release(v1.11.0-alpha.1): prepare release
* [`58a868e68`](https://github.com/siderolabs/talos/commit/58a868e68833e94d691e7ed029dce629446fecc3) chore: fix renovate config, add release-gate label
* [`a59aaee84`](https://github.com/siderolabs/talos/commit/a59aaee84bcceb20792bc4782748449ad93b0530) feat: bump dependencies, Linux 6.12.31
* [`e954ee30a`](https://github.com/siderolabs/talos/commit/e954ee30add42de6f42cbb7d96927722102afdb7) docs: typo correction: LongHorn -> Longhorn
* [`aab053394`](https://github.com/siderolabs/talos/commit/aab053394bafdf718196133e38be010d847db0ad) fix: mashal resource byte slices as strings in YAML
* [`c7d4191e7`](https://github.com/siderolabs/talos/commit/c7d4191e78bf0a455ab596f46d4cf212dce694a4) fix: rework the way CRI config generation is waited for
* [`0114183de`](https://github.com/siderolabs/talos/commit/0114183de62e4ab930ff0f10dd156f935d57cf10) docs: update `lastRelease` to 1.10.3
* [`938b0760a`](https://github.com/siderolabs/talos/commit/938b0760abdb41be1be4da02b877e2c902d594be) docs: update issue template
* [`2a7b735b2`](https://github.com/siderolabs/talos/commit/2a7b735b264ebcfa22dc2d6044c9d5cd3057b5c2) feat: drop IMA support
* [`2d5a805b0`](https://github.com/siderolabs/talos/commit/2d5a805b0ebabb804b3c32be18db1d718a91070f) fix: typo in DiscoverdVolume spec
* [`60c12bad9`](https://github.com/siderolabs/talos/commit/60c12bad93b422db2784b0203d94ca69fa31957c) feat: support nocloud include url userdata directive
* [`0fd622c82`](https://github.com/siderolabs/talos/commit/0fd622c825ba1fbb833a4b8920ac4c4e56f08a1f) fix(talosctl): correct --help output for dashboard command
* [`a90c936a1`](https://github.com/siderolabs/talos/commit/a90c936a16756cfe5fe451258f0022b808be17d2) feat: support qemu provisioner on darwin
* [`5322ca0d3`](https://github.com/siderolabs/talos/commit/5322ca0d372aa20ad90e66f04699b75debb0ab80) docs: update overlay docs
* [`a60b6322d`](https://github.com/siderolabs/talos/commit/a60b6322d1e8fbd75394e0bdb4435af605b32bbb) fix(ci): drop nebula from extensions test
* [`dbbb59a67`](https://github.com/siderolabs/talos/commit/dbbb59a6781f79ee34a6e91a72575802561c58b6) docs: add note for default `dataDirHostPath` for Rook
* [`e26054378`](https://github.com/siderolabs/talos/commit/e2605437826911cd60a6a4d9ee760a6a242e244b) docs: macos qemu provider
* [`5d0224093`](https://github.com/siderolabs/talos/commit/5d022409357d41831fa1bfd34ccdcfceecca42df) docs: use the cilium-cli image repo in the job installation manifest
* [`ff80e4cca`](https://github.com/siderolabs/talos/commit/ff80e4cca086fa01d84ceb750111dc9e31ccc978) docs: fix CIDR name
* [`a5fd15e8b`](https://github.com/siderolabs/talos/commit/a5fd15e8bd4a4547e3658981543401fd9eb8cd80) fix(ci): reproducibility test
* [`8f8963e50`](https://github.com/siderolabs/talos/commit/8f8963e50d7b05d1361fd44040c0f1ffb94693af) docs: update Nexxen brand
* [`c6b86872d`](https://github.com/siderolabs/talos/commit/c6b86872dc0d62aef5ad70fce00c411080911ace) fix(ci): iso reproducibility file permissions
* [`995a1dec4`](https://github.com/siderolabs/talos/commit/995a1dec4a34f49d84daff16b30f8920275a439d) chore: add a check for unsupported darwin flags
* [`9db5d0c97`](https://github.com/siderolabs/talos/commit/9db5d0c97ac31c7f6ce0b23d999126fc6cc094ec) fix: nocloud metadata for hostname
* [`3cf325654`](https://github.com/siderolabs/talos/commit/3cf325654e4a7f73196241e59e3ca6b5f24c3e19) feat: modularize more arm64 kernel
* [`3524745cc`](https://github.com/siderolabs/talos/commit/3524745cc49c51e4f13da954a57ab56d467fd26e) fix: allow any PKI in Talos API
* [`f438cdb09`](https://github.com/siderolabs/talos/commit/f438cdb0993b17f0e540ecefa39cde09f89730f4) chore: use custom dhcpd server on macos qemu
* [`11c17fb9a`](https://github.com/siderolabs/talos/commit/11c17fb9aad2443b10e15295069b8e24e0d514e2) fix: metal-iso reproducibility
* [`7fcb89ee3`](https://github.com/siderolabs/talos/commit/7fcb89ee385fdbf47dae4a8308299c00488df84a) chore: add darwin vmnet qemu support
* [`fc1237343`](https://github.com/siderolabs/talos/commit/fc1237343f79a1be907c43ac3ce116168409ed17) chore: clean up `/usr/bin`
* [`b551f32ce`](https://github.com/siderolabs/talos/commit/b551f32ce550f2bc3c679a9857f28d604a297bbf) feat: update containerd to v2.1.1
* [`67f4154f9`](https://github.com/siderolabs/talos/commit/67f4154f920fc0c58a9a832e14fbc7f9430747b3) docs: update disk-management.md
* [`0cb137ad7`](https://github.com/siderolabs/talos/commit/0cb137ad7366e2386f49a99aee0a3c5ffb7223f6) fix: make disk size check work on old Talos
* [`7c057edd5`](https://github.com/siderolabs/talos/commit/7c057edd5f3636dff6932ad9fbd7c51867b0c2c8) fix: use vmdk-convert istead of qemu-img to create VMDK for OVA files
* [`cd618dad0`](https://github.com/siderolabs/talos/commit/cd618dad0feb1390e5945e2bba1d20bcecf30c2a) chore: update the go-blockdevice package
* [`0b99631a0`](https://github.com/siderolabs/talos/commit/0b99631a0b64ce8d65ddcf7f40b2168debf11a62) fix: bump apid memory limit
* [`5451f35b1`](https://github.com/siderolabs/talos/commit/5451f35b148a630c6ab011dce44b52fd2ad327ba) docs: update virtualbox
* [`bd4d202a5`](https://github.com/siderolabs/talos/commit/bd4d202a5a67c56b6c6e6bc962f6bd51c729759f) refactor: bring owned.State from COSI to simplify tests
* [`0b96df574`](https://github.com/siderolabs/talos/commit/0b96df57476af86a37bcfdbf28a479444a9e6e5c) feat: update containerd to 2.1.0
* [`e1a939144`](https://github.com/siderolabs/talos/commit/e1a939144f25acc6a2715feedb30a56a47f6793d) docs: fix formatting in disk encryption
* [`7a817df1c`](https://github.com/siderolabs/talos/commit/7a817df1cce58de2a16b72b37a54ffc0103af79a) docs: fix typo
* [`f35b213b2`](https://github.com/siderolabs/talos/commit/f35b213b2b448c2e0065d4698095a843dd2f5268) test: fix DHCP unicast failures in QEMU environment
* [`7064bbf05`](https://github.com/siderolabs/talos/commit/7064bbf056f083de0f7174c9d3c600871189b4e5) docs: fix vmware factory URL
* [`78c33bcdb`](https://github.com/siderolabs/talos/commit/78c33bcdb9a30195ce401311e82b2e189faf33f3) feat: update default Kubernetes to v1.33.1
* [`da6795266`](https://github.com/siderolabs/talos/commit/da67952666d2db2b8b5636bd4cae8af09a139410) fix: disable automatic MAC assignment to bridge interfaces
* [`ca34adf58`](https://github.com/siderolabs/talos/commit/ca34adf585bfe04d2d1b84f186cb87aa77fc8e00) chore(ci): drop azure keys
* [`ea5de19fa`](https://github.com/siderolabs/talos/commit/ea5de19fad3f62889899c0d89d08b8b73dfa75da) fix: selinux detection
* [`52c76ea3a`](https://github.com/siderolabs/talos/commit/52c76ea3a61a4a3cbd963dc2ff0d6d21b4210bcd) fix: consistently apply dynamic grpc proxy dialer
* [`aa9569e5d`](https://github.com/siderolabs/talos/commit/aa9569e5d8c59b762dfd64a4e9ef42cfdc6f9d51) chore: refactor cluster create cmd flags
* [`1161faa05`](https://github.com/siderolabs/talos/commit/1161faa0594c033bf032852b880439b2082c9722) docs: fix typo in Cilium docs
* [`164745e44`](https://github.com/siderolabs/talos/commit/164745e44334146b8a6f696640692c25b731414a) docs: remove `preserve` flag mention in upgrade notes
* [`9a2ecbaaf`](https://github.com/siderolabs/talos/commit/9a2ecbaaf7b7a3f393dd29272aca34e069a24c6e) fix: makefile operating system param
* [`118aa69d6`](https://github.com/siderolabs/talos/commit/118aa69d6f6e71b88747db1e8234d478daa54ab4) chore: update cloud-image-uploader dependencies
* [`acdd721cf`](https://github.com/siderolabs/talos/commit/acdd721cfa62f9888a9ceea1693c17348c0d663a) chore: dump qemu pachine ipam records on darwin
* [`bb9094534`](https://github.com/siderolabs/talos/commit/bb90945344f02b9cdae6e0e01821792dca25096b) chore: rotate aws iam credentials
* [`0bfa4ae1b`](https://github.com/siderolabs/talos/commit/0bfa4ae1b06e1e6330adf331e1a97651bbe39b4a) chore: update deps for cloud-image-uploader
* [`956d7c71b`](https://github.com/siderolabs/talos/commit/956d7c71bcdff639b8261cf6cf1a5d19cf702f75) chore: update sops keys
* [`e2f819d88`](https://github.com/siderolabs/talos/commit/e2f819d880373102f8a8c7f0ff549e37ba75a08e) test: fix the process runner log collection
* [`fdac4cfb9`](https://github.com/siderolabs/talos/commit/fdac4cfb9143853eb21d38e1b3d517455b0ba0f2) fix: upgrade go-kubernetes for DRA flag bug
* [`09d88e1e8`](https://github.com/siderolabs/talos/commit/09d88e1e8374ef19e5730994d9b098333347f0b7) test: fix some flaky tests
* [`ec1f41a94`](https://github.com/siderolabs/talos/commit/ec1f41a948b1bda02096434e47f2a2a767951fe9) chore: make qemu config server bind work on darwin
* [`980f4d2b9`](https://github.com/siderolabs/talos/commit/980f4d2b936cfdc3ebc9882f7c25fbf2d2aa49f8) feat: bump dependencies
* [`95259337e`](https://github.com/siderolabs/talos/commit/95259337ee0ccb22d7e9125074818ac8f9afa7af) fix: k8s 1.32->1.33 upgrade check
* [`c3c326b40`](https://github.com/siderolabs/talos/commit/c3c326b405804c258b68f19b8d7dacca32535e9b) fix: improve volume mounter automaton
* [`918b94d9a`](https://github.com/siderolabs/talos/commit/918b94d9a0b71b759073f8f7eb0f5dc7fdff413f) refactor: rewrite disk size check
* [`ab7e693d7`](https://github.com/siderolabs/talos/commit/ab7e693d76500b6cdc2068221bdfce16633a8b01) chore: make qemu lb address bind work on darwin
* [`97ceab001`](https://github.com/siderolabs/talos/commit/97ceab001c1bb79407c40d8fff867342656187b9) fix: multiple logic issues in platform network config controller
* [`46349a9df`](https://github.com/siderolabs/talos/commit/46349a9df5d026a4e4b807a94865d5b3c371d32a) docs: remove azure image gallery instructions
* [`0cfcdd3de`](https://github.com/siderolabs/talos/commit/0cfcdd3de1a20690ce47d63bb56b3d33d11c1474) docs: fix search on base talos.dev
* [`78646b4e0`](https://github.com/siderolabs/talos/commit/78646b4e050358b930d27e4eddcfb22c4c825b0c) docs: add registryd debug command
* [`c6824c211`](https://github.com/siderolabs/talos/commit/c6824c211438a3fb663f4233e8663732ab2ddf44) fix: deny apply config requests without v1alpha1 in "normal" mode
* [`7df0408e4`](https://github.com/siderolabs/talos/commit/7df0408e460ebc392c6927c7b23e3795b9bd2140) fix: interactive installer config gen
* [`881c5d62b`](https://github.com/siderolabs/talos/commit/881c5d62bf0d1f3311b3cf946b7801f97c1fb94b) fix: suppress duplicate platform config updates
* [`66d77888e`](https://github.com/siderolabs/talos/commit/66d77888e42798995ddc73db3869d16959e53376) fix: replace downloaded asset paths correctly in cluster create cmd
* [`6bd6c9b5a`](https://github.com/siderolabs/talos/commit/6bd6c9b5a08ca3b0e9574e1a61edc54c6ff722bb) fix: generate iso greater than 4 gig
* [`ac140324e`](https://github.com/siderolabs/talos/commit/ac140324ebfb54f580c9b9bbbb55549bd5ffa11e) fix: skip PCR extension if TPM1.2 is found
* [`09ef1f8a4`](https://github.com/siderolabs/talos/commit/09ef1f8a41c84e6a16729e6b6aff81788da0e3f5) fix: ignore http proxy on grpc socket dial
* [`22a72dc80`](https://github.com/siderolabs/talos/commit/22a72dc80f2037a4cc7ad696d8dff504deb22630) chore: split options between three structs
* [`22c34a50f`](https://github.com/siderolabs/talos/commit/22c34a50fc66edd174ab4a65961257de28a6daa0) fix(ci): provision cron jobs
* [`b3b20eff3`](https://github.com/siderolabs/talos/commit/b3b20eff3a29f74d18df634cbb01f41bde17f2c8) fix: containerd crashing with sigsegv
* [`f7891c301`](https://github.com/siderolabs/talos/commit/f7891c3018de248c7c66483562227b614689413c) chore: calculate vmnet interface name preemptively
* [`ae87edffb`](https://github.com/siderolabs/talos/commit/ae87edffbcdaed12fef41541622f27882ed63755) fix: drop libseccomp from rootfs
* [`f74a805bb`](https://github.com/siderolabs/talos/commit/f74a805bb067f55619cae7aebb92f00bb8173c92) fix: do correct backoff for nocloud reconcile
* [`01bb294af`](https://github.com/siderolabs/talos/commit/01bb294af63f193dafa12cb623ea77ad67b698fb) fix(ci): provision tests
* [`e4945be3b`](https://github.com/siderolabs/talos/commit/e4945be3bc43cbc275e2ea5f399a0188c5e16ad8) docs: add registryd debug command
* [`d8c670ad3`](https://github.com/siderolabs/talos/commit/d8c670ad3ecba32c70ff365eaf7a5a4ccb5d721a) release(v1.11.0-alpha.0): prepare release
* [`ace44ea61`](https://github.com/siderolabs/talos/commit/ace44ea6169d419f188e0a2456c31f420e61ae77) test: update hydrophone to 0.7.0
* [`3a1163692`](https://github.com/siderolabs/talos/commit/3a1163692da7b41b17f263ab43d0fd81abafc4f8) chore: cross platform qemu preflight checks
* [`7914fb104`](https://github.com/siderolabs/talos/commit/7914fb10412d31a1b75c74b0c66578e55fb77bc7) chore: move the create command to it's own package
* [`c8e619608`](https://github.com/siderolabs/talos/commit/c8e619608dc8898be71a17c54503085ef38abf37) chore: prepare for release 1.11
* [`1299aaa45`](https://github.com/siderolabs/talos/commit/1299aaa45d997dd23aed380f858cec3bc6b975e4) chore(ci): add extensions test for Youki runtime
* [`e50ceb221`](https://github.com/siderolabs/talos/commit/e50ceb221e56f0760d5f2fc9e4b821d6b29add05) docs: activate Talos 1.10 docs
* [`9d12aaeb1`](https://github.com/siderolabs/talos/commit/9d12aaeb19d68c5e692921b938d72347f6129f65) test: improve config patch test
* [`106a656b6`](https://github.com/siderolabs/talos/commit/106a656b6132e766e9e9ef7b1c12b97a413b5de6) chore: make qemu provider build on darwin
* [`8013aa06c`](https://github.com/siderolabs/talos/commit/8013aa06cd338f1dd11061d3455767fee4b9783c) test: replace platform metadata test
* [`2b89c2810`](https://github.com/siderolabs/talos/commit/2b89c2810551ab52678e62fcbf5355dd05c72030) fix: relax etcd APIs RBAC requirements
* [`1e677587c`](https://github.com/siderolabs/talos/commit/1e677587c0e6c61f724a85f18ee9d436ae6da038) fix: preserve kubelet image suffix
* [`62ab8af45`](https://github.com/siderolabs/talos/commit/62ab8af459475cbd24a2f34d8923ce70d1fda3db) fix: disk image generation with image cache
* [`d60626f01`](https://github.com/siderolabs/talos/commit/d60626f017ef495210939ee4f8ef7f623dd325f9) fix: handle encryption type mismatch
* [`a9109ebd0`](https://github.com/siderolabs/talos/commit/a9109ebd00fcd300bf4262142ade77df6788852b) feat: allow SideroLink unique token in machine config
* [`2ff3a6e40`](https://github.com/siderolabs/talos/commit/2ff3a6e4079a29b6b45770204fd8cb30369518e9) feat(kernel): add bcache kernel module to core talos
* [`fa95a2146`](https://github.com/siderolabs/talos/commit/fa95a2146056bfe1ae322cb574fd8d432745b5c9) fix(ci): bios provision test
* [`f7c5b86be`](https://github.com/siderolabs/talos/commit/f7c5b86be7e2b28906cb66b466a017887ac5e2b6) fix: sync PCR extension with volume provisioning lifecycle
* [`f90c79474`](https://github.com/siderolabs/talos/commit/f90c79474b50da35ab8e285ee9723957e4b6cf00) chore: show bound driver in pcidevices info
* [`8db34624c`](https://github.com/siderolabs/talos/commit/8db34624c6ed9707ba1165da790f5b389bd1c92f) fix: handle correctly changing platform network config
* [`77c7a075b`](https://github.com/siderolabs/talos/commit/77c7a075bbba7ffd24dbd9d5e069ccb50f8143b4) feat: update Kubernetes to 1.33.0
* [`74f0c48c7`](https://github.com/siderolabs/talos/commit/74f0c48c738b0b80278667c3e5a1c5e1ecd5a078) feat: add version compatibility for Talos 1.11
* [`c4fb7dad0`](https://github.com/siderolabs/talos/commit/c4fb7dad0ec390781cca54e2348f116cb1cf1866) fix: force DNS runner shutdown on timeout
* [`c49b4836e`](https://github.com/siderolabs/talos/commit/c49b4836e46725940f4731e182475905ebee6019) docs: hetzner: add note about public iso
* [`16ea2b113`](https://github.com/siderolabs/talos/commit/16ea2b113fad0c81a96dbcfdf4fd1b9f43bb1282) docs: add what is new for 1.10
* [`be3f0c018`](https://github.com/siderolabs/talos/commit/be3f0c018c50da3d920ed8fe36d4f31c5d3edfac) fix: fix Gvisor tests with containerd patch
* [`37db132b3`](https://github.com/siderolabs/talos/commit/37db132b3b3e6c58f15228c64b023e77c15cf012) chore(ci): add provision test with bios
* [`ec60b70e7`](https://github.com/siderolabs/talos/commit/ec60b70e7245f49f6ac1d48cd4292b85f1d6f79e) fix: set media type to OCI for image cache layer
* [`a471eb31b`](https://github.com/siderolabs/talos/commit/a471eb31b87b393ee9fc57fbc725801d08386ad4) feat: update Linux 6.12.24, containerd 2.0.5
* [`54ad5b872`](https://github.com/siderolabs/talos/commit/54ad5b8729c7d54da2efa6baf7886163741176ed) fix: extension services logging to console
* [`601f036ba`](https://github.com/siderolabs/talos/commit/601f036ba9cc762d6a3c6ae819654005f1d49527) docs: correct flannel extra args example
* [`ae94377d1`](https://github.com/siderolabs/talos/commit/ae94377d15a3b70248fbb446d13d7ae96bb04e82) feat: support encryption config for user volumes
* [`9616f6e8d`](https://github.com/siderolabs/talos/commit/9616f6e8d280e64815fe3e1ba324df1dd5d2122d) docs: add caveat for kubespan and host ports
* [`a1d08a362`](https://github.com/siderolabs/talos/commit/a1d08a3624c7c8b5213b8e9dee1cf9289d6719dc) docs: fixes typo at OpenEBS Mayastor worker patches
* [`a91e8726e`](https://github.com/siderolabs/talos/commit/a91e8726e433be9db58f1a7a09a4cca422b2b50c) docs: add a dark theme
* [`c76189c58`](https://github.com/siderolabs/talos/commit/c76189c58a2fe65954924168d7077350974829dd) fix: grub EFI mount point
* [`4ca985c65`](https://github.com/siderolabs/talos/commit/4ca985c656c1924e550d06c073a7c1b6cb03f392) fix: grub efi platform install
* [`b31260281`](https://github.com/siderolabs/talos/commit/b31260281dba752e06fcfc645bb020872602d898) docs: update storage.md
* [`396a29040`](https://github.com/siderolabs/talos/commit/396a290408eff5bda4ad31fafc33496bea9aa899) feat: add new SBCs
* [`a902f6580`](https://github.com/siderolabs/talos/commit/a902f6580f8e104977521a335a41c0cd70256906) feat: update Flannel to v0.26.7
* [`2bbefec1a`](https://github.com/siderolabs/talos/commit/2bbefec1abacae2952782fbd163ef52d34f09858) docs: use cache in preview
* [`6028a8d2d`](https://github.com/siderolabs/talos/commit/6028a8d2da571a8a37712f9917e24372cf5af919) docs: update kubeprism.md
* [`e51a8ef8c`](https://github.com/siderolabs/talos/commit/e51a8ef8c68bb1cfab2ac845a0b6792d7e000324) fix: prefer new `MountStatus` resource
* [`d9c7e7946`](https://github.com/siderolabs/talos/commit/d9c7e79462496d6756c55b0672994aa262eaed4f) docs: fix search
* [`b32fa029b`](https://github.com/siderolabs/talos/commit/b32fa029b3f550b3403e25e23aac889d61366389) feat: update Kubernetes to 1.33.0-rc.1
* [`f0ea478cb`](https://github.com/siderolabs/talos/commit/f0ea478cb811675a450839b8dcd351e43404efd4) feat: support address priority
* [`8cd3c8dc7`](https://github.com/siderolabs/talos/commit/8cd3c8dc77b25270ed8dea65cbbd4e87c203ee74) test: fix NVIDIA OSS tests
* [`62f2d27cd`](https://github.com/siderolabs/talos/commit/62f2d27cd44de5112055b5b47f23b001cadccaae) docs: update virtualbox.md
* [`141326ea3`](https://github.com/siderolabs/talos/commit/141326ea3bb2e471a5cb51fd565521683a9792fc) docs: fix tabpane styling
* [`134aa53cc`](https://github.com/siderolabs/talos/commit/134aa53ccaba55754544977d695ad3ca5d34e604) feat: update base CoreDNS code in host DNS to 1.12.1
</p>
</details>

### Dependency Changes

* **github.com/bougou/go-ipmi**                  v0.7.6 -> v0.8.1
* **github.com/cosi-project/runtime**            v0.10.6 -> v1.13.0
* **github.com/fullstorydev/grpchan**            v1.1.1 -> v1.1.2
* **github.com/siderolabs/talos**                v1.10.2 -> v1.12.2
* **github.com/siderolabs/talos/pkg/machinery**  v1.10.2 -> v1.12.2
* **github.com/spf13/cobra**                     v1.9.1 -> v1.10.2
* **github.com/stretchr/testify**                v1.11.1 **_new_**
* **go.uber.org/zap**                            v1.27.0 -> v1.27.1
* **golang.org/x/sync**                          v0.14.0 -> v0.19.0
* **google.golang.org/grpc**                     v1.72.1 -> v1.78.0
* **google.golang.org/protobuf**                 v1.36.6 -> v1.36.11

Previous release can be found at [v0.1.3](https://github.com/siderolabs/talos-metal-agent/releases/tag/v0.1.3)

