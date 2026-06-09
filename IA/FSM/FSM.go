package FSM

import IAMobs "github.com/Try-si/IAM"

func AddTransition(from string, transition FSMTransition) {
	IAMobs.GetFSM().AddTransition(from, transition.toCore())
}

func AddState(name string, state State) {
	IAMobs.GetFSM().AddState(name, state.toCore())
}

func AddEntity(name string, entity FSMEntity) {
	IAMobs.GetFSM().AddEntity(name, entity.toCore())
}

func GetEntity(name string) FSMEntity {
	return toEGEntity(IAMobs.GetFSM().GetEntity(name))
}

func GetState(name string) State {
	return toEGState(IAMobs.GetFSM().GetState(name))
}

func GetTransitions(from string) []FSMTransition {
	return toEGTransitions(IAMobs.GetFSM().GetTransitions(from))
}

func GetTransition(from string, index int) FSMTransition {
	return toEGTransition(IAMobs.GetFSM().GetTransition(from, index))
}

func SetState(state string, stateData State) {
	IAMobs.GetFSM().SetState(state, stateData.toCore())
}

func SetTransition(from string, index int, transition FSMTransition) {
	IAMobs.GetFSM().SetTransition(from, index, transition.toCore())
}

func SetEntity(name string, entity FSMEntity) {
	IAMobs.GetFSM().SetEntity(name, entity.toCore())
}

func DeleteTransition(from string, index int) {
	IAMobs.GetFSM().DeleteTransition(from, index)
}

func DeleteState(name string) {
	IAMobs.GetFSM().DeleteState(name)
}

func DeleteEntity(name string) {
	IAMobs.GetFSM().DeleteEntity(name)
}
