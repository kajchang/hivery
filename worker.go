package main

import (
	"math"
)

type Worker struct {
	position OrderedPair
	target   OrderedPair
	cooldown int
}

type Task struct {

}

func (worker *Worker) SetCooldown(seconds float64) {
	worker.cooldown = int(math.Round(seconds * 60))
}

func (worker *Worker) SetTarget(target OrderedPair) {
	worker.target = target
}

func (worker Worker) Show() {
	SetCell(worker.position.x, worker.position.y, 'W', Worker_.fg, Worker_.bg)
}

func (worker *Worker) Tick() {
	if worker.cooldown > 0 {
		worker.cooldown--
	}
}

func (worker *Worker) Move(position OrderedPair) {
	SetCell(worker.position.x, worker.position.y, 0, 0, 0)
	worker.position = position
}

func (worker Worker) Pathfinder() OrderedPair {
	best := worker.position
	for x := -1; x <= 1; x++ {
		pos := OrderedPair{worker.position.y, worker.position.x + x}
		cell := GameBuffer.cells[pos.y * GameBuffer.size.y + pos.x]
		if cell.Ch == 0 {
			if math.Sqrt(math.Pow(float64(worker.target.x - pos.x), 2) + math.Pow(float64(worker.target.y - pos.y), 2)) <
				math.Sqrt(math.Pow(float64(worker.target.x - best.x), 2) + math.Pow(float64(worker.target.y - best.y), 2)) {
				best = pos
			}
		}
	}
	for y := -1; y <= 1; y++ {
		pos := OrderedPair{worker.position.y + y, worker.position.x}
		cell := GameBuffer.cells[pos.y * GameBuffer.size.y + pos.x]
		if cell.Ch == 0 {
			if math.Sqrt(math.Pow(float64(worker.target.x - pos.x), 2) + math.Pow(float64(worker.target.y - pos.y), 2)) <
				math.Sqrt(math.Pow(float64(worker.target.x - best.x), 2) + math.Pow(float64(worker.target.y - best.y), 2)) {
				best = pos
			}
		}
	}
	return best
}
