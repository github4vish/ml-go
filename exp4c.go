package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Create a sample data set
	n := 100
	data := make(plotter.Values, n)
	for i := range data {
		data[i] = rand.NormFloat64()
	}

	// Create a separate box plot
	pBox := plot.New()
	pBox.Title.Text = "Box Plot Example"
	pBox.Y.Label.Text = "Values"

	// Create the box plot
	boxPlot, err := plotter.NewBoxPlot(vg.Points(50), 0, data)
	if err != nil {
		panic(err)
	}

	pBox.Add(boxPlot)

	// Save the box plot to a PNG file
	if err := pBox.Save(6*vg.Inch, 6*vg.Inch, "boxplot.png"); err != nil {
		panic(err)
	}

	// Create a separate histogram plot
	pHist := plot.New()
	pHist.Title.Text = "Histogram Example"
	pHist.Y.Label.Text = "Values"

	// Create the histogram
	h, err := plotter.NewHist(data, 16)
	if err != nil {
		panic(err)
	}
	h.Normalize(1)

	// Add histogram to the histogram plot
	pHist.Add(h)

	// Save the histogram plot to a PNG file
	if err := pHist.Save(6*vg.Inch, 6*vg.Inch, "histogram.png"); err != nil {
		panic(err)
	}
}
