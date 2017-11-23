# Pine <a target="_blank" href="https://opensource.org/licenses/MIT" title="License: MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg"></a> <a target="_blank" href="http://makeapullrequest.com" title="PRs Welcome"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg"></a>

Pine provides a way to track coin prices as well overall holding values, all in your console.

The data comes from [coinmarketcap.com](https://coinmarketcap.com/).

## Install

```
go get -u github.com/aranair/pine
```

## Configuration

- Update `configs.yaml` with the coins you want to track

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
```

## Usage

```
go run main.go
```


## License
MIT
