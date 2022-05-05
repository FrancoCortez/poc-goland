package dependecy

import "cenco-pim/app/router"

type InitDependencyImpl struct {
	family router.Family
}

func NewInitDependencyImpl(family router.Family) InitDependencyImpl {
	return InitDependencyImpl{
		family: family,
	}
}

func (i *InitDependencyImpl) GetFamilyRoute() *router.Family {
	return &(i.family)
}
