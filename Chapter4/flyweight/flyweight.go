package flyweight

import "time"

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type teamFlyweightFactory struct {
	createdTeams map[string]*Team
}

func (t *teamFlyweightFactory) GetTeam(name string) *Team {
	return nil
}
func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return 0
}
