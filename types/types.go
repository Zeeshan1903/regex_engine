//In this file I'll define the types for my regex_engine

package types;

import (
		"strconv"
)

type Data interface {
	String() string;
}

type Dot 			struct {};
type Asterisk 		struct {};
type Qus_mark 		struct {} ;
type Plus 			struct {};
type Character 		struct {
	Value rune;
}
type Range 			struct {
	N int;
}
type Open_paran 	struct {}
type Close_paran 	struct {}
type Open_literal 	struct {}
type Close_literal 	struct {}
type Dash 			struct {}
type Escape 		struct {
	Value rune;
}
type Comma 			struct {};
type Open_Range 	struct {};
type Close_Range 	struct {};
//This will contain all this above defined quantifiers 
type Quantifier 	struct {
	Data Data;
}



//Defining the functions 
func (d Dot) String() string {
	return ".";
}
func (d Open_Range) String() string{
	return "{";
}

func (d Close_Range) String() string {
	return "}";
}
func (a Asterisk) String() string {
	return "*";
}

func (d Qus_mark) String() string{
	return "?";
}
func (d Plus) String() string {
	return "+";
}
func (c Character)  String() string{
	return string(c.Value);
}

func (r Range)  String() string{
	return (strconv.Itoa(r.N));	
}

func (o Open_paran)  String() string{
	return "(";
} 
func (o Close_paran) String() string {
	return ")";
} 
func (o Open_literal) String() string {
	return "[";
}
func (o Close_literal) String() string {
	return "]";
}
func (d Dash) String() string {
	return "-";
}
func (e Escape) String() string {
	return string(e.Value);
}

func (c Comma) String() string{
	return ",";
}

