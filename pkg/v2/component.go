package v2

type Component interface {
	Render()
	GetType() string
	Add(component ...Component)
}
