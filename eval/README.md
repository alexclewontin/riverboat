## Benchmarks

Riverboat is faster than  
- [chehsunliu/poker](github.com/chehsunliu/poker)
    - by 1.7x for 5-card evaluation
    - by 1.2x for 6-card evaluation*
    - by 1.4x for 7-card evaluation*
- [notnil/joker](github.com/notnil/joker/hand)
    - by 325x for 5-card evaluation
    - by 150x for 6-card evaluation*
    - by 74x for 7-card evaluation*

<sup><sub>*Riverboat and [notnil/joker's](github.com/notnil/joker/hand) 6- and 7-card hand evaluation are not directly comparable to [chehsunliu/poker's](github.com/chehsunliu/poker), as Riverboat and Joker return both the best 5 cards, as well as their absolute ranking, whereas Poker only provides the absolute ranking.

All benchmarks were measured over the same set of cards, and the timing did not include the conversion of cards from human-readable string form to native representation.

This benchmark was run on a 2018 Macbook Pro with an Intel i9-8950HK 2.9GHz processor. The precise benchmarks run can be found in [benchmarks_test.go](./benchmarks_test.go)
```shell
$ go test -bench=. -benchmem -benchtime 5s
goos: darwin
goarch: amd64
pkg: github.com/alexclewontin/riverboat/eval
BenchmarkFiveJoker-12             190854             30169 ns/op           14430 B/op        657 allocs/op
BenchmarkFivePoker-12           36639379               156 ns/op               0 B/op          0 allocs/op
BenchmarkFiveRiverboat-12       59392126                92.6 ns/op             0 B/op          0 allocs/op
BenchmarkSixJoker-12               38364            158847 ns/op           67960 B/op       2923 allocs/op
BenchmarkSixPoker-12             4252222              1298 ns/op             288 B/op          9 allocs/op
BenchmarkSixRiverboat-12         5695171              1055 ns/op             288 B/op          9 allocs/op
BenchmarkSevenJoker-12              9788            564280 ns/op          265363 B/op      10231 allocs/op
BenchmarkSevenPoker-12            555128             10397 ns/op            2304 B/op         72 allocs/op
BenchmarkSevenRiverboat-12        746978              7663 ns/op            2016 B/op         63 allocs/op
PASS
ok      github.com/alexclewontin/riverboat/eval 56.545s
                                                              
```