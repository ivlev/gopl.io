// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings" // Можно не танцевать с бубном
) // если задача - всего лишь склеить вместе аргументы командной строки

//!+
func main() {
	fmt.Println(strings.Join(os.Args[1:], " ")) // вопреки наивным ожиданиям string.Join воспринимает только два аргумента
}

//!-
