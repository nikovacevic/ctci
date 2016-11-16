package array

// RotateMatrix (1.7)
// Given an image represented by an NxN matrix, where each pixel in the image is
// 4 bytes, write a function to rotate the image by 90 degrees. In place?
func RotateMatrix(in [][]int) [][]int {
	// Assume we're given NxN matrix (not verifying)
	l := len(in)
	x, y := l/2, l/2
	if l%2 != 0 {
		x++
	}
	l--
	for a := 0; a < x; a++ {
		for b := 0; b < y; b++ {
			// Visualize each rotation
			// fmt.Printf("%v -> %v -> %v -> %v\n", in[0+a][0+b], in[0+b][l-a], in[l-a][l-b], in[l-b][0+a])
			t := in[0+a][0+b]
			in[0+a][0+b] = in[l-b][0+a]
			in[l-b][0+a] = in[l-a][l-b]
			in[l-a][l-b] = in[0+b][l-a]
			in[0+b][l-a] = t
		}
	}
	return in
}

// 1.8 Zero Matrix
// Write an alorithm such that if an element in an MxN matrix is 0, its entire
// row and columns are set to 0.
