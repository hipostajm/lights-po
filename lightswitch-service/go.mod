module lightswitch-service

go 1.26.1

require (
	github.com/google/uuid v1.6.0
	google.golang.org/grpc v1.80.0
	proto v0.0.0
)

require (
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260401024825-9d38bb4040a9 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace proto => ../proto
