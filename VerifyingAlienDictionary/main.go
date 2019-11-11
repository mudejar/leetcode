package main

func main() {

}

func isAlienSorted(words []string, order string) bool {
	indeces := make([]int, 123)
	for i, c := range order {
		indeces[c] = i
	}

	less := func(i, j int) bool {
		si, sj := len(words[i]), len(words[j])
		for k := 0; k < si && k < sj; k++ {
			ii, ij := indeces[words[i][k]], indeces[words[j][k]]
			switch {
			case ii < ij:
				return true
			case ii > ij:
				return false
			}
		}
		return si <= sj
	}

	for i := 1; i < len(words); i++ {
		if !less(i-1, i) {
			return false
		}
	}

	return true
}