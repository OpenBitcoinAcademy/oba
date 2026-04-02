## The Implementation Landscape

Several teams maintain independent Bitcoin implementations in different programming languages.

**btcd** (Go) is a full-node implementation written from scratch. It powers infrastructure at several Bitcoin companies and is the basis of the Lightning Network's lnd.

**bcoin** (JavaScript) is a modular full-node implementation designed for the Node.js ecosystem, with a built-in wallet and HTTP API.

**Bitcoin Knots** is a fork of Bitcoin Core maintained by Luke Dashjr. It includes additional configuration options and stricter policy defaults.

**rust-bitcoin** is a library for working with Bitcoin data structures in Rust. It provides serialization, parsing, and scripting tools without running a full node.

**libbitcoin** (C++) is an independent toolkit for building Bitcoin applications, including a full-node implementation called libbitcoin-node.

Other implementations exist in Python, Java, Scala, and C#. Each brings a different developer community into the Bitcoin ecosystem.
