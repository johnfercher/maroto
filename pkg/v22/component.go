package v22

type Component interface {
	Render()
	GetType() string
	Add(component ...Component)
}
