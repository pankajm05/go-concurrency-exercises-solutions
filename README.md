# Go Concurrency Exercises [![Build Status](https://travis-ci.org/loong/go-concurrency-exercises.svg?branch=main)](https://travis-ci.org/loong/go-concurrency-exercises) [![Go Report Card](https://goreportcard.com/badge/github.com/loong/go-concurrency-exercises)](https://goreportcard.com/report/github.com/loong/go-concurrency-exercises)
My Solutions to the exercises for Golang's concurrency patterns by [loong](https://github.com/loong).

## Why
Golang is very popular for its built-in concurrency support. Recently, I started learning about concurrency in Go.
The best way to learn is by solving challenges/problems so this is my attempt to do the same.

![Image of excited gopher](https://golang.org/doc/gopher/pkg.png)

## My solutions

### Problem 0: [Limit your Crawler](https://github.com/pankajm05/go-concurrency-exercises-solutions/blob/main/0-limit-crawler/main.go)
- The problem is quite straightforward, where we have a recursive crawler, and we want to add delay before visiting the next URL.
- I like to think of this as a dfs tree where we can add delay before visiting neighbours.
- This would be a pre-order traversal where processing a node would mean calling the URL on the same.
- To solve this I would use a ticker with the duration of 10 second and receive from the channel before visiting neighbour.

Execution flow:
```
dfs(url):
    process(url) // call URL and fetch children
    for child in children:
        <-tick
        dfs(child)
```
