# GraphQL Go

## Standalone Installation

### Production

Build:

```
$ docker build . -t graphql_go_alt:latest
```

Run:

```
$ docker run -d -p 8080:8080 --name graphql_go_alt graphql_go_alt:latest
```

### Development

Build:

```
$ docker build . -t graphql_go_alt:dev -f Dockerfile.dev
```

Run:

```
$ docker run -d -p 8080:8080 --name graphql_go_alt -v `pwd`/:/go/src/app graphql_go_alt:dev
```

```
$ docker exec -it graphql_go_alt bash
$ make
$ ./app
```

### Confirmation

Query:

```
$ curl -X POST -d '
{
 "query": "query {echo(text:\"hello\"){text}}"
}
' http://localhost:8080/graphql
```

## References

- https://github.com/neelance/graphql-go
- https://github.com/klud1/graphql-docker-api

