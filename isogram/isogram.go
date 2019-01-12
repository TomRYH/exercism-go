/*Package isogram contains functions about
isograms. An isogram is a word that does not
contain any repeating letters.*/
package isogram

import "strings"

// IsIsogram checks whether a word is an isogram
func IsIsogram(word string) bool {

	r := strings.NewReplacer("-", "", " ", "")
	word = r.Replace(strings.ToUpper(word))

	checkStr := ""

	for _, l := range word {
		if strings.Contains(checkStr, string(l)) {
			return false
		}
		checkStr += string(l)

	}

	return true

}
