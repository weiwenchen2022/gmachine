// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

import "errors"

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

const (
	OpHALT uint64 = iota
	OpNOOP
	OpINCA
	OpDECA
	OpSETA
)

type Machine struct {
	P      uint64
	Memory []uint64
	A      uint64
}

func New() *Machine {
	return &Machine{
		Memory: make([]uint64, DefaultMemSize),
	}
}

func (m *Machine) Run() {
	for {
		opCode := m.Memory[m.P]
		m.P++
		if int(m.P) >= len(m.Memory) {
			m.P = 0
		}

		switch opCode {
		case OpHALT:
			return
		case OpNOOP:
		case OpINCA:
			m.A++
		case OpDECA:
			m.A--
		case OpSETA:
			m.A = m.Memory[m.P]
			m.P++
		}
	}
}

func (m *Machine) RunProgram(input []uint64) error {
	if len(input) >= len(m.Memory) {
		return errors.New("out of memory")
	}

	n := copy(m.Memory, input)
	m.Memory[n] = OpHALT

	m.P = 0
	m.A = 0
	m.Run()
	return nil
}
