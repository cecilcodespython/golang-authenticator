package main

import (
	"fmt"
	"strings"
)

func main(){
	for {
		var userPassword string
		fmt.Println("Enter Your Password \n(Must be longer than 8 characters): ")
		fmt.Scan(&userPassword)

		isShort := len(userPassword) <8
		hasCharacter := strings.ContainsAny(userPassword,"!@#$%^&*()+")

		if isShort{
			fmt.Println("Short Password,Enter a longer password")
			continue
		}else if hasCharacter {
			fmt.Println("Password is Secure")
			break	
		}else if !isShort && !hasCharacter {
			fmt.Println("You Password Must Include A Character")

		}



	}
}