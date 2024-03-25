package ISP

type EatAnimalAction interface {
	eat()
}

type FlyAnimalAction interface {
	fly()
}

type SwimAnimalAction interface {
	swim()
}

/*
接口隔离原则
使用多个转移接口，而不是使用但一的总接口
*/
