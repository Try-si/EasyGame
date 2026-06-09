package BeavioursTree

import "github.com/Try-si/IAM/Core"

type BehaviourNode struct {
	Condition func(WorldState) (bool, any)
	TrueNode  *BehaviourNode
	FalseNode *BehaviourNode
	Action    func(any)
}

func (bn BehaviourNode) toCore() *Core.BehaviourNode {
	return &Core.BehaviourNode{
		Condition: conditionToCore(bn.Condition),
		TrueNode:  bn.TrueNode.toCore(),
		FalseNode: bn.FalseNode.toCore(),
		Action:    bn.Action,
	}
}

func conditionToCore(condition func(WorldState) (bool, any)) func(Core.WorldState) (bool, any) {
	return func(ws Core.WorldState) (bool, any) {
		return condition(worldStateFromCore(ws))
	}
}

func behaviourNodeFromCore(core *Core.BehaviourNode) *BehaviourNode {
	return &BehaviourNode{
		Condition: conditionFromCore(core.Condition),
		TrueNode:  behaviourNodeFromCore(core.TrueNode),
		FalseNode: behaviourNodeFromCore(core.FalseNode),
		Action:    core.Action,
	}
}

func conditionFromCore(condition func(Core.WorldState) (bool, any)) func(WorldState) (bool, any) {
	return func(ws WorldState) (bool, any) {
		return condition(ws.toCore())
	}
}

type WorldState struct {
}

func (ws WorldState) toCore() Core.WorldState {
	return Core.WorldState{}
}

func worldStateFromCore(core Core.WorldState) WorldState {
	return WorldState{}
}
