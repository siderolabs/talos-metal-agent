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

