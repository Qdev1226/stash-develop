package static

import "embed"

//go:embed performer
var Performer embed.FS

//go:embed performer_male
var PerformerMale embed.FS

//go:embed scene
var Scene embed.FS

//go:embed image
var Image embed.FS

//go:embed tag
var Tag embed.FS

//go:embed studio
var Studio embed.FS
