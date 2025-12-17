FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/orders .

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/orders /orders
EXPOSE 50051
USER 65532:65532
ENTRYPOINT ["/orders"]
