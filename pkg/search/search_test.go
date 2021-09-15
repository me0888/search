package search

import (
	"context"
	"testing"
)

func TestAll_found(t *testing.T) {
	ch := All(context.Background(), "HTTP", []string{"1.txt", "2.txt", "3.txt"})
	results, ok := <-ch
	if !ok {
		t.Errorf("error: %v", ok)
		t.Fatal("result: ", results)
	}
}

func TestAll_notFound(t *testing.T) {
	ch := All(context.Background(), "_HTTP", []string{"1.txt", "2.txt", "3.txt"})
	results, ok := <-ch
	if ok {
		t.Errorf("error: %v", ok)
		t.Fatal("result: ", results)
	}
}


