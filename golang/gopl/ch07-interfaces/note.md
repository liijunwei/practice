### ch7.3 interface satisfaction

Since interface satisfaction depends only on the methods of the two types involved, there is no need to declare the relationship between a concrete type and the interfaces it satisfies. That said, it is occasionally useful to document and assert the relationship when it is intended but not otherwise enforced by the program. The declaration below asserts at compile time that a value of type *bytes.Buffer satisfies io.Writer:

```go
// *bytes.Buffer must satisfy io.Writer
var w io.Writer = new(bytes.Buffer)
```

We needn’t allocate a new variable since any value of type `*bytes.Buffer` will do, even nil, which we write as `(*bytes.Buffer)(nil)` using an explicit conversion.

And since we never intend to refer to w, we can replace it with the blank identifier. Together, these changes give us this more frugal variant:

(use this to assert x type implemented y interface in compile time)
```go
// *bytes.Buffer must satisfy io.Writer
var _ io.Writer = (*bytes.Buffer)(nil)
var _ MyInterface = (*MyType)(nil)
```

