package structs

import "fmt"

type Platzi interface {
	Subscribe(name string)
}

func deferTest() {
	fmt.Println("La Funcion interface ha culminado")
}

func InterfaceTest() {
	defer deferTest()

	platziCourse := PlatziCourse{
		Name:   "Go",
		Slug:   "go",
		Skills: []string{"Backend"},
	}
	platziCareer := new(PlatziCareer)
	platziCareer.Name = "Go Career"
	platziCareer.Slug = "go-career"
	platziCareer.Skills = []string{"Backend"}

	callSubscribe(platziCourse)
	callSubscribe(platziCareer)
}

func callSubscribe(p Platzi) {
	p.Subscribe("Moises")
}
