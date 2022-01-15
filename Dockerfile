FROM golang:1.17-buster AS builder
LABEL maintainer="thedondope@hey.com"
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . ./
RUN go build -v -o /bin/server cmd/server/*.go

FROM debian:buster-slim
RUN set -x && apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends ca-certificates=20190110 && \
  rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /bin/server ./

CMD ["./server"]