FROM rhel8/go-toolset:latest AS builder
WORKDIR /app
ADD . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/app main.go

FROM ubi8/ubi-minimal:8.9-1029 AS production
WORKDIR /app
COPY --from=builder /tmp/app .
COPY docs /app/docs
CMD ["./app"]