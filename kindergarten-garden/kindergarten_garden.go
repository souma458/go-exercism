package kindergarten

import (
	"slices"
	"sort"
	"strings"
)

type Garden struct {
	plants   [][]string
	children []string
}

type InvalidDiagramError struct{}

func (m *InvalidDiagramError) Error() string {
	return "invalid diagram"
}

var plantEncodingMap = map[rune]string{'V': "violets", 'C': "clover", 'G': "grass", 'R': "radishes"}

func NewGarden(diagram string, children []string) (*Garden, error) {
	plants := make([][]string, 2)
	rows := strings.Split(diagram, "\n")

	if len(rows) < 3 {
		return nil, &InvalidDiagramError{}
	}

	if len(rows[1]) != len(rows[2]) || len(rows[1])%2 != 0 {
		return nil, &InvalidDiagramError{}
	}

	for i, v := range children {
		if slices.Contains(children[i+1:], v) {
			return nil, &InvalidDiagramError{}
		}
	}

	for _, v := range rows[1] {
		cup, ok := plantEncodingMap[v]
		if !ok {
			return nil, &InvalidDiagramError{}
		}
		plants[0] = append(plants[0], cup)
	}

	for _, v := range rows[2] {
		cup, ok := plantEncodingMap[v]
		if !ok {
			return nil, &InvalidDiagramError{}
		}
		plants[1] = append(plants[1], cup)
	}

	return &Garden{plants, children}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if !slices.Contains(g.children, child) {
		return nil, false
	}

	childrenSorted := make([]string, len(g.children))
	copy(childrenSorted, g.children)
	sort.Strings(childrenSorted)
	index := slices.Index(childrenSorted, child)

	return []string{g.plants[0][index*2], g.plants[0][index*2+1], g.plants[1][index*2], g.plants[1][index*2+1]}, true
}
