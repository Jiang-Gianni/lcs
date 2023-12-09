# lcs
LeetCode Scraper

## Installation

```bash
go install github.com/Jiang-Gianni/lcs@latest
```

## Usage

```
lcs scrape && lcs gen -f 0 -t 10t
```

`lcs scrape` gets the questions data from LeetCode API https://leetcode.com/graphql/ and it stores it in a local db (`store.db`) using sqlite.

It only scrapes a problem page set for each run (maximum 50 questions). Use `-s` to skip a number of questions that is multiple of 50 (example `lcs scrape -s 100` will scrape questions from 100 to 150).

Premium questions are skipped.

`lcs get` generates the file of the stored question that have an id between the values set by the `-f` (from) and `-t` (to) flags: `lcs gen -f 0 -t 10t` will consider questions with an ID between 0 and 10.

As of now `golang` is the only supported language and SQL only questions are not generated.