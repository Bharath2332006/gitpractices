package main
import "fmt"

func main(){
	var n int
	fmt.Scanln(&n)
	 arr  :=make([]int, n)
	for i:=0;i<n;i++{
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
}