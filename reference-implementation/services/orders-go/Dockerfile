FROM golang:1.22-alpine AS build
WORKDIR /src

# Copy shared modules
COPY proto/ proto/

# Build service
WORKDIR /src/services/orders-go
COPY services/orders-go/go.mod services/orders-go/go.sum* ./
RUN go mod download
COPY services/orders-go/ .
RUN CGO_ENABLED=0 go build -o /out/orders .

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/orders /orders
EXPOSE 50051
USER 65532:65532
ENTRYPOINT ["/orders"]
