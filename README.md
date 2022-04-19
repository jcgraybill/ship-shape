# Ship shape
Supply chain management indie game ... IN SPACE! Current state is beyond preliminary - there isn't actually a game yet. 

Built in [golang](https://go.dev/) using the [ebiten](https://ebiten.org/) 2D game library.

# Objectives
* It's real time. The objective is to produce a surplus of materials within a time limit.
* You make progress even when you lose.
* There isn't a combat mechanic.
* The player's direct control is limited: most of what happens is autonomous decisions by game entities.
* There's a deep tech tree. The supply chain is the tech tree. Distance matters.
* The user interface is shallow: everything shows what it's doing and uses cultural referents.

![screenshot](https://github.com/jcgraybill/ship-shape/blob/main/screenshot.png)

# Running it

With go installed, download and run the game with:

```
go run -tags=deploy github.com/jcgraybill/ship-shape@latest
```

