<h1 align="center">
  <br>
  <img src="res/banner.png" alt="Rubik's Cube" width="400" height="200">
  <br>
  TSP Rubik's Cube
  <br>
</h1>

This is a little game I wrote for my Trimui Smart Pro. It's optimized for 1280x720 resolution.

It uses [Golang bindings](https://github.com/gen2brain/raylib-go) for [RayLib](https://github.com/raysan5/raylib). My initial motivation was to learn some 3D programming.

It has 8 tutorials that show how to assemble some non-trivial states:
<img src="res/tutorial.png" alt="Tutorial" width="536" height="300">

You can also play it on your desktop computer using a mouse and a keyboard.

## Installation on Trimui Smart Pro

1. Download the zip file from https://github.com/geniot/tsp-rubik/releases Unpack it locally. 
2. Add a folder to your SD card. Eg. /mnt/SDCARD/Apps/Rubik and put the files from zip there. 
3. There is also a config.json file in /mnt/SDCARD/Apps/ which you need to modify: add a record for the new app, set 'show' to 1 (true). 
4. Reboot your TSP to pick up the changes. 
5. For updates: just rewrite the files in /mnt/SDCARD/Apps/Rubik

## Build

I use a Docker image with TSP SDK installed: https://github.com/geniot/trimui-smart-pro-toolchain

See the Makefile for more details.

For my desktop testing I use GoLand and Ubuntu.

## Deploy to TrimUI

1. On TrimUI find System Tools > Network > Display IP
2. Type `export IP=<IP>`
3. Optional: start Dummy app on TrimUI (https://github.com/geniot/tsp-dummy-banner)
4. Run `make all`


### WASM

`GOOS=js GOARCH=wasm go build -o ./Raylib-Go-Wasm/index/main.wasm tsp-rubik/src tsp-rubik/src`