FROM jinxprotocol-base

COPY ./testing/e2etest-local/local.sh /jinxprotocol/local.sh
COPY ./testing/genesis.sh /jinxprotocol/genesis.sh
COPY ./testing/start.sh /jinxprotocol/start.sh
COPY ./daemons/pricefeed/client/constants/testdata /jinxprotocol/exchange_config
COPY ./testing/delaymsg_config /jinxprotocol/delaymsg_config

RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.4.0

RUN /jinxprotocol/local.sh

ENV DAEMON_NAME=jinxprotocold
ENTRYPOINT ["cosmovisor", "run"]
