package piscine

func ContSub(indOfSub [][]int, j int) bool {
	for i := 0; i < len(indOfSub); i++ {
		if indOfSub[i][0] <= j && indOfSub[i][1] > j {
			return true
		}
	}
	return false
}
