package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var pkgdir = flag.String("pkgdir", "smchin", "dir of package containing embedded files")

func Test_main(t *testing.T)  {
	fmt.Println(os.Args)
}
