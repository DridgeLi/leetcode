package solutions

func MaxArea(height []int) int {
	area := 0
	tmpArea := 0
	if len(height) < 2 {
		return area
	}

	for i, j := 0, len(height)-1; i < j; {
		tmpArea = (j - i) * min(height[i], height[j])
		if tmpArea > area {
			area = tmpArea
		} else if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return area
}

/*
1. Make as many as variables change linearly, this will reduce the complexity of the problem.
2. Find the key point of the problem.
	Here area is determined by shorter height, so keep the higher height there, and move the shorter height to find a higher

*/
