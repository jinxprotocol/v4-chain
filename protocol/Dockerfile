# NB: This is a digest for a multi-arch manifest list, you will want to get this by running
# `docker buildx imagetools inspect golang:1.21-alpine`
ARG GOLANG_1_21_ALPINE_DIGEST="926f7f7e1ab8509b4e91d5ec6d5916ebb45155b0c8920291ba9f361d65385806"

# This Dockerfile is a stateless build of the `jinxprotocold` binary as a Docker container.
# It does not include any configuration, state, or genesis information.

# --------------------------------------------------------
# Builder
# --------------------------------------------------------

FROM golang@sha256:${GOLANG_1_21_ALPINE_DIGEST} as builder
ARG VERSION
ARG COMMIT

RUN set -eux; apk add --no-cache ca-certificates build-base; apk add git linux-headers bash

# Download go dependencies
WORKDIR /jinxprotocol
COPY go.* ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Copy the remaining files
COPY . .

# Build jinxprotocold binary
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go build \
      -mod=readonly \
      -tags "netgo,ledger,muslc" \
      -ldflags "-X github.com/cosmos/cosmos-sdk/version.Name="jinxprotocol" \
              -X github.com/cosmos/cosmos-sdk/version.AppName="jinxprotocold" \
              -X github.com/cosmos/cosmos-sdk/version.Version=$VERSION \
              -X github.com/cosmos/cosmos-sdk/version.Commit=$COMMIT \
              -X github.com/cosmos/cosmos-sdk/version.BuildTags='netgo,ledger,muslc' \
              -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
      -trimpath \
      -o /jinxprotocol/build/ \
      ./...

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

FROM golang@sha256:${GOLANG_1_21_ALPINE_DIGEST}

RUN apk add --no-cache bash

COPY --from=builder /jinxprotocol/build/jinxprotocold /bin/jinxprotocold

ENV HOME /jinxprotocol
WORKDIR $HOME

# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657
# rest server
EXPOSE 1317
# grpc
EXPOSE 9090

ENTRYPOINT ["jinxprotocold"]
