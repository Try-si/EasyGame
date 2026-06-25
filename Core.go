package EasyGame

import (
	"math"

	ETEngine "github.com/Try-si/ETE"
	"github.com/Try-si/ETE/ETECore"
	ETEhelper "github.com/Try-si/ETE/ETEHelper"
	"github.com/Try-si/EasyGame/IA"
	IAs "github.com/Try-si/EasyGame/IA"
	"github.com/Try-si/EasyGame/Inputs"
	"github.com/Try-si/IAM/Core"
	"github.com/hajimehoshi/ebiten/v2"
)

var updateFunc func(float32) error
var confi config

func NewGame(update func(float32) error, IAInit func(), debug bool) {
	confi = ETEhelper.JsonToStruct[config]("config.json")
	updateFunc = update

	Inputs.Init(confi.InputsPath)

	ETEngine.Init(Update, "config.json") // Initialize the engine
	ETEngine.Game.Debug = debug
	IA.Init(IAInit)
	ETEngine.GameLoop() // Start the game loop
}

func Update(deltaTime float32) error {
	Inputs.Update()
	err := IAs.Update(deltaTime, Maps.GetMapName())
	if err != nil {
		return err
	}
	err = updateFunc(deltaTime)
	if err != nil {
		return err
	}
	return nil
}

func (e *entitys) MoveElement(elementID string, position [3]float32) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID].Pos[0] += position[0]
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID].Pos[1] -= position[1]
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID].Z -= position[2]
}

func (e *entitys) RotateElement(elementID string, angle float32) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID].Rotation = angle
}

func (e *entitys) ScaleElement(elementID string, scale [2]int) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[elementID].Size = scale
}

func (e *entitys) AddEntity(id string, name string, pos [2]float32, rotation float32, Z float32, metaData map[string]string) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[id] = &ETECore.Element{
		Name:     name,
		Pos:      pos,
		Rotation: rotation,
		Z:        Z,
		MetaData: metaData,
	}
}

func (e *entitys) SetAnimation(entityID string, animation string) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Animation = animation
}

func (e *entitys) RemoveEntity(entityID string) {
	delete(ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements, entityID)
}

func (e *entitys) GetEntity(entityID string) *ETECore.Element {
	return ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID]
}

func (e *entitys) GetEntities() map[string]*ETECore.Element {
	return ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements
}

func (e *entitys) GetEntityPosition(entityID string) [3]float32 {
	return [3]float32{
		ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Pos[0],
		-ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Pos[1],
		ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Z,
	}
}

func (e *entitys) SetVisible(entityID string, visible bool) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Visible = visible
}

func (e *entitys) GetVisible(entityID string) bool {
	return ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Visible
}

func (e *entitys) SetMetaData(entityID string, key string, value string) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].MetaData[key] = value
}

func (e *entitys) GetMetaData(entityID string, key string) string {
	return ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].MetaData[key]
}

func (e *entitys) GetFrame(entityID string) int {
	return ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Frame
}

func (e *entitys) SetFrame(entityID string, frame int) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Frame = frame
}

func (e *entitys) AddFrame(entityID string, frame int) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Frame += frame
}

func (e *entitys) GetFFrame(entityID string) int {
	return ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].FFrame
}

func (e *entitys) SetFFrame(entityID string, frame int) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].FFrame = frame
}

func (e *entitys) AddFFrame(entityID string, frame int) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].FFrame += frame
}

func (e *entitys) GetRotation(entityID string) float32 {
	return ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Rotation
}

func (e *entitys) SetRotation(entityID string, rotation float32) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Rotation = rotation
}

func (e *entitys) GetZAxis(entityID string) float32 {
	return ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Z
}

func (e *entitys) SetElementPosition(entityID string, position [3]float32) {
	if position[0] != math.MaxFloat32 {
		ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Pos[0] = position[0]
	}
	if position[1] != math.MaxFloat32 {
		ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Pos[1] = position[1]
	}
	if position[2] != math.MaxFloat32 {
		ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Z = position[2]
	}
}

func (e *entitys) GetPresanceOfTag(entityID, tag string, tagId int) bool {
	return ETEngine.Game.Elements[ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Name].Tags[tagId] == tag
}

func (e *entitys) GetTags(entityID string) []string {
	return ETEngine.Game.Elements[ETEngine.Game.Maps[ETEngine.Game.Config.Map].Elements[entityID].Name].Tags
}

func (e *entitys) GetAllTags() map[string][]string {
	result := make(map[string][]string)
	for id, element := range ETEngine.Game.Elements {
		result[id] = element.Tags
	}
	return result
}

func (M *maps) GetMapName() string {
	return ETEngine.Game.Config.Map
}

func (M *maps) GetMap() *ETECore.Map {
	return ETEngine.Game.Maps[ETEngine.Game.Config.Map]
}

func (M *maps) GetWorldState() WorldState {
	return WorldState{}
}

func SaveGame(path string) {
	ETEhelper.SaveFile(path, ETEhelper.StructToJson(ETEngine.Game))
}

func LoadGame(path string) {
	ETEngine.Game = ETEhelper.JsonToStruct[*ETECore.Game](path)
}

func (c *camera) MoveCamera(position [3]float32) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Offset[0] -= position[0]
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Offset[1] -= position[1]
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Z -= position[2]
}

func (c *camera) GetCamera() *ETECore.Camera {
	return &ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam
}

func (c *camera) GetCameraPosition() [3]float32 {
	return [3]float32{
		ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Offset[0],
		ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Offset[1],
		ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Z,
	}
}

func (c *camera) SetCameraPosition(position [3]float32) {
	if position[0] != math.MaxFloat32 {
		ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Offset[0] = position[0]
	}
	if position[1] != math.MaxFloat32 {
		ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Offset[1] = position[1]
	}
	if position[2] != math.MaxFloat32 {
		ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Z = position[2]
	}
}

func (c *camera) ZoomCamera(zoom float32) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Zoom += zoom
}

func (c *camera) SetZoomCamera(zoom float32) {
	ETEngine.Game.Maps[ETEngine.Game.Config.Map].Cam.Zoom = zoom
}

func Quit() {
	ETEngine.Game.Quite = true
}

func (i *inputs) BindingIsPressed(binding string) bool {
	return Inputs.BindingIsPressed(binding)
}

func (i *inputs) BindingIsJustPressed(binding string) bool {
	return Inputs.IsBindingJustPressed(binding)
}

func (i *inputs) BindingIsJustReleased(binding string) bool {
	return Inputs.IsBindingJustReleased(binding)
}

func (i *inputs) SetBinding(key string, binding int) {
	Inputs.SetBinding(key, binding)
}

func (i *inputs) BindingExists(binding string) bool {
	return Inputs.BindingExists(binding)
}

func (ia *ias) AddEntityBT(element string, node BehaviourNode) {
	if IAs.World[Maps.GetMapName()] == nil {
		IAs.World[Maps.GetMapName()] = IA.INITWORLD()
	}
	IAs.World[Maps.GetMapName()].AddEntityBT(element, bTtoCoreBT(&node, element))
}

func bTtoCoreBT(bt *BehaviourNode, element string) *Core.BehaviourNode {
	if bt == nil {
		return nil
	}
	return &Core.BehaviourNode{
		Condition: func(state Core.WorldState) (bool, any) {
			if bt.Condition == nil {
				return true, nil // Default behavior when condition is nil
			}
			return bt.Condition(Maps.GetWorldState())
		},
		Action: func(a any) {
			if bt.Action != nil { // Consider adding this check too
				bt.Action(element, a)
			}
		},
		TrueNode:  bTtoCoreBT(bt.TrueNode, element),
		FalseNode: bTtoCoreBT(bt.FalseNode, element),
	}
}

func (ia *ias) AddEntityBrain(element string, brain Brain) {
	if IAs.World[Maps.GetMapName()] == nil {
		IAs.World[Maps.GetMapName()] = IA.INITWORLD()
	}
	IAs.World[Maps.GetMapName()].AddEntityBrain(element, brainToCoreBrain(brain, element))
}

func brainToCoreBrain(brain Brain, element string) *Core.Brain {
	t := make(map[string][]Core.BrainTransition)
	for state, transitions := range brain.Transition {
		t[state] = make([]Core.BrainTransition, len(transitions))
		for i, transition := range transitions {
			t[state][i] = Core.BrainTransition{
				Weight: transition.Weight,
				To:     transition.To,
				Condition: func(ws Core.WorldState) (bool, any) {
					if transition.Condition == nil {
						return true, nil // Default behavior when condition is nil
					}
					return transition.Condition(Maps.GetWorldState())
				},
			}
		}
	}

	s := make(map[string]Core.State)
	for state, stateData := range brain.States {
		s[state] = Core.State{
			Action: func(name string, a any) {
				stateData.Action(name, a)
			},
		}
	}

	e := make(map[string]Core.BrainEntity)
	for entity, entityData := range brain.Entity {
		e[entity] = Core.BrainEntity{
			CurrentState: entityData.CurrentState,
			PréTrans: Core.BrainTransition{
				Weight: entityData.PréTrans.Weight,
				To:     entityData.PréTrans.To,
				Condition: func(ws Core.WorldState) (bool, any) {
					if entityData.PréTrans.Condition == nil {
						return true, nil // Default behavior when condition is nil
					}
					return entityData.PréTrans.Condition(Maps.GetWorldState())
				},
			},
		}
	}

	return &Core.Brain{
		Transition: t,
		States:     s,
		Entity:     e,
	}
}

func (ia *ias) AddEntityFSM(element string, fsm FSM) {
	if IAs.World[Maps.GetMapName()] == nil {
		IAs.World[Maps.GetMapName()] = IA.INITWORLD()
	}
	IAs.World[Maps.GetMapName()].AddEntityFSM(element, fSMtoCoreFSM(fsm, element))
}

func fSMtoCoreFSM(fsm FSM, element string) *Core.FSM {
	t := make(map[string][]Core.FSMTransition)
	for state, transitions := range fsm.Transition {
		t[state] = make([]Core.FSMTransition, len(transitions))
		for i, transition := range transitions {
			t[state][i] = Core.FSMTransition{
				To: transition.To,
				Condition: func(ws Core.WorldState) (bool, any) {
					if transition.Condition == nil {
						return true, nil // Default behavior when condition is nil
					}
					return transition.Condition(Maps.GetWorldState())
				},
			}
		}
	}

	s := make(map[string]Core.State)
	for state, stateData := range fsm.State {
		s[state] = Core.State{
			Action: func(name string, a any) {
				stateData.Action(name, a)
			},
		}
	}

	e := make(map[string]Core.FSMEntity)
	for entity, entityData := range fsm.Entity {
		e[entity] = Core.FSMEntity{
			CurrentState: entityData.CurrentState,
			PréTrans: Core.FSMTransition{
				To: entityData.PréTrans.To,
				Condition: func(ws Core.WorldState) (bool, any) {
					if entityData.PréTrans.Condition == nil {
						return true, nil // Default behavior when condition is nil
					}
					return entityData.PréTrans.Condition(Maps.GetWorldState())
				},
			},
		}
	}

	return &Core.FSM{
		Transition: t,
		States:     s,
		Entity:     e,
	}
}

func GetFPS() float32 {
	return float32(ebiten.ActualFPS())
}

func GetDeltaTime() float32 {
	return ETEngine.Game.DeltaTime
}
