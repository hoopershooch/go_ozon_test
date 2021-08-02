package main

import (
    
    "fmt"
)

func main(){
    var arr [] uint
    var start,stop uint
    start,stop = findUnsortedSubarray(arr)
    fmt.Printf("Start=%d, Stop=%d", start,stop)
}


func findUnsortedSubarray(array []uint) (left uint, right uint){
    type diff struct{
        startInd uint
        stopInd uint
    }    

    var negativeDiffs []diff
    var positiveDiffs []diff    
    var prevVal uint

    for i,val := range array{
        if i==0 {
            prevVal = val 
            continue
        }
        if prevVal <= val {
            fmt.Printf("Positive Difference: prev val = %d, val=%d\n",prevVal, val)  

            if positiveDiffs!=nil{
               if positiveDiffs[len(positiveDiffs)-1].stopInd < uint(i-1) {
                   positiveDiffs = append(positiveDiffs, diff{startInd: uint(i-1), stopInd: uint(i)})
               } else {
                   positiveDiffs[len(positiveDiffs)-1].stopInd = uint(i)
               }   
            } else{
                positiveDiffs = append(positiveDiffs, diff{startInd: uint(i-1), stopInd: uint(i) })   
            }

        }else{
            fmt.Printf("Negative Difference: prev val = %d, val=%d\n",prevVal, val)  
            if negativeDiffs!=nil{
               if negativeDiffs[len(negativeDiffs)-1].stopInd < uint(i-1) {
                   negativeDiffs = append(negativeDiffs, diff{startInd: uint(i-1), stopInd: uint(i) })
               } else {
                   negativeDiffs[len(negativeDiffs)-1].stopInd = uint(i)
               }   
            } else{
                negativeDiffs = append(negativeDiffs, diff{startInd: uint(i-1), stopInd: uint(i) })   
            }

        } 
        prevVal = val
    }

    for i, val:=range positiveDiffs{
        fmt.Printf("i=%d, pos_diff.start=%d, pos_diff.stop=%d\n",i,val.startInd,val.stopInd)
    } 

    for i, val:=range negativeDiffs{
        fmt.Printf("i=%d, neg_diff.start=%d, neg_diff.stop=%d\n",i,val.startInd,val.stopInd)
    } 

    if negativeDiffs!=nil{
        return negativeDiffs[0].startInd+1, negativeDiffs[len(negativeDiffs)-1].stopInd
    } else{
        return 0,0
    }   

}