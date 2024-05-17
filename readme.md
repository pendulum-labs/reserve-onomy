# reserve

The Onomy Reserve

## Reserve Concepts

`denom` is a minted representation of a real-world asset
`collateral` is a coin that is deposited in a vault to mint denoms
`vault` is storage of collateral that is being or will be used for minting of denoms
`collateralization ratio` is the ratio of value of denoms minted to the collateral stored within a vault
`liquidator` is a user that liquidates a vault that is in default because of too low of collateralization ratio

## Reserve Parameters

`vault ratio` is the collateralization ratio that of a vault at a given moment
`minting ratio` is the minimum vault ratio at which a specific denom may be minted
`liquidation ratio` is the vault ratio at which the vault for a specific denom may be liquidated by a liquidator
`burn ratio` is parameter that defines the ratio of burned vs stored NOM at liquidation of a vault

## Reserve Transactions

`deposit` collateral into a vault
`mint` denoms based on vault value and ratios
`redeem` collateral from a vault by surrendering loaned denoms
`liquidate` vaults whose collateral ratios drop below the liquidation ratio

## Reserve Governance

`create-denom` is a governance proposal that must be submitted to initialize a new denom
`update-minting-ratio` is a governance proposal that updates the minting ratio for a denom
`update-liquidation-ratio` is a governance proposal that updates the liquidation ratio for a denom

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