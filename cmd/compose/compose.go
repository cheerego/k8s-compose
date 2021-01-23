package compose

type Service struct {
	Name string
}

type Compose struct {
	Envs     *map[string]string
	Services *[]Service
}

func NewC() *Compose {
	return &Compose{}
}
