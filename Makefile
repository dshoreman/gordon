PREFIX ?= /usr

all:
	@echo "Building Gordon..."
	@go build -i -v -o bin/gordon && echo " [DONE]"

install:
	@echo "Preparing package structure"
	@mkdir -p "$(DESTDIR)$(PREFIX)/bin"

	@echo "Installing Gordon..."
	@cp bin/gordon "$(DESTDIR)$(PREFIX)/bin/gordon"

uninstall:
	@echo "Uninstalling Gordon..."
	@rm -v "$(DESTDIR)$(PREFIX)/bin/gordon"
	@echo "Uninstall complete"
