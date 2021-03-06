// Code generated by counterfeiter. DO NOT EDIT.
package v3actionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"github.com/cloudfoundry/sonde-go/events"
)

type FakeNOAAClient struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	RecentLogsStub        func(appGuid string, authToken string) ([]*events.LogMessage, error)
	recentLogsMutex       sync.RWMutex
	recentLogsArgsForCall []struct {
		appGuid   string
		authToken string
	}
	recentLogsReturns struct {
		result1 []*events.LogMessage
		result2 error
	}
	recentLogsReturnsOnCall map[int]struct {
		result1 []*events.LogMessage
		result2 error
	}
	SetOnConnectCallbackStub        func(cb func())
	setOnConnectCallbackMutex       sync.RWMutex
	setOnConnectCallbackArgsForCall []struct {
		cb func()
	}
	TailingLogsStub        func(appGuid, authToken string) (<-chan *events.LogMessage, <-chan error)
	tailingLogsMutex       sync.RWMutex
	tailingLogsArgsForCall []struct {
		appGuid   string
		authToken string
	}
	tailingLogsReturns struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}
	tailingLogsReturnsOnCall map[int]struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNOAAClient) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *FakeNOAAClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeNOAAClient) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNOAAClient) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNOAAClient) RecentLogs(appGuid string, authToken string) ([]*events.LogMessage, error) {
	fake.recentLogsMutex.Lock()
	ret, specificReturn := fake.recentLogsReturnsOnCall[len(fake.recentLogsArgsForCall)]
	fake.recentLogsArgsForCall = append(fake.recentLogsArgsForCall, struct {
		appGuid   string
		authToken string
	}{appGuid, authToken})
	fake.recordInvocation("RecentLogs", []interface{}{appGuid, authToken})
	fake.recentLogsMutex.Unlock()
	if fake.RecentLogsStub != nil {
		return fake.RecentLogsStub(appGuid, authToken)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.recentLogsReturns.result1, fake.recentLogsReturns.result2
}

func (fake *FakeNOAAClient) RecentLogsCallCount() int {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	return len(fake.recentLogsArgsForCall)
}

func (fake *FakeNOAAClient) RecentLogsArgsForCall(i int) (string, string) {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	return fake.recentLogsArgsForCall[i].appGuid, fake.recentLogsArgsForCall[i].authToken
}

func (fake *FakeNOAAClient) RecentLogsReturns(result1 []*events.LogMessage, result2 error) {
	fake.RecentLogsStub = nil
	fake.recentLogsReturns = struct {
		result1 []*events.LogMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeNOAAClient) RecentLogsReturnsOnCall(i int, result1 []*events.LogMessage, result2 error) {
	fake.RecentLogsStub = nil
	if fake.recentLogsReturnsOnCall == nil {
		fake.recentLogsReturnsOnCall = make(map[int]struct {
			result1 []*events.LogMessage
			result2 error
		})
	}
	fake.recentLogsReturnsOnCall[i] = struct {
		result1 []*events.LogMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeNOAAClient) SetOnConnectCallback(cb func()) {
	fake.setOnConnectCallbackMutex.Lock()
	fake.setOnConnectCallbackArgsForCall = append(fake.setOnConnectCallbackArgsForCall, struct {
		cb func()
	}{cb})
	fake.recordInvocation("SetOnConnectCallback", []interface{}{cb})
	fake.setOnConnectCallbackMutex.Unlock()
	if fake.SetOnConnectCallbackStub != nil {
		fake.SetOnConnectCallbackStub(cb)
	}
}

func (fake *FakeNOAAClient) SetOnConnectCallbackCallCount() int {
	fake.setOnConnectCallbackMutex.RLock()
	defer fake.setOnConnectCallbackMutex.RUnlock()
	return len(fake.setOnConnectCallbackArgsForCall)
}

func (fake *FakeNOAAClient) SetOnConnectCallbackArgsForCall(i int) func() {
	fake.setOnConnectCallbackMutex.RLock()
	defer fake.setOnConnectCallbackMutex.RUnlock()
	return fake.setOnConnectCallbackArgsForCall[i].cb
}

func (fake *FakeNOAAClient) TailingLogs(appGuid string, authToken string) (<-chan *events.LogMessage, <-chan error) {
	fake.tailingLogsMutex.Lock()
	ret, specificReturn := fake.tailingLogsReturnsOnCall[len(fake.tailingLogsArgsForCall)]
	fake.tailingLogsArgsForCall = append(fake.tailingLogsArgsForCall, struct {
		appGuid   string
		authToken string
	}{appGuid, authToken})
	fake.recordInvocation("TailingLogs", []interface{}{appGuid, authToken})
	fake.tailingLogsMutex.Unlock()
	if fake.TailingLogsStub != nil {
		return fake.TailingLogsStub(appGuid, authToken)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.tailingLogsReturns.result1, fake.tailingLogsReturns.result2
}

func (fake *FakeNOAAClient) TailingLogsCallCount() int {
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	return len(fake.tailingLogsArgsForCall)
}

func (fake *FakeNOAAClient) TailingLogsArgsForCall(i int) (string, string) {
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	return fake.tailingLogsArgsForCall[i].appGuid, fake.tailingLogsArgsForCall[i].authToken
}

func (fake *FakeNOAAClient) TailingLogsReturns(result1 <-chan *events.LogMessage, result2 <-chan error) {
	fake.TailingLogsStub = nil
	fake.tailingLogsReturns = struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeNOAAClient) TailingLogsReturnsOnCall(i int, result1 <-chan *events.LogMessage, result2 <-chan error) {
	fake.TailingLogsStub = nil
	if fake.tailingLogsReturnsOnCall == nil {
		fake.tailingLogsReturnsOnCall = make(map[int]struct {
			result1 <-chan *events.LogMessage
			result2 <-chan error
		})
	}
	fake.tailingLogsReturnsOnCall[i] = struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeNOAAClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	fake.setOnConnectCallbackMutex.RLock()
	defer fake.setOnConnectCallbackMutex.RUnlock()
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNOAAClient) recordInvocation(key string, args []interface{}) {
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

var _ v3action.NOAAClient = new(FakeNOAAClient)
