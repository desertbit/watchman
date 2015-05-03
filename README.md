# Watchman

Watchman is a simple HTTP Reverse Proxy with authentication.
HTTP basic authentication is used as authenticaton mechanism.
You should only provide a secure HTTPS connection to the reverse proxy!

Run watchman and set the configuration path.

```
./watchman -config="sample/watchman.conf"
```

Then go to http://localhost/ and log in as `foo`, password: `bar`.

It is also possible to pass the configuration path with an environment variable. Check the section below.


## Configuration

- `ListenHost`: The host to listen for HTTP requests. Default: empty (All hosts)
- `ListenPort`: The host port to listen on. Default: 80 (HTTP port)
- `DestinationHost`: The host to redirect requests to. Default: 127.0.0.1
- `DestinationPort`: The destination host port. Default: 8080
- `Description`: A short description of the secured area. This is optional. Default: Secured Area
- `PasswdFile`: The path to the htpasswd file.

Check the sample configuration in the sample directory.


## Environment variables

- `WATCHMAN_CONFIG` Sets the path to the watchman configuration.
- `WATCHMAN_DIR` Sets the lookup directory path. (Config and passwd files)


## Manage Users

Create an initial htpasswd file and add a user:

```
htpasswd -c /path/to/watchman.passwd foo
```

Add or update an user:

```
htpasswd /path/to/watchman.passwd foo
```


## Docker

Pull the docker image from desertbit/watchman

```
docker pull desertbit/watchman
```

There is one volume `/config`. Place the watchman configuration and htpasswd file to that location.
The run the image.

```
docker run -p 80:80 -v /path/to/configdir:/config desertbit/watchman
```