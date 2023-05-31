package main

import (
	"os"
	"testing"
)

const TestFileName = "topic_id.dat"

func TestCheckComment(t *testing.T) {
	os.Remove(TestFileName)
	teamName := "Койнфлип"
	want := "https://vk.com/topic-173798358_49268096?post=2245"
	comment, status := check(3, TestFileName, teamName)

	if comment != want || status != true {
		t.Fatalf(`check(3, %v, %q) = %q, %v, 
want match for %q, true`, TestFileName, teamName, comment, status, want)
	}
}
