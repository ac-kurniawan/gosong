# GOSONG
Tag based dependency injection (DI) resolver for Go.

## How to use

### 1. Defining "A" struct

```go
type A struct {
	BStruct *B
}
```
If you see above, struct `A` need struct `B` to be injected.

### 2. Defining "B" struct

```go
type B struct {
	Name string
}
```

### 3. Update "A" struct
To injecting `B` struct to `A` struct we just need to add tag called `import` with value is struct name.

```go
type A struct {
    BStruct *B `import:"B"`
}
```

### 4. Create Apps instance
you can create app instance wherever you want, for example in `main.go`
```go
func main() {
    apps := gosong.Application{
        Name: "SimpleApps",
    }
	apps.AddSingleton("B", &B{Name: "John"})
	apps.AddSingleton("A", &A{})
}
```

## Examples

1. [Simple example](https://github.com/ac-kurniawan/gosong/tree/main/example/simple)
2. [REST API]()