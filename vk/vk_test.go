package vk

import (
	"os"
	"testing"
)

const FileName = "topic_id.dat"

func TestCheckComment(t *testing.T) {
	os.Remove(FileName)
	teamName := "Койнфлип"
	want := "https://vk.com/topic-173798358_49268096?post=2245"
	comment, status := CheckComment(3, FileName, teamName)

	if comment != want || status != true {
		t.Fatalf(`CheckComment(3, %v, %q) = %q, %v, 
want match for %q, true`, FileName, teamName, comment, status, want)
	}
}
