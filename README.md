`readval` is a command line tool which is designed for one specific purpose--
to prompt for a value to be included in another command line. `readval` prints
the prompt on the TTY, reads the input also from the TTY, and echoes the entered
input to stdout.

Examples
-------

``` sh
$ readval prompt </dev/null | od -c
prompt: foo
0000000   f   o   o  \n
0000004
```

Or a more "real" example:

``` sh
$ echo "foo=bar" \
  | curl -s -XPOST -d @- "https://httpbin.org/post?user=$( readval Enter username )&pass=$( readval -s Enter password )" \
  | jq .
Enter username: testuser
Enter password:
{
  "args": {
    "pass": "testpass",
    "user": "testuser"
  },
  "data": "",
  "files": {},
  "form": {
    "foo": "bar"
  },
...
```

Installation
------------

``` sh
go install github.com/venkytv/readval@latest
```
