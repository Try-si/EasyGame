# Description

EasyGame est un framework de jeu vidéo fait en go pour faciliter la création de jeux vidéo en permettant de ne faire presque que de la configuration.

## Installation

```bash
go get github.com/Try-si/EasyGame
```

## Utilisation 

### Exemple minimal

``` files

/Mon jeu
 Textures/
    Dog.png // c'est un exemple
    Skeleton.png // c'est un exemple
    nos textures
 Sounds/ // non utilisé pour le moment
    nos sons
 Maps/
    Maps/*
        Overworld.tmx // c'est un exemple
        nos maps tiled (format tmx)
    tilesets/ // non obligatoire
        nos tilesets tiled (format tsx)
    Json/*
        Overworld.json // c'est un exemple
        nos maps json (format json)
    Tiles.json
    Elements.json
    Maps.json
 IAs/
    IAs.go
    Dog.go
    Skeleton.go
    Guard.go
 Animations/
    Exemple.json
    Animations.json
 config.json // obligatoire
 main.go
 go.mod
 go.sum

 * == tous les fichiers qui sont dans ces dossiers on tous un fichier avec le même nom dans les autres dossiers
```

Pour main.go :

```go
package main

import "github.com/Try-si/EasyGame"
import "my_game/IAs"

func main() {
    IAs.Init()
    EasyGame.NewGame()
}
```

Pour config.json :

```json
{
    "ScreenWidth": 800,
    "ScreenHeight": 600,
    "Title": "My Game",
    "Map": "Overworld",         // nom de la map à charger au démarrage
    "AdaptativeSize": false,     // si les element doive retrécire quand la taille de la fenêtre est augmenté

    "SpritePath": "Textures",
    "AnimationPath": "Animations/Animations.json",
    "SoundPath": "Sounds",
    "MapsPath": "Maps/Maps.json",
    "IAsPath": "IAs/IAs.go"
}
```

Pour Maps.json :

```json
{
    "Maps": ["Overworld"],
    "JsonMap": "Json",
    "TiledMap": "Maps",
    "Elements": "Elements.json",
    "Tiles": "Tiles.json",
    "Parrallax": true,
    "ParrallaxFactor": 1.0
}
```

Pour Elements.json :

```json
{
    "Elements": {
        "Player": {
            "Image": "Player.png",
            "Size": [32, 32],
            "Box": [0, 0, 0, 0] // width, height (si il est == a 0 alors c'est un cercle et width = rayon), box pos x, box pos y
        }
    }
}
```

Pour Tiles.json :

```json
{
    "0":{// is a tile id in map of tiled of grass
        "Animation": "nil",// cela peut etre nil
        "Box": [32,32,0,0],
        "Tags": ["Grass"]
    },
    "1":{// is a tile id in map of tiled of water
        "Animation": "Water_idle",
        "Box": [32,32,0,0],
        "Tags": ["Water"]
    }
}
```

Pour Animations.json :

```json
["Exemple"]
```

Pour Overworld.json :

```json
{
    "Map": "Overworld",         // nom de la map
    "CellSize": 1,              // taille de la cellule
    "Unité": 32,                // taille d'une unité en pixels
    "Cam": {
        "Z": 1.0,
        "Offset": [0.0, 0.0]
    },

    "Elements": {
        "Player": {
            "Name": "Player",  // nom de l'élément dans Elements.json
            "Pos": [0.0, 0.0], // position de l'élément
            "Rotation": 0,
            "Z": 1,
            "Visible": true,
            "MetaData": {
                "Nom de la variable": "valeur de la variable"
            }
        }
    }   // éléments dans la map
}
```

Pour IAs.go :
```go
package IAs

import (
    "github.com/EasyGame/IA"
    "my_game/IAs"
)

func Init() {
    IA.Init()

    IAs.InitDog()
    IAs.InitSkeleton()
    IAs.InitGuard()
}
```

Pour Dog.go (c'est encore un exemple):
```go
package IAs

import "github.com/Try-si/EasyGame/IA/FSM"

func InitDog() {
    FSM.AddState("idle_dog", FSM.State{
        Action: func(data any) {
            
        },
    })
}
```

Pour Exemple.json :
```json 
{
    "Player_idle": {
        "Frames": [
            {"Frame": "Player_idle_1.png", "Duration": "rand", "Box": [32, 32, 0, 0]},
            {"Frame": "Player_idle_2.png", "Duration": "rand", "Box": [32, 32, 0, 0]},
            {"Frame": "Player_idle_3.png", "Duration": "rand", "Box": [32, 32, 0, 0]}
        ],
        "Speed": 1
    }
}
```