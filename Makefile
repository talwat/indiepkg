build:
	@go build

all:
	@-rm -rf output
	@-mkdir output

	@env GOOS=darwin GOARCH=arm64 go build
	@mv indiepkg output/indiepkg-macOS_m1
	@env GOOS=darwin GOARCH=amd64 go build
	@mv indiepkg output/indiepkg-macOS_intel

	@env GOOS=linux GOARCH=arm64 go build
	@mv indiepkg output/indiepkg-linux_arm64
	@env GOOS=linux GOARCH=arm go build
	@mv indiepkg output/indiepkg-linux_arm
	@env GOOS=linux GOARCH=amd64 go build
	@mv indiepkg output/indiepkg-linux_amd64
	@env GOOS=linux GOARCH=386 go build
	@mv indiepkg output/indiepkg-linux_i386
