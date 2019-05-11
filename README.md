# cache

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Patreon][patreon-image]][patreon-url]

Cache arbitrary data with an expiration time.

## Features

* 0 dependencies
* Less than 100 lines of code
* 100% test coverage

## Usage

```go
// New cache
c := cache.New(5 * time.Minute)

// Put something into the cache
c.Set("a", "b", 1 * time.Minute)

// Read from the cache
obj, found := c.Get("a")

// Convert the type
fmt.Println(obj.(string))
```

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars2.githubusercontent.com/u/438936?s=70&v=4)](https://twitter.com/eduardurbach) |
| --- | --- |
| [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://www.patreon.com/eduardurbach)

[godoc-image]: https://godoc.org/github.com/akyoto/cache?status.svg
[godoc-url]: https://godoc.org/github.com/akyoto/cache
[report-image]: https://goreportcard.com/badge/github.com/akyoto/cache
[report-url]: https://goreportcard.com/report/github.com/akyoto/cache
[tests-image]: https://cloud.drone.io/api/badges/akyoto/cache/status.svg
[tests-url]: https://cloud.drone.io/akyoto/cache
[coverage-image]: https://codecov.io/gh/akyoto/cache/graph/badge.svg
[coverage-url]: https://codecov.io/gh/akyoto/cache
[patreon-image]: https://img.shields.io/badge/patreon-donate-green.svg
[patreon-url]: https://www.patreon.com/eduardurbach
