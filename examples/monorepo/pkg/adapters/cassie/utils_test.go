package cassie_test

import (
	"testing"

	"{{.moduleName}}/pkg/adapters/cassie"

	"github.com/stretchr/testify/assert"
)

func TestParseCassieUrl(
	t *testing.T,
) {
	sample := "cassandra://localhost:9042/sss?username=sss&password=sss&ssl=true"

	username, password, host, keyspace, port, ssl, err := cassie.ParseUri(sample)

	assert.NoError(t, err)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "sss", username)
	assert.Equal(t, "sss", password)
	assert.Equal(t, "localhost", host)
	assert.Equal(t, 9042, port)
	assert.Equal(t, "sss", keyspace)
	assert.True(t, ssl)
}
