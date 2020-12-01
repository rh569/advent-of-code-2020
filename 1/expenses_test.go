package main

import "testing"

func TestFindEntries(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	want1, want2 := 1721, 299

	err, found1, found2 := findEntries(2020, entries)

	if err != nil {
		t.Fatalf("Error was encountered")
	}

	if found1 != want1 && found1 != want2 {
		t.Fatalf("Found1: %v, was not want1: %v or want2: %v", found1, want1, want2)
	}

	if found2 != want1 && found2 != want2 {
		t.Fatalf("Found2: %v, was not want1: %v or want2: %v", found2, want1, want2)
	}

	if want1 != want2 && found1 == found2 {
		t.Fatalf("found1: %v, matched found2: %v, when want1: %v and want2: %v", found1, found2, want1, want2)
	}
}

func TestFindThreeEntries(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	want := []int{366, 675, 979}

	err, found := findThreeEntries(2020, entries)

	if err != nil {
		t.Fatalf("Error was encountered")
	}

	if len(want) != len(found) {
		t.Fatalf("Number found: %v, wanted %v", len(found), len(want))
	}

	for i := 0; i < len(want); i++ {

		if want[i] != found[i] {
			t.Fatalf("found: %v, was not wanted: %v", found[i], want[i])
		}
	}
}
