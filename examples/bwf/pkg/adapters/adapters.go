package adapters

type Options struct {
	ServiceName string
}

type Results struct{}

// SetupAdapters initializes the adapters package.
// It is called by the main package.
// Extend this function to add your own adapters.
// Pass the adapters dependencies as parameters.
// Add your adapters to the function return type
func SetupAdapters(
	opts *Options,
) (*Results, error) {
	res := &Results{}
	return res, nil
}
