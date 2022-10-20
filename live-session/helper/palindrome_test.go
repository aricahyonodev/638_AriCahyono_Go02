package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccessPalindrome(t *testing.T) {
	result := isPalindrome("kasur ini rusak")
	require.Equal(t, false, result, "Result isPalindrome")
}

func TestFailPalindrome(t *testing.T) {
	result := isPalindrome("kasu222r")
	require.Equal(t, false, result, "Result isPalindrome")
}