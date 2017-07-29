// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ { // в этом примере может заинтересовать разница
										// между выполнением в интерпретаторе и
										// откомпилированном бинарнике. Если указать
										// i :=0 то можно увидеть первый элемент
										// среза командной строки - временный каталог с исполняемым файлом
										// /tmp/go-build295781861/command-line-arguments/_obj/exe/main
										// обратите внимание на "_obj/exe/main"
		s += sep + os.Args[i]			// += обычное переприсваивание в итеррации
		sep = " "						// разделитель объявляется после первой итеррации
	}
	fmt.Println(s)
}

//!-
