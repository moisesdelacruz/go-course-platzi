package structs

type Platzi interface {
	Subscribe(name string)
}

func InterfaceTest() {
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
