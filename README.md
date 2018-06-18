# HeyRussia

`heyrussia` is a command line interface tool for fetching updates from The 2018 FIFA World Cup in Russia. You can fetch groups, teams, standings, stadiums and feed from all pasta and upcoming matches. 

<!-- TOC -->

- [Heyrussia](#heyrussia)
  - [Usage](#usage)
  - [Install](#install)
  - [Licence](#licence)
  - [Author](#author)

<!-- /TOC -->

## Usage
To list all groups in the world cup, use `groups` or `g` command:

```bash
$ heyrussia groups
```

To list all teams in the world cup, use `teams` or `t` command:

```bash
$ heyrussia teams
```

To list all stadiums in the world cup, use `stadiums` or `std` command:

```bash
$ heyrussia stadiums
```

To get current groups standings in the world cup, use `standings` or `s` command:

```bash
$ heyrussia standings
```

To fetch feed with details of all past and upcoming matches, including scores, dates, scorers etc ,
 use `feed` or `f` command:

```bash
$ heyrussia feed
```

You can filter a feed based on dates with the `date` or `-d` flag.:

```bash
$ heyrussia feed -d 2018-06-18
```

Date format is yyyy-mm-dd


You can use the special 'today' to fetch feed for the current day.

```bash
$ heyrussia feed -d today
```

More features to be available soon.

## Install

To install, use `go get`:

```bash
$ go get -d github.com/geeknat/heyrussia
```

or for 32 bit Windows , [download here](https://github.com/geeknat/heyrussia/releases/heyrussia_windows_386.exe).

or for 64 bit Windows , [download here](https://github.com/geeknat/heyrussia/releases/heyrussia_windows_amd64.exe).

or for 32 bit Linux , [download here](https://github.com/geeknat/heyrussia/releases/heyrussia_linux_386).

or for 64 bit Linux , [download here](https://github.com/geeknat/heyrussia/releases/heyrussia_linux_amd64).

or for RISC Linux , [download here](https://github.com/geeknat/heyrussia/releases/heyrussia_linux_arm).

or for 32 bit Mac OS, [download here](https://github.com/geeknat/heyrussia/releases/heyrussia_darwin_386).

or for 64 bit Mac OS, [download here](https://github.com/geeknat/heyrussia/releases/heyrussia_darwin_amd64).

or download for any platform from the [releases page](https://github.com/geeknat/heyrussia/releases).

## Licence

[MIT](https://choosealicense.com/licenses/mit/)

## Author

[Geek Nat](https://geeknat.com)

