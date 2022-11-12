package main

import (
	"fmt"
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

		time.Sleep(time.Second * 3)
		out, err := exec.Command("curl", url, "-w", template, "-o", "/dev/null").Output()
		if err != nil {
			fmt.Println(err)
		}

		float, err := strconv.ParseFloat(string(out), 64)

		data = append(data, float)

		fmt.Print("\033[H\033[2J")
		fmt.Println(float)
		graph := asciigraph.Plot(data, asciigraph.Height(10), asciigraph.Precision(6))

		fmt.Println("-------------------------------------------------------")
		fmt.Println(graph)
		fmt.Println("-------------------------------------------------------")

	}
}
