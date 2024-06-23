package main

import (
"fmt"
"strings"
"strconv"
//"math"
)

var new string //стока введенная пользователем
var r int = 0 //количество римских цифр во веденной строке
var o int = 0 //количество исполняемых операторов во введенной строке
var oper []string //операнды
var type_o int //тип оператора

var rom = map[int]string{
1:"I",
2:"II",
3:"III",
4:"IV",
5:"V",
6:"VI",
7:"VII",
8:"VIII",
9:"IX",
10:"X",
}

var rom_d = map[int]string{
1:"X",
2:"XX",
3:"XXX",
4:"XL",
5:"L",
6:"LX",
7:"LXX",
8:"LXXX",
9:"XC",
10:"C",
}

func operator(){ //поиск исполняемых операторов
	new_copy := new
	for i:=1; i<len([]rune(new_copy)); i++{
		switch(true) {
			case strings.Contains(new_copy, "+"): 
				new_copy = strings.Replace(new_copy, "+", " ", 1)
				o=o+1
				type_o=0
			case strings.Contains(new_copy, "-"): 
				new_copy = strings.Replace(new_copy, "-", " ", 1)
				o=o+1
				type_o=1
			case strings.Contains(new_copy, "*"): 
				new_copy = strings.Replace(new_copy, "*", " ", 1)
				o=o+1
				type_o=2
			case strings.Contains(new_copy, "/"): 
				new_copy = strings.Replace(new_copy, "/", " ", 1)
				o=o+1
				type_o=3
		}
		if o>1 {
			return
		}
	}
	oper = strings.Split(new_copy, " ")//деление строки на операнды
	
}

func compare(X string) bool{ //сверка операндов
	switch(X) {
		case string("1"): 
			return true
		case string("2"): 
			return true
		case string("3"): 
			return true
		case string("4"): 
			return true
		case string("5"): 
			return true
		case string("6"): 
			return true
		case string("7"): 
			return true
		case string("8"): 
			return true
		case string("9"): 
			return true
		case string("10"): 
			return true
		default: 
			return false
	}
}

func calc(x int, y int) int{ //подсчет
	var result int
	if type_o==0{
		result=x+y
	}else if type_o==1{
		result=x-y
	}else if type_o==2{
		result=x*y
	}else if type_o==3{
		result=x/y
	}
	return result
}

func replace(X string) string{ // Замена римских цифр на арабские
	for i:=0; i<len([]rune(X)); i++{
		switch(true) {
			case strings.Contains(X, "IX"): 
				X = strings.Replace(X, "IX", "9", 1)
				r=r+1
			case strings.Contains(X, "X"): 
				X = strings.Replace(X, "X", "10", 1)
				r=r+1
			case strings.Contains(X, "VIII"): 
				X = strings.Replace(X, "VIII", "8", 1)
				r=r+1
			case strings.Contains(X, "VII"): 
				X = strings.Replace(X, "VII", "7", 1)
				r=r+1
			case strings.Contains(X, "VI"): 
				X = strings.Replace(X, "VI", "6", 1)
				r=r+1
			case strings.Contains(X, "IV"): 
				X = strings.Replace(X, "IV", "4", 1)
				r=r+1
			case strings.Contains(X, "III"): 
				X = strings.Replace(X, "III", "3", 1)
				r=r+1
			case strings.Contains(X, "II"): 
				X = strings.Replace(X, "II", "2", 1)
				r=r+1
			case strings.Contains(X, "V"): 
				X = strings.Replace(X, "V", "5", 1)
				r=r+1
			case strings.Contains(X, "I"): 
				X = strings.Replace(X, "I", "1", 1)
				r=r+1
			default: 
				return X
		}
	}
	return X
}

func replace_rom(x int) string{//перевод арабских чисел в римские
	var result string = rom_d[x/10]
	result = result+rom[x%10]
	return result
}

func main() {
	
	fmt.Println("Введите пример состоящий из одного оператора и двух операндов римскими или арабскими числами")
	fmt.Scanln(&new)
	operator()
	
	if o==0 {
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
	}else if o>1{ 
		fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}else{
		oper[0]=replace(oper[0])//преобразование римских цифр первого операнда в арабские str
		oper[1]=replace(oper[1])//преобразование римских цифр второго операнда в арабские str
		
		if compare(oper[0])==false{ //сравнение операндов
			fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — введен не подходящий первый операнд.")
		}else{
			if compare(oper[1])==false{
				fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — введен не подходящий второй операнд.")
			}else{
		
				oper1, err1 := strconv.Atoi(oper[0])//преобразование первого операнда в int
				oper2, err2 := strconv.Atoi(oper[1])//преобразование второго операнда в int
				_ = err1
				_ = err2
				result:= calc(oper1, oper2)//счет
		
				if r==1{ // определение римский счет или арабский
					fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
				}else if r>1{//счет римский
					if result <= 0{
						fmt.Println("Выдача паники, так как в римской системе только натуральные числа.")
					}else{
						result_rom := replace_rom(result)
						fmt.Println("Ответ")
						fmt.Println(new,"=", result_rom)
					}
				}else if r==0{//счет арабский
					fmt.Println("Ответ")
					fmt.Println(new,"=", result)
				}
			}
		}
	}
			
}