# Adapters

The purpose of the adpaters package is to set up the connections required to the
external data sources.

This can be in the form of:

- Database connection (e.g. MySQL, PostgreSQL, MongoDB, etc.)
- API connection (e.g. REST, GraphQL, etc.)
- Cache connection (e.g. Redis, Memcached, etc.)
- Message broker connection (e.g. RabbitMQ, Kafka, etc.)
- Some SDK (e.g. AWS SDK, Kubernetes SDK, etc.)
- Telemetry connection (e.g. Prometheus, Grafana, etc.)

What is being done with the connection is not the responsibility of the adapters
rather the responsibility of the infra package or any other package that is dealing
with the adapter

> TL;DR: The adapters package is responsible for setting up the connections
The infra package is responsible for the implementation of the connections

## How To Extend the Adapters Package

The adapters package is designed to be extended by adding new adapters to it.

To add a new adapter to the adapters package, you need to follow the following steps:

1. Create a new package in the adapters package with the name of the adapter
2. Set up the `constructos` for the adapter
3. Update the `SetupAdapters` function with the following:
    - Add the new adapter's dependencies to functions' parameters
    - Update the function's body to set up the new adapter
    - Include the new adapter in the return type
