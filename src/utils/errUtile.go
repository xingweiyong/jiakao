package utils

import (
	"fmt"
	"os"
)

func HandlerError(err error,when string)  {
	if err != nil{
		fmt.Println(when,err)
		os.Exit(1)
	}
}
