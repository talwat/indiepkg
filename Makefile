build:
	@go build

install:
	@mv indiepkg $$HOME/.local/bin/indiepkg

all:
	@scripts/dev/build.py