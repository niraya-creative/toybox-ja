package main

import (
	"flag"
	"fmt"
	"os"

	imgconv "github.com/gohandson/toybox-ja/solution/section06/step01"
)

var (
	flagTo   = imgconv.PNG
	flagFrom = imgconv.JPEG
)

func init() {
	flag.Var(&flagTo, "to", "after format")
	flag.Var(&flagFrom, "from", "before format")
}

func main() {
	// TODO: convertAllはパッケージ名をつけてエクスポートされたものに変える
	if err := imgconv.convertAll(os.Args[1], flagFrom, flagTo); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}
