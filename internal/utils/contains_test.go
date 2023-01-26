package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindString(t *testing.T) {
	slice := []string{"string1", "string2", "string3"}
	expected := "string3"

	assert.Equal(t, true, Contains(slice, expected))
}

func TestStringNotFound(t *testing.T) {
	slice := []string{"string1", "string2", "string3"}
	expected := "string4"

	assert.Equal(t, false, Contains(slice, expected))
}

func TestGivenEmptyStringSlice(t *testing.T) {
	slice := []string{}
	expected := "string3"

	assert.Equal(t, false, Contains(slice, expected))
}

func TestGivenEmptyString(t *testing.T) {
	slice := []string{"string1", "string2", "string3"}
	expected := ""

	assert.Equal(t, false, Contains(slice, expected))
}
