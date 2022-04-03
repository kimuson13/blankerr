package main

import (
	"github.com/kimuson13/blankerr"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(blankerr.Analyzer) }
