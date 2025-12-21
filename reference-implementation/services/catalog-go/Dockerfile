FROM golang:1.22-alpine AS build
WORKDIR /src

# Copy shared modules (if any in future)
# COPY proto/ proto/

# Copy the service and its go.mod/go.sum
COPY services/catalog-go/ services/catalog-go/

# Build service
WORKDIR /src/services/catalog-go
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /out/catalog .

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/catalog /catalog
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/catalog"]
