module github.com/appliedgocourses/mastergo/2_LanguageBasics/2-12-Modules/2-12-7_how_go_selects_a_module_version/selectversion

go 1.14

require (
	github.com/appliedgocourses/B v1.0.0
	github.com/appliedgocourses/C v1.0.0
)

exclude github.com/appliedgocourses/D v1.1.0

replace github.com/appliedgocourses/B@v1.0.0 => ./B

replace github.com/appliedgocourses/C@v1.0.0 => ./C
