cover:
	# go install github.com/nikolaydubina/go-cover-treemap@latest
	go test -coverprofile cover.out ./...
	go-cover-treemap -coverprofile cover.out > out.svg