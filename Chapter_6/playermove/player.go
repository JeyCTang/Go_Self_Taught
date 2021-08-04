package main

type Player struct {
	currPos   Vec2    // current position
	targetPos Vec2    // target position
	speed     float32 // moving speed
}

// MoveTo set up the target position the player wants to move ot
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// Pos Get the current position
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// IsArrived estimate if the player arrived the target position
func (p *Player) IsArrived() bool {
	// calculate the distance between currentPos and targetPos is less than the moving steps to do the estimation
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

// Update update the position of player
func (p *Player) Update() {
	if !p.IsArrived() {
		// calculate the direction of current position
		dir := p.targetPos.Sub(p.currPos).Normalize()
		// add the speed vector to generate the new position
		newPos := p.currPos.Add(dir.Scale(p.speed))
		// after the movement, update the current position
		p.currPos = newPos
	}
}

// NewPlayer Generate a new player
func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}
