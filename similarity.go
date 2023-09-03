package main

import (
	"math"
)

// Calculate the dot product of two vectors
func dotProduct(a, b []GenreSummary) float64 {
	var sum float64
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i].GenreName == b[j].GenreName {
				sum += float64(a[i].Count * b[j].Count)
			}
		}
	}
	return sum
}

// Calculate the magnitude of a vector
func magnitude(a []GenreSummary) float64 {
	var sum float64
	for _, v := range a {
		sum += float64(v.Count * v.Count)
	}
	return math.Sqrt(sum)
}

// Calculate the cosine similarity between two vectors
func CosineSimilarity(a, b []GenreSummary) float64 {
	return dotProduct(a, b) / (magnitude(a) * magnitude(b))
}
