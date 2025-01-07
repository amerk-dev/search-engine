package index

import (
	"searchSystem/internal/models"
	"searchSystem/internal/tokenize"
)

// index is an inverted index. It maps tokens to document IDs.
type Index map[string][]int

// добавление документа в индекс
func (idx Index) Add(docs []models.Document) {
	for _, doc := range docs {
		for _, token := range tokenize.Analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Don't add same ID twice.
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func Intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// поиск
func (idx Index) Search(text string) []int {
	var r []int
	for _, token := range tokenize.Analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = Intersection(r, ids)
			}
		} else {
			// Token doesn't exist.
			return nil
		}
	}
	return r
}
