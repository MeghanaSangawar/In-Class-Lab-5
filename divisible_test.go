package main

import "testing"
import "fmt"

func TestIsDivisibleByTwo(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected bool
    }{
        {"Even number", 4, true},
        {"Odd number", 3, false},
        {"Zero", 0, true},
        {"Negative even number", -8, true},
        {"Negative odd number", -7, false},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            result := IsDivisibleByTwo(tc.input)
            fmt.Printf("Testing %d: got %v, expected %v\n", tc.input, result, tc.expected)
            if result != tc.expected {
                t.Errorf("IsDivisibleByTwo(%d) = %v; expected %v", tc.input, result, tc.expected)
            }
        })
    }
}
