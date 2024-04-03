# market
**market** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://starport.com).

Note: in order to regenerate proto files, install Ignite CLI v0.23.0 (do not use a later version until we update past SDK v0.45), `go mod tidy`, `rm -rf vendor` (you actually have to remove it for some reason), run `ignite generate proto-go`, then `go mod tidy` after that to clean up

## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.com).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Devnet Linux Install

First download binary
```curl https://github.com/pendulum-labs/market/releases/download/v0.0.4dev/marketd```
Second make the binary executable
```chmod +x marketd```
Third move to /usr/local/bin
```cp marketd /usr/local/bin```


## Learn more

- [Starport](https://starport.com)
- [Tutorials](https://docs.starport.com/guide)
- [Starport docs](https://docs.starport.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/H6wGTY8sxw)
