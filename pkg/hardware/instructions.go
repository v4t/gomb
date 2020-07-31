package hardware

// Instructions represents basic cpu instructions
var Instructions = [0x100]func(*CPU){
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,

	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
	nop, nop, nop, nop,
}

func nop(cpu *CPU) {

}
