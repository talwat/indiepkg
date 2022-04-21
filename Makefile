build:
	@go build

install:
	@mkdir $$HOME/.local/bin/indiepkg
	@mv indiepkg $$HOME/.local/bin/indiepkg

all:
	@scripts/dev/build.py