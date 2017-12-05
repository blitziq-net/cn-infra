// Code generated by protoc-gen-gogo.
// source: status.proto
// DO NOT EDIT!

/*
Package status is a generated protocol buffer package.

Package status provides data model for status information of the agent.

It is generated from these files:
	status.proto

It has these top-level messages:
	AgentStatus
	PluginStatus
	InterfaceStats
*/
package status

import proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type OperationalState int32

const (
	OperationalState_INIT  OperationalState = 0
	OperationalState_OK    OperationalState = 1
	OperationalState_ERROR OperationalState = 2
)

var OperationalState_name = map[int32]string{
	0: "INIT",
	1: "OK",
	2: "ERROR",
}
var OperationalState_value = map[string]int32{
	"INIT":  0,
	"OK":    1,
	"ERROR": 2,
}

func (x OperationalState) String() string {
	return proto.EnumName(OperationalState_name, int32(x))
}

type AgentStatus struct {
	BuildVersion   string           `protobuf:"bytes,1,opt,name=build_version,proto3" json:"build_version,omitempty"`
	BuildDate      string           `protobuf:"bytes,2,opt,name=build_date,proto3" json:"build_date,omitempty"`
	State          OperationalState `protobuf:"varint,3,opt,name=state,proto3,enum=status.OperationalState" json:"state,omitempty"`
	StartTime      int64            `protobuf:"varint,4,opt,name=start_time,proto3" json:"start_time,omitempty"`
	LastChange     int64            `protobuf:"varint,5,opt,name=last_change,proto3" json:"last_change,omitempty"`
	LastUpdate     int64            `protobuf:"varint,6,opt,name=last_update,proto3" json:"last_update,omitempty"`
	InterfaceStats *InterfaceStats  `protobuf:"bytes,7,opt,name=interface_stats" json:"interface_stats,omitempty"`
}

func (m *AgentStatus) Reset()         { *m = AgentStatus{} }
func (m *AgentStatus) String() string { return proto.CompactTextString(m) }
func (*AgentStatus) ProtoMessage()    {}

func (m *AgentStatus) GetInterfaceStats() *InterfaceStats {
	if m != nil {
		return m.InterfaceStats
	}
	return nil
}

type PluginStatus struct {
	State      OperationalState `protobuf:"varint,1,opt,name=state,proto3,enum=status.OperationalState" json:"state,omitempty"`
	LastChange int64            `protobuf:"varint,4,opt,name=last_change,proto3" json:"last_change,omitempty"`
	LastUpdate int64            `protobuf:"varint,5,opt,name=last_update,proto3" json:"last_update,omitempty"`
	Error      string           `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *PluginStatus) Reset()         { *m = PluginStatus{} }
func (m *PluginStatus) String() string { return proto.CompactTextString(m) }
func (*PluginStatus) ProtoMessage()    {}

type InterfaceStats struct {
	Interface []*InterfaceStats_Interface `protobuf:"bytes,1,rep,name=interface" json:"interface,omitempty"`
}

func (m *InterfaceStats) Reset()         { *m = InterfaceStats{} }
func (m *InterfaceStats) String() string { return proto.CompactTextString(m) }
func (*InterfaceStats) ProtoMessage()    {}

func (m *InterfaceStats) GetInterface() []*InterfaceStats_Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type InterfaceStats_Interface struct {
	InternalName string `protobuf:"bytes,1,opt,name=internal_name,proto3" json:"internal_name,omitempty"`
	Index        uint32 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Status       string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	IpAddress    string `protobuf:"bytes,5,opt,name=ip_address,proto3" json:"ip_address,omitempty"`
	MacAddress   string `protobuf:"bytes,6,opt,name=mac_address,proto3" json:"mac_address,omitempty"`
}

func (m *InterfaceStats_Interface) Reset()         { *m = InterfaceStats_Interface{} }
func (m *InterfaceStats_Interface) String() string { return proto.CompactTextString(m) }
func (*InterfaceStats_Interface) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("status.OperationalState", OperationalState_name, OperationalState_value)
}
