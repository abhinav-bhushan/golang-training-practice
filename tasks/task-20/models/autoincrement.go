package models

import "strconv"

type AutoIncrement struct {
	id int
}

func (a *AutoIncrement)Increment()string{
	a.id+=1
	return strconv.Itoa(a.id-1)
}

