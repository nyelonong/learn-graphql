# learn-graphql

Example GraphQL API implemented in Go.

## Run

```bash
go install && learn-graphql
```

## Test

```bash
curl -XPOST http://localhost:31337/graphql -d '{user(id:1){name, shop{name, product{name, price}}}}'
```
