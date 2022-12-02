package day02

import (
	"errors"
	"fmt"
	"strings"
)

type Game struct {
	Rounds    []string
	Responses []Response
	Requests  []Request
}

type Response struct {
	Name  string
	Score int
	Win   string
	Draw  string
	Lose  string
}

type Request struct {
	Name    string
	Outcome string
}

func (g *Game) Init() {
	g.Responses = []Response{
		{Name: "X", Score: 1, Win: "C", Draw: "A", Lose: "B"},
		{Name: "Y", Score: 2, Win: "A", Draw: "B", Lose: "C"},
		{Name: "Z", Score: 3, Win: "B", Draw: "C", Lose: "A"},
	}
	g.Requests = []Request{
		{Name: "X", Outcome: "Lose"},
		{Name: "Y", Outcome: "Draw"},
		{Name: "Z", Outcome: "Win"},
	}
}

func (g Game) ScoreGame() int {
	totalscore := 0
	for _, value := range g.Rounds {
		plays := strings.Fields(value)
		totalscore += g.ScoreRound(plays[0], plays[1])
	}
	return totalscore
}

func (g Game) ScoreRound(opponent string, player string) int {
	resp, err := g.GetResponse(player)
	if err != nil {
	}
	score := 0
	score += resp.Score

	if opponent == resp.Win {
		score += 6
	} else if opponent == resp.Draw {
		score += 3
	}
	return score
}

func (g Game) Process(values []string) []string {
	processed := []string{}
	for _, v := range values {
		plays := strings.Fields(v)
		processed = append(processed, fmt.Sprintf("%s %s", plays[0], g.Decide(plays[0], plays[1])))
	}
	return processed
}

func (g Game) Decide(opponent string, player string) string {
	request := []Request{
		{Name: "X", Outcome: "Lose"},
		{Name: "Y", Outcome: "Draw"},
		{Name: "Z", Outcome: "Win"},
	}

	req, err := g.GetRequest(request, player)
	if err != nil {
	}

	for _, v := range g.Responses {
		if req.Outcome == "Lose" {
			if v.Lose == opponent {
				return v.Name
			}
		} else if req.Outcome == "Draw" {
			if v.Draw == opponent {
				return v.Name
			}
		} else if req.Outcome == "Win" {
			if v.Win == opponent {
				return v.Name
			}
		}
	}
	return ""
}

func (g Game) GetResponse(key string) (Response, error) {
	for _, v := range g.Responses {
		if v.Name == key {
			return v, nil
		}
	}
	return Response{}, errors.New("No response found")
}

func (g Game) GetRequest(resp []Request, key string) (Request, error) {
	for _, v := range resp {
		if v.Name == key {
			return v, nil
		}
	}
	return Request{}, errors.New("No request found")
}

func Day2_1(values []string) int {
	game := Game{}
	game.Init()
	game.Rounds = values
	return game.ScoreGame()
}

func Day2_2(values []string) int {
	game := Game{}
	game.Init()
	game.Rounds = game.Process(values)
	return game.ScoreGame()
}
