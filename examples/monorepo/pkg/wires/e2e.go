package wires

type (
	InfraOption int64
)

const (
	IDIER InfraOption = iota
	PSQL
	REDIS
)

type InfraTuple struct {
	Name   string
	Option InfraOption
}

type AdapterFlags struct {
	Idier bool
}

type InfraOptions struct {
	PSQL           bool
	Redis          bool
	Rsync          bool
	Extra          []InfraTuple
	MigrationsFlag bool
}

type Config struct {
	ServiceName string
	Adapters    AdapterFlags
	Infra       InfraOptions
}
