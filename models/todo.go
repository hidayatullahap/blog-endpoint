package models

import (
	"time"
)

type Todo struct {
	ID          int32
	Title       string
	Description string
	Due         time.Time
	Created     time.Time
	Created     int32
	Assigned    []int32
	Status      Status
	Deleted     bool
	List        int32
}

type State int

const (
	Incomplete State = iota
	In_progress
	Completed
)

type Status struct {
	State State
	Prior Status
	Time  time.Time
	User  int32
}

type Comment struct {
	ID         int32
	Parent     int32
	ParentType ObjectType
	Text       string
	Created    time.Time
	Creator    int32
}

type Reaction struct {
	ID         int32
	Parent     int32
	ParentType ObjectType
}

type ObjectType int

const (
	Todo ObjectType = iota
	TodoList
	Status
	Comment
	Reaction
)
