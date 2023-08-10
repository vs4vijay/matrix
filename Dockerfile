# Builder stage
FROM golang:1.14.1-alpine AS builder

# Secure against running as root
RUN adduser -D -u 10000 matrix
RUN mkdir /app/
RUN chown matrix /app/
USER matrix

WORKDIR /app/
ADD . /app/

# Compile the binary, but without CGO
RUN CGO_ENABLED=0 go build -o /app/matrix .

# Deploy stage
FROM alpine:3.18.3

# Secure against running as root
RUN adduser -D -u 10000 matrix
USER matrix

WORKDIR /
COPY --from=builder /app/matrix /

EXPOSE 9090

ENTRYPOINT ["/matrix"]