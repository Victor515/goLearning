package main

import (
	"regexp"
	"fmt"
)

const text = `My email is victor861@gmail.com@gmail.com
email 1 is      abc@example.org
email2 is xyz@haha. com
email3 is victor@qq.com.cn`

func main() {
	// ` ` is for regex string
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	subtext := re.FindAllString(text, -1)
	fmt.Println(subtext)

	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
