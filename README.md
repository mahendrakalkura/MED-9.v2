How to install?
===============

```
$ psql -c 'CREATE DATABASE "MED-9"' -d postgres
$ mkdir MED-9
$ cd MED-9
$ git clone --recursive git@bitbucket.org:kalkura/MED-9.v2.git .
$ cp settings.toml.sample settings.toml
$ go get
```

How to run?
===========

```
$ cd MED-9
$ go build
$ ./MED-9.v2 --action=bootstrap
    $ ./MED-9.v2 --action=insert
    $ ./MED-9.v2 --action=test_source_1
$ ./MED-9.v2 --action=test_source_2
    $ ./MED-9.v2 --action=queue_1
    $ ./MED-9.v2 --action=queue_2
    $ ./MED-9.v2 --action=workers_1
    $ ./MED-9.v2 --action=workers_2
    $ ./MED-9.v2 --action=report
```
