package embedded

import (
	"embed"
)

//go:embed dir1
//go:embed dir2
var Dir12 embed.FS

//go:embed dir3
var Dir3 embed.FS

//go:embed A/B/C
var ABC embed.FS
