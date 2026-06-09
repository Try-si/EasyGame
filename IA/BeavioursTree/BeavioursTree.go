package BeavioursTree

import IAMobs "github.com/Try-si/IAM"

func AddRoot(name string, node *BehaviourNode) {
	IAMobs.GetBeavioursTree().AddRoot(name, node.toCore())
}

func RemoveRoot(name string) {
	IAMobs.GetBeavioursTree().RemoveRoot(name)
}

func GetRoot(name string) *BehaviourNode {
	return behaviourNodeFromCore(IAMobs.GetBeavioursTree().GetRoot(name))
}

func GetRoots() map[string]*BehaviourNode {
	cores := IAMobs.GetBeavioursTree().GetRoots()
	result := make(map[string]*BehaviourNode)
	for name, core := range cores {
		result[name] = behaviourNodeFromCore(core)
	}
	return result
}
