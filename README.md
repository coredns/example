# example

The example middleware prints "example" on every query received. It can be used as documentation for
writing external middleware and to test if external middleware compiles with CoreDNS.

## Syntax

~~~ txt
example
~~~

## Examples

```
example.com {
  file example.com.db {
    upstream 8.8.8.8
  }
  example
}
```

## How to Enable

Follow [these](https://coredns.io/2017/07/25/compile-time-enabling-or-disabling-plugins/) steps,
*example* should be put relatively early in the middleware chain.
