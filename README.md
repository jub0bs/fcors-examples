# fcors-examples

A collection of examples showing how to use [github.com/jub0bs/fcors][fcors],
a principled CORS middleware library,
in conjunction with popular Web frameworks and HTTP toolkits for [Go][go].

| Library                                | Major version (pre-v2 if unspecified) | Path to example                                                      | Comment                        |
| -------------------------------------- | ------------------------------------- | -------------------------------------------------------------------- | ------------------------------ |
| [chi][chi]                             |                                       | [chi/main.go](chi/main.go)                                           |                                |
| [Echo][echo-v4]                        | v4                                    | [echo-v4/main.go](echo-v4/main.go)                                   |                                |
| [Fiber][fiber]                         | v2                                    | [fiber-v2/main.go](fiber-v2/main.go)                                 |                                |
| [Gin][gin]                             |                                       | [gin/main.go](gin/main.go)                                           |                                |
| [gorilla/mux][gorilla]                 |                                       | [gorilla-mux/main.go](gorilla-mux/main.go)                           | unmaintained as of Dec 9, 2022 |
| [julienschmidt/httprouter][httprouter] |                                       | [julienschmidt-httprouter/main.go](julienschmidt-httprouter/main.go) |                                |
| [net/http][net-http]                   |                                       | [net-http/main.go](net-http/main.go)                                 | standard-library package       |

[chi]: https://go-chi.io/#/
[echo-v4]: https://echo.labstack.com/
[fcors]: https://pkg.go.dev/github.com/jub0bs/fcors
[fiber]: https://gofiber.io/
[gin]: https://gin-gonic.com/
[go]: https://go.dev/
[gorilla]: https://github.com/gorilla/mux
[httprouter]: https://pkg.go.dev/github.com/julienschmidt/httprouter
[net-http]: https://pkg.go.dev/net/http
