module github.com/akishichinibu/goadt/pkg/tuple

go 1.24.1

replace github.com/akishichinibu/goadt/pkg/runtime => ../runtime

require (
	github.com/akishichinibu/goadt/pkg/runtime v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
