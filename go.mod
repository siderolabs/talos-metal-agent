module github.com/siderolabs/talos-metal-agent

go 1.23.4

replace (
	github.com/bougou/go-ipmi => github.com/utkuozdemir/go-ipmi v0.5.12-0.20241216144512-73a3fe8e6f78
	github.com/pensando/goipmi => github.com/siderolabs/goipmi v0.0.0-20211214143420-35f956689e67
)

require (
	github.com/bougou/go-ipmi v0.5.11
	github.com/cosi-project/runtime v0.7.6
	github.com/jhump/grpctunnel v0.3.0
	github.com/pensando/goipmi v0.0.0-20240603174436-eb122d901c23
	github.com/planetscale/vtprotobuf v0.6.1-0.20241121165744-79df5c4772f2
	github.com/siderolabs/go-procfs v0.1.2
	github.com/siderolabs/talos v1.9.0-beta.1
	github.com/siderolabs/talos/pkg/machinery v1.9.0-beta.1
	github.com/spf13/cobra v1.8.1
	go.uber.org/zap v1.27.0
	golang.org/x/sync v0.10.0
	google.golang.org/grpc v1.69.0
	google.golang.org/protobuf v1.35.2
)

require (
	cel.dev/expr v0.19.1 // indirect
	github.com/ProtonMail/go-crypto v1.1.3 // indirect
	github.com/ProtonMail/go-mime v0.0.0-20230322103455-7d82a3887f2f // indirect
	github.com/ProtonMail/gopenpgp/v2 v2.8.1 // indirect
	github.com/adrg/xdg v0.5.3 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cloudflare/circl v1.5.0 // indirect
	github.com/containerd/go-cni v1.1.11 // indirect
	github.com/containernetworking/cni v1.2.3 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fullstorydev/grpchan v1.1.1 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/google/cel-go v0.22.1 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.24.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/josharian/native v1.1.0 // indirect
	github.com/jsimonetti/rtnetlink/v2 v2.0.2 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mdlayher/ethtool v0.2.0 // indirect
	github.com/mdlayher/genetlink v1.3.2 // indirect
	github.com/mdlayher/netlink v1.7.2 // indirect
	github.com/mdlayher/socket v0.5.1 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/opencontainers/runtime-spec v1.2.0 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/siderolabs/crypto v0.5.1 // indirect
	github.com/siderolabs/gen v0.7.0 // indirect
	github.com/siderolabs/go-api-signature v0.3.6 // indirect
	github.com/siderolabs/go-blockdevice/v2 v2.0.7 // indirect
	github.com/siderolabs/go-pointer v1.0.0 // indirect
	github.com/siderolabs/net v0.4.0 // indirect
	github.com/siderolabs/protoenc v0.2.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/exp v0.0.0-20241210194714-1829a127f884 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.8.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241209162323-e6fa225c2576 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241209162323-e6fa225c2576 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
