FROM golang:1.22-alpine AS build
WORKDIR /src

# Copy shared modules (if any in future)
# COPY proto/ proto/

# Build service
WORKDIR /src/services/catalog-go
COPY services/catalog-go/go.mod services/catalog-go/go.sum* ./
RUN go mod download
COPY services/catalog-go/ .
RUN CGO_ENABLED=0 go build -o /out/catalog .

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/catalog /catalog
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/catalog"]
