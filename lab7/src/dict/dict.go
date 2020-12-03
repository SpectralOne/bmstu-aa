package dict

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/brianvoe/gofakeit"
)

// Dictionary - dictionary type
type Dictionary map[string]string

// FA - Frequent Analysis struct
type FA struct {
	letter string
	count  int
	array  []Dictionary
}

func isInSlice(slice []Dictionary, i Dictionary) bool {
	for _, v := range slice {
		if v["name"] == i["name"] {
			return false
		}
	}
	return true
}

// Create - creates sized array of dictionaries
func Create(size int) (d []Dictionary) {
	d = make([]Dictionary, size)

	for i := 0; i < size; i++ {
		res := false
		var company Dictionary
		// govno := 0
		for res != true {
			company = Dictionary{
				"name":  gofakeit.LoremIpsumParagraph(2, 10, 10, " "),
				"email": gofakeit.Email(),
				"city":  gofakeit.City(),
			}
			res = isInSlice(d[:i], company)
		}

		d[i] = company
	}

	return
}

func out(d []Dictionary) {
	for _, v := range d {
		fmt.Printf("Name: %s\nEmail: %s\nCity: %s\n\n", v["name"], v["email"], v["city"])
	}
}

func outSingle(v Dictionary) {
	fmt.Printf("Name: %s\nEmail: %s\nCity: %s\n\n", v["name"], v["email"], v["city"])
}

func pick(d []Dictionary, w string) string {
	i := rand.Int() % len(d)
	for _, v := range d[i:] {
		if v["name"][:1] == w {
			return v["name"]
		}
	}

	for _, v := range d {
		if v["name"][:1] == w {
			return v["name"]
		}
	}

	return d[i]["name"]
}

// Brute - Linear bruteforce
func Brute(d []Dictionary, name string) (res Dictionary) {
	res = d[0]
	for _, v := range d {
		if v["name"] == name {
			res = v
			return
		}
	}
	return
}

// Binary - Binary search
func Binary(d []Dictionary, name string) (result Dictionary) {
	mid := len(d) / 2
	switch {
	case len(d) == 0:
		return
	case d[mid]["name"] > name:
		result = Binary(d[:mid], name)
	case d[mid]["name"] < name:
		result = Binary(d[mid+1:], name)
	default:
		result = d[mid]
	}
	return
}

// FreqAnal - Frequent Analysis of names
func FreqAnal(d []Dictionary) (st []FA) {
	alf := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	st = make([]FA, len(alf))
	for i, v := range alf {
		a := FA{
			letter: string(v),
			count:  0,
			array:  make([]Dictionary, 0),
		}
		st[i] = a
	}

	for _, v := range d {
		l := v["name"][:1]
		for i := range st {
			if st[i].letter == l {
				st[i].count++
			}
		}
	}

	sort.Slice(st, func(i, j int) bool {
		return st[i].count > st[j].count
	})

	for i := range st {
		for j := range d {
			if d[j]["name"][:1] == st[i].letter {
				st[i].array = append(st[i].array, d[j])
			}
		}
		sort.Slice(st[i].array, func(l, m int) bool {
			return st[i].array[l]["name"] < st[i].array[m]["name"]
		})
	}

	return
}

// Combined - Frequent Analysis + Binary search
func Combined(fa []FA, n string) (r Dictionary) {
	letter := n[:1]
	r = fa[0].array[0]
	for _, v := range fa {
		if v.letter == letter {
			r = Binary(v.array, n)
		}
	}

	return r
}
