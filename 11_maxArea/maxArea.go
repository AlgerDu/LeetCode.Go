package leetcode

func maxArea(height []int) int {

	max := 0
	index := make([]int,len(height))
	indexCount := 1;

	for i := 1; i < len(height) ; i++{
		for j:=0; j < indexCount; j++{
			area := 0

			if height[j] < height[i] {
				area = height[j]
			}else{
				area = height[i]
			}

			area = area * ( i -index[j] )

			if area > max {
				max = area
			}
		}

	 	if height[i] > height[indexCount -1]{
			 height[indexCount] = height[i]
			 index[indexCount] = i
			 indexCount++
		 }
	}

	return max
}