module github.com/tinkerbell/pbnj

go 1.16

require (
	github.com/bmc-toolbox/bmclib v0.4.16-0.20211230160158-5afdbf3b6a65
	github.com/cristalhq/jwt/v3 v3.0.9
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/equinix-labs/otel-init-go v0.0.4
	github.com/fatih/color v1.7.0
	github.com/go-logr/logr v1.2.2
	github.com/go-logr/zapr v1.2.2
	github.com/go-test/deep v1.0.7
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/manifoldco/promptui v0.8.0
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/onsi/gomega v1.10.4
	github.com/packethost/pkg/grpc/authz v0.0.0-20210106215246-8e2e62dc8f0c
	github.com/packethost/pkg/log/logr/v2 v2.0.0
	github.com/philippgille/gokv v0.6.0
	github.com/philippgille/gokv/freecache v0.6.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0
	github.com/rs/xid v1.2.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.26.1
	go.opentelemetry.io/otel v1.1.0
	go.opentelemetry.io/otel/trace v1.1.0
	go.uber.org/zap v1.19.0
	goa.design/goa v2.2.5+incompatible
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/packethost/pkg/log/logr/v2 v2.0.0 => github.com/joelrebel/pkg/log/logr/v2 v2.0.0
