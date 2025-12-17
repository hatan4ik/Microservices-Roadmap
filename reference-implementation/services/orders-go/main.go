package main

import (
  "context"
  "log"
  "net"
  "os"

  ordersv1 "github.com/example/microservices-zero-to-hero/reference-implementation/proto/ordersv1"
  "google.golang.org/grpc"
)

type server struct {
  ordersv1.UnimplementedOrdersServiceServer
}

func (s *server) CreateOrder(ctx context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {
  // In a real system:
  // - validate invariants (customer exists, items exist, stock policy)
  // - write to DB transactionally
  // - publish event via outbox
  id := "ord_" + randID()
  log.Printf("CreateOrder customer=%s items=%d id=%s", req.GetCustomerId(), len(req.GetItemIds()), id)
  return &ordersv1.CreateOrderResponse{OrderId: id}, nil
}

func main() {
  addr := getenv("ORDERS_GRPC_ADDR", ":50051")
  lis, err := net.Listen("tcp", addr)
  if err != nil {
    log.Fatalf("listen: %v", err)
  }
  grpcServer := grpc.NewServer()
  ordersv1.RegisterOrdersServiceServer(grpcServer, &server{})

  log.Printf("orders gRPC listening on %s", addr)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("serve: %v", err)
  }
}

func getenv(k, def string) string {
  if v := os.Getenv(k); v != "" {
    return v
  }
  return def
}

// randID is intentionally simple for training. Do not use for production IDs.
func randID() string {
  const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
  b := make([]byte, 8)
  for i := range b {
    b[i] = alphabet[i%len(alphabet)]
  }
  return string(b)
}
