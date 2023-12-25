<p align="center"><img src="https://jinx.exchange/icon.svg?" width="256" /></p>

<h1 align="center">jInX Chain</h1>

<div align="center">
  <a href='https://github.com/jinxprotocol/v4-chain/blob/main/LICENSE'>
    <img src='https://img.shields.io/badge/License-AGPL_v3-blue.svg' alt='License' />
  </a>
</div>

The jInX v4 software (the ”jInX Chain”) is a sovereign blockchain software built using Cosmos SDK and CometBFT. It powers a robust decentralized perpetual futures exchange with a high-performance orderbook and matching engine for a feature-rich, self-custodial perpetual trading experience on any market.

This repository contains the source code for the Cosmos SDK application responsible for running the chain and the associated indexer services. The indexer services are a read-only layer that indexes on-chain and off-chain data from a node and serves it to users in a more performant and user-friendly way.

# Getting Started

[Architecture Overview](https://jinx.exchange/blog/v4-technical-architecture-overview)

[Indexer Deep Dive](https://jinx.exchange/blog/v4-deep-dive-indexer)

[Front End Deep Dive](https://jinx.exchange/blog/v4-deep-dive-front-end)

# Resources and Documentation

[Official Documentation](https://docs.jinx.exchange/)

[jInX Academy](https://jinx.exchange/crypto-learning#)

[jInX Blog](https://jinx.exchange/blog#)

# Directory Structure

`audits` — Audit reports live here.

`docs` — Home for documentation pertaining to the entire repo.

`indexer` — Contains source code for indexer services. See its [README](https://github.com/jinxprotocol/v4-chain/blob/main/indexer/README.md) for developer documentation.

`proto` — Contains all protocol buffers used by protocol and indexer.

`protocol` — Contains source code for the Cosmos SDK app that runs the protocol. See its [README](https://github.com/jinxprotocol/v4-chain/blob/main/indexer/README.md) for developer documentation.

`v4-proto-js` — Scripts for publishing proto package to [npm](https://www.npmjs.com/package/@dydxprotocol/v4-proto).

`v4-proto-py` — Scripts for publishing proto package to [PyPI](https://pypi.org/project/v4-proto/).

# Contributing

If you encounter a bug, please file an [issue](https://github.com/jinxprotocol/v4-chain/issues) to let us know. Alternatively, please feel free to contribute a bug fix directly. If you are planning to contribute changes that involve significant design choices, please open an issue for discussion instead.

# License and Terms

The jInX Chain is licensed under AGPLv3 and subject to the [v4 Terms of Use](https://jinx.exchange/v4-terms). The full license can be found [here](https://github.com/jinxprotocol/v4-chain/blob/main/LICENSE). jInX does not deploy or run v4 software for public use, or operate or control any jInX Chain infrastructure. jInX products and services are not available to persons or entities who reside in, are located in, are incorporated in, or have registered offices in the United States or Canada, or Restricted Persons (as defined in the jInX [Terms of Use](https://jinx.exchange/terms)).
