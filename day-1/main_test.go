package main

import (
	"reflect"
	"testing"
)

func TestElfCalorieParseAndSum(t *testing.T) {
	resultArr := ElfCalorieParseAndSum("test-data.txt")
	expectedArr := []int{11000, 6000, 4000}

	if !reflect.DeepEqual(resultArr, expectedArr) {
		t.Fatalf("unexpected result!")
	}
}
