FROM jinxprotocol-base

COPY ./testing/testnet-staging/staging.sh /jinxprotocol/staging.sh
COPY ./testing/genesis.sh /jinxprotocol/genesis.sh
COPY ./testing/start.sh /jinxprotocol/start.sh
COPY ./daemons/pricefeed/client/constants/testdata /jinxprotocol/exchange_config
COPY ./testing/delaymsg_config /jinxprotocol/delaymsg_config

RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.4.0

RUN /jinxprotocol/staging.sh

ENV DAEMON_NAME=jinxprotocold
ENTRYPOINT ["cosmovisor", "run"]
