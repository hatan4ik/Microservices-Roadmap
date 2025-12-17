module github.com/example/microservices-zero-to-hero/reference-implementation/services/orders-go

go 1.22

require (
  google.golang.org/grpc v1.66.0
  google.golang.org/protobuf v1.34.2
)

replace github.com/example/microservices-zero-to-hero/reference-implementation/proto/ordersv1 => ../../proto/ordersv1
