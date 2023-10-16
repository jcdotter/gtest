package gtest

import (
	"testing"
)

var config = &Config{
	PrintTest:   true,
	PrintTrace:  true,
	PrintDetail: true,
	Msg:         "%s",
}

func TestAll(t *testing.T) {
	gt := New(t, config)
	gt.Equal(1, 1, "1 == 1")
	gt.NotEqual(1, 2, "1 != 2")
	gt.True(true, "true is true")
	gt.False(false, "false is false")
}

func TestTable(t *testing.T) {
	data := [][]string{
		{"Col1", "Col2", "Col3"},
		{"1", "2", "3"},
		{"4", "5", "6"},
	}
	PrintTable(data, true)
}
