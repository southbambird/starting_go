package main

import "fmt"

type Stringify interface {
    ToString() string
}

type Person struct {
    Name string
    Age int
}

func (p *Person) ToString() string {
    return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
    Number string
    Model string
}

func (c *Car) ToString() string {
    return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

func main() {
    vs := []Stringify{
        &Person{Name: "Taro", Age: 21},
        &Car{Number: "XXX-1234", Model: "PX512"},
    }
    for _, v := range vs {
        fmt.Println(v.ToString())
    }
}
