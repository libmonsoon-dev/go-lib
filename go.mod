module github.com/libmonsoon-dev/go-lib

go 1.23.1

require (
	github.com/deckarep/gosx-notifier v0.0.0-20180201035817-e127226297fb
	github.com/paulmach/orb v0.11.1
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c
	golang.org/x/sync v0.8.0
	golang.org/x/sys v0.26.0
	golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da
	gopkg.in/toast.v1 v1.0.0-20180812000517-0a84660828b2
)

require github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect

replace gopkg.in/toast.v1 => github.com/libmonsoon-dev/toast v0.0.0-20241014200443-d0a04cde6d5c

replace github.com/deckarep/gosx-notifier => github.com/libmonsoon-dev/gosx-notifier v0.0.0-20241014200646-92526e117586
