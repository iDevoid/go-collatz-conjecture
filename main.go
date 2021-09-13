package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

var (
	num   int
	print bool
	delay int
)

func init() {
	flag.IntVar(&num, "num", 0, "start number of collatz conjecture")
	flag.BoolVar(&print, "print", false, "Prints all values of collatz conjecture")
	flag.IntVar(&delay, "delay", 10, "delay of printing in millisecond")
	flag.Parse()
}

func main() {
	if num == 0 {
		return
	}

	pin := float64(num)
	yAxis := []float64{
		pin,
	}
	ctr := float64(1)
	xAxis := []float64{
		ctr,
	}
	// var labels []chart.Value2

	for {
		if print {
			fmt.Printf("Counter-%d Value: %d\n", int(ctr), int(pin))
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}

		if pin == 1.0 {
			break
		}

		ctr++
		switch int(pin) % 2 {
		case 1: // odd
			pin = (3.0 * pin) + 1.0
		case 0: // even
			pin = pin / 2.0
		}

		yAxis = append(yAxis, pin)
		xAxis = append(xAxis, ctr)
		// labels = append(labels, chart.Value2{
		// 	XValue: ctr,
		// 	YValue: pin,
		// 	Label:  fmt.Sprint(pin),
		// })
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "Step",
			Ticks: []chart.Tick{
				{
					Value: 1,
					Label: "1",
				},
				{
					Value: ctr,
					Label: fmt.Sprint(ctr),
				},
			},
		},
		YAxis: chart.YAxis{
			Name: "Value",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xAxis,
				YValues: yAxis,
			},
			// chart.AnnotationSeries{
			// 	Annotations: labels,
			// },
		},
	}

	f, _ := os.Create(fmt.Sprint("output-", num, ".png"))
	defer f.Close()
	graph.Render(chart.PNG, f)

	if print {
		fmt.Println(yAxis)
	}
}
