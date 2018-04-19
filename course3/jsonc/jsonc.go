package main

import "encoding/json"
import "fmt"

func main() {
	m := map[string]string{
		"foo": "bar",
	}

	buf, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	s := string(buf)
	fmt.Println(s)

	var mm map[string]string
	err = json.Unmarshal([]byte(s), &mm)
	if err != nil {
		panic(err)
	}
	fmt.Println(mm)

	var foo Foo
	err = json.Unmarshal(buf, &foo)
	if err != nil {
		panic(err)
	}
	fmt.Println(foo)

	foo2 := Foo{
		Bar: "Foo.Bar",
		Foo: Bar{
			Foo: "Bar.Foo",
		},
	}
	buf, err = json.Marshal(foo2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}

type Foo struct {
	Bar string `json:"foo"`
	Foo Bar    `json:"bar"`
}

type Bar struct {
	Foo string `json:"bar"`
}
