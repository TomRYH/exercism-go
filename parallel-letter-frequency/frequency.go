package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency makes concurrent calls to Frequency to count
// the frequency of words in a series of texts concurrently, it combines
// each and returns a single FreqMap
func ConcurrentFrequency(texts []string) FreqMap {

	ch := make(chan FreqMap)

	for _, t := range texts {
		go func(t string) {
			ch <- Frequency(t)
		}(t)
	}

	freq := FreqMap{}
	for range texts {
		for k, v := range <-ch {
			freq[k] += v
		}
	}

	return freq
}
