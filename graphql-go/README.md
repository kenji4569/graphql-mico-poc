# GraphQL Go

## Standalone Installation

### Production

Build:

```
$ docker build . -t graphql_go:latest
```

Run:

```
$ docker run -d -p 8080:8080 --name graphql_go graphql_go:latest
```

### Development

Build:

```
$ docker build . -t graphql_go:dev -f Dockerfile.dev
```

Run:

```
$ docker run -d -p 8080:8080 --name graphql_go -v `pwd`/:/go/src/app graphql_go:dev
```

```
$ docker exec -it graphql_go bash
$ make
$ ./app
```

### Confirmation

Query:

```
$ curl -g 'http://localhost:8080/graphql??query={echo(text:"hello"){text}}'
```

## References

- https://github.com/graphql-go/graphql

