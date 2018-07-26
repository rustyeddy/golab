# gotto

Go otto - The Macro and Master controller.

## Preparing Arduino for firmata

```bash
$ gort scan serial
$ gort arduino install
$ gort arduino upload firmata /dev/tty.usbmodemXXXX
```

Linux machines (at least the debian derivates) are likely to have a serial
port something like: ```/dev/ttyACM0```.
