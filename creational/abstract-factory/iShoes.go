package main

type IShoes interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoes struct {
	logo string
	size int
}

func (s *Shoes) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoes) getLogo() string {
	return s.logo
}

func (s *Shoes) setSize(size int) {
	s.size = size
}

func (s *Shoes) getSize() int {
	return s.size
}
