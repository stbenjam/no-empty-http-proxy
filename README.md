# no-empty-http-proxy

When creating an HTTP Client with a custom transport, it's common to leave Proxy unspecified like this:

```go
	c := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
```

However, it's good practice to set Proxy to `http.ProxyFromEnvironment` like this:

```go
	c := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
			Proxy:           http.ProxyFromEnvironment,
		},
	}
```

Leaving Proxy setting blank will cause problems later when someone inevitably tries to use your software with the
standard proxy environment variables set, and finds out it doesn't work.

This linter forces a definition of the Proxy variable on an `http.Transport.` If you really absolutely do not want a
proxy, explicitly set the field to `nil`.
