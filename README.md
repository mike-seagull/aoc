aoc - adventofcode.com cli and data
===
### Get AdventofCode.com data in Go
```golang
import (
    aoc "github.com/mike-seagull/aoc"
)
func main() {
    var session = aoc.Data{token: "MY_SESSION_TOKEN"}
    todaysInput := session.GetTodaysInput() // Returns raw puzzle input
    todaysInputArray := session.GetTodaysInputLines() // Returns todays input as an array of strings
}
```
### AdventofCode.com data to a file
Coming soon...
___
## Quickstart
Get package:
```bash
go get github.com/mike-seagull/aoc
```
Install cli
```bash
go install github.com/mike-seagull/aoc
```
___
### Why Go
Because its fast! Also it can be compiled to a shared library to be used in other languages. Examples of this coming soon...
