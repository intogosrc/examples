package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	replacertest()

}

func replacertest() {
	src := "hello leo, hello jack, hello tom"

	r := strings.NewReplacer("hello", "hi", "leo", "somebody")

	dist := r.Replace(src)

	fmt.Println(dist)
}

func doreadertest() {
	fp,err := os.Open("./iotest")

	if err!=nil {
		panic(err)
	}

	defer fp.Close()

	readertest(fp)

	strfp := strings.NewReader("hi leo!")

	readertest(strfp)
}

func readertest(r io.Reader) {
	result,err := ioutil.ReadAll(r)

	if err!=nil {
		panic(err)
	}

	fmt.Println(string(result))
}

func buildertest() {
	bstr := new(strings.Builder)

	fmt.Println("the cap of builder is ",bstr.Cap())

	bstr.Grow(10) // 给 builder 增加容量，不建议使用，内核会在需要的时候自动分配

	fmt.Println("the cap of builder is ",bstr.Cap())

	bstr.Write([]byte("hello"))
	bstr.WriteByte(' ')
	bstr.WriteString("world")
	bstr.WriteRune('!')

	fmt.Println(fmt.Sprintf("cap of Builder is %d, len of Builder is %d, str of Builder is %s", bstr.Cap(), bstr.Len(), bstr.String()))

	bstr.Reset() // reset 对 Builder 重新初始化
	bstr.WriteString("hello golang!")

	fmt.Println(fmt.Sprintf("cap of Builder is %d, len of Builder is %d, str of Builder is %s", bstr.Cap(), bstr.Len(), bstr.String()))
}

func strtest() {
	str := "hello world"

	str2 := str[0:3]

	fmt.Println(str2)
}