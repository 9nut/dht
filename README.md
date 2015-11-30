# DHT
This is a trivial Go package for reading the _DHT11_ and _DHT22_
temperature and humidity sensors through a GPIO pin on Raspberry Pi.

It uses (via cgo) low level C drivers which are part of the
[Adafruit Python DHT](https://github.com/adafruit/Adafruit_Python_DHT)
library

# Example

See the **example** directory for usage.  The example may need *root*
permission to access the GPIO pins.

# Setup

Here's a typical hardware setup:
![file_000](https://cloud.githubusercontent.com/assets/616755/11465492/fcdf23e4-96ed-11e5-8325-ba12bbbdc248.jpeg)
# Copyright

Each file contains its own copyright notice by the authors.



