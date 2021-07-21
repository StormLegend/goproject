package main

func Sum(numbers []int) int {
    sum := 0
    for _,i := range numbers{
        sum += i
    }
    return sum
}

func SumAllTails(arrs ...[]int)[]int{
    var sums []int
    for _,number := range arrs {
	if len(number)  == 0{
	    sums = append(sums,0)
	}else{
            tails := number[1:]
	    sums = append(sums, Sum(tails))
        }
    }
    return sums
}


func Sumall(numbersToSum ...[]int)[]int{
    var sums []int
    for _,number := range numbersToSum{
        sums = append(sums,Sum(number))
    }
    return sums
}

func main(){
    Sumall([]int{1,1,1})
}
