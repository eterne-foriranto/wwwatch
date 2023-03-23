package vk

import "testing"

func TestCheckComment(t *testing.T) {
	teamName := "Койнфлип"
	want := "Койнфлип"
	comment, status := CheckComment(3, "topic_id.dat", teamName)

	if comment != want || status != true {
		t.Fatalf(`CheckComment(3, "topic_id.dat", %q) = %q, %v, 
want match for %q, true`, teamName, comment, status, want)
	}
}
