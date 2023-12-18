# Doorbell

This is the web server and web page that handles the PIN input. It handles rudimentary logging for audit reasons. It logs successful PIN inputs (without displaying the PIN), and then logs unsuccessful attempts.

After X amount of failures, the system stops processing inputs for some amount of time.

This website is meant to be viewed in a landscape size of 360x240.

## Passwd file

There needs to be a mapping of a pin to a known user. This is a file located in the home directory of the user running doorbell. The file is `/etc/doorbell.passwd`. It has a format of:

```
<numerical pin>:<authorized username>
```

An example:

```
666:satan
777:oh lawdy
```

If you have a local file in `src/doorbell/passwd`, when you run `make` this file will be copied to the appropriate location. Otherwise `make` will fail.
