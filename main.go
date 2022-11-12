package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/guptarohit/asciigraph"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	template := "%{time_total}"

	url := os.Args[1]
	data := []float64{}

	for {

		time.Sleep(time.Second * 1)
		out, err := exec.Command("curl", url, "-w", template, "-o", "/dev/null").Output()
		if err != nil {
			fmt.Println(err)
		}

		float, err := strconv.ParseFloat(string(out), 64)

		if len(data) > 69 {
			data = data[1:]
		}

		data = append(data, float)

		fmt.Print("\033[H\033[2J")
		graph := asciigraph.Plot(data, asciigraph.Height(10), asciigraph.Precision(6))

		title := figure.NewFigure("CloudWatch.sh", "", true)
		title.Print()

		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println(graph)
		fmt.Println("--------------------------------------------------------------------------------")

	}
}
