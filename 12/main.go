package main

import (
	"github.com/PsyonixMonroe/AOCLib/lib"
)

type Plot struct {
	plant      rune
	posX       int
	posY       int
	fenceUp    bool
	fenceDown  bool
	fenceLeft  bool
	fenceRight bool
	scored     bool
}

func CreatePlot(r rune, row int, column int) Plot {
	return Plot{plant: r, fenceUp: false, fenceDown: false, fenceLeft: false, fenceRight: false, scored: false, posX: row, posY: column}
}

func MarkGrid(grid *lib.Grid[Plot]) {
	for i := range grid.Rows() {
		for j := range grid.Columns() {
			currentPlot := grid.Get(i, j).Get()
			plotUp := grid.Get(i-1, j)
			if !plotUp.Exists() || (plotUp.Exists() && plotUp.Get().plant != currentPlot.plant) {
				// up
				currentPlot.fenceUp = true
			}
			plotDown := grid.Get(i+1, j)
			if !plotDown.Exists() || (plotDown.Exists() && plotDown.Get().plant != currentPlot.plant) {
				// down
				currentPlot.fenceDown = true
			}
			plotLeft := grid.Get(i, j-1)
			if !plotLeft.Exists() || (plotLeft.Exists() && plotLeft.Get().plant != currentPlot.plant) {
				// left
				currentPlot.fenceLeft = true
			}
			plotRight := grid.Get(i, j+1)
			if !plotRight.Exists() || (plotRight.Exists() && plotRight.Get().plant != currentPlot.plant) {
				// right
				currentPlot.fenceRight = true
			}
		}
	}
}

func ScoreGrid(grid lib.Grid[Plot]) int {
	score := 0
	for i := range grid.Rows() {
		for j := range grid.Columns() {
			currentPlot := grid.Get(i, j).Get()
			if !currentPlot.scored {
				// process field
				workQueue := lib.NewQueue[*Plot]()
				numPlots := 0
				numFence := 0
				currentPlot.scored = true
				workQueue.Enqueue(currentPlot)
				for plotOpt := workQueue.Dequeue(); plotOpt.Exists(); {
					numPlots++
					plot := plotOpt.Get()
					if plot.fenceUp {
						numFence++
					} else {
						adjPlot := grid.Get(plot.posX-1, plot.posY).Get()
						if adjPlot.plant == currentPlot.plant && !adjPlot.scored {
							adjPlot.scored = true
							workQueue.Enqueue(adjPlot)
						}
					}
					if plot.fenceDown {
						numFence++
					} else {
						adjPlot := grid.Get(plot.posX+1, plot.posY).Get()
						if adjPlot.plant == currentPlot.plant && !adjPlot.scored {
							adjPlot.scored = true
							workQueue.Enqueue(adjPlot)
						}
					}
					if plot.fenceLeft {
						numFence++
					} else {
						adjPlot := grid.Get(plot.posX, plot.posY-1).Get()
						if adjPlot.plant == currentPlot.plant && !adjPlot.scored {
							adjPlot.scored = true
							workQueue.Enqueue(adjPlot)
						}
					}
					if plot.fenceRight {
						numFence++
					} else {
						adjPlot := grid.Get(plot.posX, plot.posY+1).Get()
						if adjPlot.plant == currentPlot.plant && !adjPlot.scored {
							adjPlot.scored = true
							workQueue.Enqueue(adjPlot)
						}
					}

					plotOpt = workQueue.Dequeue()
				}
				score += numPlots * numFence
			}
		}
	}

	return score
}
