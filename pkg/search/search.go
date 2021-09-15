package search

import (
	"context"
	"log"
	"os"
	"strings"
)

type Result struct {
	Phrase  string
	Line    string
	LineNum int64
	ColNum  int64
}

func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result, 10000)
	for _, file := range files {
		result, err := lines(file, phrase)
		if err != nil {
			log.Println(err)
			return nil
		}
		if len(result) > 0 {
			ch <- result
		}
	}

	close(ch)
	return ch
}

func lines(file, phrase string) ([]Result, error) {
	result := []Result{}

	data, err := os.ReadFile(file)
	if err != nil {
		d, _ := os.Getwd()
		log.Print(d, err)
	}
	s := string(data)
	lines := strings.Split(s, "\n")

	for i, nLine := range lines {
		col := strings.Index(nLine, phrase)
		if col != -1 {
			result = append(result, Result{
				Phrase:  phrase,
				Line:    nLine,
				LineNum: int64(i + 1),
				ColNum:  int64(col + 1)})
		}
	}

	return result, nil
}