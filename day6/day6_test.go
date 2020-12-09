package day6_test

import (
	"HSteffensen/AoC2020/day6"
	"reflect"
	"testing"
)

func TestGroups(t *testing.T) {
	groups := day6.Groups(day6.ExampleInput)
	expected := [][]string{{"abc"}, {"a", "b", "c"}, {"ab", "ac"}, {"a", "a", "a", "a"}, {"b"}}
	t.Log(groups)
	t.Log(expected)
	if !reflect.DeepEqual(groups, expected) {
		t.Error("Incorrect group parsing")
	}
}

func TestGroupAnswers(t *testing.T) {
	groups := day6.Groups(day6.ExampleInput)
	expected := []string{"abc", "abc", "abc", "a", "b"}
	for i, group := range groups {
		if answer := day6.GroupAnswers(group); answer != expected[i] {
			t.Errorf("Incorrect group answers: expected %#v, got %#v", expected[i], answer)
		}
	}
}
