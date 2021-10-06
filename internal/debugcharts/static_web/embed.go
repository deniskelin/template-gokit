package static_web

import "embed"

//go:embed static
var StaticFiles embed.FS

//go:embed index.html
var IndexHTMLPage []byte
