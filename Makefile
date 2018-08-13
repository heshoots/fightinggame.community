build:
		CGO_ENABLED=0 GOOS=linux go build -o fgcmvp -a -installsuffix cgo cmd/main.go

package:
		zip fgcmvp.zip fgcmvp assets/static/*

