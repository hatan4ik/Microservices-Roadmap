package main

import (
  "encoding/json"
  "log"
  "net/http"
  "os"
  "time"
)

type Product struct {
  ID    string `json:"id"`
  Name  string `json:"name"`
  Price int64  `json:"price_cents"`
}

func main() {
  addr := getenv("CATALOG_ADDR", ":8080")

  mux := http.NewServeMux()
  mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
  })
  mux.HandleFunc("/v1/products", func(w http.ResponseWriter, r *http.Request) {
    // In a real system, this is a query model (CQRS read side).
    // For training, return a stable in-memory list.
    products := []Product{
      {ID: "p1", Name: "Flight Simulator Session", Price: 9900},
      {ID: "p2", Name: "Pro Pilot Prep", Price: 19900},
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
  })

  srv := &http.Server{
    Addr:              addr,
    Handler:           logMiddleware(mux),
    ReadHeaderTimeout: 5 * time.Second,
  }

  log.Printf("catalog http listening on %s", addr)
  if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    log.Fatalf("serve: %v", err)
  }
}

func getenv(k, def string) string {
  if v := os.Getenv(k); v != "" {
    return v
  }
  return def
}

func logMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    next.ServeHTTP(w, r)
    log.Printf("%s %s dur=%s", r.Method, r.URL.Path, time.Since(start))
  })
}
