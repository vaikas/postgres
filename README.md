# Example SQLBinding usage

This repository contains a simple example on how to use SQL Binding from
[github.com/mattmoor/bindings](https://github.com/mattmoor/bindings).

It assumes there's a database with pretty standard golang example, you
obvs might want to modify the example that suits your needs.

You have to modify samples/example.yaml to point to your database.

To run:

```shell
ko apply -f ./config/example.yaml
```


