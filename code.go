package main 

import ()

type Code struct {
	Id		int			`json:"id"`
	Name	string		`json:"name"`
	Field	string		`json:"field"`
	Module	string		`json:"module"`
	Active	bool		`json:"active"`
}

type Codes []Code