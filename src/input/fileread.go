// 파일 리드에 관한 것

package input

import(
	"bufio"
	"os"
	"strconv"
)

//string 읽어오기
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

//float64 읽어오기
func FilereadFloat(filename string)([]float64,error){
	var numbers []float64
	file,err:=os.Open(filename)

	if err!=nil{
		return nil,err
	}
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		number,err:=strconv.ParseFloat(scanner.Text(),64)
		if err!=nil{
			return nil, err
		}
		numbers=append(numbers,number)
	}
	err=file.Close()
	if err!=nil{
		return nil, err
	}
	if scanner.Err()!=nil{
		return nil, scanner.Err()
	}
	return numbers,nil
}

//int 읽어오기
func FilereadInt(filename string)([]int,error){
	var numbers []int
	file,err:=os.Open(filename)

	if err!=nil{
		return nil,err
	}
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		number,err:=strconv.Atoi(scanner.Text())
		if err!=nil{
			return nil, err
		}
		numbers=append(numbers,number)
	}
	err=file.Close()
	if err!=nil{
		return nil, err
	}
	if scanner.Err()!=nil{
		return nil, scanner.Err()
	}
	return numbers,nil
}

