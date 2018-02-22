# example

## Name

*example* - prints "example" on every query received.

## Description

The example plugin prints "example" on every query received. It servesused as documentation for
writing CoreDNS plugins.

## Syntax

~~~ txt
example
~~~

## Examples

In this configuration, we forward all queries to 9.9.9.9 and print "example" whenever we recieve
a query.

``` corefile
. {
  forward . 9.9.9.9
  example
}
```
