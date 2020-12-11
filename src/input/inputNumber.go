package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 에러 체크
func ErrCheck(e error){
	if e!=nil {
		log.Fatal(e)
	}
}

//키보드로부터 입력받은 것을 int형으로 바꿔주는 것
func InputInt()int{
	fmt.Print("\n숫자를 입력하세요 : ")
	reader :=bufio.NewReader(os.Stdin)
	input,err:=reader.ReadString('\n')
	ErrCheck(err)
	input=strings.TrimSpace(input)

	//선택하고자 하는 답
	intvalue,err :=strconv.Atoi(input)
	ErrCheck(err)
	return intvalue
}

func InputFloat()float64{
	fmt.Print("\n숫자를 입력하세요 : ")
	reader :=bufio.NewReader(os.Stdin)
	input,err:=reader.ReadString('\n')
	ErrCheck(err)
	input=strings.TrimSpace(input)

	//선택하고자 하는 답
	floatvalue,err :=strconv.ParseFloat(input,64)
	ErrCheck(err)
	return floatvalue
}


