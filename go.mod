module github.com/doctorj/gobatis

go 1.12

require (
	github.com/antonmedv/expr v1.1.4
	github.com/go-sql-driver/mysql v1.4.1
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.0.0-00010101000000-000000000000
)

replace gopkg.in/yaml.v2 => github.com/go-yaml/yaml v2.1.0+incompatible
