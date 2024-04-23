package main
import ("fmt")


/////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////// PRINT STATEMENT //////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////
/* 
	fmt.Println needs double quote adds a new line
   	Println coma adds up a space in output
	\n for new line
	in print statement , doesn't add a space (applies only for string outputs)
	printf is used to mention the formating of the data type
	printf: %v - prints the value of variable using default format
			%T - prints the type of variable
			%d, %x, %o - integer in decimal(10), hexadecimal(16)(lowercase), octal(8) respectively
			%f, %e - floating point numbers
			%#v - for go syntax format
			%% - prints %
			%b - base 2
			%+d - decimal with sign
			%O - octal with 0o at front
			%X - upper case Hexadecimal
			%#x - base 16 with 0x in front
			%4d - padding with 4 places
			%-4d - padding with -4 places
			%04d - padding with 0 and 4 places
			%s - string
			%q - string with double quote
			%8s and %-8s - left justified or right justified with width of 8 spaces
			%x - hex value of string
			% x - hex with spaces
			%t - boolean
			%e - exponent
			%.2f - float number with precision upto .2
			%6.2f -  total width of 6 and precision upto .2
			%g - imp part of long decimal
*/


/////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////// VARIABLES /////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////
/*
	my god this variables, if we use "var" no : needed
	when declaring and initialising at the same time variable_name := value and we can't performe this out of a function
	ya obvi worst thing u can't keep a variable unused
	you can create untyped variables
	default values string "", int 0, bool false
	variable name must start with letter or underscore _
	cannot contain special characters except for underscore
	an const (constant) variable can be created, whether it's going to be used or not doesn't matter
*/


/////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////// FUNTIONS //////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////
// main funtion just like c c++ and java


func main(){

	fmt.Println("\nhello - fmt.Println")

	var insideMainVarInt int8 = 2
	fmt.Println("\nint8 variable using fmt.Println :",insideMainVarInt)

	insideMainVarFloat := 5.5
	fmt.Println("\nfloat variable (:=) using fmt.Println :",insideMainVarFloat)

	var a,b,c,d = 1,2,3,false
	fmt.Println("\nmultiple declaration :",a+b,c,d)

	var(
		emptint int
		valuedfloat = 2.5
		stringA = "A string"
	)
	fmt.Println("\nThis line of output from var declaration block. default value of int",emptint,"float valur",valuedfloat,"this ia",stringA)

	var _ap1_ple1 int = 0
	fmt.Println(_ap1_ple1)

	const constint int = 10
	fmt.Println("\nconstant integer: ",constint)

	var printStatementExpStr string = "hello"
	fmt.Print("\na string using fmt.Print :",printStatementExpStr)
	fmt.Printf("\na string using fmt.Printf :",printStatementExpStr)
	fmt.Println("\na string using fmt.Println :",printStatementExpStr)

	var(
		printfValA = "hiiiii"
		printfValB = 56
	)
	fmt.Printf("\nprintfValA value : %v, type : %T",printfValA,printfValA)
	fmt.Printf("\nprintfValB value : %v, type : %T",printfValB,printfValB)
	fmt.Printf("\n%%g : %g", 3.145298492484294849849249)

}
