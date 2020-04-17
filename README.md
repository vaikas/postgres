# Example SQLBinding usage

This repository contains a simple example on how to use SQL Binding from
[github.com/mattmoor/bindings](https://github.com/mattmoor/bindings).

## Create a secret with your db credentials

You have to modify samples/dbsecret.yaml to point to your database.

Once you have modified that, create the secret that we'll then use
in our binding examples.

```shell
ko apply -f ./config/dbsecret.yaml
```

## Create a binding

Then create bindings that will bind this secret to `Job`s and `Deployment`s
that have label sql-inject="true".

```shell
ko apply -f ./config/dbbinding.yaml
```

## (Optional) Create a sample database

You can then initialize the example database, or modify the examples
to fit the schema to your particular database. Notice you must use
create here instead of apply.

```shell
ko create -f ./config/initdb.yaml
```


## Deploy sample application (read)

Now deploy the application for returning the users in our db.

```shell
ko apply -f ./config/getusers.yaml
```

Then you can curl it for example and you should see:

```shell
vaikas-a01:postgres vaikas$ curl http://ville-test.default.10.185.144.217.xip.io
"Ville" "Aikas" "vaikas@vmware.com"
"Scotty" "Nicholson" "snichols@vmware.com"
"Matt" "Moore" "mattmoor@vmware.com"
```

## Deploy sample application (write)

Now deploy the application.

```shell
ko apply -f ./config/insertuser.yaml
```

Then you can curl it to insert a new user.

```shell
vaikas-a01:postgres vaikas$ curl 'http://ville-test.default.10.185.144.217.xip.io?first=spongebob&last=squarepants&email=sponge99@example.com'
```

And now re-running the query for users will return the new entry.

```shell
vaikas-a01:postgres vaikas$ curl http://ville-test.default.10.185.144.217.xip.io
"Ville" "Aikas" "vaikas@vmware.com"
"Scotty" "Nicholson" "snichols@vmware.com"
"Matt" "Moore" "mattmoor@vmware.com"
"spongebob" "squarepants" "sponge99@example.com"
```

