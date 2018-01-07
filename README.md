# GraphQL Micro POC

The purpose of this project is to evaluate GraphQL-based services including API gateway with GraphQL remote schema.

There exist the following services each of which just implement an echo API:

- apollo-server: A Node.js service using [apollo-server-express](https://github.com/apollographql/apollo-server)
- graphq-go: A Golang service using [graphql-go/graphql](https://github.com/graphql-go/graphql)
- graphq-go-alt: A Golang service using [neelance/graphql-go](https://github.com/neelance/graphql-go)

An API gateway is also introduced for POC of [GraphQL API Gateway](https://blog.graph.cool/graphql-api-gateway-graphql-native-1e46e4f179f7).

- graphql-gateway: A Node.js service using GraphQL remote schema to proxy requests to an above echo API service (Currently I couldn't find any Golang implementation for GraphQL remote schema).

```
+-------------+   +----------+
|             |   |          |
+ API Gateway +---+ Echo API |
|  (GraphQL)  |   | (GraphQL)|
+-------------+   +----------+
```

## Installation

- Just install [Docker](https://www.docker.com/) somehow

- Run one of the echo services (apollo-server/graphq-go/graphq-go-alt):

```
$ cd apollo-server
$ docker build . -t apollo_server:latest
$ docker run -d -p 3000:3000 --name apollo_server apollo_server:latest
```

- Confirm the API:

```
$ curl -g 'http://localhost:3000/graphql?query={echo(text:"hello"){text}}'
```

- Run the API gateway with informing the echo service IP:

```
$ cd graphql-gateway
$ docker build . -t graphql_gateway:latest
$ SERVICE1_IP=`docker inspect --format '{{ .NetworkSettings.IPAddress }}' apollo_server`
$ docker run -d -p 80:3000 \
--name graphql_gateway \
-e SERVICE1="http://$SERVICE1_IP:3000" \
graphql_gateway:latest
```
- Confirm the API Gateway:

```
$ curl -g 'http://localhost:80/graphql?query={echo(text:"hello"){text}}'
```

More details are self-explained in each target directory.

## Benchmark

The benchmark results are summarized in `bench` directory.
Here we just shows mean Latencies (on 2 GHz Intel Core i5 on MacBook Pro).
Note that `baseline` shows a result for responses without GraphQL operations.

```
apollo-server                      2.949122ms (baseline: 2.342966ms)
graphql-go                         2.478427ms (baseline: 2.215219ms)
graphql-go-alt                     2.126513ms (baseline: 1.890965ms)
graphql-gateway -> apollo-server   6.241109ms (baseline: 4.466211ms)
graphql-gateway -> graphql-go      5.636672ms (baseline: 4.233277ms)
graphql-gateway -> graphql-go-alt  4.718619ms (baseline: 4.140846ms)
```


