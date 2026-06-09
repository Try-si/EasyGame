package Brain

import IAMobs "github.com/Try-si/IAM"

func AddTransition(from string, transition BrainTransition) {
	IAMobs.GetBrain().AddTransition(from, transition.toCore())
}

func AddState(name string, state State) {
	IAMobs.GetBrain().AddState(name, state.toCore())
}

func AddEntity(name string, entity BrainEntity) {
	IAMobs.GetBrain().AddEntity(name, entity.toCore())
}

func GetEntity(name string) BrainEntity {
	return toEGBrainEntity(IAMobs.GetBrain().GetEntity(name))
}

func GetState(name string) State {
	return toEGState(IAMobs.GetBrain().GetState(name))
}

func GetTransitions(from string) []BrainTransition {
	transitions := IAMobs.GetBrain().GetTransitions(from)
	result := make([]BrainTransition, len(transitions))
	for i, transition := range transitions {
		result[i] = toEGTransition(transition)
	}
	return result
}

func GetTransition(from string, index int) BrainTransition {
	return toEGTransition(IAMobs.GetBrain().GetTransition(from, index))
}

func SetState(state string, stateData State) {
	IAMobs.GetBrain().SetState(state, stateData.toCore())
}

func SetTransition(from string, index int, transition BrainTransition) {
	IAMobs.GetBrain().SetTransition(from, index, transition.toCore())
}

func SetEntity(name string, entity BrainEntity) {
	IAMobs.GetBrain().SetEntity(name, entity.toCore())
}

func DeleteTransition(from string, index int) {
	IAMobs.GetBrain().DeleteTransition(from, index)
}

func DeleteTransitions(from string) {
	IAMobs.GetBrain().DeleteTransitions(from)
}

func DeleteState(name string) {
	IAMobs.GetBrain().DeleteState(name)
}

func DeleteEntity(name string) {
	IAMobs.GetBrain().DeleteEntity(name)
}
