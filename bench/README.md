# Benchmark

## Usage

Install vegeta:

```
# For mac
$ brew install vegeta

# For others
$ go get github.com/tsenart/vegeta
$ go install github.com/tsenart/vegeta
```

Attack:

```
vegeta attack -rate=100 -duration=10s -targets=graphql-go.targets.txt > results.txt
```

Report:

```
vegeta report -inputs=results.txt -reporter=text
```

## Results

- Executed under 2 GHz Intel Core i5 on MacBook Pro with VirtualBox
- Baseline shows a result for responses without GraphQL operations

### Echo query in GraphQL

apollo-server:

```
Latencies     [mean, 50, 95, 99, max]  2.949122ms, 2.947025ms, 3.748597ms, 6.150521ms, 14.021012ms

(Baseline)
Latencies     [mean, 50, 95, 99, max]  2.342966ms, 2.433707ms, 2.737196ms, 3.000383ms, 8.8029ms
```

graphql-go:

```
Latencies     [mean, 50, 95, 99, max]  2.478427ms, 2.569106ms, 2.891536ms, 3.687911ms, 13.407998ms

(Baseline)
Latencies     [mean, 50, 95, 99, max]  2.215219ms, 2.387624ms, 2.614172ms, 2.733234ms, 3.26209ms
```

graphql-go-alt:

```
Latencies     [mean, 50, 95, 99, max]  2.126513ms, 2.285669ms, 2.482378ms, 2.762952ms, 8.116577ms

(Baseline)
Latencies     [mean, 50, 95, 99, max]  1.890965ms, 2.036998ms, 2.28754ms, 2.476887ms, 2.599275ms
```

### Echo query in GraphQL with GraphQL Gateway

graphql-gateway -> apollo-server:

```
Latencies     [mean, 50, 95, 99, max]  6.241109ms, 5.811276ms, 9.166503ms, 19.317874ms, 34.239311ms

(Baseline)
Latencies     [mean, 50, 95, 99, max]  4.466211ms, 4.40836ms, 5.715209ms, 8.338639ms, 18.38526ms
```

graphql-gateway -> graphql-go:

```
Latencies     [mean, 50, 95, 99, max]  5.636672ms, 5.447959ms, 7.525099ms, 13.788379ms, 27.029716ms

(Baseline)
Latencies     [mean, 50, 95, 99, max]  4.233277ms, 4.19605ms, 5.225457ms, 7.995379ms, 17.0992ms
```

graphql-gateway -> graphql-go-alt:

```
Latencies     [mean, 50, 95, 99, max]  4.718619ms, 4.426526ms, 6.460594ms, 12.726997ms, 22.235654ms

(Baseline)
Latencies     [mean, 50, 95, 99, max]  4.140846ms, 4.069605ms, 5.426539ms, 7.825691ms, 15.481355ms
```