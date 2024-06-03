# Simple usage with a mounted data directory:
# > docker build -t market .
# > docker run -it -v ~/.reserve:/reserve/.reserve onomy/reserve-dev init market --home /reserve/.reserve
# Copy genesis.json from tools/config/devnet to ~/.reserve/config and Dealer and Validator keys are in dev/config
# > docker run -it -v ~/.reserve:/reserve/.reserve onomy/reserve-dev keys add dealer --recover --home /reserve/.reserve
# > docker run -it -v ~/.reserve:/reserve/.reserve onomy/reserve-dev keys add validator --recover --home /reserve/.reserve
# > docker run -it -v ~/.reserve:/reserve/.reserve onomy/reserve-dev gentx validator 10000000000000000000stake --chain-id market --home /reserve/.reserve
# > docker run -it -v ~/.reserve:/reserve/.reserve onomy/reserve-dev collect-gentxs --home /reserve/.reserve
# > docker run -it -p 26656:26656 -p 26657:26657 -p 1317:1317 -p 9090:9090 -p 9091:9091 -d -v ~/.reserve:/reserve/.reserve onomy/reserve-dev start --home /reserve/.reserve
FROM golang:1.19-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/pendulum-labs/market

# Add source files
COPY . .
RUN pwd
RUN ls

RUN go version

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES
RUN make install

# Final image
FROM alpine:edge

ENV MARKET /market

# Install ca-certificates
RUN apk add --update ca-certificates

WORKDIR $MARKET

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/reserved /usr/bin/reserved

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317
EXPOSE 9090
EXPOSE 9091

# Run reserved by default, omit entrypoint to ease using container with marketcli
ENTRYPOINT ["reserved"]