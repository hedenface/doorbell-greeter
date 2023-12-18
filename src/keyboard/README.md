# Keyboard

This is the badly named keyboard application. It's badly named because it should probably be called rotary-keyboard. Or something like that.

Regardless, this takes Raspberry PI GPIO inputs from a specific rotary dial from a phone from the fourteen hundreds and converts them into keyboard key presses using the [robotgo package](https://github.com/go-vgo/robotgo) (this is a really neat piece of software - the robotgo, not the keyboard application).

This works well because after we start this application, we start the web browser that launches the doorbell website, which automatically handles the focus of the keyboard.

## Build Dependencies

```bash
sudo apt-get install gcc libc6-dev libx11-dev xorg-dev libxtst-dev
```
