package main 

import ()

type jsonErr struct {
	Code	int 	`json: "code"`
	Text	string	`json: "text"`
}