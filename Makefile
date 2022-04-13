build:
	@go build

install:
	@mv indiepkg $$HOME/.local/bin/indiepkg

all:
	@scripts/build.py