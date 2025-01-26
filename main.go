package main

import (
	"fmt"
	_"regex_engine/types"
	"regex_engine/parsing"
	_"reflect"
	"regex_engine/evaluate"
)

func main(){

	pattern_token , input_token := parsing.Print_Tokens();


	fmt.Printf("\nWe have successfully parser %v and %v\n",pattern_token,input_token);

	ans := evaluate.Evaluate(&pattern_token,&input_token);
	fmt.Printf("Ans : %v\n",ans)
}
