package main

import (
	"fmt"
	"net/url"
	"regexp"
)

func main() {
	pkgRe := regexp.MustCompile("^(?:@([^/]+?)[/])?([^/]+?)$")
	match := pkgRe.FindStringSubmatch("@myor?g/mypackage")
	fmt.Println(url.PathEscape(match[1]))

}
