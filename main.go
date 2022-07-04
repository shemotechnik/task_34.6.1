package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
	"strconv"
)

func main(){
	str := ReadFile("./input.txt")
	ParseStrAndWrite(str, "./output.txt")
}

func ParseStrAndWrite(s string, fn string){
	_ = os.Remove(fn) // удаляем файл на диске
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	re := regexp.MustCompile(`([\d]+)([\+\-]{1})([\d]+)=\?`)
	submatches := re.FindAllStringSubmatch(s, -1)
	for _,v := range submatches{
		x1,_:=strconv.Atoi(v[1])
		x2,_:=strconv.Atoi(v[3])
		y:=0
		switch v[2] {
		case "+" : y = x1+x2
		case "-" : y = x1-x2
		}
		newS := strings.Replace(v[0], "?", strconv.Itoa(y), 1)
		_, _ = f.WriteString(newS+"\n")
	}
	fmt.Println("Done")
}

func ReadFile(n string) string{
	var ret string
	filename := n

	f, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fileReader := bufio.NewReader(f)

	for {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			break
		}
		ret+=string(line)+"\n"
	}
	return ret
}
