package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)
var hold [][]int
var amnt, con, change []int
var id []string
var cnt int
func main() {
	var name, ttl string
	fmt.Println("File Name")
	fmt.Scanln(&name)
	fmt.Println("Total amount in dollars")
	fmt.Scanln(&ttl)

	lf,e := strconv.ParseFloat(ttl,64)
	if e != nil {
		fmt.Println(ttl, " Not a number bud")
	}
	total := fltint(lf*100)
	inputs(name)
	for i := 0; i < len(amnt); i++ {
		fmt.Println(id[i],amnt[i],con[i])
	}
	sum :=0
	fmt.Println(dyn(total,amnt,con))
	fmt.Println(change)
	for i := 0; i < len(id); i++ {
		if change[i] != 0 {
			fmt.Println(id[i],amnt[i],change[i])
			sum += amnt[i]*change[i]
		}
	}
	fmt.Println(sum)
}


func dyn(amount int, coins []int, limits []int) int{

	coinsUsed := make([][]int,amount+1)
	for i := 0; i <= amount; i++ {
		coinsUsed[i]= make([]int,len(coins))
	}

	minCoins := make([]int,amount+1)
	for i := 1; i <= amount; i++ {
		minCoins[i]=math.MaxInt32
	}

	limitsCopy := make([]int,len(limits))
	copy(limitsCopy,limits)


	for i := 0; i < len(coins); i++ {
		for limitsCopy[i]>0 {
			for j := amount; j >= 0; j-- {
				currAmount := j+coins[i]
				if currAmount <= amount {
					if minCoins[currAmount] > minCoins[j] + 1  {
						//fmt.Println("AYAdadYAY")
						minCoins[currAmount] = minCoins[j] + 1
						copy(coinsUsed[currAmount],coinsUsed[j])

						coinsUsed[currAmount][i] += 1
					}
				}
			}
			limitsCopy[i] -= 1
		}
	}
	if minCoins[amount] == math.MaxInt {
		change = nil
		return 0
		//fmt.Println("Oh noes")

	}

	change = coinsUsed[amount]
	//fmt.Println(change)
	return minCoins[amount]
}


//Data struct imports
func inputs(ams string) {

	file, err := os.Open(ams+".csv")
	if err != nil{
		fmt.Println("Unable 2 read")
		log.Fatal(err)
	}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		spit := strings.Split(scan.Text(),",")
		id = append(id, spit[0])
		amnt = append(amnt, fltint(strflt(spit[1])*100))
		con = append(con, s2i(spit[2]))
	}

}
//Conversions
func strflt(s string) float64 {
	ret, e := strconv.ParseFloat(s,64)
	if e!=nil{
		fmt.Println(s, " is not a float")
	}
	return ret

}
func str2ints(s []string) []int {
	var ret []int
	for i := 0; i < len(s); i++ {
		ret = append(ret, s2i(s[i]))
	}
	return ret

}
func s2i(s string) int{
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("error with ", s)
	}
	return i
}
func fltint(a float64) int  {

	return int(math.RoundToEven(a))
}
