package main

import (
	"fmt"
	"github.com/oikomi/go_h264/h264"
	"github.com/oikomi/go_h264/nalu"
	"log"
	"os"
)

func help() {
	fmt.Printf("go_h264 filename \n")
}

func main() {
	if len(os.Args) != 2 {
		help()
		os.Exit(0)
	}
	fh := h264.New()
	if err := fh.Open(os.Args[1]); err != nil {
		log.Fatalln(err.Error())
	}

	nalu := &nalu.Nalu{}
	fh.GetAnnexbNALU(nalu)
	fmt.Println(nalu)
}
