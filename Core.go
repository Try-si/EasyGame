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
	// TODO: Implement element movement
}

func RotateElement(elementID string, angle float32) {
	// TODO: Implement element rotation
}

func ScaleElement(elementID string, scale [2]float32) {
	// TODO: Implement element scaling
}

func AddEntity(Animation string, Size [2]int, Box [4]float32, Tags []string, Name string, Pos [2]float32, Rotation float32, Layer int, MetaData map[string]string) {

}

func RemoveEntity(entityID string) {

}

func SaveGame(path string) {
	ETEhelper.SaveFile(path, ETEhelper.StructToJson(ETEngine.Game))
}

func LoadGame(path string) {
	ETEngine.Game = ETEhelper.JsonToStruct[*ETECore.Game](path)
}
