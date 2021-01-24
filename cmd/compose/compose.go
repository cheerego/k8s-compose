package compose

type Service struct {
	Name string
}

type Compose struct {
	Envs      *map[string]string
	Services  *[]Service
	Workspace string
}

func NewC() *Compose {
	c := &Compose{}
	return c
}
