build:
	@go build

install:
	@mkdir -p $$HOME/.local/bin/
	@mv indiepkg $$HOME/.local/bin/indiepkg

all:
	@scripts/dev/build.py