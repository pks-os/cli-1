// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v6"
)

type FakeDeleteOrphanedRoutesActor struct {
	GetOrphanedRoutesBySpaceStub        func(spaceGUID string) ([]v2action.Route, v2action.Warnings, error)
	getOrphanedRoutesBySpaceMutex       sync.RWMutex
	getOrphanedRoutesBySpaceArgsForCall []struct {
		spaceGUID string
	}
	getOrphanedRoutesBySpaceReturns struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	getOrphanedRoutesBySpaceReturnsOnCall map[int]struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	DeleteRouteStub        func(routeGUID string) (v2action.Warnings, error)
	deleteRouteMutex       sync.RWMutex
	deleteRouteArgsForCall []struct {
		routeGUID string
	}
	deleteRouteReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	deleteRouteReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteOrphanedRoutesActor) GetOrphanedRoutesBySpace(spaceGUID string) ([]v2action.Route, v2action.Warnings, error) {
	fake.getOrphanedRoutesBySpaceMutex.Lock()
	ret, specificReturn := fake.getOrphanedRoutesBySpaceReturnsOnCall[len(fake.getOrphanedRoutesBySpaceArgsForCall)]
	fake.getOrphanedRoutesBySpaceArgsForCall = append(fake.getOrphanedRoutesBySpaceArgsForCall, struct {
		spaceGUID string
	}{spaceGUID})
	fake.recordInvocation("GetOrphanedRoutesBySpace", []interface{}{spaceGUID})
	fake.getOrphanedRoutesBySpaceMutex.Unlock()
	if fake.GetOrphanedRoutesBySpaceStub != nil {
		return fake.GetOrphanedRoutesBySpaceStub(spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getOrphanedRoutesBySpaceReturns.result1, fake.getOrphanedRoutesBySpaceReturns.result2, fake.getOrphanedRoutesBySpaceReturns.result3
}

func (fake *FakeDeleteOrphanedRoutesActor) GetOrphanedRoutesBySpaceCallCount() int {
	fake.getOrphanedRoutesBySpaceMutex.RLock()
	defer fake.getOrphanedRoutesBySpaceMutex.RUnlock()
	return len(fake.getOrphanedRoutesBySpaceArgsForCall)
}

func (fake *FakeDeleteOrphanedRoutesActor) GetOrphanedRoutesBySpaceArgsForCall(i int) string {
	fake.getOrphanedRoutesBySpaceMutex.RLock()
	defer fake.getOrphanedRoutesBySpaceMutex.RUnlock()
	return fake.getOrphanedRoutesBySpaceArgsForCall[i].spaceGUID
}

func (fake *FakeDeleteOrphanedRoutesActor) GetOrphanedRoutesBySpaceReturns(result1 []v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.GetOrphanedRoutesBySpaceStub = nil
	fake.getOrphanedRoutesBySpaceReturns = struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDeleteOrphanedRoutesActor) GetOrphanedRoutesBySpaceReturnsOnCall(i int, result1 []v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.GetOrphanedRoutesBySpaceStub = nil
	if fake.getOrphanedRoutesBySpaceReturnsOnCall == nil {
		fake.getOrphanedRoutesBySpaceReturnsOnCall = make(map[int]struct {
			result1 []v2action.Route
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getOrphanedRoutesBySpaceReturnsOnCall[i] = struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteRoute(routeGUID string) (v2action.Warnings, error) {
	fake.deleteRouteMutex.Lock()
	ret, specificReturn := fake.deleteRouteReturnsOnCall[len(fake.deleteRouteArgsForCall)]
	fake.deleteRouteArgsForCall = append(fake.deleteRouteArgsForCall, struct {
		routeGUID string
	}{routeGUID})
	fake.recordInvocation("DeleteRoute", []interface{}{routeGUID})
	fake.deleteRouteMutex.Unlock()
	if fake.DeleteRouteStub != nil {
		return fake.DeleteRouteStub(routeGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteRouteReturns.result1, fake.deleteRouteReturns.result2
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteRouteCallCount() int {
	fake.deleteRouteMutex.RLock()
	defer fake.deleteRouteMutex.RUnlock()
	return len(fake.deleteRouteArgsForCall)
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteRouteArgsForCall(i int) string {
	fake.deleteRouteMutex.RLock()
	defer fake.deleteRouteMutex.RUnlock()
	return fake.deleteRouteArgsForCall[i].routeGUID
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteRouteReturns(result1 v2action.Warnings, result2 error) {
	fake.DeleteRouteStub = nil
	fake.deleteRouteReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteRouteReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.DeleteRouteStub = nil
	if fake.deleteRouteReturnsOnCall == nil {
		fake.deleteRouteReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.deleteRouteReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrphanedRoutesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrphanedRoutesBySpaceMutex.RLock()
	defer fake.getOrphanedRoutesBySpaceMutex.RUnlock()
	fake.deleteRouteMutex.RLock()
	defer fake.deleteRouteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteOrphanedRoutesActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v6.DeleteOrphanedRoutesActor = new(FakeDeleteOrphanedRoutesActor)
