package main

import (
	"fmt"
	"sort"
)

// g[10, 9, 8, 7] s[5, 6, 7, 8]
func FindContentChildren(g []int, s []int) int {
	count := 0
	// 1. check if g and s just have one element and compare them
	if len(s) == 1 && len(g) == 1 {
		if g[0] <= s[0] {
			return 1
		} else {
			return 0
		}
	}

	// 2. sort g and s in non-decreasing order
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	i, j := 0, 0

	for i < len(g) && j < len(s) {
		// 3. compare g and s if g[i] <= s[j] count++
		if s[j] >= g[i] {
			count++
			i++
			j++
		} else { //
			j++
		}
	}
	return count
}

/*
2125. Number of Laser Beams in a Bank

Anti-theft security devices are activated inside a bank.
You are given a 0-indexed binary string array bank representing the floor plan of the bank
, which is an m x n 2D matrix. bank[i] represents the ith row, consisting of '0's and '1's.

	'0' means the cell is empty, while'1' means the cell has a security device.

There is one laser beam between any two security devices if both conditions are met:

The two devices are located on two different rows: r1 and r2, where r1 < r2.
For each row i where r1 < i < r2, there are no security devices in the ith row.
Laser beams are independent, i.e., one beam does not interfere nor join with another.

Return the total number of laser beams in the bank
*/

// func checkSecurity(s1 string) int {
// 	ls1 := 0
// 	for i := 0; i < len(s1); i++ {
// 		if s1[i] == '1' {
// 			ls1++
// 		}
// 	}
// 	return ls1
// }

// func removeItem(s []string, index int) []string {
// 	return append(s[:index], s[index+1:]...)
// }

// func NumberOfBeams(bank []string) int {
// 	if len(bank) == 0 {
// 		return 0
// 	}

// 	m, n := len(bank), len(bank[0])
// 	if m < 1 || n < 1 || m > 500 || n > 500 {
// 		return 0
// 	}

// 	beams := 0

// 	for i := 0; i < m; i++ {
// 		ls := checkSecurity(bank[i])
// 		if ls == 0 {
// 			bank = removeItem(bank, i)
// 			i--
// 			m--
// 		}
// 	}

// 	i, j := 0, 1
// 	for i < m && j < m {
// 		ls1, ls2 := checkSecurity(bank[i]), checkSecurity(bank[j])
// 		fmt.Println("bank[i], bank[j]", bank[i], bank[j])
// 		fmt.Println("ls1, ls2", ls1, ls2)
// 		if ls1 != 0 && ls2 != 0 {
// 			beams += ls1 * ls2
// 			i++
// 			j++
// 		}

// 	}
// 	return beams
// }

// // best solution
// func numberOfBeams(bank []string) int {
// 	n := make([]int, len(bank)) // this is the number of 1s in each row
// 	for i := range bank {
// 		c := 0
// 		for _, laser := range bank[i] {
// 			if laser == '1' {
// 				c++
// 			}
// 		}
// 		n[i] = c // number of 1s in each row
// 	}
// 	res := 0
// 	last := n[0]
// 	for i := 1; i < len(n); i++ {
// 		if n[i] != 0 {
// 			res += last * n[i]
// 			last = n[i]
// 		}
// 	}
// 	return res
// }

/*
2870. Minimum Number of Operations to Make Array Empty Medium

You are given a 0-indexed array nums consisting of positive integers.

There are two types of operations that you can apply on the array any number of times:

Choose two elements with equal values and delete them from the array.
Choose three elements with equal values and delete them from the array.
Return the minimum number of operations required to make the array empty, or -1 if it is not possible.

Example 1:

Input: nums = [2,3,3,2,2,4,2,3,4]
Output: 4
Explanation: We can apply the following operations to make the array empty:
- Apply the first operation on the elements at indices 0 and 3. The resulting array is nums = [3,3,2,4,2,3,4].
- Apply the first operation on the elements at indices 2 and 4. The resulting array is nums = [3,3,4,3,4].
- Apply the second operation on the elements at indices 0, 1, and 3. The resulting array is nums = [4,4].
- Apply the first operation on the elements at indices 0 and 1. The resulting array is nums = [].
It can be shown that we cannot make the array empty in less than 4 operations.
Example 2:

Input: nums = [2,1,2,2,3,3]
Output: -1
Explanation: It is impossible to empty the array.

Constraints:

2 <= nums.length <= 105
1 <= nums[i] <= 106
*/

// func removeIndex(s []int, index int) []int {
// 	if index < 0 || index >= len(s) {
// 		return s
// 	}
// 	return append(s[:index], s[index+1:]...)
// }

// func remove(s []int, value int, quality int) []int {
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] == value {
// 			s = append(s[:i], s[i+1:]...)
// 			quality--
// 			if quality == 0 {
// 				break
// 			}
// 		}
// 	}
// 	return s
// }

// func minOperations(nums []int) int {
// 	if len(nums) < 2 || len(nums) > 100000 {
// 		return -1
// 	}

// 	sort.Ints(nums)

// 	count := 0
// 	m := make(map[int]int)

// 	sort.Ints(make([]int, 0))

// 	for _, v := range nums {
// 		m[v]++
// 	}

// 	l := len(nums)
// 	i := 0
// 	for l > 0 {
// 		dup := m[nums[i]]
// 		if dup <= 1 {
// 			return -1
// 		}

// 		if dup == 3 || dup == 2 {
// 			count++
// 			l -= dup
// 			m[nums[i]] -= dup
// 			// nums = remove(nums, nums[0], dup)
// 			i += dup
// 		} else {
// 			if dup-3 != 1 {
// 				m[nums[i]] -= 3
// 				l -= 3
// 				// nums = remove(nums, nums[0], 3)
// 				i += 3
// 			} else {
// 				m[nums[i]] -= 2
// 				l -= 2
// 				i += 2
// 				// nums = remove(nums, nums[0], 2)
// 			}
// 			count++
// 		}
// 	}
// 	// if len(m) == 1 {
// 	// 	dup := len(nums)
// 	// 	if dup <= 1 {
// 	// 		return -1
// 	// 	}
// 	// 	for dup > 0 {
// 	// 		if dup == 3 || dup == 2 {
// 	// 			count++
// 	// 			dup -= dup
// 	// 		} else {
// 	// 			if dup-3 != 1 {
// 	// 				dup -= 3
// 	// 			} else {
// 	// 				dup -= 2
// 	// 			}
// 	// 			count++

// 	// 		}
// 	// 	}
// 	// 	return count
// 	// } else {
// 	// 	for l > 0 {
// 	// 		dup := m[nums[0]]
// 	// 		if dup <= 1 {
// 	// 			return -1
// 	// 		}
// 	// 		if dup == 3 || dup == 2 {
// 	// 			count++
// 	// 			l -= dup
// 	// 			m[nums[0]] -= dup
// 	// 			nums = remove(nums, nums[0], dup)
// 	// 		} else {
// 	// 			if dup-3 != 1 {
// 	// 				m[nums[0]] -= 3
// 	// 				l -= 3
// 	// 				nums = remove(nums, nums[0], 3)

// 	// 			} else {
// 	// 				m[nums[0]] -= 2
// 	// 				l -= 2
// 	// 				nums = remove(nums, nums[0], 2)
// 	// 			}
// 	// 			count++
// 	// 		}
// 	// 	}
// 	// }

// 	return count

// }

/*
300. Longest Increasing Subsequence Medium

Given an integer array nums, return the length of the longest strictly increasing
subsequence
.



Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4
Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1


Constraints:

1 <= nums.length <= 2500
-104 <= nums[i] <= 104


Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?
*/

// func lengthOfLIS(nums []int) int {
// 	n := []int{}
// 	i, j := 0, 1
// 	n = append(n, nums[0])
// 	for i < len(nums) && j < len(nums) {
// 		fmt.Println("i j", i, j)
// 		fmt.Println("nums[i], nums[j], nums[0]", nums[i], nums[j])
// 		lsv := n[len(n)-1]
// 		fmt.Println("lsv", len(n), len(nums)/2)

// 		if nums[i] < n[0] {
// 			if len(n) < len(nums)/2 || len(n) == 1 {
// 				n = []int{}
// 				n = append(n, nums[i])
// 			}
// 		}

// 		if nums[i] > nums[j] {
// 			if lsv < nums[i] && nums[i] > n[0] {
// 				n = append(n, nums[i])
// 			}

// 		} else {
// 			if lsv > nums[i] && lsv <= nums[j] && nums[i] > n[0] {
// 				pre := n[len(n)-1] + n[len(n)-2]
// 				nex := nums[i] + nums[j]
// 				if pre <= nex {
// 					n[len(n)-1] = nums[i]
// 				}
// 			}

// 			if (lsv > nums[i] && nums[i] > n[0]) && (lsv > nums[j] && nums[j] > n[0]) {
// 				pre := n[len(n)-2]
// 				if pre < nums[i] {
// 					n[len(n)-1] = nums[i]
// 				}
// 			}

// 			if lsv < nums[i] && lsv < nums[j] {
// 				n = append(n, nums[i])
// 			}
// 			if (j == len(nums)-1) && n[len(n)-1] < nums[j] {
// 				n = append(n, nums[j])
// 			}
// 		}

//			i++
//			j++
//			fmt.Println("n", n)
//		}
//		return len(n)
//	}
// func lengthOfLIS(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	n := []int{nums[0]}
// 	for i := 1; i < len(nums); i++ {
// 		fmt.Println("n", n)
// 		if nums[i] > n[len(n)-1] {
// 			n = append(n, nums[i])
// 		} else {
// 			// Find the position to insert the current number in the sorted sequence
// 			left, right := 0, len(n)-1
// 			for left < right {
// 				fmt.Println("left right", left, right)
// 				fmt.Println("n[left] n[right]", n[left], n[right])
// 				mid := left + (right-left)/2
// 				fmt.Println("mid", mid)
// 				if n[mid] < nums[i] {
// 					left = mid + 1
// 				} else {
// 					right = mid
// 				}
// 			}

// 			// Replace the element at the found position with the current number
// 			n[left] = nums[i]
// 		}
// 	}

// 	return len(n)
// }

// func lengthOfLIS(nums []int) int {
// 	list := []int{nums[0]}
// 	index := 1
// 	for i := 1; i < len(nums); i++ {
// 		index = binarySearch(list, 0, len(list), nums[i])
// 		if index == len(list) {
// 			list = append(list, nums[i])
// 		} else {
// 			list[index] = nums[i]
// 		}
// 	}
// 	return len(list)
// }

//	func binarySearch(list []int, left, right, target int) int {
//		for left < right {
//			mid := (left + right) / 2
//			if list[mid] == target {
//				return mid
//			}
//			if list[mid] > target {
//				right = mid
//			} else {
//				left = mid + 1
//			}
//		}
//		return left
//	}

/*
872 Leaf-Similar Trees Easy

Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLeafNode(root *TreeNode) []int {
	leaf := []int{}
	if root == nil {
		return leaf
	}

	if root.Left == nil && root.Right == nil {
		leaf = append(leaf, root.Val)
	}

	leaf = append(leaf, getLeafNode(root.Left)...)
	leaf = append(leaf, getLeafNode(root.Right)...)
	return leaf
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}

	// 1. get leaf node of root1
	leaf1 := getLeafNode(root1)
	leaf2 := getLeafNode(root2)

	fmt.Println(leaf1)
	fmt.Println(leaf2)

	if len(leaf1) != len(leaf2) {
		return false
	}

	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}

	return true
}

func main() {
	// fmt.Println("minOperations", minOperations([]int{3, 14, 3, 14, 3, 14, 14, 3, 3, 14, 14, 14, 3, 14, 14, 3, 14, 14, 14, 3}))
	// fmt.Println("minOperations", minOperations([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}))
	// fmt.Println("minOperations", minOperations([]int{2, 1, 2, 2, 3, 3}))
	// fmt.Println("minOperations", minOperations([]int{2, 3, 3, 2, 2, 4, 2, 3, 4}))
	// fmt.Println("minOperations", minOperations([]int{16, 16, 16, 19, 16, 3, 16, 8, 16, 16, 16, 19, 3, 16, 16}))
	// fmt.Println("lengthOfLIS", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	// fmt.Println("lengthOfLIS", lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	// fmt.Println("lengthOfLIS", lengthOfLIS([]int{10, 9, 2, 5, 3, 4}))
	// fmt.Println("lengthOfLIS", lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	// fmt.Println("lengthOfLIS", lengthOfLIS([]int{2, 15, 3, 7, 8, 6, 18}))
	// fmt.Println("lengthOfLIS", lengthOfLIS([]int{11, 12, 13, 14, 15, 6, 7, 8, 101, 18}))
	// fmt.Println("lengthOfLIS", lengthOfLIS([]int{10, 22, 9, 33, 21, 50, 41, 60, 80}))
	// root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
	fmt.Println("Similar", leafSimilar(&TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}))
	fmt.Println("Similar", leafSimilar(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}))
	fmt.Println("Similar", leafSimilar(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}))
}
