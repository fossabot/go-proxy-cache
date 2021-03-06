// +build all unit

package slice_test

//                                                                         __
// .-----.-----.______.-----.----.-----.--.--.--.--.______.----.---.-.----|  |--.-----.
// |  _  |  _  |______|  _  |   _|  _  |_   _|  |  |______|  __|  _  |  __|     |  -__|
// |___  |_____|      |   __|__| |_____|__.__|___  |      |____|___._|____|__|__|_____|
// |_____|            |__|                   |_____|
//
// Copyright (c) 2020 Fabio Cicerchia. https://fabiocicerchia.it. MIT License
// Repo: https://github.com/fabiocicerchia/go-proxy-cache

import (
	"testing"

	"github.com/fabiocicerchia/go-proxy-cache/utils/slice"
	"github.com/stretchr/testify/assert"
)

// --- ContainsInt

func TestContainsIntEmpty(t *testing.T) {
	match := slice.ContainsInt([]int{}, 1)

	assert.False(t, match)
}

func TestContainsIntNoValue(t *testing.T) {
	match := slice.ContainsInt([]int{1, 2, 3}, 4)

	assert.False(t, match)
}

func TestContainsIntValue(t *testing.T) {
	match := slice.ContainsInt([]int{1, 2, 3}, 3)

	assert.True(t, match)
}

// --- ContainsString

func TestContainsStringEmpty(t *testing.T) {
	match := slice.ContainsString([]string{}, "d")

	assert.False(t, match)
}

func TestContainsStringNoValue(t *testing.T) {
	match := slice.ContainsString([]string{"a", "b", "c"}, "d")

	assert.False(t, match)
}

func TestContainsStringValue(t *testing.T) {
	match := slice.ContainsString([]string{"a", "b", "c"}, "c")

	assert.True(t, match)
}

// --- Unique

func TestUniqueEmpty(t *testing.T) {
	input := []string{}
	value := slice.Unique(input)

	assert.Equal(t, []string{}, value)
}

func TestUniqueOneElement(t *testing.T) {
	input := []string{"a"}
	value := slice.Unique(input)

	assert.Equal(t, []string{"a"}, value)
}

func TestUniqueTwoElements(t *testing.T) {
	input := []string{"a", "b"}
	value := slice.Unique(input)

	assert.Equal(t, []string{"a", "b"}, value)
}

func TestUniqueTwoElementsWithDuplicates(t *testing.T) {
	input := []string{"a", "b", "c", "b", "a"}
	value := slice.Unique(input)

	assert.Equal(t, []string{"a", "b", "c"}, value)
}

// --- LenSliceBytes

func TestLenSliceByteEmpty(t *testing.T) {
	input := [][]byte{}
	value := slice.LenSliceBytes(input)

	assert.Equal(t, 0, value)
}

func TestLenSliceBytesOneItem(t *testing.T) {
	input := make([][]byte, 0)
	input = append(input, []byte("testing"))

	value := slice.LenSliceBytes(input)

	assert.Equal(t, 7, value)
}

func TestLenSliceBytesTwosItems(t *testing.T) {
	input := make([][]byte, 1)
	input = append(input, []byte("testing"))
	input = append(input, []byte("sample"))

	value := slice.LenSliceBytes(input)

	assert.Equal(t, 13, value)
}
