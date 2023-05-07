module example/hello

go 1.20

require (
	example/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.9.0 // indirect
	rsc.io/sampler v1.99.99 // indirect
)

replace example/greetings => ../greetings
