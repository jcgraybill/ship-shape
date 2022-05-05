# Ship shape
Supply chain management indie game ... IN SPACE! Current state is preliminary - there's a six-level tutorial, about an hour's worth of gameplay. You can try it out in your browser [here](https://jcgraybill.github.io/ship-shape/).

Built in [golang](https://go.dev/) using the [ebiten](https://ebiten.org/) 2D game library and [gg](https://github.com/fogleman/gg) 2D graphics library.

# Objectives
* It's real time. The objective is to produce a surplus of materials within a time limit.
* You make progress even when you lose.
* There isn't a combat mechanic.
* The player's direct control is limited: most of what happens is autonomous decisions by game entities.
* There's a deep tech tree. The supply chain is the tech tree. Distance matters.
* The user interface is shallow: everything shows what it's doing and uses cultural referents.

![screenshot](https://github.com/jcgraybill/ship-shape/blob/main/screenshot.png)

# Building it locally

With go installed, download and run the game with:

```
go run github.com/jcgraybill/ship-shape@latest
```

To build on Ubuntu, install additional packages:
```
apt install libgl1-mesa-dev xorg-dev libasound2-dev
```
