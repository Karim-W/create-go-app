package apps

type Application interface {
	Setup() error
	Close()
}
