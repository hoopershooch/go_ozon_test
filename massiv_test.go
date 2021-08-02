package main

import (
    "fmt"
    "testing"
)


func TestFindUnsortedSubarray(t *testing.T){
    type testCase struct{
       arr [] uint   
       resStart uint
       resStop uint
    }   
    var start,stop uint
    testCases := []testCase {
        {
            arr: [] uint {1,2,3,4,5,6,7,8},
            resStart: 0, 
            resStop: 0,   
        },
        {
            arr: [] uint {100,9,8,7,6,5,20,30,40,50,60},
            resStart: 0, 
            resStop: 5,   
        },
        {
            arr: [] uint {10,9,8,7,6,5,20,30,40,50,60},
            resStart: 1, 
            resStop: 5,   
        },

        {
            arr: [] uint {2,3,4,5,6},
            resStart: 0, 
            resStop: 0,   
        },
    }  

    for i, testCaseVal := range testCases{
        start,stop = findUnsortedSubarray(testCaseVal.arr)
        fmt.Printf("TEST N%d: start=%d, stop=%d", i,start, stop)     
        if start != testCaseVal.resStart && stop!=testCaseVal.resStart {
            t.Errorf("HERE %d", 17)
        } 

    }
}