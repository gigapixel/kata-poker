
package main

import "testing"
import "github.com/stretchr/testify/assert"

type Person struct {
	name string
}

func TestFrame1DoubleStrike(t *testing.T) {
	scores := [...] string{
		"X", "-",
		"X", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-", "-", 
	}
	result := board(scores);
	expected := int64(30);
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 30")
}
