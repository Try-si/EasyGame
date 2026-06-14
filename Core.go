package EasyGame

import (
	ETEngine "github.com/Try-si/ETE"
	"github.com/Try-si/ETE/ETECore"
	ETEhelper "github.com/Try-si/ETE/ETEHelper"
	IAs "github.com/Try-si/EasyGame/IA"
	"github.com/Try-si/EasyGame/Inputs"
)

var updateFunc func(float32) error
var config Config

func NewGame(update func(float32) error) {
	config = ETEhelper.JsonToStruct[Config]("config.json")
	updateFunc = update

	ETEngine.Init(Update, "config.json") // Initialize the engine
	ETEngine.GameLoop()                  // Start the game loop
}

func Update(deltaTime float32) error {
	Inputs.Update()
	err := IAs.Update(deltaTime)
	if err != nil {
		return err
	}
	err = updateFunc(deltaTime)
	if err != nil {
		return err
	}
	return nil
}

func MoveElement(elementID string, position [2]float32) {
	element := ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID]
	element.Pos = position
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID] = element
}

func RotateElement(elementID string, angle float32) {
	element := ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID]
	element.Rotation = angle
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID] = element
}

func ScaleElement(elementID string, scale [2]int) {
	element := ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID]
	element.Size = scale
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID] = element
}

func AddEntity(id string, name string, pos [2]float32, rotation float32, layer int, metaData map[string]string) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[id] = &ETECore.Element{
		Name:     name,
		Pos:      pos,
		Rotation: rotation,
		Layer:    layer,
		MetaData: metaData,
	}
}

func RemoveEntity(entityID string) {
	delete(ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements, entityID)
}

func SaveGame(path string) {
	ETEhelper.SaveFile(path, ETEhelper.StructToJson(ETEngine.Game))
}

func LoadGame(path string) {
	ETEngine.Game = ETEhelper.JsonToStruct[*ETECore.Game](path)
}
