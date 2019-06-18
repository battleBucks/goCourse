package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type user struct {
	First string
	Age   int
}
type people struct {
	First   string
	Age     int
	Sayings []string
}

type ByName []people
type ByAge []people
type BySaying []people

func (a ByName) Len() int   { return len(a) }
func (a ByAge) Len() int    { return len(a) }
func (a BySaying) Len() int { return len(a) }

func (a ByName) Swap(i, j int)   { a[i], a[j] = a[j], a[i] }
func (a ByAge) Swap(i, j int)    { a[i], a[j] = a[j], a[i] }
func (a BySaying) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByName) Less(i, j int) bool   { return a[i].First < a[j].First }
func (a ByAge) Less(i, j int) bool    { return a[i].Age < a[j].Age }
//func (a BySaying) Less(i, j int) bool { return a[i].Sayings < a[j].Sayings }

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}
	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}
	u3 := user{
		First: "M",
		Age:   54,
	}
	users := []user{u1, u2, u3}
	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	js := `[{"First":"James","Age":32,"Sayings":["Shaken, not stirred", "blah blah blaaaaaah"]},{"First":"Moneypenny","Age":27,"Sayings":["MoneyMoneyMoney"]},{"First":"M","Age":54,"Sayings":["Hello, I'm M", "Another M saying"]}]`

	jd := []people{}
	err = json.Unmarshal([]byte(js), &jd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jd)

	fmt.Println("sorting by name.....")
	sort.Sort(ByName(jd))
	fmt.Println(jd)
	fmt.Println("sorting by age.....")
	sort.Sort(ByAge(jd))
	fmt.Println(jd)
	//fmt.Println("sorting by saying.....")
	//sort.Sort(BySaying(jd))
	//fmt.Println(jd)
}
