FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go mod download
RUN go build -o ./dental_clinic_go ./cmd/server/main.go


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/dental_clinic_go .
ENTRYPOINT ["./dental_clinic_go"]