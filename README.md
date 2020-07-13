# Riverboat

![Go](https://github.com/alexclewontin/riverboat/workflows/Go/badge.svg)

A full-service Go Texas hold'em library, featuring an ultra-fast poker hand evaluation module.

![Str. Toronto, a riverboat that was not a casino, on the St. Lawrence River](https://cdn.loc.gov/service/pnp/det/4a30000/4a31000/4a31700/4a31769v.jpg)


## Table of Contents

- [Riverboat](#riverboat)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [How-To](#how-to)
    - [Installation](#installation)
    - [Usage](#usage)
  - [Documentation](#documentation)
  - [Contributing](#contributing)
  - [Acknowledgements](#acknowledgements)
  - [License](#license)

## Features

Riverboat plays No-limit Texas Hold'em. It's a full-service game-management library including:

- **Illegal move rejection** - disallows illegal plays, including moves out of turn, or raises that don't meet the minimum
- **Information hiding** - calculates exactly what any player can see at any moment, including edge cases like when all players are all-in
- **Winner determination and pot allocation** - correctly allocates pots at the end, including arbitrary numbers of sidepots and splits
- **Configurable** - buy-in limits and blinds can be set on a game-by-game basis

Riverboat also includes an evaluation submodule, which as of July 2020 is the *[fastest](./eval#benchmarks)* 5-, 6-, and 7-card poker hand evaluator in Go on Github.


## How-To

### Installation

To install:

```shell
$ go get github.com/alexclewontin/riverboat
```

To use in your project:

```go
import (
    // For the base library:
    "github.com/alexclewontin/riverboat"

    // For direct access to the hand evaluation module:
    "github.com/alexclewontin/riverboat/eval"
)
```

### Usage

Create a game:
```go
    g := riverboat.NewGame()
```

Add players, buy-in, ready up:
```go
    pNum := g.AddPlayer()

    err = riverboat.BuyIn(g, pNum, 1000)
    if err != nil {
        fmt.Println(err)
    }

    err = riverboat.ToggleReady(g, pNum, 0)
    // ... check for error
```

Play:

```go
    err = riverboat.Deal(g, pNum, 0)
    // ... check for error
    err = riverboat.Bet(g, pNum, 25)
    // ... check for error
    err = riverboat.Fold(g, pNum, 0)
    // ... check for error
```

Inspect the game's state:

```go
    // This doesn't contain opponents' cards, or the contents of the deck
    playerView := g.GenerateView(pNum)

    // This is the entire game state, for easy serialization and storage in a persistence layer
    godView := g.GenerateOmniView(pNum)
```
## Documentation

Full documentation can be found [here](https://pkg.go.dev/github.com/alexclewontin/riverboat) (TODO: flesh this out).

## Contributing

Contributions to Riverboat are more than welcome!

- If the contribution is a minor fix, go ahead and open a PR.
- If the contribution is larger (e.g. supporting an additional ruleset or variation like Omaha or limit hold'em) open an issue to coordinate development efforts.

## Acknowledgements

TODO: Acknowledgements

## License

Riverboat and its component pieces are provided under the **[BSD-2-Clause license](http://opensource.org/licenses/mit-license.php)** except where otherwise specified. 

Copyright 2020 © Alex Lewontin.