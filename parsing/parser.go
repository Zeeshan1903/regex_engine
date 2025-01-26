package parsing

import (
	"fmt"
	"bufio"
	"os"
	_"reflect"
	"regex_engine/types"
)

//We will take the raw input here and will return the queue with the parsed section in it


//Always remember that Strings are immutable datatype in golang


func Take_input() (string, string){

	reader := bufio.NewReader(os.Stdin);
	fmt.Printf("Give Pattern:\t");

	//taking the input from the user (pattern)
	input_pattern,_ := reader.ReadString('\n');
	input_pattern = input_pattern[:len(input_pattern)-1];
	fmt.Printf("String:\t\t");

	//taking the input stirng to search in 
	input_string,_ := reader.ReadString('\n');
	input_string = input_string[:len(input_string)-1];	

	//here i will be returning the string from the input 
	return input_pattern,input_string;
}

//what ill do is make a list of type Quantifiers and make a fn which will take this 2 input string and will convert them into the particular Struct and will add that 
//strut to that list 

func Generate_Type(c rune) types.Data{

	switch c {
		case '.': 	return types.Dot{};
		case '*': 	return types.Asterisk{}
		case '?':	return types.Qus_mark{}
		case '+':	return types.Plus{}
		case '(':	return types.Open_paran{}
		case ']':	return types.Close_literal{}
		case '[':	return types.Open_literal{}
		case ')': 	return types.Close_paran{}
		case '-': 	return types.Dash{}
		case '\\':	return types.Escape{Value:'\\'};
		case ',': 	return types.Comma{}
		case '{':	return types.Open_Range{};
		case '}':	return types.Close_literal{}
		case '0','1','2','3','4','5','6','7','8','9':	return types.Range{N:int(c)}
		default: 	return types.Character{Value: c}
	}
}
func Parse_Input(a,b string)([]types.Quantifier, []types.Quantifier){

	var pattern []types.Quantifier;
	var input []types.Quantifier;

	for _,char := range a {
		pattern = append(pattern, types.Quantifier{Data:Generate_Type(char)});
	} 
	for _,char := range b{
		input = append(input,types.Quantifier{Data:Generate_Type(char)});
	}
	return pattern, input;
}


//A helper fn to return types to check while solving them 
func Get_TypeOf(c types.Quantifier) string{
	switch c.Data.(type){
			case types.Dot:					return "Dot"
			case types.Asterisk:			return "Asterisk"
			case types.Qus_mark:			return "Qus_mark"
			case types.Plus:				return "Plus"
			case types.Character:			return "Character" 
			case types.Range:				return "Range"
			case types.Open_paran:			return "Open_paran"
			case types.Close_Range:			return "Close_Range"
			case types.Close_paran:			return "Close_paran"
			case types.Open_literal:		return "Open_literal"
			case types.Close_literal:		return "Close_literal"
			case types.Open_Range:			return "Open_Range"
			case types.Dash:				return "Dash"
			case types.Escape:				return "Escape"
			case types.Comma:				return "Comma"
			default:						return "Unknown"
	}
}


func Print_Tokens()([]types.Quantifier, []types.Quantifier){ 
	pattern , input := Take_input();
	a,b := Parse_Input(pattern,input);
	return a,b;
}