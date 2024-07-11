package main

import (
    "fmt"
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
    "math"
    "os"
    "strconv"
    "time"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run main.go <frequency in Hz> <duration in seconds>")
        return
    }

    frequency, err := strconv.ParseFloat(os.Args[1], 64)
    if err != nil {
        fmt.Printf("Invalid frequency: %s\n", os.Args[1])
        return
    }

    duration, err := strconv.ParseFloat(os.Args[2], 64)
    if err != nil {
        fmt.Printf("Invalid duration: %s\n", os.Args[2])
        return
    }

    ticker := time.NewTicker(1 * time.Millisecond) // Adjust the interval as needed
    defer ticker.Stop()

    startTime := time.Now()
    var pts plotter.XYs

    for now := range ticker.C {
        elapsed := now.Sub(startTime).Seconds()
        if elapsed > duration {
            break
        }

        // Calculate the current point in the sinusoidal wave
        sineValue := math.Sin(2 * math.Pi * frequency * elapsed)
        pts = append(pts, plotter.XY{X: elapsed, Y: sineValue})
    }

    // Create a new plot
    p := plot.New()

    p.Title.Text = "Sinusoidal Waveform"
    p.X.Label.Text = "Time (s)"
    p.Y.Label.Text = "Amplitude"

    // Make a line plotter and add it to the plot
    line, err := plotter.NewLine(pts)
    if err != nil {
        panic(err)
    }
    p.Add(line)

    // Save the plot to a PNG file
    if err := p.Save(10*vg.Inch, 4*vg.Inch, "sinusoidal.png"); err != nil {
        panic(err)
    }

    fmt.Println("Plot saved as sinusoidal.png")
}
