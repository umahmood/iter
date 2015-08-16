// Package iter a memory efficient way to iterate over a sequence of integers.
package iter

// N returns a slice of length n. n must be greater than or equal to 0. Apart
// from a slice header does not make any more allocations, regardless the size
// of n.
func N(n int) []struct{} {
	return make([]struct{}, n)
}
