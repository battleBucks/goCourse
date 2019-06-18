package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	First   string
	Age     int
	Sayings []string
}

func main() {

	js := `[{"First":"James","Age":32,"Sayings":["Shaken, not stirred", "blah blah blaaaaaah"]},{"First":"Moneypenny","Age":27,"Sayings":["MoneyMoneyMoney"]},{"First":"M","Age":54,"Sayings":["Hello, I'm M", "Another M saying"]}]`

	jd := []people{}
	err = json.Unmarshal([]byte(js), &jd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jd)
}
