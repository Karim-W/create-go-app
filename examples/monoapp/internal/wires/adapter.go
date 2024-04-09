package wires

type AdapterOptions struct{}

type AdapterResults struct{}

// SetupAdapters initializes the adapters package.
// It is called by the main package.
// Extend this function to add your own adapters.
// Pass the adapters dependencies as parameters.
// Add your adapters to the function return type
func SetupAdapters(
	opts AdapterOptions,
) (AdapterResults, error) {
	res := AdapterResults{}

	return res, nil
}
