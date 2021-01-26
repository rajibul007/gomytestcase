package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	input := os.Args[1]
	size := len(input)
	rep := strings.Repeat("!", size)
	fmt.Println(rep + strings.ToUpper(input) + rep)
	//raw string
	html := `
	 
<html>
		 <body>
		 This is test 
		 </body>

</html>	 `

	fmt.Println(html)
	fmt.Println(utf8.RuneCountInString("hétérogénéité"))
	fmt.Println(len("hétérogénéité"))

	//path := "c:\\program files\\duper super\\fun.txt\n" +
	//	"c:\\program files\\really\\funny.png"
	path := `
c:\program files\duper super\fun.txt
c:\program files\really\funny.png`

	fmt.Println(path)

	//trim the string

	msg := `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(strings.TrimSpace(msg))
	msg = `   The Bang`
	fmt.Println(strings.TrimLeft(msg, " "))

}
