FROM golang:1.19 as builder
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app apps/echo/main.go

FROM gcr.io/distroless/static-debian11 as runner
COPY --from=builder /go/bin/app /
CMD ["/app"]
