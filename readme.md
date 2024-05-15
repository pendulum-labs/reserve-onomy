# reserve

The Onomy Reserve

## Reserve Concepts

`denom` is a minted representation of a real-world asset
`collateral` is a coin that is deposited in a vault to mint denoms
`vault` is storage of collateral that is being or will be used for minting of denoms
`collateralization ratio` is the ratio of value of denoms minted to the collateral stored within a vault

## Reserve Parameters

`minting ratio` is the minimum value collateralization ratio at which denoms may be minted
`liquidation ratio` is the collateralization ratio at which the vault may be liquidated
`vault ratio` is the collateralization ratio that a vault currently has

## Reserve Transactions

`deposit` collateral into a vault
`mint` denoms based on vault value and ratios
`redeem` collateral from a vault by surrendering loaned denoms
`reap` vaults whose collateral ratios drop below the liquidation ratio

## Reserve Governance

`create-denom` governance proposal is how denoms are initiated

## Development

**reserve** is a blockchain module built using Cosmos SDK and Tendermint and created with [Ignite](https://github.com/ignite/cli).

Note: in order to regenerate proto files, install Ignite CLI v0.23.0 (do not use a later version until we update past SDK v0.45), `go mod tidy`, `rm -rf vendor` (you actually have to remove it for some reason), run `ignite generate proto-go`, then `go mod tidy` after that to clean up

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.com).

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.