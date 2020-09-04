package demo

import (
	"fmt"
	"github.com/lu4p/cat"
)

func ExtractText(){
	txt, _ := cat.File("/home/devadmin/redisbook.pdf")
	//txt, _ := cat.File("/home/devadmin/test.py")
	fmt.Println(txt)
}

