package main
import "fmt"
func linearSearchSorted(data []int, value int) bool {
size := len(data)
for i := 0; i < size; i++ {
if value == data[i] {
return true
} else if value < data[i] {
return false
}
}
return false
}


func main() {
    items := []int{1,2,3,4,5,6,7,8,9,10}
    fmt.Println(linearSearchSorted(items,6))
}