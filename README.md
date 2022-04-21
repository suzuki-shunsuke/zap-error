# zap-error

Embed [zap](https://github.com/uber-go/zap).Field into error.

This is a Go's small library to embed [zap](https://github.com/uber-go/zap)'s Fields into error.

This inspired to [suzuki-shunsuke/logrus-error](https://github.com/suzuki-shunsuke/logrus-error).

## Motivation

With [fmt.Errorf](https://pkg.go.dev/fmt#Errorf), you can add additional context to error.

e.g.

```go
fmt.Errorf("get a user: %w", err)
```

[uber-go/zap](https://github.com/uber-go/zap) is one of most popular structured logging library.

e.g.

```go
logger.Info("get a user",
  zap.String("username", username),
  zap.Error(err),
)
```

`fmt.Errorf` is very useful, but you can add only a string to error as context. You can't add structured data to error.
If you use zap, you may want to add structured data to error.

`zap-error` is a small library to add structured data to error and get structured data from error for logging.

Mainly zap-error provides only two simple API.

* [WithFields](https://pkg.go.dev/github.com/suzuki-shunsuke/zap-error/logerr#WithFields): Add structured data to error
* [ExpandError](https://pkg.go.dev/github.com/suzuki-shunsuke/zap-error/logerr#ExpandError): Get structured data from error and return []zap.Field

AS IS (without zap-error)

```go
return fmt.Errorf("get a user (username: %s): %w", username, err)
```

```go
zap.Error("add a member to a group", zap.Error(err))
```

TO BE (with zap-error)

```go
return logerr.WithFields(fmt.Errorf("get a user: %w", err), zap.String("username", username))
```

```go
logger.Error("add a member to a group", logerr.ExpandError(err)...)
```

Using zap-error, you can add structured data to error as context. You don't have to construct a string with [fmt's format](https://pkg.go.dev/fmt#hdr-Printing).

## Document

Please see https://pkg.go.dev/github.com/suzuki-shunsuke/zap-error/logerr

## LICENSE

[MIT](LICENSE)
