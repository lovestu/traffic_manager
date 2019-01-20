package main

import (
    "engo.io/engo"
    "engo.io/ecs"
    "engo.io/engo/common"
    "log"
)

type myScene struct {
}

func (*myScene) Type() string {
    return "myGame"
}

func (*myScene) Preload() {
    engo.Files.Load("textures/city.png")
}

func (*myScene) Setup(u engo.Updater) {
    world, _ := u.(*ecs.World)
    world.AddSystem(&common.RenderSystem{})

    city := City{BasicEntity: ecs.NewBasic()}
    city.SpaceComponent = common.SpaceComponent{
        Position: engo.Point{10, 10},
        Width: 303,
        Height: 641,
    }

    texture, err := common.LoadedSprite("textures/city.png")
    if err != nil {
        log.Println("Unable to load texture: ", err.Error())
    }

    city.RenderComponent = common.RenderComponent{
        Drawable: texture,
        Scale: engo.Point{1, 1},
    }

    for _, system := range world.Systems() {
        switch sys := system.(type) {
        case *common.RenderSystem:
            sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
        }
    }
}

type City struct {
    ecs.BasicEntity
    common.RenderComponent
    common.SpaceComponent
}

func main() {
    opts := engo.RunOptions {
        Title: "Hello World",
        Width: 400,
        Height: 400,
    }
    engo.Run(opts, &myScene{})
}
