package set

func isAnagram(s string, t string) bool {
	hash := make([]rune, 'z'-'a'+1)
	for _, v := range s {
		hash[v-'a']++
	}
	for _, v := range t {
		hash[v-'a']--
	}
	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}
