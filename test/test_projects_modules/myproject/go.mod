module dumpplane.io/myproject

go 1.20

replace dumpplane.io/myproject/module1 => ./module1

replace dumpplane.io/myproject/module2 => ./module2

require (
	dumpplane.io/myproject/module1 v0.0.0-00010101000000-000000000000
	dumpplane.io/myproject/module2 v0.0.0-00010101000000-000000000000
)
