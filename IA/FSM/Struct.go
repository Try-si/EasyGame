package FSM

import "github.com/Try-si/IAM/Core"

type FSMTransition struct {
	To        string
	Condition func(WorldState) (bool, any)
}

func conditionToCore(condition func(WorldState) (bool, any)) func(Core.WorldState) (bool, any) {
	return func(worldState Core.WorldState) (bool, any) {
		return condition(WorldState(worldState))
	}
}

func conditionToEG(condition func(Core.WorldState) (bool, any)) func(WorldState) (bool, any) {
	return func(worldState WorldState) (bool, any) {
		return condition(Core.WorldState(worldState))
	}
}

func (t FSMTransition) toCore() Core.FSMTransition {
	return Core.FSMTransition{
		To:        t.To,
		Condition: conditionToCore(t.Condition),
	}
}

func toEGTransition(transition Core.FSMTransition) FSMTransition {
	return FSMTransition{
		To:        transition.To,
		Condition: conditionToEG(transition.Condition),
	}
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

func toEGTransitions(transitions []Core.FSMTransition) []FSMTransition {
	result := make([]FSMTransition, len(transitions))
	for i, transition := range transitions {
		result[i] = toEGTransition(transition)
	}
	return result
}

func toEGEntity(entity Core.FSMEntity) FSMEntity {
	return FSMEntity{
		CurrentState: entity.CurrentState,
		PréTrans:     toEGTransition(entity.PréTrans),
	}
}

type FSMEntity struct {
	CurrentState string
	PréTrans     FSMTransition
}

func (e FSMEntity) toCore() Core.FSMEntity {
	return Core.FSMEntity{
		CurrentState: e.CurrentState,
		PréTrans:     e.PréTrans.toCore(),
	}
}

type WorldState struct {
}

func (w WorldState) toCore() Core.WorldState {
	return Core.WorldState(w)
}
