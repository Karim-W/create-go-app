package cassie_test

import (
	"ams/pkg/adapters/cassie"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCassieUrl(
	t *testing.T,
) {
	sample := "cassandra://localhost:9042/ams?username=ams&password=ams&ssl=true"

	username, password, host, keyspace, port, ssl, err := cassie.ParseUri(sample)

	assert.NoError(t, err)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "ams", username)
	assert.Equal(t, "ams", password)
	assert.Equal(t, "localhost", host)
	assert.Equal(t, 9042, port)
	assert.Equal(t, "ams", keyspace)
	assert.True(t, ssl)
}
