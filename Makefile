baud_rate := 115200

.PHONY: build flash deps pip-deps

flash: deps
	tinygo flash -target=esp32-mini32 -port="${port}" .

build: main
	tinygo build .

deps: pip-deps

pip-deps: pip-deps.done

pip-deps.done:
	pip install esptool
	touch pip-deps.done

serial:
	screen $(port) ${baud_rate}