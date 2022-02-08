/*
1.Begin by importing the Go standard libraries

2. Next, add a loop to display the program’s default
location on your system, together with any passed-in
arguments

3.Now, display the final passed-in argument

(e.g)
Run the program as: go run command-flag.go a b c 

*/
package main

import (
	"fmt"
	"os"
)

func main(){
	//[]string Arg holds the arguments
	for i,v:=range os.Args{
		fmt.Printf("Arguments for index %d is %s\n",i,v)
	}

	//what is the final passed-in argument?
	fmt.Printf("The final passed arguments is %v\n",os.Args[len(os.Args)-1])

}