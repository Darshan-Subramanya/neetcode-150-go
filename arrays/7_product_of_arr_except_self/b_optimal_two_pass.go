package main

import "fmt"

func productExceptSelfOptimalTwoPass(nums []int) []int {
    prefix := make([]int, len(nums))
    postfix := make([]int, len(nums))
    output := make([]int, len(nums))

    prefix[0] = 1
    for i := 1; i < len(nums); i++ {
        prefix[i] = prefix[i-1] * nums[i-1]
    }

    postfix[len(nums)-1] = 1
    for i := len(nums) - 2; i >= 0; i-- {
        postfix[i] = postfix[i+1] * nums[i+1]
    }

    for i := range output {
        output[i] = prefix[i] * postfix[i]
    }

    return output
}

func main() {
	nums := []int{1,2,4,6}
	res := productExceptSelfOptimalTwoPass(nums)
	fmt.Println(res)
}