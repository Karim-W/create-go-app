package cassie

import (
	"fmt"
	"strconv"
	"strings"
)

/*
ParseUri parses a Cassandra URI and returns the username, password, host, port, and keyspace.

The URI should be in the format:
cassandra://host:port/keyspace?username=username&password=password&ssl=true
*/
func ParseUri(
	uri string,
) (username, password, host, keyspace string, port int, ssl bool, err error) {
	if !strings.HasPrefix(uri, "cassandra://") {
		err = fmt.Errorf("invalid URI: %s", uri)
		return
	}

	uri = strings.TrimPrefix(uri, "cassandra://")

	parts := strings.Split(uri, "/")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid URI %s while looking up keyspace", uri)
		return
	}

	upstream := parts[0]
	keyspace = parts[1]

	hostParts := strings.Split(upstream, ":")
	if len(hostParts) != 2 {
		err = fmt.Errorf("invalid URI: %s while looking up host and port", uri)
		return
	}

	host = hostParts[0]
	port, err = strconv.Atoi(hostParts[1])
	if err != nil {
		err = fmt.Errorf("invalid URI: %s while looking up port", uri)
		return
	}

	keyspaceParts := strings.Split(keyspace, "?")
	if len(keyspaceParts) != 2 {
		err = fmt.Errorf("invalid URI: %s while looking up credentials", uri)
		return
	}

	keyspace = keyspaceParts[0]
	creds := keyspaceParts[1]

	parts = strings.Split(creds, "&")

	for _, part := range parts {
		kv := strings.Split(part, "=")
		partsCount := len(parts)
		if partsCount < 2 {
			return
		}

		val := kv[1]
		if len(kv) != 2 {
			// join the rest of the parts
			val = strings.Join(kv[1:], "=")
		}

		switch kv[0] {
		case "username":
			username = val
		case "password":
			password = val
		case "ssl":
			ssl = val == "true"
		}

	}

	return
}
