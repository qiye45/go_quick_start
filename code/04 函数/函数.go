package main

import (
	"fmt"
	"math/big"
)

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

// addAll 多数相加
func addAll(a int, numbers ...int) int {
	sum := a
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// interface多数相加
func addAll2(numbers ...interface{}) *big.Int {
	result := big.NewInt(0) // Initialize with 0
	//fmt.Println("numbers",numbers)
	for _, v := range numbers {
		val, ok := v.(int) // Type assertion to get int64 value
		if !ok {
			// Handle error if value is not int64
			continue
		}
		result.Add(result, big.NewInt(int64(val))) // Use Add method correctly
	}
	//fmt.Println(result)
	return result
}

// interface多数相加 类型判断
func addAll3(numbers ...interface{}) *big.Int {
	result := big.NewInt(0) // Initialize with 0
	for _, v := range numbers {
		switch val := v.(type) {
		case int64:
			result.Add(result, big.NewInt(val))
		case int:
			result.Add(result, big.NewInt(int64(val)))
		case *big.Int:
			result.Add(result, val)
		default:
			fmt.Printf("Unsupported type: %T\n", v)
		}
	}
	return result
}

func main() {
	fmt.Println("lesson7 函数")
	kelvin := 294.0

	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "°K is", celsius, "°C") //294 °K is 20.850000000000023 °C

	sum := addAll(3, 4, 5, 6, 7)
	fmt.Println(sum) //25

	tests := []struct {
		numbers []any
		want    string
	}{
		{[]any{1, 2, 3}, "6"},
		{[]any{10, 20, 30, 40}, "100"},
		{[]any{1000000000000000000, 1}, "1000000000000000001"},
		{[]any{"abc", 1, 2}, "3"}, // Error case: non-int64 value
	}

	for _, tt := range tests {
		//got := addAll2(tt.numbers) 变成传入数组[[]]
		//got := addAll2(tt.numbers...)
		//fmt.Println(got.String())
		//fmt.Printf("addAll2(%v) = %v, want %v\n", tt.numbers, got, tt.want)
		got := addAll3(tt.numbers...)
		fmt.Printf("addAll3(%v) = %v, want %v\n", tt.numbers, got, tt.want)

	}

}
