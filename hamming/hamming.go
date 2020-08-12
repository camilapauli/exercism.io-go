package hamming

import "errors"

///Distance returns the distance between two DNA strands
func Distance(firstStrand, secondStrand string) (int, error) {
	if len(firstStrand) != len(secondStrand) {
		return 0, errors.New("different strands size")
	}
	differences := 0
	for i := 0; i < len(firstStrand); i++ {
		if firstStrand[i] != secondStrand[i] {
			differences++
		}
	}
	return differences, nil
}
