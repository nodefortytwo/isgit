# isgit

I created this simple package to find out if a directory is part of a git repo, 
the only other way I could find to do this was by calling `git rev-parse --show-top-level`
this is a lot simpler and faster, but doesn't actually check the state of git, instead it
just looks for a `.git` directory present in any parent directory.

## Benchmark
```
BenchmarkGetRootDirWD-12                   30321             38386 ns/op
BenchmarkGetRootDirWithBinary-12             184           6591809 ns/op
```
