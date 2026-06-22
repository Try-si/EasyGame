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

func NewGame(update func(float32) error, debug bool) {
	config = ETEhelper.JsonToStruct[Config]("config.json")
	updateFunc = update

	ETEngine.Init(Update, "config.json") // Initialize the engine
	ETEngine.Game.Debug = debug
	ETEngine.GameLoop() // Start the game loop
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

func MoveElement(elementID string, position [3]float32) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID].Pos[0] -= position[0]
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID].Pos[1] += position[1]
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID].Z = int(position[2])
}

func RotateElement(elementID string, angle float32) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID].Rotation = angle
}

func ScaleElement(elementID string, scale [2]int) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID].Size = scale
}

func AddEntity(id string, name string, pos [2]float32, rotation float32, Z int, metaData map[string]string) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[id] = &ETECore.Element{
		Name:     name,
		Pos:      pos,
		Rotation: rotation,
		Z:        Z,
		MetaData: metaData,
	}
}

func SetAnimation(entityID string, animation string) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Animation = animation
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

func MoveCamera(position [3]float32) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Offset[0] += position[0]
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Offset[1] -= position[1]
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Z += position[2]
}

func Quit() {
	ETEngine.Game.Quite = true
}
