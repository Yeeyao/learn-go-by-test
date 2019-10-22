package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
		{ "Name": "Cleo", "Wins": 10 },
		{ "Name": "Chris", "Wins": 33 }]`)

		store := FileSystemStore{database}

		got := store.GetPlayerScore("Chris")

		want := 33

		assertScoreEquals(t, got, want)

	})

	t.Run("/get player socre", func(t *testing.T) {
		database := strings.NewReader(`[
		{ "Name": "Cleo", "Wins": 10 },
		{ "Name": "Chris", "Wins": 33 }]`)

		store := FileSystemStore{database}

		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})
}

func assertLeague(t *testing.T, got, want []Player) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertScoreEquals(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
