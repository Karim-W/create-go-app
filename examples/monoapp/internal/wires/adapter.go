package wires

type AdapterOptions struct{}

type AdapterResults struct{}

func (r *AdapterResults) Close() error {
	// Close all the adapters

	return nil
}

// SetupAdapters initializes the adapters package.
// It is called by the main package.
// Extend this function to add your own adapters.
// Pass the adapters dependencies as parameters.
// Add your adapters to the function return type
func SetupAdapters(
	opts AdapterOptions,
) (res AdapterResults, err error) {
	return
}
