# yarn-lock
parse yarn lock file

```bash
go get github.com/hellojukay/yarn-lock
```

### example

```golang
	lock, err := lock.FromFile("./yarn.lock")
	if err != nil {
		panic(err)
	}
	components := lock.Component()
	for _, c := range components {
		fmt.Printf("%20s%30s\n", c.Name, c.Version)
	}
```
