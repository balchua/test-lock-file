## How to run

Precompiled binary:

1. Download the [test-lock](test-lock)
2. Copy it to some location, ex. `/data/`
3. Make the file executable, `chmod +x test-lock`

```shell
$ ./test-lock filename 
```

### write to files while it is locked.

Will write some timestamp in the file.

```shell
$ ./test-lock filename true
```