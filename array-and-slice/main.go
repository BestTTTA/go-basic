package main

// import "fmt"

func main() {
	//สามารถเก็บข้อมูลได้หลายข้อมูลแต่ต้องเป็นชนิดเดียวกัน

	// var a [3]int = [3]int{1, 2, 3}

	//print value position array
	// for v := range a {
	// 	fmt.Println(a[v])
	// }

	// print index of array
	// for v := range a {
	// 	fmt.Println(v)
	// }

	//slice have size
	// b := []string{"best1", "best2", "best3"}
	// fmt.Println(b)
	// fmt.Println(b[2])

	//slice no size
	// b := make([]int, 0)
	// b = append(b, 1, 2, 3)
	// fmt.Println(b)
	// b[2] = 5
	// b = append(b, 4)
	// fmt.Println(b)

	//have size
	// a := [3]int{1, 2, 3}
	// for i := range a {
	// 	fmt.Printf("address in main function %v\n", &a[i])
	// }
	// address(a)

	//don't have size
	// a := []int{1, 2, 3}
	// for i := range a {
	// 	fmt.Printf("address in main function %v\n", &a[i])
	// }
	// address(a)

}

//have max size array
// func address(num [3]int) {
// 	for i := range num {
// 		fmt.Printf("address inside func address %v\n", &num[i])
// 	}
// }

//have size
// func double(nums []int) []int {
// 	newNums := make([]int, len(nums))
// 	for i := range nums {
// 		newNums[i] = nums[i] * 2
// 	}
// 	return newNums
// }

//don't have size
// func double(nums []int) []int {
// 	newNums := make([]int, len(nums))

// 	for i := range nums {
// 		newNums[i] = nums[i] * 2
// 	}
// 	return newNums
// }

//don't have max size array
// func address(num []int) {
// 	for i := range num{
// 		fmt.Printf("address inside func address %v\n", &num[i])
// 	}
// }
