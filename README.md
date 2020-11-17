# shukujitsu

shukujitsu datermin japanese holiday.
Holidays are collected from https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv

When csv is updated, you can updated as follow command.

```
$ go run ./internal/gen/gen.go
```

## HOW TO USE

```go
if shukujitsu.IsShukujitsu(time.Now()) {
    fmt.Println("shukujitsu!")
}
```
