package main 

import "testing"
import "reflect"

func TestSumall(t *testing.T){
   got := Sumall([]int{1,2}, []int{0,9})
   want := []int{3,9}
   if !reflect.DeepEqual(got, want){
       t.Errorf("want %v but got %v",want,got)
   }
}

func TestSumAllTails(t *testing.T){
    t.Run("make the sum of slices", func(t *testing.T){
        got := SumAllTails([]int{1,2},[]int{0,9})
        want := []int{2,9}
        if !reflect.DeepEqual(got, want){
            t.Errorf("want %v but got %v", want, got)
        }
    })
    t.Run("make the sum with empty slices",func(t *testing.T){
        got := SumAllTails([]int{},[]int{3,4,5})
        want := []int{0,9}
        if !reflect.DeepEqual(got, want){
            t.Errorf("want %v but got %v", want, got)
        }
    })
}

func TestSum(t *testing.T){
    t.Run("collection of 5 numbers:",func (t *testing.T){
    
        numbers := []int{1,2,3,4,5} 
        got := Sum(numbers)
        want := 15
        if got != want {
            t.Errorf("want %d but got %d",want,got)
        }
    })

    t.Run("collection of any numbers:",func(t *testing.T){
    
        numbers := []int{1,2,3} 
        got := Sum(numbers)
        want := 6
        if got != want {
            t.Errorf("want %d but got %d",want,got)
        }
    })
}
