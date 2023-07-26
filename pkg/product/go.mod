module github.com/pwnedgod/tanshogyo/pkg/product

go 1.20

replace github.com/pwnedgod/tanshogyo/pkg/common => ../common

require (
	github.com/golang/mock v1.7.0-rc.1
	github.com/pwnedgod/tanshogyo/pkg/common v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
)
