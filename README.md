# Go expr lang playground

```go
	env := map[string]interface{}{
		"foo":      1,
		"double":   func(i int) int { return i * 2 },
		"users":    getUsers(),
		"getUsers": getUsers,
		"getName":  getUserName,
	}

	out, _ := expr.Eval("double(foo)", env)
	fmt.Println(out)
	out, _ = expr.Eval("users[1]", env)
	fmt.Println(out)
	out, _ = expr.Eval("users[0].Name", env)
	fmt.Println(out)
	out, _ = expr.Eval("getUsers()[0].Email", env)
	fmt.Println(out)
	out, _ = expr.Eval("getName(users[2])", env)
	fmt.Println(out)
}
```
```bash
$ go run main.go

2
{shellycjpk Aryana Renn crosby_swiftb@casting.gd Controversy educational workshop website waters habitat decline, nightmare argue departure eye shapes humidity chelsea, settled dig. }
Tirzah Menges
elicia_lewteree1@alternative.fbu
Kyleigh Manes
```

see [main.go](./main.go) for how it works