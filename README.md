# iprangetree

Simple IP range tree collection

```go
tree := iprangetree.New(2)

if item, err := iprangetree.ItemByString("86.100.32.0-86.100.32.255"); nil == err {
  tree.Add(item)
} else {
  log.Fatal(err)
}

fmt.Println(tree.Lookup(net.ParseIP("86.100.32.10")))
```
