package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type TeamStats struct {
	matchesPlayed int
	wins          int
	draws         int
	losses        int
	points        int
}

func (ts *TeamStats) AddWin() {
	ts.matchesPlayed++
	ts.wins++
	ts.points += 3
}

func (ts *TeamStats) AddDraw() {
	ts.matchesPlayed++
	ts.draws++
	ts.points++
}

func (ts *TeamStats) AddLoss() {
	ts.matchesPlayed++
	ts.losses++
}

func Tally(reader io.Reader, writer io.Writer) error {
	leagueInfo := map[string]TeamStats{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")

		if len(line) > 0 && line[0] == '#' {
			continue
		}

		if len(parts) != 3 {
			if len(line) > 0 {
				return errors.New("invalid input")
			}
			continue
		}

		team1 := parts[0]
		team2 := parts[1]
		result := parts[2]

		if team1[0] == '#' {
			continue
		}

		if team1 == team2 {
			return errors.New("invalid input")
		}
		stats1 := leagueInfo[team1]
		stats2 := leagueInfo[team2]

		if result == "win" {
			stats1.AddWin()
			stats2.AddLoss()
		} else if result == "draw" {
			stats1.AddDraw()
			stats2.AddDraw()
		} else if result == "loss" {
			stats1.AddLoss()
			stats2.AddWin()
		} else {
			return errors.New("invalid input")
		}

		leagueInfo[team1] = stats1
		leagueInfo[team2] = stats2
	}

	sortedTeams := make([]string, 0, len(leagueInfo))
	for team := range leagueInfo {
		sortedTeams = append(sortedTeams, team)
	}
	sort.Slice(sortedTeams, func(i, j int) bool {
		stats1 := leagueInfo[sortedTeams[i]]
		stats2 := leagueInfo[sortedTeams[j]]
		if stats1.points == stats2.points {
			return sortedTeams[i] < sortedTeams[j]
		}
		return stats1.points > stats2.points
	})

	fmt.Fprintf(writer, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	for _, team := range sortedTeams {
		stats := leagueInfo[team]
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n", team, stats.matchesPlayed, stats.wins, stats.draws, stats.losses, stats.points)
	}

	return nil
}
