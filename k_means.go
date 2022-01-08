package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

var iterationNum int = 0
var colors []color.RGBA
const doPlotting bool = false
const doLogging bool = false

func plotCurrentIteration(points, centroids []vec2d, pointOwnerIds []int) {
    if !doPlotting { return }
    
    p := plot.New()

    // groups := make([]plotter.XYs, len(centroids))
    for i := range centroids {
        currentPoints := make(plotter.XYs, 0)

        for j := range points {
            if pointOwnerIds[j] != i { continue }

            currentPoints = append(currentPoints, plotter.XY{X: points[j].x, Y: points[j].y})
        }

        log.Println(currentPoints)

        s, _ := plotter.NewScatter(currentPoints)
        s.GlyphStyle.Color = colors[i]
        p.Add(s)
    }

    // centroidsScatter := make(plotter.XYs, len(centroids))
    for i := range centroids {
        centroidScatter := make(plotter.XYs, 1)
        centroidScatter[0] = plotter.XY{X: centroids[i].x, Y: centroids[i].y}
        
        l, _ := plotter.NewScatter(centroidScatter)
        l.GlyphStyle.Color = colors[i]
        l.Shape = draw.PyramidGlyph{}
        p.Add(l)
    }

    p.Save(7*vg.Inch, 7*vg.Inch, "./output/" + fmt.Sprint(iterationNum) + ".png")
}

func calculateNewOwners(points, centroids [] vec2d, pointOwnerIds []int) bool {
    change := false
    
    for pointId := range points {

        oldOwner := pointOwnerIds[pointId]
        minDistance := math.MaxFloat64
        
        // Find centroid with minimal distance
        for centroidId := range centroids {
            distance := getDistance(centroids[centroidId], points[pointId])

            if distance < minDistance {
                minDistance = distance
                pointOwnerIds[pointId] = centroidId
            }
        }

        if oldOwner != pointOwnerIds[pointId] {
            change = true
        }
    }

    return change
}

func calculateNewCentroids(points, centroids []vec2d, pointOwnerIds []int) {
    for centroidId := range centroids {
        sum := vec2d{0, 0}
        count := 0

        for pointId := range points {
            // Calculate only owned points
            if pointOwnerIds[pointId] != centroidId { continue }

            sum = sum.add(points[pointId])            
            count++
        }

        sum = sum.div(float64(count))

        centroids[centroidId] = vec2d{sum.x, sum.y}
    }
}

func iteration(points, centroids []vec2d, pointOwnerIds []int) bool {
    // Steps:
    // 1. Change ownership
    // 2. Calculate new centroids

    // If ownership didn't change, end program

    log.Println("=======================")
    log.Println("Step: ", iterationNum)
    
    change := calculateNewOwners(points, centroids, pointOwnerIds)
    log.Println("New owners: ", pointOwnerIds)

    calculateNewCentroids(points, centroids, pointOwnerIds)
    log.Println("New centroids: ", centroids)

    log.Println("Step ", iterationNum, " done")
    log.Println("=======================")

    iterationNum++

    plotCurrentIteration(points, centroids, pointOwnerIds)

    return change
}

func initCentroids(centroids []vec2d, points []vec2d) {
    for i := range centroids {
        centroids[i] = vec2d{x: points[i].x, y: points[i].y}
    }
}

func initPoints(points []vec2d) {
    for i := range points {
        points[i] = randVec2d(100)
    }
}

func main() {
    if !doLogging { log.SetOutput(ioutil.Discard) }

    if len(os.Args) < 4 {
        fmt.Println("Usage: ./k_means -- <number of points> <number of centroids>")
    }
    
    arraySize, _ := strconv.Atoi(os.Args[2])
    k, _ := strconv.Atoi(os.Args[3])

    if arraySize < k {
        fmt.Println("Array must be larger than k")
        return
    }

    // Checks done, proceeding with running the program

    rand.Seed(10100)

    os.RemoveAll("./output/")
    os.Mkdir("./output/", 0755)

    // Declarations
    points := make([]vec2d, arraySize)
    pointOwnerIds := make([]int, arraySize)
    centroids := make([]vec2d, k)

    for range centroids {
        colors = append(colors, color.RGBA{R: uint8(rand.Float32() * 255), G: uint8(rand.Float32() * 255), B: uint8(rand.Float32() * 255), A: 255})
    }

    // Initialization    
    initPoints(points)
    initCentroids(centroids, points)
    
    // Status log
    fmt.Println("Number of points: ", arraySize)
    fmt.Println("Number of centroids: ", k)
    log.Println("Points: ", points)
    log.Println("Centroids: ", centroids)
    log.Println("Owners: ", pointOwnerIds)

    fmt.Println("Initialized")

    plotCurrentIteration(points, centroids, pointOwnerIds)

    start := time.Now()

    // Main loop
    for {
        change := iteration(points, centroids, pointOwnerIds)

        if !change { break }
    }

    elapsed := time.Since(start)

    fmt.Println("Time: ", elapsed)

    // Exit log
    fmt.Println("Converged")
    fmt.Println("Steps: ", iterationNum)
    fmt.Println("Centroids: ", centroids)
}
