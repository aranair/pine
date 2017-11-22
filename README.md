# Pine <a target="_blank" href="https://opensource.org/licenses/MIT" title="License: MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg"></a> <a target="_blank" href="http://makeapullrequest.com" title="PRs Welcome"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg"></a>

## What is this?

Pine provides a way to look at coin prices in your console. All the data comes from [coinmarketcap.com](https://coinmarketcap.com/).

## Install

```
go get -u github.com/aranair/pine
```

## Usage

Update configs.yaml with some tickers, then:

```
go run main.go
```

## Configuration

1. `cp configs.toml.sample configs.toml`
2. Add or remove coins that you want to track
3. If you want to keep track of the price you entered at, feel free to add that in as well. (WIP)

## License
MIT
