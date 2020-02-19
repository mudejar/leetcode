package main

func main() {

}

func removeVowels(S string) string {
	// set with vowels
	vowelSet := map[byte]struct{}{
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
	}

	for i := 1; i < len(S)-1; i++ {
		_, ok := vowelSet[S[i]]
		if ok {
			S = S[:i] + S[i+1:]
			i--
		}
	}

	_, ok := vowelSet[S[0]]
	if ok {
		if len(S) > 1 {
			S = S[1:]
		} else {
			S = ""
		}
	}

	_, ok = vowelSet[S[len(S)-1]]
	if ok {
		if len(S) > 1 {
			S = S[:len(S)-1]
		} else {
			S = ""
		}
	}

	return S
}
