package main 

import "fmt"

var currentId int

var codes Codes

//Seed data
func init() {
	RepoCreateCode(Code{
			Name: 	"Male",
			Field:	"con_gender",
			Module:	"CON",
			Active: true,
			})
	RepoCreateCode(Code{
			Name:	"Female",
			Field:	"con_gender",
			Module:	"CON",
			Active:	true,
		})
	RepoCreateCode(Code{
			Name:	"Nurse",
			Field:	"con_jobrole",
			Module:	"CON",
			Active: true,
		})
}

func RepoFindCode(id int) Code {
	for _, c := range codes {
		if c.Id == id {
			return c
		}
	}
	//return empty code if not found
	return Code{}
}

func RepoFindCodeByModule(module string) Codes {
	moduleCodes := Codes{}
	for _, c := range codes {
		if c.Module == module {
			moduleCodes = append(moduleCodes, c)
		}
	}
	//return empty code if not found
	return moduleCodes
}

func RepoFindCodeByField(field string) Codes {
	fieldCodes := Codes{}
	for _, c := range codes {
		if c.Field == field {
			fieldCodes = append(fieldCodes, c)
		}
	}
	//return empty code if not found
	return fieldCodes
}

func RepoCreateCode(c Code) Code {
	currentId += 1
	c.Id = currentId
	codes = append(codes, c)
	return c
}

func RepoDestroyCode(id int) error {
	for i, c := range codes {
		if c.Id == id {
			codes = append(codes[:i], codes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Cound not find Code with id of %d to delete", id)
}