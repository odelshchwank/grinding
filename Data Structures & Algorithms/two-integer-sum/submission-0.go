func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, num := range nums {
        need := target - num
        if v, ok := seen[need]; ok {
            return []int{v, i}
        }

        seen[num] = i
    }

    return nil
}