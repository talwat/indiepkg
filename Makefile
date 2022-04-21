build:
	@go build

install:
	@mkdir -p $$HOME/.local/bin/indiepkg
	@mv indiepkg $$HOME/.local/bin/indiepkg

all:
	@scripts/dev/build.py