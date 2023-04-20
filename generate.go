package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
)

const (
	rootFolder = "./break/"

	start = 1000
	nMultiplier = 2
	nStep =10
)

func main() {
	if err := os.MkdirAll(rootFolder, 0777); err != nil {
		panic(err)
	}

	for i := 1; i <= nStep; i++ {
		nn := start * int((math.Pow(nMultiplier, float64(i))))
		pkgName := "drink"+strconv.Itoa(nn)+"cups"
		pkgName = strings.Replace(pkgName, "000000cups", "Mcups", -1)
		pkgName = strings.Replace(pkgName, "000cups", "Kcups", -1)

		if err := os.MkdirAll(rootFolder + pkgName, 0777); err != nil {
			panic(err)
		}

		generateGenerics(rootFolder, pkgName, nn)
	}
}

func generateGenerics(root, pkgName string, nn int) {
	s := "package "+pkgName+"\n\n"
	s += `
type number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint
}
`
	generate(root+pkgName+"/generics.go", s, nn,
		func(i int) string {

			return `
func MaxOfType` + strconv.Itoa(i) + `[T number](a, b T) T {
	if a > b {
		return a
	}

	return b
}

`
		})
}

func generate(fileName, beginning string, nn int, f func(int) string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(beginning)
	if err != nil {
		panic(err)
	}
	defer writer.Flush()

	for i := 0; i < nn; i++ {
		_, err = writer.WriteString(f(i))
		if err != nil {
			panic(err)
		}
	}
}

type xhf string

func (s xhf) String() string {
	return ""
}

type Stringer interface {
	String() string
	xhf
}

func Tos[T Stringer](s []T) []string {
	var ret []string
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}
