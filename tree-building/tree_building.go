package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	nodeMap := make(map[int]*Node)
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	if records[0].ID != 0 {
		return nil, errors.New("no root node")
	}

	for i, record := range records {
		if record.ID != i {
			return nil, errors.New("non-continuous tree")
		}
		if _, ok := nodeMap[record.ID]; ok {
			return nil, fmt.Errorf("duplicate node ID %d", record.ID)
		}
		if record.ID == 0 && record.Parent != 0 {
			return nil, errors.New("root node has parent")
		}
		nodeMap[record.ID] = &Node{ID: record.ID}

		if record.ID != record.Parent {
			if record.Parent > record.ID {
				return nil, errors.New("parent must be smaller than child")
			}
			parentNode := nodeMap[record.Parent]
			parentNode.Children = append(parentNode.Children, nodeMap[record.ID])
		} else {
			if record.ID != 0 {
				return nil, errors.New("direct cycle found")
			}
		}
	}

	return nodeMap[0], nil
}
