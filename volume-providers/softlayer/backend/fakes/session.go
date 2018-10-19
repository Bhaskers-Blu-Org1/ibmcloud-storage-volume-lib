// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.ibm.com/alchemy-containers/armada-cluster/iaas/softlayer/backend"
)

type Session struct {
	GetAccountServiceStub        func() backend.AccountService
	getAccountServiceMutex       sync.RWMutex
	getAccountServiceArgsForCall []struct{}
	getAccountServiceReturns     struct {
		result1 backend.AccountService
	}
	getAccountServiceReturnsOnCall map[int]struct {
		result1 backend.AccountService
	}
	GetBillingItemServiceStub        func() backend.BillingItemService
	getBillingItemServiceMutex       sync.RWMutex
	getBillingItemServiceArgsForCall []struct{}
	getBillingItemServiceReturns     struct {
		result1 backend.BillingItemService
	}
	getBillingItemServiceReturnsOnCall map[int]struct {
		result1 backend.BillingItemService
	}
	GetBillingOrderServiceStub        func() backend.BillingOrderService
	getBillingOrderServiceMutex       sync.RWMutex
	getBillingOrderServiceArgsForCall []struct{}
	getBillingOrderServiceReturns     struct {
		result1 backend.BillingOrderService
	}
	getBillingOrderServiceReturnsOnCall map[int]struct {
		result1 backend.BillingOrderService
	}
	GetConfigurationStorageGroupArrayTypeServiceStub        func() backend.ConfigurationStorageGroupArrayTypeService
	getConfigurationStorageGroupArrayTypeServiceMutex       sync.RWMutex
	getConfigurationStorageGroupArrayTypeServiceArgsForCall []struct{}
	getConfigurationStorageGroupArrayTypeServiceReturns     struct {
		result1 backend.ConfigurationStorageGroupArrayTypeService
	}
	getConfigurationStorageGroupArrayTypeServiceReturnsOnCall map[int]struct {
		result1 backend.ConfigurationStorageGroupArrayTypeService
	}
	GetHardwareComponentPartitionOperatingSystemServiceStub        func() backend.HardwareComponentPartitionOperatingSystemService
	getHardwareComponentPartitionOperatingSystemServiceMutex       sync.RWMutex
	getHardwareComponentPartitionOperatingSystemServiceArgsForCall []struct{}
	getHardwareComponentPartitionOperatingSystemServiceReturns     struct {
		result1 backend.HardwareComponentPartitionOperatingSystemService
	}
	getHardwareComponentPartitionOperatingSystemServiceReturnsOnCall map[int]struct {
		result1 backend.HardwareComponentPartitionOperatingSystemService
	}
	GetHardwareServerServiceStub        func() backend.HardwareServerService
	getHardwareServerServiceMutex       sync.RWMutex
	getHardwareServerServiceArgsForCall []struct{}
	getHardwareServerServiceReturns     struct {
		result1 backend.HardwareServerService
	}
	getHardwareServerServiceReturnsOnCall map[int]struct {
		result1 backend.HardwareServerService
	}
	GetNetworkSubnetServiceStub        func() backend.NetworkSubnetService
	getNetworkSubnetServiceMutex       sync.RWMutex
	getNetworkSubnetServiceArgsForCall []struct{}
	getNetworkSubnetServiceReturns     struct {
		result1 backend.NetworkSubnetService
	}
	getNetworkSubnetServiceReturnsOnCall map[int]struct {
		result1 backend.NetworkSubnetService
	}
	GetNetworkVlanServiceStub        func() backend.NetworkVlanService
	getNetworkVlanServiceMutex       sync.RWMutex
	getNetworkVlanServiceArgsForCall []struct{}
	getNetworkVlanServiceReturns     struct {
		result1 backend.NetworkVlanService
	}
	getNetworkVlanServiceReturnsOnCall map[int]struct {
		result1 backend.NetworkVlanService
	}
	GetProductOrderServiceStub        func() backend.ProductOrderService
	getProductOrderServiceMutex       sync.RWMutex
	getProductOrderServiceArgsForCall []struct{}
	getProductOrderServiceReturns     struct {
		result1 backend.ProductOrderService
	}
	getProductOrderServiceReturnsOnCall map[int]struct {
		result1 backend.ProductOrderService
	}
	GetProductPackageServiceStub        func() backend.ProductPackageService
	getProductPackageServiceMutex       sync.RWMutex
	getProductPackageServiceArgsForCall []struct{}
	getProductPackageServiceReturns     struct {
		result1 backend.ProductPackageService
	}
	getProductPackageServiceReturnsOnCall map[int]struct {
		result1 backend.ProductPackageService
	}
	GetVirtualGuestServiceStub        func() backend.VirtualGuestService
	getVirtualGuestServiceMutex       sync.RWMutex
	getVirtualGuestServiceArgsForCall []struct{}
	getVirtualGuestServiceReturns     struct {
		result1 backend.VirtualGuestService
	}
	getVirtualGuestServiceReturnsOnCall map[int]struct {
		result1 backend.VirtualGuestService
	}
	GetVirtualGuestBlockDeviceTemplateGroupServiceStub        func() backend.VirtualGuestBlockDeviceTemplateGroupService
	getVirtualGuestBlockDeviceTemplateGroupServiceMutex       sync.RWMutex
	getVirtualGuestBlockDeviceTemplateGroupServiceArgsForCall []struct{}
	getVirtualGuestBlockDeviceTemplateGroupServiceReturns     struct {
		result1 backend.VirtualGuestBlockDeviceTemplateGroupService
	}
	getVirtualGuestBlockDeviceTemplateGroupServiceReturnsOnCall map[int]struct {
		result1 backend.VirtualGuestBlockDeviceTemplateGroupService
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Session) GetAccountService() backend.AccountService {
	fake.getAccountServiceMutex.Lock()
	ret, specificReturn := fake.getAccountServiceReturnsOnCall[len(fake.getAccountServiceArgsForCall)]
	fake.getAccountServiceArgsForCall = append(fake.getAccountServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetAccountService", []interface{}{})
	fake.getAccountServiceMutex.Unlock()
	if fake.GetAccountServiceStub != nil {
		return fake.GetAccountServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getAccountServiceReturns.result1
}

func (fake *Session) GetAccountServiceCallCount() int {
	fake.getAccountServiceMutex.RLock()
	defer fake.getAccountServiceMutex.RUnlock()
	return len(fake.getAccountServiceArgsForCall)
}

func (fake *Session) GetAccountServiceReturns(result1 backend.AccountService) {
	fake.GetAccountServiceStub = nil
	fake.getAccountServiceReturns = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *Session) GetAccountServiceReturnsOnCall(i int, result1 backend.AccountService) {
	fake.GetAccountServiceStub = nil
	if fake.getAccountServiceReturnsOnCall == nil {
		fake.getAccountServiceReturnsOnCall = make(map[int]struct {
			result1 backend.AccountService
		})
	}
	fake.getAccountServiceReturnsOnCall[i] = struct {
		result1 backend.AccountService
	}{result1}
}

func (fake *Session) GetBillingItemService() backend.BillingItemService {
	fake.getBillingItemServiceMutex.Lock()
	ret, specificReturn := fake.getBillingItemServiceReturnsOnCall[len(fake.getBillingItemServiceArgsForCall)]
	fake.getBillingItemServiceArgsForCall = append(fake.getBillingItemServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetBillingItemService", []interface{}{})
	fake.getBillingItemServiceMutex.Unlock()
	if fake.GetBillingItemServiceStub != nil {
		return fake.GetBillingItemServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getBillingItemServiceReturns.result1
}

func (fake *Session) GetBillingItemServiceCallCount() int {
	fake.getBillingItemServiceMutex.RLock()
	defer fake.getBillingItemServiceMutex.RUnlock()
	return len(fake.getBillingItemServiceArgsForCall)
}

func (fake *Session) GetBillingItemServiceReturns(result1 backend.BillingItemService) {
	fake.GetBillingItemServiceStub = nil
	fake.getBillingItemServiceReturns = struct {
		result1 backend.BillingItemService
	}{result1}
}

func (fake *Session) GetBillingItemServiceReturnsOnCall(i int, result1 backend.BillingItemService) {
	fake.GetBillingItemServiceStub = nil
	if fake.getBillingItemServiceReturnsOnCall == nil {
		fake.getBillingItemServiceReturnsOnCall = make(map[int]struct {
			result1 backend.BillingItemService
		})
	}
	fake.getBillingItemServiceReturnsOnCall[i] = struct {
		result1 backend.BillingItemService
	}{result1}
}

func (fake *Session) GetBillingOrderService() backend.BillingOrderService {
	fake.getBillingOrderServiceMutex.Lock()
	ret, specificReturn := fake.getBillingOrderServiceReturnsOnCall[len(fake.getBillingOrderServiceArgsForCall)]
	fake.getBillingOrderServiceArgsForCall = append(fake.getBillingOrderServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetBillingOrderService", []interface{}{})
	fake.getBillingOrderServiceMutex.Unlock()
	if fake.GetBillingOrderServiceStub != nil {
		return fake.GetBillingOrderServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getBillingOrderServiceReturns.result1
}

func (fake *Session) GetBillingOrderServiceCallCount() int {
	fake.getBillingOrderServiceMutex.RLock()
	defer fake.getBillingOrderServiceMutex.RUnlock()
	return len(fake.getBillingOrderServiceArgsForCall)
}

func (fake *Session) GetBillingOrderServiceReturns(result1 backend.BillingOrderService) {
	fake.GetBillingOrderServiceStub = nil
	fake.getBillingOrderServiceReturns = struct {
		result1 backend.BillingOrderService
	}{result1}
}

func (fake *Session) GetBillingOrderServiceReturnsOnCall(i int, result1 backend.BillingOrderService) {
	fake.GetBillingOrderServiceStub = nil
	if fake.getBillingOrderServiceReturnsOnCall == nil {
		fake.getBillingOrderServiceReturnsOnCall = make(map[int]struct {
			result1 backend.BillingOrderService
		})
	}
	fake.getBillingOrderServiceReturnsOnCall[i] = struct {
		result1 backend.BillingOrderService
	}{result1}
}

func (fake *Session) GetConfigurationStorageGroupArrayTypeService() backend.ConfigurationStorageGroupArrayTypeService {
	fake.getConfigurationStorageGroupArrayTypeServiceMutex.Lock()
	ret, specificReturn := fake.getConfigurationStorageGroupArrayTypeServiceReturnsOnCall[len(fake.getConfigurationStorageGroupArrayTypeServiceArgsForCall)]
	fake.getConfigurationStorageGroupArrayTypeServiceArgsForCall = append(fake.getConfigurationStorageGroupArrayTypeServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetConfigurationStorageGroupArrayTypeService", []interface{}{})
	fake.getConfigurationStorageGroupArrayTypeServiceMutex.Unlock()
	if fake.GetConfigurationStorageGroupArrayTypeServiceStub != nil {
		return fake.GetConfigurationStorageGroupArrayTypeServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getConfigurationStorageGroupArrayTypeServiceReturns.result1
}

func (fake *Session) GetConfigurationStorageGroupArrayTypeServiceCallCount() int {
	fake.getConfigurationStorageGroupArrayTypeServiceMutex.RLock()
	defer fake.getConfigurationStorageGroupArrayTypeServiceMutex.RUnlock()
	return len(fake.getConfigurationStorageGroupArrayTypeServiceArgsForCall)
}

func (fake *Session) GetConfigurationStorageGroupArrayTypeServiceReturns(result1 backend.ConfigurationStorageGroupArrayTypeService) {
	fake.GetConfigurationStorageGroupArrayTypeServiceStub = nil
	fake.getConfigurationStorageGroupArrayTypeServiceReturns = struct {
		result1 backend.ConfigurationStorageGroupArrayTypeService
	}{result1}
}

func (fake *Session) GetConfigurationStorageGroupArrayTypeServiceReturnsOnCall(i int, result1 backend.ConfigurationStorageGroupArrayTypeService) {
	fake.GetConfigurationStorageGroupArrayTypeServiceStub = nil
	if fake.getConfigurationStorageGroupArrayTypeServiceReturnsOnCall == nil {
		fake.getConfigurationStorageGroupArrayTypeServiceReturnsOnCall = make(map[int]struct {
			result1 backend.ConfigurationStorageGroupArrayTypeService
		})
	}
	fake.getConfigurationStorageGroupArrayTypeServiceReturnsOnCall[i] = struct {
		result1 backend.ConfigurationStorageGroupArrayTypeService
	}{result1}
}

func (fake *Session) GetHardwareComponentPartitionOperatingSystemService() backend.HardwareComponentPartitionOperatingSystemService {
	fake.getHardwareComponentPartitionOperatingSystemServiceMutex.Lock()
	ret, specificReturn := fake.getHardwareComponentPartitionOperatingSystemServiceReturnsOnCall[len(fake.getHardwareComponentPartitionOperatingSystemServiceArgsForCall)]
	fake.getHardwareComponentPartitionOperatingSystemServiceArgsForCall = append(fake.getHardwareComponentPartitionOperatingSystemServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetHardwareComponentPartitionOperatingSystemService", []interface{}{})
	fake.getHardwareComponentPartitionOperatingSystemServiceMutex.Unlock()
	if fake.GetHardwareComponentPartitionOperatingSystemServiceStub != nil {
		return fake.GetHardwareComponentPartitionOperatingSystemServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getHardwareComponentPartitionOperatingSystemServiceReturns.result1
}

func (fake *Session) GetHardwareComponentPartitionOperatingSystemServiceCallCount() int {
	fake.getHardwareComponentPartitionOperatingSystemServiceMutex.RLock()
	defer fake.getHardwareComponentPartitionOperatingSystemServiceMutex.RUnlock()
	return len(fake.getHardwareComponentPartitionOperatingSystemServiceArgsForCall)
}

func (fake *Session) GetHardwareComponentPartitionOperatingSystemServiceReturns(result1 backend.HardwareComponentPartitionOperatingSystemService) {
	fake.GetHardwareComponentPartitionOperatingSystemServiceStub = nil
	fake.getHardwareComponentPartitionOperatingSystemServiceReturns = struct {
		result1 backend.HardwareComponentPartitionOperatingSystemService
	}{result1}
}

func (fake *Session) GetHardwareComponentPartitionOperatingSystemServiceReturnsOnCall(i int, result1 backend.HardwareComponentPartitionOperatingSystemService) {
	fake.GetHardwareComponentPartitionOperatingSystemServiceStub = nil
	if fake.getHardwareComponentPartitionOperatingSystemServiceReturnsOnCall == nil {
		fake.getHardwareComponentPartitionOperatingSystemServiceReturnsOnCall = make(map[int]struct {
			result1 backend.HardwareComponentPartitionOperatingSystemService
		})
	}
	fake.getHardwareComponentPartitionOperatingSystemServiceReturnsOnCall[i] = struct {
		result1 backend.HardwareComponentPartitionOperatingSystemService
	}{result1}
}

func (fake *Session) GetHardwareServerService() backend.HardwareServerService {
	fake.getHardwareServerServiceMutex.Lock()
	ret, specificReturn := fake.getHardwareServerServiceReturnsOnCall[len(fake.getHardwareServerServiceArgsForCall)]
	fake.getHardwareServerServiceArgsForCall = append(fake.getHardwareServerServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetHardwareServerService", []interface{}{})
	fake.getHardwareServerServiceMutex.Unlock()
	if fake.GetHardwareServerServiceStub != nil {
		return fake.GetHardwareServerServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getHardwareServerServiceReturns.result1
}

func (fake *Session) GetHardwareServerServiceCallCount() int {
	fake.getHardwareServerServiceMutex.RLock()
	defer fake.getHardwareServerServiceMutex.RUnlock()
	return len(fake.getHardwareServerServiceArgsForCall)
}

func (fake *Session) GetHardwareServerServiceReturns(result1 backend.HardwareServerService) {
	fake.GetHardwareServerServiceStub = nil
	fake.getHardwareServerServiceReturns = struct {
		result1 backend.HardwareServerService
	}{result1}
}

func (fake *Session) GetHardwareServerServiceReturnsOnCall(i int, result1 backend.HardwareServerService) {
	fake.GetHardwareServerServiceStub = nil
	if fake.getHardwareServerServiceReturnsOnCall == nil {
		fake.getHardwareServerServiceReturnsOnCall = make(map[int]struct {
			result1 backend.HardwareServerService
		})
	}
	fake.getHardwareServerServiceReturnsOnCall[i] = struct {
		result1 backend.HardwareServerService
	}{result1}
}

func (fake *Session) GetNetworkSubnetService() backend.NetworkSubnetService {
	fake.getNetworkSubnetServiceMutex.Lock()
	ret, specificReturn := fake.getNetworkSubnetServiceReturnsOnCall[len(fake.getNetworkSubnetServiceArgsForCall)]
	fake.getNetworkSubnetServiceArgsForCall = append(fake.getNetworkSubnetServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetNetworkSubnetService", []interface{}{})
	fake.getNetworkSubnetServiceMutex.Unlock()
	if fake.GetNetworkSubnetServiceStub != nil {
		return fake.GetNetworkSubnetServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getNetworkSubnetServiceReturns.result1
}

func (fake *Session) GetNetworkSubnetServiceCallCount() int {
	fake.getNetworkSubnetServiceMutex.RLock()
	defer fake.getNetworkSubnetServiceMutex.RUnlock()
	return len(fake.getNetworkSubnetServiceArgsForCall)
}

func (fake *Session) GetNetworkSubnetServiceReturns(result1 backend.NetworkSubnetService) {
	fake.GetNetworkSubnetServiceStub = nil
	fake.getNetworkSubnetServiceReturns = struct {
		result1 backend.NetworkSubnetService
	}{result1}
}

func (fake *Session) GetNetworkSubnetServiceReturnsOnCall(i int, result1 backend.NetworkSubnetService) {
	fake.GetNetworkSubnetServiceStub = nil
	if fake.getNetworkSubnetServiceReturnsOnCall == nil {
		fake.getNetworkSubnetServiceReturnsOnCall = make(map[int]struct {
			result1 backend.NetworkSubnetService
		})
	}
	fake.getNetworkSubnetServiceReturnsOnCall[i] = struct {
		result1 backend.NetworkSubnetService
	}{result1}
}

func (fake *Session) GetNetworkVlanService() backend.NetworkVlanService {
	fake.getNetworkVlanServiceMutex.Lock()
	ret, specificReturn := fake.getNetworkVlanServiceReturnsOnCall[len(fake.getNetworkVlanServiceArgsForCall)]
	fake.getNetworkVlanServiceArgsForCall = append(fake.getNetworkVlanServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetNetworkVlanService", []interface{}{})
	fake.getNetworkVlanServiceMutex.Unlock()
	if fake.GetNetworkVlanServiceStub != nil {
		return fake.GetNetworkVlanServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getNetworkVlanServiceReturns.result1
}

func (fake *Session) GetNetworkVlanServiceCallCount() int {
	fake.getNetworkVlanServiceMutex.RLock()
	defer fake.getNetworkVlanServiceMutex.RUnlock()
	return len(fake.getNetworkVlanServiceArgsForCall)
}

func (fake *Session) GetNetworkVlanServiceReturns(result1 backend.NetworkVlanService) {
	fake.GetNetworkVlanServiceStub = nil
	fake.getNetworkVlanServiceReturns = struct {
		result1 backend.NetworkVlanService
	}{result1}
}

func (fake *Session) GetNetworkVlanServiceReturnsOnCall(i int, result1 backend.NetworkVlanService) {
	fake.GetNetworkVlanServiceStub = nil
	if fake.getNetworkVlanServiceReturnsOnCall == nil {
		fake.getNetworkVlanServiceReturnsOnCall = make(map[int]struct {
			result1 backend.NetworkVlanService
		})
	}
	fake.getNetworkVlanServiceReturnsOnCall[i] = struct {
		result1 backend.NetworkVlanService
	}{result1}
}

func (fake *Session) GetProductOrderService() backend.ProductOrderService {
	fake.getProductOrderServiceMutex.Lock()
	ret, specificReturn := fake.getProductOrderServiceReturnsOnCall[len(fake.getProductOrderServiceArgsForCall)]
	fake.getProductOrderServiceArgsForCall = append(fake.getProductOrderServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetProductOrderService", []interface{}{})
	fake.getProductOrderServiceMutex.Unlock()
	if fake.GetProductOrderServiceStub != nil {
		return fake.GetProductOrderServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getProductOrderServiceReturns.result1
}

func (fake *Session) GetProductOrderServiceCallCount() int {
	fake.getProductOrderServiceMutex.RLock()
	defer fake.getProductOrderServiceMutex.RUnlock()
	return len(fake.getProductOrderServiceArgsForCall)
}

func (fake *Session) GetProductOrderServiceReturns(result1 backend.ProductOrderService) {
	fake.GetProductOrderServiceStub = nil
	fake.getProductOrderServiceReturns = struct {
		result1 backend.ProductOrderService
	}{result1}
}

func (fake *Session) GetProductOrderServiceReturnsOnCall(i int, result1 backend.ProductOrderService) {
	fake.GetProductOrderServiceStub = nil
	if fake.getProductOrderServiceReturnsOnCall == nil {
		fake.getProductOrderServiceReturnsOnCall = make(map[int]struct {
			result1 backend.ProductOrderService
		})
	}
	fake.getProductOrderServiceReturnsOnCall[i] = struct {
		result1 backend.ProductOrderService
	}{result1}
}

func (fake *Session) GetProductPackageService() backend.ProductPackageService {
	fake.getProductPackageServiceMutex.Lock()
	ret, specificReturn := fake.getProductPackageServiceReturnsOnCall[len(fake.getProductPackageServiceArgsForCall)]
	fake.getProductPackageServiceArgsForCall = append(fake.getProductPackageServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetProductPackageService", []interface{}{})
	fake.getProductPackageServiceMutex.Unlock()
	if fake.GetProductPackageServiceStub != nil {
		return fake.GetProductPackageServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getProductPackageServiceReturns.result1
}

func (fake *Session) GetProductPackageServiceCallCount() int {
	fake.getProductPackageServiceMutex.RLock()
	defer fake.getProductPackageServiceMutex.RUnlock()
	return len(fake.getProductPackageServiceArgsForCall)
}

func (fake *Session) GetProductPackageServiceReturns(result1 backend.ProductPackageService) {
	fake.GetProductPackageServiceStub = nil
	fake.getProductPackageServiceReturns = struct {
		result1 backend.ProductPackageService
	}{result1}
}

func (fake *Session) GetProductPackageServiceReturnsOnCall(i int, result1 backend.ProductPackageService) {
	fake.GetProductPackageServiceStub = nil
	if fake.getProductPackageServiceReturnsOnCall == nil {
		fake.getProductPackageServiceReturnsOnCall = make(map[int]struct {
			result1 backend.ProductPackageService
		})
	}
	fake.getProductPackageServiceReturnsOnCall[i] = struct {
		result1 backend.ProductPackageService
	}{result1}
}

func (fake *Session) GetVirtualGuestService() backend.VirtualGuestService {
	fake.getVirtualGuestServiceMutex.Lock()
	ret, specificReturn := fake.getVirtualGuestServiceReturnsOnCall[len(fake.getVirtualGuestServiceArgsForCall)]
	fake.getVirtualGuestServiceArgsForCall = append(fake.getVirtualGuestServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetVirtualGuestService", []interface{}{})
	fake.getVirtualGuestServiceMutex.Unlock()
	if fake.GetVirtualGuestServiceStub != nil {
		return fake.GetVirtualGuestServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getVirtualGuestServiceReturns.result1
}

func (fake *Session) GetVirtualGuestServiceCallCount() int {
	fake.getVirtualGuestServiceMutex.RLock()
	defer fake.getVirtualGuestServiceMutex.RUnlock()
	return len(fake.getVirtualGuestServiceArgsForCall)
}

func (fake *Session) GetVirtualGuestServiceReturns(result1 backend.VirtualGuestService) {
	fake.GetVirtualGuestServiceStub = nil
	fake.getVirtualGuestServiceReturns = struct {
		result1 backend.VirtualGuestService
	}{result1}
}

func (fake *Session) GetVirtualGuestServiceReturnsOnCall(i int, result1 backend.VirtualGuestService) {
	fake.GetVirtualGuestServiceStub = nil
	if fake.getVirtualGuestServiceReturnsOnCall == nil {
		fake.getVirtualGuestServiceReturnsOnCall = make(map[int]struct {
			result1 backend.VirtualGuestService
		})
	}
	fake.getVirtualGuestServiceReturnsOnCall[i] = struct {
		result1 backend.VirtualGuestService
	}{result1}
}

func (fake *Session) GetVirtualGuestBlockDeviceTemplateGroupService() backend.VirtualGuestBlockDeviceTemplateGroupService {
	fake.getVirtualGuestBlockDeviceTemplateGroupServiceMutex.Lock()
	ret, specificReturn := fake.getVirtualGuestBlockDeviceTemplateGroupServiceReturnsOnCall[len(fake.getVirtualGuestBlockDeviceTemplateGroupServiceArgsForCall)]
	fake.getVirtualGuestBlockDeviceTemplateGroupServiceArgsForCall = append(fake.getVirtualGuestBlockDeviceTemplateGroupServiceArgsForCall, struct{}{})
	fake.recordInvocation("GetVirtualGuestBlockDeviceTemplateGroupService", []interface{}{})
	fake.getVirtualGuestBlockDeviceTemplateGroupServiceMutex.Unlock()
	if fake.GetVirtualGuestBlockDeviceTemplateGroupServiceStub != nil {
		return fake.GetVirtualGuestBlockDeviceTemplateGroupServiceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getVirtualGuestBlockDeviceTemplateGroupServiceReturns.result1
}

func (fake *Session) GetVirtualGuestBlockDeviceTemplateGroupServiceCallCount() int {
	fake.getVirtualGuestBlockDeviceTemplateGroupServiceMutex.RLock()
	defer fake.getVirtualGuestBlockDeviceTemplateGroupServiceMutex.RUnlock()
	return len(fake.getVirtualGuestBlockDeviceTemplateGroupServiceArgsForCall)
}

func (fake *Session) GetVirtualGuestBlockDeviceTemplateGroupServiceReturns(result1 backend.VirtualGuestBlockDeviceTemplateGroupService) {
	fake.GetVirtualGuestBlockDeviceTemplateGroupServiceStub = nil
	fake.getVirtualGuestBlockDeviceTemplateGroupServiceReturns = struct {
		result1 backend.VirtualGuestBlockDeviceTemplateGroupService
	}{result1}
}

func (fake *Session) GetVirtualGuestBlockDeviceTemplateGroupServiceReturnsOnCall(i int, result1 backend.VirtualGuestBlockDeviceTemplateGroupService) {
	fake.GetVirtualGuestBlockDeviceTemplateGroupServiceStub = nil
	if fake.getVirtualGuestBlockDeviceTemplateGroupServiceReturnsOnCall == nil {
		fake.getVirtualGuestBlockDeviceTemplateGroupServiceReturnsOnCall = make(map[int]struct {
			result1 backend.VirtualGuestBlockDeviceTemplateGroupService
		})
	}
	fake.getVirtualGuestBlockDeviceTemplateGroupServiceReturnsOnCall[i] = struct {
		result1 backend.VirtualGuestBlockDeviceTemplateGroupService
	}{result1}
}

func (fake *Session) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAccountServiceMutex.RLock()
	defer fake.getAccountServiceMutex.RUnlock()
	fake.getBillingItemServiceMutex.RLock()
	defer fake.getBillingItemServiceMutex.RUnlock()
	fake.getBillingOrderServiceMutex.RLock()
	defer fake.getBillingOrderServiceMutex.RUnlock()
	fake.getConfigurationStorageGroupArrayTypeServiceMutex.RLock()
	defer fake.getConfigurationStorageGroupArrayTypeServiceMutex.RUnlock()
	fake.getHardwareComponentPartitionOperatingSystemServiceMutex.RLock()
	defer fake.getHardwareComponentPartitionOperatingSystemServiceMutex.RUnlock()
	fake.getHardwareServerServiceMutex.RLock()
	defer fake.getHardwareServerServiceMutex.RUnlock()
	fake.getNetworkSubnetServiceMutex.RLock()
	defer fake.getNetworkSubnetServiceMutex.RUnlock()
	fake.getNetworkVlanServiceMutex.RLock()
	defer fake.getNetworkVlanServiceMutex.RUnlock()
	fake.getProductOrderServiceMutex.RLock()
	defer fake.getProductOrderServiceMutex.RUnlock()
	fake.getProductPackageServiceMutex.RLock()
	defer fake.getProductPackageServiceMutex.RUnlock()
	fake.getVirtualGuestServiceMutex.RLock()
	defer fake.getVirtualGuestServiceMutex.RUnlock()
	fake.getVirtualGuestBlockDeviceTemplateGroupServiceMutex.RLock()
	defer fake.getVirtualGuestBlockDeviceTemplateGroupServiceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Session) recordInvocation(key string, args []interface{}) {
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

var _ backend.Session = new(Session)