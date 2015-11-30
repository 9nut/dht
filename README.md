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
![alt text](https://drive.google.com/open?id=0B0sQhgOyZZBsWGdDdkdpY2xPNUE)

# Copyright

Each file contains its own copyright notice by the authors.



