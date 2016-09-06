// Copyright 2016 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vsphere

import (
	"github.com/vmware/vic/lib/portlayer/event/events"

	"github.com/vmware/govmomi/vim25/types"
)

type VmEvent struct {
	*events.BaseEvent
}

func NewVMEvent(be types.BaseEvent) *VmEvent {
	var ee string
	// vm events that we care about
	switch be.(type) {
	case *types.VmPoweredOnEvent:
		ee = events.ContainerPoweredOn
	case *types.VmPoweredOffEvent:
		ee = events.ContainerPoweredOff
	case *types.VmSuspendedEvent:
		ee = events.ContainerSuspended
	case *types.VmRemovedEvent:
		ee = events.ContainerRemoved
	case *types.VmGuestShutdownEvent:
		ee = events.ContainerShutdown
	}
	e := be.GetEvent()
	return &VmEvent{
		&events.BaseEvent{
			Event:  ee,
			ID:     int(e.Key),
			Detail: e.FullFormattedMessage,
			Ref:    e.Vm.Vm.String(),
		},
	}

}

func (vme *VmEvent) Topic() string {
	if vme.Type == "" {
		vme.Type = events.NewEventType(vme)
	}
	return vme.Type.Topic()
}