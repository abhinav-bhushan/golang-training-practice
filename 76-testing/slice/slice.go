package slice



func Reverse(s []int )[]int {
	output:=make([]int,0)

	for i:=len(s)-1;i>-1;i--{
		output = append(output, s[i])
	}
	// output[len(output)-1]=-1

	return output
}

func Smallest(s []int)(output int){
	output=s[0]

	for _,v:=range(s){
		if v<output{
			output=v
		}
	}
	return

}

func Biggest(s []int)(output int){
	output=s[0]

	for _,v:=range(s){
		if v>output{
			output=v
		}
	}
	return

}

func Sum(s []int)(output int){
	output=0

	for _,v:=range(s){
		output+=v
	}

	return
}







