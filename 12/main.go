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

func ScoreGrid(grid lib.Grid[Plot], scoreFunc func(grid lib.Grid[Plot], plot Plot) int) int {
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
					numFence += scoreFunc(grid, *plot)
					if !plot.fenceUp {
						adjPlot := grid.Get(plot.posX-1, plot.posY).Get()
						if adjPlot.plant == currentPlot.plant && !adjPlot.scored {
							adjPlot.scored = true
							workQueue.Enqueue(adjPlot)
						}
					}
					if !plot.fenceDown {
						adjPlot := grid.Get(plot.posX+1, plot.posY).Get()
						if adjPlot.plant == currentPlot.plant && !adjPlot.scored {
							adjPlot.scored = true
							workQueue.Enqueue(adjPlot)
						}
					}
					if !plot.fenceLeft {
						adjPlot := grid.Get(plot.posX, plot.posY-1).Get()
						if adjPlot.plant == currentPlot.plant && !adjPlot.scored {
							adjPlot.scored = true
							workQueue.Enqueue(adjPlot)
						}
					}
					if !plot.fenceRight {
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

func ScoreFence(grid lib.Grid[Plot], plot Plot) int {
	score := 0
	if plot.fenceUp {
		score++
	}
	if plot.fenceDown {
		score++
	}
	if plot.fenceLeft {
		score++
	}
	if plot.fenceRight {
		score++
	}

	return score
}

func ScoreCorner(grid lib.Grid[Plot], plot Plot) int {
	score := 0
	adjUp := grid.Get(plot.posX-1, plot.posY)
	adjDown := grid.Get(plot.posX+1, plot.posY)
	adjLeft := grid.Get(plot.posX, plot.posY-1)
	adjRight := grid.Get(plot.posX, plot.posY+1)

	// exterior corners
	if plot.fenceUp && plot.fenceLeft {
		// topleft corner
		score++
	}
	if plot.fenceUp && plot.fenceRight {
		// topright corner
		score++
	}
	if plot.fenceDown && plot.fenceLeft {
		// bottomleft corner
		score++
	}
	if plot.fenceDown && plot.fenceRight {
		// bottomright corner
		score++
	}

	// interior corners
	if !plot.fenceUp && !plot.fenceRight && adjUp.Get().fenceRight && adjRight.Get().fenceUp {
		// topright corner
		score++
	}

	if !plot.fenceUp && !plot.fenceLeft && adjUp.Get().fenceLeft && adjLeft.Get().fenceUp {
		// topleft corner
		score++
	}

	if !plot.fenceDown && !plot.fenceRight && adjDown.Get().fenceRight && adjRight.Get().fenceDown {
		// bottomright corner
		score++
	}

	if !plot.fenceDown && !plot.fenceLeft && adjDown.Get().fenceLeft && adjLeft.Get().fenceDown {
		// bottomleft corner
		score++
	}

	return score
}
