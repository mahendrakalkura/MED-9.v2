How to install?
===============

```
$ psql -c 'CREATE DATABASE "MED-9"' -d postgres
$ mkdir MED-9
$ cd MED-9
$ git clone --recursive git@bitbucket.org:kalkura/MED-9.v2.git .
$ cp settings.toml.sample settings.toml
$ go get
$ iconv --from-code iso-8859-1 --to-code utf-8 --output 1.csv ADR461.CSV
```

How to run?
===========

```
$ cd MED-9
$ go build
$ ./MED-9.v2 --action=bootstrap
$ ./MED-9.v2 --action=insert
$ ./MED-9.v2 --action=source_1_one
$ ./MED-9.v2 --action=source_2_one
$ ./MED-9.v2 --action=source_1_all
$ ./MED-9.v2 --action=source_2_all
$ ./MED-9.v2 --action=progress
$ ./MED-9.v2 --action=report
```
