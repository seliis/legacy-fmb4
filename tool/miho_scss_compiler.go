package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	scss "github.com/wellington/go-libsass"
)

func find(s string) []string {
	var a []string
	r := path.Dir(s)

	filepath.Walk(r, func(p string, i os.FileInfo, e error) error {
		if filepath.Ext(p) == ".scss" {
			d := path.Dir(p)
			if d != r {
				a = append(a, d)
			}
		}
		return nil
	})

	fmt.Println(a)
	return a
}

func main() {
	dst := "./misc/miho.css"
	src := "./addr/"

	o, _ := os.Create(dst)
	f, _ := os.Open(src + "addr.scss")
	c, _ := scss.New(o, f)

	c.Option(
		scss.OutputStyle(
			scss.COMPRESSED_STYLE,
		),
		scss.IncludePaths(
			find(src),
		),
	)

	if e := c.Run(); e != nil {
		log.Fatal(e)
	}
}
