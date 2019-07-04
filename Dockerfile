# === Builder ===
FROM golang:1.12 AS builder

RUN apt-get update && apt-get install -yq apt-utils gettext-base

# Reconstruct source tree inside docker
WORKDIR $GOPATH/src/github.com/altinity/clickhouse-operator
ADD . .
ENV GO111MODULE on
RUN go get .

# Build operator binary with explicitly specified output
RUN OPERATOR_BIN=/tmp/clickhouse-operator ./dev/binary_build.sh

# === Runner ===

FROM alpine:3.10
RUN apk add --no-cache ca-certificates
WORKDIR /
COPY --from=builder /tmp/clickhouse-operator .
ENTRYPOINT ["/clickhouse-operator"]
CMD ["-alsologtostderr=true", "-v=1"]
