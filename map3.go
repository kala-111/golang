package main
import "fmt"
func main(){
	marks:=[]int{98,99,78,23,11,98}
	total:=total(marks)
	markslength:=len(marks)
	fmt.Println("total marks are",total)
	fmt.Println("average marks are",average(total,markslength))
}
func total(marks[]int)int{
	total:=0
	for_, v := range marks {
		total +=v
	
	}
	return total
}
func average(total int,markslength int)int{
	return total/markslength

}