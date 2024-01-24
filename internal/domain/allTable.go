package domain

type AllTables struct {
	Name       string      `json:"name"`
	Data       interface{} `json:"data"`
	IsCollapse bool        `json:"isCollapse"`
}
