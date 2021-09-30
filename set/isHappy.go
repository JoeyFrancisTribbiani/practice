package set

func IsHappy(n int) bool {
	hash := make(map[int]int)
	sum := 0
	for {
		sum = 0
		for n > 0 {
			t := n % 10
			sum += t * t
			n = n / 10
		}
		_, ok := hash[sum]
		if ok {
			return false
		}
		if sum == 1 {
			return true
		} else {
			n = sum
			hash[sum] = 1
		}
	}
}
