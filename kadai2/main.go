package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/gopherdojo-studyroom/kadai2/komazz/imgconv"
)

var (
	srcExt, dstExt          string
	errTooFewArgument       = errors.New("Too Few Arguments")
	errUnsupportedExtension = errors.New("Unsupported extension")
)

func init() {
	flag.StringVar(&srcExt, "s", "jpg", "Optional: Extension of Source Image.")
	flag.StringVar(&dstExt, "d", "png", "Optional: Extension of Destination Image.")
}

func exec() error {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		return errTooFewArgument
	}

	if !imgconv.ValidExt(srcExt) || !imgconv.ValidExt(dstExt) {
		return errUnsupportedExtension
	}

	c := imgconv.NewConverter(srcExt, dstExt)

	srcPath := args[0]
	dir, err := os.Stat(srcPath)
	if err != nil {
		return err
	}

	if dir.IsDir() {
		// ディレクトリ指定の場合
		fileList, err := c.SrcFileList(srcPath)
		if err != nil {
			return err
		}
		for _, src := range fileList {
			if err := c.Convert(src); err != nil {
				return err
			}
		}
	} else {
		// ファイル指定の場合
		if filepath.Ext(srcPath) != c.SrcExt {
			return errUnsupportedExtension
		}
		if err := c.Convert(srcPath); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := exec(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
