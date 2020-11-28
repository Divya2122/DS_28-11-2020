package main
import "fmt"
func linearSearch(numbers []int, item int) int {

	if numbers != nil && len(numbers) > 0 {

		for i := 0; i< len(numbers); i++ {

			if numbers[i] == item {
				return numbers[i]
			}

		}

	}

	return -1
}


func main() {
    items := []int{1,2,3,4,5,6,7,8,9,10}
    fmt.Println(linearSearch(items,6))
}