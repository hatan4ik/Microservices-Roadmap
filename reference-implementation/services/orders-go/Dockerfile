FROM golang:1.22-alpine AS build
WORKDIR /src

# Copy shared modules
COPY proto/ proto/

# Copy the service and its go.mod/go.sum
COPY services/orders-go/ services/orders-go/

# Build service
WORKDIR /src/services/orders-go
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /out/orders .

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/orders /orders
EXPOSE 50051
USER 65532:65532
ENTRYPOINT ["/orders"]
