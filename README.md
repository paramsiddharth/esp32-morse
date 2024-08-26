# Morse Code Generator (ESP32)
A simple text to morse code generator that runs on an ESP-32 microcontroller.

# Instructions
Follow the given instructions to deploy this app to your own ESP32 microcontroller.

For the examples, assume the port to be accessible at `/dev/ttyUSB0`.

## Installing Dependencies
Install all Python and locally.
```bash
make deps
```

## Compile
Build the app.
```bash
make build port=/dev/ttyUSB0
```

## Upload
Burn the app with the firmware to the microcontroller.
```bash
make flash port=/dev/ttyUSB0
```

You're done! You now have a morse code generator running on your microcontroller, accessible via a serial console connection.

If you connect to the serial monitor and enter some text, you'll see morse code being printed to the console as well as being rendered via the build-in LED.
```bash
make serial port=/dev/ttyUSB0
```

# Made with ❤ by [Param](https://www.paramsid.com).