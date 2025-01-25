package main

import (
	"fmt"
	"regex_engine/types"
	"regex_engine/parsing"
	_"reflect"
)

func main(){
	fmt.Printf("Hello World\n");
	fmt.Printf("Value %s\n",types.Open_literal{}.String());

	pattern_token , input_token := parsing.Print_Tokens();


	fmt.Printf("\nWe have successfully parser %v and %v\n",pattern_token,input_token);
}