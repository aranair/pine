# Pine <a target="_blank" href="https://opensource.org/licenses/MIT" title="License: MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg"></a> <a target="_blank" href="http://makeapullrequest.com" title="PRs Welcome"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg"></a>

Pine provides a way to not only track a list of coin prices, but also your portfolio values. All in your console!

The data comes from [coinmarketcap.com](https://coinmarketcap.com/).

## Screenshot

![Demo](https://github.com/aranair/pine/blob/master/demo.png?raw=true "Demo")

## Building from code

```
go get -u github.com/aranair/pine
dep ensure
./scripts/build.sh
```

## Configuration

- Update `configs.yaml` with the coins you want to track
- If you want to track the `Gain`, `% Gain` and `Overall Profit`, fill in the Cost and Unit.
- Currency is still WIP

Sample:

```yaml
coins:
  - ticker: "bitcoin"
    cost: 1000
    units: 5
  - ticker: "ethereum"
    cost: 100
    units: 1
  - ticker: "bitcoin-cash"
    cost: 100
    units: 1

    # If no cost/units provided, it'll just track prices, with no portfolio tracking
  - ticker: "neo"
  - ticker: "ethereum-classic"
```

## Usage

If you're on mac, just run the binary that I've built.

```
pine
```

```
go run main.go
```


It automatically refreshes every minute, but if you want to refresh it manually, press r.

## License
MIT
