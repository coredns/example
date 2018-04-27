# example

## Name

*example* - prints "example" on every query received.

## Description

The example plugin prints "example" on every query received. It serves as documentation for
writing CoreDNS plugins.

## Syntax

~~~ txt
example
~~~

## Metrics

If monitoring is enabled (via the *prometheus* directive) the following metric is exported:

* `coredns_example_request_count_total{server}` - query count to the *example* plugin.

The `server` label indicated which server handled the request, see the *metrics* plugin for details.

## Health

This plugin implements dynamic health checking. It will always return healthy though.

## Examples

In this configuration, we forward all queries to 9.9.9.9 and print "example" whenever we receive
a query.

``` corefile
. {
  forward . 9.9.9.9
  example
}
```

## Also See

See the [manual](https://coredns.io/manual).
