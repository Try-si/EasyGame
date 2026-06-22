package IA

import (
	EGMath "github.com/Try-si/EasyGame/Math"
	IAMobs "github.com/Try-si/IAM"
)

func Init(addings func()) {
	IAMobs.InitFSM()
	IAMobs.InitBrain()
	IAMobs.InitBeavioursTree()
	addings()
}

func Update(deltaTime float32) error {
	IAMobs.UpdateBeavioursTree()
	IAMobs.UpdateBrain(EGMath.Random)
	IAMobs.UpdateFSM()
	return nil
}
