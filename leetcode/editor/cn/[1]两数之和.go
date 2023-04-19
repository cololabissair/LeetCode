
//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	var index []int
	for i := 0;i<len(nums);i++{
		for j := i+1;j<len(nums);j++{
			if nums[i] + nums[j] == target{
				index = append(index,i,j)
			}
		}
	}
	return index
}
//leetcode submit region end(Prohibit modification and deletion)
