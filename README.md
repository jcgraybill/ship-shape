# Ship shape
Supply chain management indie game ... IN SPACE! Current state is preliminary - there's a six-level tutorial, about an hour's worth of gameplay. You can try it out in your browser [here](https://jcgraybill.github.io/ship-shape/). 

**Ship shape was mentioned in the [Ebitengine 2022 year in review](https://ebitengine.org/en/blog/2022.html#Other_Games_and_Applications)** :) 

Built with [Golang](https://go.dev/) using the [Ebitengine](https://ebitengine.org/) game engine and [gg](https://github.com/fogleman/gg) graphics library.

# Design principles
The game in its current state meets *some* but not yet all of these.
* It's real time. The objective is to produce a surplus of materials within a time limit.
* It conveys a sense that you're playing in a real physical place. You build and progress even when you don't meet all the objectives for an individual level.
* There isn't a combat mechanic.
* It's a management game. The player's direct control is limited: most of what happens is autonomous decisions by game entities.
* There's a deep tech tree. The supply chain is the tech tree. Distance matters.
* The user interface is shallow: everything shows what it's doing and uses already well-understood cultural referents (ice, water, computers) rather than requiring players to learn game-specific terminology.

![screenshot](https://github.com/jcgraybill/ship-shape/blob/main/screenshot.png)

# Building it locally

With go installed, download and run the game with:

```
go run github.com/jcgraybill/ship-shape@latest
```

To build on Ubuntu, first install additional packages:
```
apt install libgl1-mesa-dev xorg-dev libasound2-dev
```
