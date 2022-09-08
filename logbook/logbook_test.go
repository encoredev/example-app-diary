package logbook

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestListEntries_OkWithNoRecords(t *testing.T) {
	resp, err := ListEntries(context.Background(), "1993-01-01")
	if err != nil {
		t.Fatal(err)
	}

	wantCount := 0
	if !reflect.DeepEqual(len(resp.Entries), wantCount) {
		t.Fatalf("got len(resp.Entries) %v, want %v", len(resp.Entries), wantCount)
	}
}

func TestListEntries_OkWithOneRecord(t *testing.T) {
	TestCreateEntry(t)

	resp, err := ListEntries(context.Background(), time.Now().Format("2006-01-02"))
	if err != nil {
		t.Fatal(err)
	}

	wantCount := 1
	if !reflect.DeepEqual(len(resp.Entries), wantCount) {
		t.Fatalf("got len(resp.Entries) %v, want %v", len(resp.Entries), wantCount)
	}

	firstEntry := resp.Entries[0]
	want := "I built this amazing demo using Encore"
	if !reflect.DeepEqual(firstEntry.Text, want) {
		t.Fatalf("got resp.Entries[0].Text %v, want %v", firstEntry.Text, want)
	}
}

func TestCreateEntry(t *testing.T) {
	resp, err := CreateEntry(context.Background(), &CreateParams{Text: "I built this amazing demo using Encore"})
	if err != nil {
		t.Fatal(err)
	}

	want := "I built this amazing demo using Encore"
	if !reflect.DeepEqual(resp.Text, want) {
		t.Fatalf("got resp %v, want %v", resp.Text, want)
	}
}
