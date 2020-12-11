package main

import (
	"../input"
	"bufio"
	"fmt"
	"log"
	"os"
)

func FilereadString(filename string)([]string,error){
	var strs []string
	file,err:=os.Open(filename)

	if err!=nil{
		return nil,err
	}
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		str:=scanner.Text()
		strs=append(strs,str)
	}
	err=file.Close()
	if err!=nil{
		return nil, err
	}
	if scanner.Err()!=nil{
		return nil, scanner.Err()
	}
	return strs,nil
}

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

	//실제 경로를 확인하기 위해 - 이렇게 안 하면 오류남
	path,err:=os.Getwd()
	path=path+"\\src\\"

	// string
	strings,err:=FilereadString(path+"string.txt")
	errCheck(err)
	fmt.Println(strings)

	// float64
	floats,err:=input.FilereadFloat(path+"float.txt")
	errCheck(err)
	fmt.Println(floats)

	// int
	ints,err:=input.FilereadInt(path+"int.txt")
	errCheck(err)
	fmt.Println(ints)



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
