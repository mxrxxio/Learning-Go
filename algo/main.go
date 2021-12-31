package main

import "fmt"

func main() {
	fmt.Println(collatz(6))
	fmt.Println(fibonacci(10))
	fmt.Println(insertionSort([]int{10, 5, 3, 4, 8, 2, 9}))
	fmt.Println(binarySearch(34, []int{1, 4, 9, 11, 34, 36, 77, 120}))
	fmt.Println(isPrime(19))
	fmt.Println(GCD(4, 12))
}

func collatz(n int) []int {
	result := []int{n}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n *= 3
			n += 1
		}
		result = append(result, n)
	}
	return result
}

func fibonacci(n int) []int {
	result := []int{1, 1}
	for i := 0; i < n-2; i++ {
		result = append(result, result[i]+result[i+1])
	}
	return result
}

func insertionSort(arr []int) []int {
	/* Is an efficient algorithm for sorting small numbers of elements */
	for i := 1; i < len(arr)-1; i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j = j - 1
		}
	}
	return arr
}

func binarySearch(target int, arr []int) int {
	L := 0
	R := len(arr) - 1
	for L <= R {
		mid := L + (R-L)/2 // never use (L+R) / 2 'cause overflow
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	return -1
}

func isPrime(n int) bool {
	for i := 2; i < n-2; i++ {
		if n%i == 0 {
			// fmt.Printf("%v / %v = %v\n", n, i, n/i)
			return true
		}
	}
	return false
}

func GCD(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}
