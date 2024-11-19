package entity

type Task struct {
	ID         int     `json:"id,omitempty"`
	Title      string  `json:"title,omitempty"`
	Content    string  `json:"content,omitempty"`
	AuthorID   int     `json:"authorID,omitempty"`
	AssignedID int     `json:"assignedID,omitempty"`
	Opened     int64   `json:"opened,omitempty"`
	Closed     int64   `json:"closed,omitempty"`
	Labels     []Label `json:"label,omitempty"`
}

type Label struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
