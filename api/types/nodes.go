// Package types provides shared types and structs.
package types

// Nodes holds list of Node type
type Nodes []Node

// Node structure to hold node details like role and machine id
type Node struct {
	Name      string `json:"name" yaml:"name"`
	Role      string `json:"role" yaml:"role"`
	MachineID int    `json:"machineid" yaml:"machineid"`
}
