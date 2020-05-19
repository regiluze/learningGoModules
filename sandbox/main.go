package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	_, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		logrus.Fatal()
	}
}
