# Apollo Server

## Standalone Installation

### Production


Build:

```
$ docker build . -t apollo_server:latest
```

Run:

```
$ docker run -d -p 3000:3000 --name apollo_server apollo_server:latest
```

### Development

Build:

```
$ docker build . -t apollo_server:dev -f Dockerfile.dev
```

Run:

```
$ docker run -d -p 3000:3000 --name apollo_server -v `pwd`/:/usr/src/app apollo_server:dev
```


```
$ docker exec -it apollo_server sh
$ yarn install
$ npm run dev
```

### Confirmation

Query:

```
$ curl -g 'http://localhost:3000/graphql?query={echo(text:"hello"){text}}'
```

## References

- https://github.com/apollographql/apollo-server

