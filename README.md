# advent-of-code-go

Resolving all [Advent of Code](https://adventofcode.com/) challenges in GO.

<div align="center">

![](https://img.shields.io/badge/2015-14_days-blue)
![](https://img.shields.io/badge/2016-0_days-red)
![](https://img.shields.io/badge/2017-0_days-red)
![](https://img.shields.io/badge/2018-0_days-red)
![](https://img.shields.io/badge/2019-0_days-red)
![](https://img.shields.io/badge/2020-0_days-red)
![](https://img.shields.io/badge/2021-0_days-red)
![](https://img.shields.io/badge/2022-0_days-red)
![](https://img.shields.io/badge/2023-14_days-blue)


![](https://img.shields.io/badge/stars%20⭐-56-yellow) 
![](https://img.shields.io/badge/days%20completed-28-green)

<img src="./assets/gopher.png" width="170" />

</div>

### Prerequisites

Make sure you have Go >= 1.21.4.

### Clone the repository

```sh
git clone git@github.com:teodorpopa/advent-of-code-go.git
```

### Test the solutions

To test the solutions for a specific year/day, you need to run the following command:

```sh
make test year=2023 day=01
```


### Compile and run

To run a solution, you need to provide some arguments.
* the `--year` flag must be set to specify the year of the challenge
* the `--day` flag must be set to specify which day's solution should run
* the `--part` flag specifies if you only want to run a specific part. by default both parts will run

make command:

```sh
make run year=2023 day=13
```

go command:

```sh
go run main.go --year=2023 --day=13
```