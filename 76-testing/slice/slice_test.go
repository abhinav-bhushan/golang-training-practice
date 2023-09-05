package slice

import "testing"

func TestSliceReverse(t *testing.T){
	input:=[]int{1,2,3,4,5}
	output:=[]int{5,4,3,2,1}

	actualOutput:=Reverse(input)

	if len(output)!=len(actualOutput){
		t.Fail()
		return
	}

	for i:=0;i<len(output);i++{
		if output[i]!=actualOutput[i]{
			t.Fail()
		}
	}

}

func TestSliceSmallest(t *testing.T){
	input:=[]int{1,2,3,4,5}
	output:=1

	actualOutput:=Smallest(input)

	if output!=actualOutput{
		t.Fail()
	}

}

func TestSliceBiggest(t *testing.T){
	input:=[]int{1,2,3,4,5}
	output:=5

	actualOutput:=Biggest(input)

	if output!=actualOutput{
		t.Fail()
	}

}

func TestSliceSum(t *testing.T){
	input:=[]int{1,2,3,4,5}
	output:=15

	actualOutput:=Sum(input)

	if output!=actualOutput{
		t.Fail()
	}

}