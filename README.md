Go PrettyPrinter
================

A simple pretty-printer for Go

```go
import(
    "mrl/pretty"
)

pretty.Pretty([]int{3, 4, 5}, "    ")
```

will produce
```
[3]int[
    3
    4
    5
]
```
