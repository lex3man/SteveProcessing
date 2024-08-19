package models

type Status struct {
	ID    string `json:"id"`
	Time  string `json:"time"`
	Core  string `json:"core"`
	Bot   string `json:"bot"`
	Admin string `json:"admin"`
}
