module github.com/gomesmf/gopl/ch07-interfaces/surface

go 1.19

replace github.com/gomesmf/gopl/ch03-basic-data-types/surf => ../../ch03-basic-data-types/surf

replace github.com/gomesmf/gopl/ch07-interfaces/eval => ../eval

require (
	github.com/gomesmf/gopl/ch03-basic-data-types/surf v0.0.0-00010101000000-000000000000
	github.com/gomesmf/gopl/ch07-interfaces/eval v0.0.0-00010101000000-000000000000
)
