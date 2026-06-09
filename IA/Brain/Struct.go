package Brain

import "github.com/Try-si/IAM/Core"

type Brain struct {
	Transition map[string][]BrainTransition
	States     map[string]State
	Entity     map[string]BrainEntity
}

type BrainTransition struct {
	Weight    float32
	To        string
	Condition func(WorldState) (bool, any)
}

type BrainEntity struct {
	CurrentState string
	PréTrans     BrainTransition
}

func (e BrainEntity) toCore() Core.BrainEntity {
	return Core.BrainEntity{
		CurrentState: e.CurrentState,
		PréTrans:     e.PréTrans.toCore(),
	}
}

func toEGBrainEntity(entity Core.BrainEntity) BrainEntity {
	return BrainEntity{
		CurrentState: entity.CurrentState,
		PréTrans:     toEGTransition(entity.PréTrans),
	}
}

func (t BrainTransition) toCore() Core.BrainTransition {
	return Core.BrainTransition{
		Weight:    t.Weight,
		To:        t.To,
		Condition: toEGCondition(t.Condition),
	}
}

func toEGTransition(transition Core.BrainTransition) BrainTransition {
	return BrainTransition{
		Weight:    transition.Weight,
		To:        transition.To,
		Condition: toCoreCondition(transition.Condition),
	}
}

func toEGCondition(condition func(WorldState) (bool, any)) func(Core.WorldState) (bool, any) {
	return func(ws Core.WorldState) (bool, any) {
		return condition(toEG(ws))
	}
}

func toCoreCondition(condition func(Core.WorldState) (bool, any)) func(WorldState) (bool, any) {
	return func(ws WorldState) (bool, any) {
		return condition(ws.toCore())
	}
}

func (b Brain) toCore() Core.Brain {
	return Core.Brain{
		Transition: toCoreTransitions(b.Transition),
		States:     toCoreStates(b.States),
		Entity:     toCoreEntities(b.Entity),
	}
}

func toEGBrain(brain Core.Brain) Brain {
	return Brain{
		Transition: toEGTransitions(brain.Transition),
		States:     toEGStates(brain.States),
		Entity:     toEGEntities(brain.Entity),
	}
}

func toEGTransitions(transitions map[string][]Core.BrainTransition) map[string][]BrainTransition {
	result := make(map[string][]BrainTransition)
	for key, value := range transitions {
		result[key] = make([]BrainTransition, len(value))
		for i, transition := range value {
			result[key][i] = toEGTransition(transition)
		}
	}
	return result
}

func toCoreTransitions(transitions map[string][]BrainTransition) map[string][]Core.BrainTransition {
	result := make(map[string][]Core.BrainTransition)
	for key, value := range transitions {
		result[key] = make([]Core.BrainTransition, len(value))
		for i, transition := range value {
			result[key][i] = transition.toCore()
		}
	}
	return result
}

func toEGStates(states map[string]Core.State) map[string]State {
	result := make(map[string]State)
	for key, value := range states {
		result[key] = toEGState(value)
	}
	return result
}

func toCoreStates(states map[string]State) map[string]Core.State {
	result := make(map[string]Core.State)
	for key, value := range states {
		result[key] = value.toCore()
	}
	return result
}

func toEGEntities(entities map[string]Core.BrainEntity) map[string]BrainEntity {
	result := make(map[string]BrainEntity)
	for key, value := range entities {
		result[key] = toEGBrainEntity(value)
	}
	return result
}

func toCoreEntities(entities map[string]BrainEntity) map[string]Core.BrainEntity {
	result := make(map[string]Core.BrainEntity)
	for key, value := range entities {
		result[key] = value.toCore()
	}
	return result
}

type WorldState struct {
}

func (w WorldState) toCore() Core.WorldState {
	return Core.WorldState(w)
}

func toEG(ws Core.WorldState) WorldState {
	return WorldState(ws)
}

type State struct {
	Action func(any)
}

func (s State) toCore() Core.State {
	return Core.State{
		Action: s.Action,
	}
}

func toEGState(state Core.State) State {
	return State{
		Action: state.Action,
	}
}
