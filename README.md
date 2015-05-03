# Watchman

Watchman is a simple HTTP Reverse Proxy with authentication.
HTTP basic authentication is used as authenticaton mechanism.
You should only provide a secure HTTPS connection to the reverse proxy!

Run watchman and set the config path.

```
./watchman -config="sample/watchman.conf"
```

Then go to http://localhost/ and log in as `foo`, password: `bar`.


## Configuration

- `ListenHost`: The host to listen for HTTP requests. Default: empty (All hosts)
- `ListenPort`: The host port to listen on. Default: 80 (HTTP port)
- `DestinationHost`: The host to redirect requests to. Default: 127.0.0.1
- `DestinationPort`: The destination host port. Default: 8080
- `Description`: A short description of the secured area. This is optional. Default: Secured Area
- `PasswdFile`: The path to the htpasswd file.

Check the sample configuration in the sample directory.


## Manage Users

Create an initial htpasswd file and add a user:

```
htpasswd -c /path/to/watchman.passwd foo
```

Add or update an user

```
htpasswd /path/to/watchman.passwd foo
```