# GraphQL Gateway

## Standalone Installation

### Production


Build:

```
$ docker build . -t graphql_gateway:latest
```

Run:

```
$ SERVICE1_IP=`docker inspect --format '{{ .NetworkSettings.IPAddress }}' apollo_server`
$ docker run -d -p 80:3000 \
--name graphql_gateway \
-e SERVICE1="http://$SERVICE1_IP:3000" \
graphql_gateway:latest
```

### Development

Build:

```
$ docker build . -t graphql_gateway:dev -f Dockerfile.dev
```

Run:

```
$ SERVICE1_IP=`docker inspect --format '{{ .NetworkSettings.IPAddress }}' apollo_server`
$ docker run -d -p 80:3000 \
--name graphql_gateway \
-v `pwd`/:/usr/src/app \
-e SERVICE1='http://$SERVICE1_IP:3000' \
graphql_gateway:dev
```


```
$ docker exec -it graphql_gateway sh
$ yarn install
$ npm run dev
```

### Confirmation

Query items in GraphQL:

```
$ curl -g 'http://localhost:80/graphql?query={echo(text:"hello"){text}}'
```

## References

- https://blog.graph.cool/graphql-api-gateway-graphql-native-1e46e4f179f7
- https://blog.graph.cool/how-do-graphql-remote-schemas-work-7118237c89d7
- https://www.apollographql.com/docs/graphql-tools/remote-schemas.html
- https://github.com/apollographql/apollo-server

