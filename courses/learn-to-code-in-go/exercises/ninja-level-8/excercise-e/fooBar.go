package main

type user struct {
	FirstName string
	LastName  string
	Age       int
	Sayings   []string
}

type ByAge []user

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByLastName []user

func (a ByLastName) Len() int {
	return len(a)
}
func (a ByLastName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByLastName) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}
