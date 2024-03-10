# PDSA - Prisoner's Dilemma Strategy Analyzer

## Overview

The Prisoner's Dilemma Strategy Analyzer is an open-source project designed to explore and analyze strategies within the framework of the Prisoner's Dilemma, a standard example of a game analyzed in game theory that shows why two completely rational individuals might not cooperate, even if it appears that it is in their best interest to do so. Inspired by the comprehensive discussion at [Stanford Encyclopedia of Philosophy](https://plato.stanford.edu/entries/prisoner-dilemma/), this tool allows users to define strategies, run simulations, and generate analytical reports in various formats to determine the best course of action in a competitive environment.

## Features

- Define and compare multiple strategies within the Prisoner's Dilemma context.
- Run simulations with a configurable number of iterations.
- Use different solvers for strategy analysis.
- Generate reports in text, JSON, and HTML formats.

## Installation

Before you can run the Strategy Analyzer, ensure you have Go installed on your system. [Download Go](https://golang.org/dl/) if you haven't already.

## Usage

To run the program, navigate to the project directory in your terminal and execute the following command:

```bash
go run cmd/main.go --config <PATH_TO_CONFIG_FILE>
```

or use an environment variable:

```bash
export CONFIG_PATH=<CONFIG_PATH>
```

The configuration file should be in YAML format. Here's an example of how to structure your configuration:

```yaml
runners:
  - tag: "Unconditional Cooperator vs Random"
    iterations: 50
    solver: smp
    strategies:
      - cu
      - rand
  - tag: "Unconditional Cooperator vs Tip for Tip"
    iterations: 50
    solver: smp
    strategies:
      - cu
      - tft
  - tag: "Random vs Tip for Tip"
    iterations: 50
    solver: smp
    strategies:
      - rand
      - tft

viewers:
  - type: html
    out: ./results/result.html
  - type: text
    out: ./results/result.text
  - type: json
    out: ./results/result.json
```

## Configuring Your Simulations

### Runners
- `tag`: A unique identifier for the simulation.
- `iterations`: The number of times the simulation will run.
- `solver`: The method used for solving the game. Currently, only `smp` (simple method) is available.
- `strategies`: The strategies to compare. Options include `cu`, `rand`, and `tft` (tit for tat).

### Viewers
- `type`: The format of the output report (`html`, `text`, or `json`).
- `out`: The output path for the report file.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- This project was inspired by the article on the Prisoner's Dilemma from the [Stanford Encyclopedia of Philosophy 🎓](https://plato.stanford.edu/entries/prisoner-dilemma/).
- Gratitude is also extended to [Strategies for the Iterated Prisoner's Dilemma](https://plato.stanford.edu/entries/prisoner-dilemma/strategy-table.html) for providing a comprehensive guide on various strategies that can be employed in the iterated version of the game.
- Special thanks to the creators of the following educational videos, which offer insightful explorations of the Prisoner's Dilemma:
  - [Video explaining the Prisoner's Dilemma in English 🇬🇧](https://www.youtube.com/watch?v=mScpHTIi-kM).
  - [Ukrainian translation of the above video, making the content accessible to a wider audience 🇺🇦](https://www.youtube.com/watch?v=HnjVSCnPYCU&t=522s).
- Thanks to all contributors and users of this project for exploring the fascinating dynamics of game theory together.
