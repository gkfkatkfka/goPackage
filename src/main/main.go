package main

import (
	"../input"
	"fmt"
	"log"
)

func errCheck(e error){
	if e!=nil{
		log.Fatal(e)
	}
}
func main(){
	// int로 입력
	fmt.Print(input.InputInt())

	//float64로 입력
	fmt.Println(input.InputFloat())

	//=============================================================================================================================//

	/*파일 읽어오기*/
	// float64
	floats,err:=input.FilereadFloat("C:\\Users\\gkfka\\Documents\\대학\\대2 2학기\\오픈소스프로그래밍\\packages\\src\\main\\float.txt")
	errCheck(err)
	fmt.Println(floats)

	// int
	ints,err:=input.FilereadInt("C:\\Users\\gkfka\\Documents\\대학\\대2 2학기\\오픈소스프로그래밍\\packages\\src\\main\\int.txt")
	errCheck(err)
	fmt.Println(ints)

	// string
	strings,err:=input.FilereadString("C:\\Users\\gkfka\\Documents\\대학\\대2 2학기\\오픈소스프로그래밍\\packages\\src\\main\\string.txt")
	errCheck(err)
	fmt.Println(strings)

	//=============================================================================================================================//

	/*선택창 만들기*/
	fmt.Println("1) 2) 3) 4) 5) ")
	fmt.Print("숫자 선택 : ")

	switch input.InputInt() {
	case 1:
		fmt.Println("1번 선택")
	case 2:
		fmt.Println("2번 선택")
	case 3:
		fmt.Println("3번 선택")
	case 4:
		fmt.Println("4번 선택")
	case 5:
		fmt.Println("5번 선택")
	case 6:
		fmt.Println("6번 선택")
	default:
		fmt.Println("제대로 다시 선택해주세요")
	}
}
