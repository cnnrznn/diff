# diff
Text diff calculator implementing the Myers algorithm

BFS to calculate the fewest number of steps to transform the first string into the other

```go
func Diff(s1, s2 string) Sequence
```

where `Sequence` is

```go
type Sequence struct {
    Steps []Step
}

type Step struct {
    Action Action // add, remove, Equal(noop)
    Seq    string // content
}
```