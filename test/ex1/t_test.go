package goef

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var pkgdir = flag.String("pkgdir", "smchin", "dir of package containing embedded files")

func TestGenerateGoPackage(t *testing.T) {
	fmt.Println("somin")
	fmt.Println(os.Args[1:])
	fmt.Printf("%v\n", pkgdir)
	flag.Parse()
	fmt.Printf("%v\n", pkgdir)

}