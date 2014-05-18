# nocache

Martini middleware/handler for removing ETag related headers and adding "no cache" headers

## Usage

~~~ go
package main

import (
    "github.com/go-martini/martini"
    "github.com/mytrile/nocache"
)

func main() {
    m := martini.Classic()
    m.Use(nocache.UpdateCacheHeaders())
    m.Get("/", func() string {
      return "Hello world!"
    })
    m.Run()
}

~~~

The middleware will remove the following headers:

  * ETag
  * If-Modified-Since
  * If-Match
  * If-None-Match
  * If-Range
  * If-Unmodified-Since

and add the following headers:

  * Cache-Control: no-cache, no-store, max-age=0, must-revalidate
  * Pragma:        no-cache
  * Expires:       Fri, 29 Aug 1997 02:14:00 EST

## Meta

* Author  : Dimitar Kostov
* Email   : mitko.kostov@gmail.com
* Website : [http://mytrile.github.com](http://mytrile.github.com)
* Twitter : [http://twitter.com/mytrile](http://twitter.com/mytrile)