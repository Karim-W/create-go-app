package rsync

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/karim-w/cafe"
)

const (
	CONFIG_KEY = "RSYNC"
)

// Cafe is the config schema for the rdb package.
// It is used by the config package to initialize the config.
var Cafe = cafe.SubSchema(
	CONFIG_KEY,
	cafe.Schema{
		"REDIS_URI": cafe.String("REDIS_URI").Require(),
	},
)

func Initialize(c *cafe.Cafe) (res *redsync.Redsync, err error) {
	config, err := c.GetSubSchema(CONFIG_KEY)
	if err != nil {
		return
	}

	uri, err := config.GetString("REDIS_URI")
	if err != nil {
		return
	}

	return Init(uri)
}
