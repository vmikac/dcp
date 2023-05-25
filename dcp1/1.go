package main
import "fmt"

func main() {
   const l = 4
   const target = 17
   var nums = [l]int{10, 5, 3, 7}
   for i:=0;i<l;i++ {
     for j:=i;j<l;j++ {
       if nums[i] + nums[j] == target {
          fmt.Printf("%d %d %d\n", nums[i], nums[j], target)
       }
     }
   }

   fmt.Println("ok")

}
