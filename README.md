# Example SQLBinding usage

This repository contains a simple example on how to use SQL Binding from
[github.com/mattmoor/bindings](https://github.com/mattmoor/bindings). Examples
assume you have Knative Serving running.

## Create a secret with your db credentials

You have to modify config/dbsecret.yaml to point to your database.

Once you have modified that, create the secret that we'll then use
in our binding examples.

```shell
kubectl apply -f ./config/dbsecret.yaml
```

## Create a binding

Then create bindings that will bind this secret to `Job`s and Knative Serving
`Services`s that have label sql-inject="true".

```shell
kubectl apply -f ./config/dbbinding.yaml
```

## Create a sample database

You then initialize the example database to fit the schema to your particular
database. Notice you must use create here instead of apply.

```shell
kubectl create -f ./config/prebuilt/initdb.yaml
```

## Deploy sample application (read)

Now deploy the application for returning the users in our db.

```shell
kubectl apply -f ./config/prebuilt/getusers.yaml
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
kubectl apply -f ./config/prebuilt/insertuser.yaml
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

## Modifying samples

The examples above used prebuilt containers to run things. If you want to modify
the examples you can change the code and redeploy your own code by using
[ko](https://github.com/google/ko) or choose whatever way suits you best to
create containers.

## Test
