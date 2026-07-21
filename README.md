# toolbox-fiber

Easy tools for [GoFiber v3](https://github.com/gofiber/fiber) — a small collection of helper utilities for common web tasks.

## Installation

```bash
go get go.gh.ink/toolbox/fiber/v3
```

Requires Go 1.26+ and Fiber v3.

## Packages

### `auth`

Parse the `Authorization` header and detect its scheme.

```go
import "go.gh.ink/toolbox/fiber/v3/auth"

func handler(c fiber.Ctx) error {
    token, authType := auth.GetAuthorization(c)

    switch authType {
    case auth.AuthorizationTypeBearer:
        // token is the value after "Bearer "
    case auth.AuthorizationTypeBasic:
        // token is the value after "Basic "
    case auth.AuthorizationTypeNone:
        // no Authorization header present
    case auth.AuthorizationTypeOther:
        // unrecognized scheme; token is the raw header value
    }
    return nil
}
```

`GetAuthorization` returns the credential with its scheme prefix stripped, along with one of the `AuthorizationType` constants: `Bearer`, `Basic`, `None`, or `Other`.

### `ip`

Get the real client IP address, preferring proxy-forwarded addresses when available.

```go
import "go.gh.ink/toolbox/fiber/v3/ip"

func handler(c fiber.Ctx) error {
    clientIP := ip.GetIP(c)
    return c.SendString(clientIP)
}
```

`GetIP` returns the first entry from `X-Forwarded-For` (via `c.IPs()`) if present, otherwise falls back to `c.IP()`.

> **Note:** `X-Forwarded-For` can be spoofed by clients. Only trust it when your app runs behind a trusted proxy — see Fiber's [`TrustProxy`](https://docs.gofiber.io/api/fiber#config) configuration.

## License

See [LICENSE](LICENSE).
