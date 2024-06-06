package solution

// https://leetcode.com/problems/container-with-most-water/description/

// two-pointers
func MaxArea(height []int) int {
	var area int = 0
	lenHeight := len(height)
	if lenHeight < 2 {
		return area
	}

	var higherIndex = 0
	// var lowerIndex = 0
	for i := range height {
		var tmpArea int
		if height[i] <= height[higherIndex] {
			tmpArea = (i - higherIndex) * height[i]
			if tmpArea > area {
				// lowerIndex = i
				area = tmpArea
			}
		} else {
			tmpArea = (i - higherIndex) * height[higherIndex]
			if tmpArea > area {
				higherIndex = i
				area = tmpArea
			}
		}
	}

	return area
}
