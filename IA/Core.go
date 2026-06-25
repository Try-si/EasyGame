package IA

import (
	EGMath "github.com/Try-si/EasyGame/Math"
	"github.com/Try-si/IAM/Core"
)

var (
	World map[string]*IAWorld
)

func Init(addings func()) {
	World = make(map[string]*IAWorld)
	addings()
}

func Update(deltaTime float32, worldName string) error {
	if World[worldName] == nil {
		return nil
	}
	w := World[worldName]

	for _, e := range w.entity {
		//fmt.Println("Updating entity:", e)
		if w.bt != nil {
			if _, exists := w.bt.Roots[e]; exists {
				w.bt.Execute(e)
			} else {
				//fmt.Println("No behaviour tree for entity:", e)
			}
		} else {
			//fmt.Println("No behaviour tree")
		}
		if w.brain != nil {
			if _, exists := w.brain.Entity[e]; exists {
				w.brain.UpdateEntity(e, EGMath.Random)
			} else {
				//fmt.Println("No brain for entity:", e)
			}
		} else {
			//fmt.Println("No brain")
		}
		if w.fsm != nil {
			if _, exists := w.fsm.Entity[e]; exists {
				w.fsm.UpdateEntity(e)
			} else {
				//fmt.Println("No fsm for entity:", e)
			}
		} else {
			//fmt.Println("No fsm")
		}
	}
	return nil
}

type IAWorld struct {
	entity []string
	brain  *Core.Brain
	bt     *Core.BeavioursTree
	fsm    *Core.FSM
}

func (w *IAWorld) AddEntityBT(element string, node *Core.BehaviourNode) {
	w.bt.AddRoot(element, node)
	w.entity = append(w.entity, element)
}

func (w *IAWorld) AddEntityBrain(element string, brain *Core.Brain) {
	for k, v := range brain.Entity {
		w.brain.Entity[k] = v
	}
	for k, v := range brain.States {
		w.brain.States[k] = v
	}
	for k, v := range brain.Transition {
		w.brain.Transition[k] = append(w.brain.Transition[k], v...)
	}
	w.entity = append(w.entity, element)
}

func (w *IAWorld) AddEntityFSM(element string, fsm *Core.FSM) {
	for k, v := range fsm.Entity {
		w.fsm.Entity[k] = v
	}
	for k, v := range fsm.States {
		w.fsm.States[k] = v
	}
	for k, v := range fsm.Transition {
		w.fsm.Transition[k] = append(w.fsm.Transition[k], v...)
	}
	w.entity = append(w.entity, element)
}

func INITWORLD() *IAWorld {
	return &IAWorld{
		entity: make([]string, 0),
		brain: &Core.Brain{
			Entity:     make(map[string]Core.BrainEntity),
			States:     make(map[string]Core.State),
			Transition: make(map[string][]Core.BrainTransition),
		},
		bt: Core.InitBeavioursTree(),
		fsm: &Core.FSM{
			Entity:     make(map[string]Core.FSMEntity),
			States:     make(map[string]Core.State),
			Transition: make(map[string][]Core.FSMTransition),
		},
	}
}
