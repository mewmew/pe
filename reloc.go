package pe

import "github.com/mewmew/pe/enum"

// BaseRelocBlock is a base relocation block descriptor.
type BaseRelocBlock struct {
	// Relative address of page. The address of a relocation is computed by
	// adding the image base, the page relative address and the offset of the
	// reloc.
	PageRelAddr uint32
	// Size of block in number of bytes.
	BlockSize uint32
	// Relocation entries of the block.
	Entries []BaseRelocEntry
}

// BaseRelocEntry is a base relocation entry.
type BaseRelocEntry struct {
	// Base relocation type.
	Type enum.BaseRelocType
	// Offset of base relocation from the start of the base relocation block.
	Offset uint16
}
