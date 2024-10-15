//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    ch chan *Tweet
)

func producer(stream Stream) {
    for {
        tweet, err := stream.Next()
        if err == ErrEOF {
            close(ch)
            return
        }
        ch <- tweet
    }
}

func consumer(wg *sync.WaitGroup, results *[]string) {
    defer wg.Done()
    for t := range ch {
        if t.IsTalkingAboutGo() {
            *results = append(*results, fmt.Sprint(t.Username, "\ttweets about golang"))
        } else {
            *results = append(*results, fmt.Sprint(t.Username, "\tdoes not tweet about golang"))
        }
    }
}

func processStreams(stream *Stream) (*[]string, time.Duration) {
    start := time.Now()
    var wg sync.WaitGroup
    ch = make(chan *Tweet)
    results := make([]string, 0)
    wg.Add(1)
    go producer(*stream)
    go consumer(&wg, &results)
    wg.Wait()
    return &results, time.Since(start)
}

func main() {
    stream := GetMockStream()

    // process streams and get result.
    results, runtime := processStreams(&stream)
    for _, result := range *results {
        fmt.Println(result)
    }
    fmt.Println("Process took:", runtime)
}
