package main

import (
    "fmt"
    "testing"
    "time"
)

func TestProcessStreams(t *testing.T) {
    stream := GetMockStream()
    consumerResults, duration := processStreams(&stream)
    if len(*consumerResults) != len(mockdata) {
        t.Errorf("Expected %d results, got %d", len(mockdata), len(*consumerResults))
    }
    for i := 0; i < len(mockdata); i++ {
        var subject string
        if mockdata[i].IsTalkingAboutGo() {
            subject = "tweets about golang"
        } else {
            subject = "does not tweet about golang"
        }
        expected := fmt.Sprintf("%s\t%s", mockdata[i].Username, subject)
        if (*consumerResults)[i] != expected {
            t.Errorf("Expected %s, got %s", expected, (*consumerResults)[i])
        }
    }
    if duration > 2*time.Second {
        t.Errorf("Expected duration to be less than 2ms, got %v", duration)
    }
}
