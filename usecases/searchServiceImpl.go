package usecases

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

// SearchServiceImpl is an implementation of the SearchService
type SearchServiceImpl struct {
	wordRegexp *regexp.Regexp
	termCounts map[string]float64
	corpus     map[string]SearchObject
}

// NewSearchServiceImpl returns a new instance of SearchServiceImpl
func NewSearchServiceImpl() *SearchServiceImpl {
	return &SearchServiceImpl{
		wordRegexp: regexp.MustCompile(`\w+`),
		termCounts: make(map[string]float64),
		corpus:     make(map[string]SearchObject),
	}
}

// Index adds the given object to the search index
func (s *SearchServiceImpl) Index(obj SearchObject) {
	// get term counts and add them to term counts list
	terms := s.getTermCounts(obj.Content)
	for term := range terms {
		s.termCounts[term] += 1.0
	}

	// add search object to document corpus
	obj.Terms = terms
	s.corpus[obj.ID] = obj
}

// Search returns the list of relevant results
func (s *SearchServiceImpl) Search(objtype string, query string) []SearchObject {
	queryTerms := s.wordRegexp.FindAllString(strings.ToLower(query), -1)

	results := make([]SearchObject, 0)
	for _, doc := range s.corpus {
		score := 0.0
		if doc.Type == objtype {
			for _, queryTerm := range queryTerms {
				if docTermFreq, ok := doc.Terms[queryTerm]; ok {
					score += docTermFreq * s.getIDF(queryTerm)
				}
			}

			// add to result if score is above cutoff
			if score > ScoreCutoff {
				doc.Score = score
				results = append(results, doc)
			}
		}
	}

	// sort search results by score
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	return results
}

// getTermCounts extracts the query terms into a list and gives them normalized weights
func (s *SearchServiceImpl) getTermCounts(query string) map[string]float64 {
	terms := s.wordRegexp.FindAllString(strings.ToLower(query), -1)

	// build list of terms and their counts
	termCounts := make(map[string]float64)
	for _, term := range terms {
		termCounts[term] += 1.0
	}

	// normalise term count list
	length := float64(len(terms))
	for key, val := range termCounts {
		termCounts[key] = val / length
	}

	return termCounts
}

// getIDF gets the inverse of the number of documents with the term in it
func (s *SearchServiceImpl) getIDF(term string) float64 {
	docCount := 0.0
	for _, doc := range s.corpus {
		for docTerm := range doc.Terms {
			if term == docTerm {
				docCount += 1.0
			}
		}
	}
	return math.Log(float64(len(s.corpus)) / (docCount + 1.0))
}
