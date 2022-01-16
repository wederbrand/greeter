module github.com/wederbrand/greeter

go 1.17

require github.com/wederbrand/greeting v1.1.0

// go mod edit --replace github.com/wederbrand/greeting=../greeting
replace github.com/wederbrand/greeting => ../greeting
