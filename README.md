# Greenweb

A very simple Go binary webserver, packaged into an rpm.

To generate an RPM:

```bash
build_rpms.sh <major version number> <minor version number>
```

RPMs will be dumped into the `rpms` directory.

An example to create a load of packages at different versions:

```bash
for maj in `seq 1 3`; do
  t=$(($RANDOM %20 +1))
  for min in `seq 0 $t`; do
    echo Building Version $maj-$min
    ./build_rpms.sh $maj $min
  done
done
```
