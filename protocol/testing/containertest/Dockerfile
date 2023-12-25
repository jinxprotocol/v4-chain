FROM jinxprotocol-base

COPY ./testing/containertest/containertest.sh /jinxprotocol/containertest.sh
COPY ./testing/containertest/preupgrade_entrypoint.sh /jinxprotocol/preupgrade_entrypoint.sh
COPY ./testing/containertest/preupgrade_genesis.json /jinxprotocol/preupgrade_genesis.json
COPY ./testing/genesis.sh /jinxprotocol/genesis.sh
COPY ./daemons/pricefeed/client/constants/testdata /jinxprotocol/exchange_config
COPY ./testing/delaymsg_config /jinxprotocol/delaymsg_config

RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.4.0

RUN /jinxprotocol/containertest.sh

COPY ./testing/containertest/config/pricefeed_exchange_config.toml /jinxprotocol/chain/.alice/config/pricefeed_exchange_config.toml
COPY ./testing/containertest/config/pricefeed_exchange_config.toml /jinxprotocol/chain/.bob/config/pricefeed_exchange_config.toml
COPY ./testing/containertest/config/pricefeed_exchange_config.toml /jinxprotocol/chain/.carl/config/pricefeed_exchange_config.toml
COPY ./testing/containertest/config/pricefeed_exchange_config.toml /jinxprotocol/chain/.dave/config/pricefeed_exchange_config.toml

ENTRYPOINT ["jinxprotocold"]
