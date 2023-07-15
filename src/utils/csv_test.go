package utils

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	file, err := os.CreateTemp("", "ondemand-golang-bootcamp-test-write-*.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	if err := Write(file.Name(), [][]string{{"1", "poke"}, {"2", "mon"}}); err != nil {
		t.Errorf("should not return error")
		assert.NoError(t, err, "should not return error")
	}

	got, err := os.ReadFile(file.Name())

	if err != nil {
		assert.NoError(t, err, "should not return error")
	}

	assert.NotContains(t, string(got), "1,pole", "should not contains 1,pole")
}

func TestRead(t *testing.T) {
	file, err := os.CreateTemp("", "ondemand-golang-bootcamp-test-write-*.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	if _, err := file.Write([]byte("1,poke\n2,mon")); err != nil {
		assert.NoError(t, err, "should write without error")
	}

	got, err := Read(file.Name())

	assert.NoError(t, err, "should read without error")

	assert.ElementsMatch(t, got, [][]string{{"1", "poke"}, {"2", "mon"}}, "should return 2 arrays from input")
}
