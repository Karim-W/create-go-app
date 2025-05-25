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

func MustInit(c *cafe.Cafe) *redsync.Redsync {
	config, err := c.GetSubSchema(CONFIG_KEY)
	if err != nil {
		panic(err)
	}

	uri, err := config.GetString("REDIS_URI")
	if err != nil {
		panic(err)
	}

	return Init(uri)
}
