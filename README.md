[![FriendsOfGo](https://img.shields.io/badge/powered%20by-Friends%20of%20Go-73D7E2.svg)](https://friendsofgo.tech)
[![Go Report Card](https://goreportcard.com/badge/github.com/friendsofgo/errors)](https://goreportcard.com/report/github.com/friendsofgo/errors)
[![GoDoc](https://godoc.org/github.com/friendsofgo/errors?status.svg)](https://godoc.org/github.com/friendsofgo/errors)

# errors

This package is a fork from [github.com/pkg/errors](https://github.com/pkg/errors) package created by
[Dave Cheney](https://github.com/davecheney). The original package has no longer accepting proposals for new functionality.

With the new errors on [go 1.13](https://godoc.org/errors), the way to using the errors on Go has some
changes that can be applied into Dave Cheney library. We want to offer one way to migrate your code to new
errors, but with the minimum refactor, for that we've created this package.

This package provide the same interface that the original library have, but using new [go 1.13](https://godoc.org/errors)
errors.

## How to start using friendsofgo/errors

If you previously was using the package [github.com/pkg/errors](https://github.com/pkg/errors), you only need
change your imports for **github.com/friendsofgo/errors**, with this simple change now you're capable to use
[go 1.13](https://godoc.org/errors) in your code, and use the new methods `As` and `Is` if you want.

Furthermore the method `Wrap` `Wrapf become compatible with `Unwrap` interface of new [go 1.13](https://godoc.org/errors) errors.

## Adding context to an error

With the original package [go 1.13](https://godoc.org/errors) if you want add context, ergo wrap your error you need to create
a new error and using the new verb `"%w" like that:

```go
_, err := ioutil.ReadAll(r)
if err != nil {
        return errors.fmt("read failed: %w", err)
}
```

Using our library you can do that forgetting to the new verb:

```go
_, err := ioutil.ReadAll(r)
if err != nil {
        return errors.Wrap(err, "read failed")
}
```

## Retrieving the cause of an error

We want to keep the compatibility with the [github.com/pkg/errors](https://github.com/pkg/errors) package, for that
our package provides a `Cause` method, but this method is not longer needed, because we can use the new methods `Is` or `As`
that provides the official package.

So previously if you needed to check an error cause, your error must be implemented the `causer` inteface:

```go
type causer interface {
        Cause() error
}
```

`errors.Cause` will recursively retrieve the topmost error which does not implement causer, which is assumed to be the original cause. For example:

```go
switch err := errors.Cause(err).(type) {
case *MyError:
        // handle specifically
default:
        // unknown error
}
```

But now you can do:

```go
var target *MyError
if errors.As(err, &target) {
    // handle specifically
} else {
   // unknown error
}
```

Or if you uses a sentinel error:

```go
var ErrMyError = errors.New("my sentinel error")
if errors.Is(err, ErrMyError) {
    // handle specifically
} else {
   // unknown error
}
```

## Disclaimer
This package is only compatible from go1.13 onwards. 

## Contributing

[Contributions](https://github.com/friendsofgo/errors/issues?q=is%3Aissue+is%3Aopen) are more than welcome, if you are interested please fork this repo and send your Pull Request.
