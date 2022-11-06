# go-struct-mapper
This is package for mapper function between two struct

# goal
I need to help for mapping between two struct like that

```
struct Hoge {
  A int `map:C`
  B int `map:D`
}

struct Piyo {
  C int
  D string
}

hogeInstance := Hoge{
  A: 1,
  B: 2,
}
piyoInstance := go-struct-mapper.mappingFrom(hogeInstance)
```
