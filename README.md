# example

## Name

*example* - prints "example" on every query handled.

## Description

The example plugin prints "example" on every query that go handled by the server. It serves as
documentation for writing CoreDNS plugins.

## Syntax

~~~ txt
example
~~~

## Metrics

If monitoring is enabled (via the *prometheus* directive) the following metric is exported:

* `coredns_example_request_count_total{server}` - query count to the *example* plugin.

The `server` label indicated which server handled the request, see the *metrics* plugin for details.

## Ready

This plugin reports readiness to the ready plugin. It will be immediately ready.

## Examples

In this configuration, we forward all queries to 9.9.9.9 and print "example" whenever we receive
a query.

~~~ corefile
. {
  forward . 9.9.9.9
  example
}
~~~

## Also See

See the [manual](https://coredns.io/manual).
