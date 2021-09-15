package search

import (
	"context"
	"log"
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

func TestAny_found(t *testing.T) {
	ch := Any(context.Background(), "HTTP", []string{"1.txt", "2.txt", "3.txt"})
	result, ok := <-ch
	if !ok {
		log.Fatal("Phrase: ", result.Phrase, 
		", Line: ", result.Line, 
		", LineNum: ", result.LineNum, 
		", ColNum: ", result.ColNum)
	}
}

func TestAny_notFound(t *testing.T) {
	ch := Any(context.Background(), "_HTTP", []string{"1.txt", "2.txt", "3.txt"})
	result, ok := <-ch
	if ok {
		log.Fatal("Phrase: ", result.Phrase, 
		", Line: ", result.Line, 
		", LineNum: ", result.LineNum, 
		", ColNum: ", result.ColNum)
	}
}
