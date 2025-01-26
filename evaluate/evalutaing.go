package evaluate

import ("fmt"
		"regex_engine/parsing"
		"regex_engine/types"
		"strings"
		"strconv"
)
type State struct {
	Previous string
	Current string
}


func Range_fn(m, n int, c rune, b *[]types.Quantifier) bool {


    new_m := m
    new_n := n - new_m

    for m != 0 {
        if len(*b) == 0 {
            return false
        }

        switch data := (*b)[0].Data.(type) {

        case types.Character:
            if data.Value != c { 
                return false
            }else{
				(*b) = (*b)[1:]
				m--;
			}
        case types.Escape:
            if data.Value != c { 
                return false
            }
        default:
            // For other types like Dot, Asterisk, etc., I'll handle them differently if needed
            return false
        }

    }

    // Now check for max n value
    // n = -1 means infinity
    // n = -2 means exact match
    if n == -1 {
        // Loop through until we can't match c anymore
        for len(*b) > 0 {

            switch data := (*b)[0].Data.(type) {
            
			case types.Character:
                if data.Value != c {
                    return true;
                }else{
					(*b) =(*b)[1:];
				}
            
			case types.Escape:
                if data.Value != c {
                    return false
                }
            
			default:
            
				return true;
            }

        }
        return true
    } else if n == -2 {
        return true
    } else {
     
        for new_n != 0 && len(*b) > 0 {
            
			switch data := (*b)[0].Data.(type) {
            
			case types.Character:
                if data.Value != c {
                    return true;
                }else{
					(*b) =(*b)[1:];
					new_n--
				}
            
			case types.Escape:
                if data.Value != c {
                    return false
                }
            
			default:
                return true;
            }
            
        }
    }

    return true
}




func Range_value(a *[]types.Quantifier) (n, m int) {
   
	
	if parsing.Get_TypeOf((*a)[1]) == "Comma" && parsing.Get_TypeOf((*a)[2]) == "Range" && parsing.Get_TypeOf((*a)[3]) == "Close_Range" {
    
		var err error
        n, err = strconv.Atoi((*a)[0].Data.String()) 
    
		if err != nil {
            return 0, 0
        }
    
		m, err = strconv.Atoi((*a)[2].Data.String())  
    
		if err != nil {
            return 0, 0
        }
    
		*a = (*a)[4:] 
        return n-48, m-48
    
	} else if parsing.Get_TypeOf((*a)[1]) == "Comma" && parsing.Get_TypeOf((*a)[2]) == "Close_Range" {
       
		var err error
        
		n, err = strconv.Atoi((*a)[0].Data.String())  
		
        if err != nil {
            return 0, 0
        }
        
		m = -1
        *a = (*a)[3:] 
        return n-48, m
    
	
	} else {
        var err error
		
        n, err = strconv.Atoi((*a)[0].Data.String()) 

		fmt.Printf("Hurray %v\n",n);
        if err != nil {
            return 0, 0
        }
        m = -2
        *a = (*a)[2:] 
        return n-48, m
    }
}

func Create_literal_array(a *[]types.Quantifier) string {
	
	str := ""

	for len(*a) > 0 && parsing.Get_TypeOf((*a)[0]) != "Close_literal" {

		if parsing.Get_TypeOf((*a)[1]) == "Dash" {
			
			if char1, ok1 := (*a)[0].Data.(types.Character); ok1 {
			
				if char2, ok2 := (*a)[2].Data.(types.Character); ok2 {
			
					for i := char1.Value; i <= char2.Value; i++ {
						str += string(i)
					}
			
					*a = (*a)[3:]
					continue
			
				}
			
			}
		}

		if char, ok := (*a)[0].Data.(types.Character); ok {
			str += string(char.Value)
		}

		*a = (*a)[1:]
	}
	return str
}





func Evaluate_literal(s string, b *[]types.Quantifier) bool {
	
	ch, ok :=  (*b)[0].Data.(types.Character);
	
	if len(*b) > 0 && ok && strings.ContainsRune(s,rune(ch.Value)) {
		*b = (*b)[1:]
		return true
	}
	
	return false
}



func Evaluate(a *[]types.Quantifier,b *[]types.Quantifier) bool {
	state := State {Current : "", Previous : ""};

	for len(*a) != 0{
	
		state.Previous = state.Current;

		fmt.Printf("State.Previous value %v and *a[0] value %v len %v\n",state.Previous,(*a)[0].Data,len(*a));
	
		switch (*a)[0].Data.(type) {
	
			case types.Dot :	
	
			ch, ok :=  (*b)[0].Data.(types.Character);
	
				if ok {
					state.Current = string(ch.Value);
				}
			
				(*a )= (*a)[1:];
				(*b) = (*b)[1:];
			
			case types.Character :
				if (*a)[0].Data.(types.Character).Value == (*b)[0].Data.(types.Character).Value {
					ch, ok := (*b)[0].Data.(types.Character)
					if ok {
						state.Current = string(ch.Value)
					}
	
					*a = (*a)[1:]
					*b = (*b)[1:]
				} else {
					return false
				}
		
			case types.Asterisk :
				is_true := Range_fn(0,-1,[]rune(state.Previous)[0],b);
				if !is_true{
					return false;
				}		
				(*a) = (*a)[1:];
				(*b) = (*b);
				state.Current = state.Previous;

		
			case types.Plus:
				is_true := Range_fn(1,-1,[]rune(state.Previous)[0],b);
				if !is_true{
					return false;
				}
				(*a) = (*a)[1:];
				(*b) = (*b);
				state.Current = state.Previous;

			case types.Qus_mark:
				is_true := Range_fn(0,1,[]rune(state.Previous)[0],b);
				if !is_true{
					return false;
				}
				(*a) = (*a)[1:]
				(*b) = (*b);
				state.Current = state.Previous;

			case types.Open_Range:
				
				(*a) = (*a)[1:];
				n,m := Range_value(a);
				
				var r rune = []rune(state.Previous)[0];
				
				is_true := Range_fn(n,m,r, b);
				
				if !is_true {
					return false;
				}
		
				(*a) = (*a);
				(*b) = (*b)[1:];
				state.Current = state.Previous;

			case types.Close_Range:
				fmt.Printf("No Bracket for Opening Range\n");

				return false;

			case types.Escape:


			
				if (*a)[0].Data.(types.Character).Value == (*b)[0].Data.(types.Character).Value {
					(*a) = (*a)[2:];
					(*b) = (*b)[1:];

				}else{
					return false;
				}

			case types.Open_literal:

				ch, ok :=  (*b)[0].Data.(types.Character);
				
				if ok {
					state.Current = string(ch.Value);
				}
		
				array := Create_literal_array(a);
				is_type := Evaluate_literal(array,b);
		
				if !is_type {
					return false;
				}	
			
			case types.Close_literal:
				fmt.Printf("No Bracket For Closing Literal\n");

			default:
				fmt.Printf("Unknown Token Types\n");

		} 


	}
	return true;	
}
























































