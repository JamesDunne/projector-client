module Projector

go 1.17

//require golang.org/x/mobile v0.0.0-20211109191125-d61a72f26a1a
require golang.org/x/mobile v0.0.0-jsd

replace golang.org/x/mobile v0.0.0-jsd => I:/Developer/mobile

require (
	//golang.org/x/exp v0.0.0-20190731235908-ec7cb31e5a56 // indirect
	golang.org/x/exp v0.0.0-jsd // indirect
	golang.org/x/image v0.0.0-20190802002840-cff245a6509b // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
)

require github.com/gobwas/ws v1.1.0

require (
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
)

replace golang.org/x/exp v0.0.0-jsd => I:/Developer/exp
