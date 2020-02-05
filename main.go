package  main

import (
	"fmt"
)

func Index (sl []int, predicate func(int)bool) int {
	 for i, value :=range sl {
	 	if predicate(value) {
	 		return i
		}}
	 return -1
}

func Any (sl []int, predicate func(int)bool) bool{
	 return Index (sl,predicate) != -1
}

func All (sl []int, predicate func(int)bool) bool{

     return Index(sl,func(i int)bool{
     	return !predicate(i)
     } ) == -1
}

func Find (sl []int, predicate func(int)bool) (int,error){
	indx := Index(sl,predicate)
	if indx==-1{
		return -1,fmt.Errorf("Current element does not exist")
	}
	return sl[indx],nil
}

func None (sl []int, predicate func(int)bool) bool {
	return Index (sl,predicate) == -1
}

func main() {

}