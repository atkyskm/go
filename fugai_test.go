package main

import "testing"

func TestTestscore(t *testing.T) {
    // if !testscore(80) {
    //     t.Errorf("Expected true for score 80")
    // }
    if testscore(50) {
        t.Errorf("Expected false for score 50")
    }
}