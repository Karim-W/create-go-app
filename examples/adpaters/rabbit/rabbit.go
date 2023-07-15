// Description: rabbitmq connection ,anager
// rabbit for rabbitmq
package rabbit

import (
	"github.com/BetaLixT/usago"
	"go.uber.org/zap" // for logging (zap is a dependency of usago) not much i can do about it
)

// CreateChannelManager creates a new ChannelManager
// with the given URI.
// panics if the URI is invalid or if the logger fails
func CreateChannelManager(
	uri string,
) *usago.ChannelManager {
	ch := usago.NewChannelManager(
		uri,
		zap.Must(zap.NewProduction()),
	)
	return ch
}
