package pe

import "github.com/mewmew/pe/enum"

// RawFileHeader is a COFF file header (in raw format).
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#coff-file-header-object-and-image
type RawFileHeader struct {
	// Target CPU type.
	//
	// offset: 0x0000 (2 bytes)
	Machine enum.MachineType
	// Number of sections.
	//
	// offset: 0x0002 (2 bytes)
	NSections uint16
	// File creation time, measured in number of seconds since Epoch.
	//
	// offset: 0x0004 (4 bytes)
	Date uint32
	// File offset of COFF symbol table.
	//
	// offset: 0x0008 (4 bytes)
	SymbolTableOffset uint32
	// Number of entries in symbol table.
	//
	// offset: 0x000C (4 bytes)
	NSymbols uint32
	// Size in bytes of optional header.
	//
	// offset: 0x0010 (2 bytes)
	OptHdrSize uint16
	// Image characteristics.
	//
	// offset: 0x0012 (2 bytes)
	Characteristics enum.Characteristics
}

// RawOptHeader32 is an optional header of a 32-bit PE file (in raw format).
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#optional-header-image-only
type RawOptHeader32 struct {
	// Magic number (PE32 or PE32+).
	//
	// offset: 0x0000 (2 bytes)
	//Magic uint16
	// Major linker version.
	//
	// offset: 0x0002 (1 bytes)
	MajorLinkerVer uint8
	// Minor linker version.
	//
	// offset: 0x0003 (1 bytes)
	MinorLinkerVer uint8
	// Size of code section(s) in number of bytes.
	//
	// offset: 0x0004 (4 bytes)
	CodeSize uint32
	// Size of initialized data section(s) in number of bytes.
	//
	// offset: 0x0008 (4 bytes)
	InitializedDataSize uint32
	// Size of uninitialized data section(s) in number of bytes.
	//
	// offset: 0x000C (4 bytes)
	UninitializedDataSize uint32
	// Relative address to entry point of the executable (relative to image
	// base).
	//
	// offset: 0x0010 (4 bytes)
	EntryRelAddr uint32
	// Relative address of code section (relative to image base).
	//
	// offset: 0x0014 (4 bytes)
	CodeBase uint32
	// Relative address of data section (relative to image base).
	//
	// offset: 0x0018 (4 bytes)
	DataBase uint32
	// Base address of image when loaded into memory (preferred address).
	//
	// offset: 0x001C (4 bytes)
	ImageBase uint32
	// Section alignment in bytes.
	//
	// offset: 0x0020 (4 bytes)
	SectionAlign uint32
	// File alignment in bytes.
	//
	// offset: 0x0024 (4 bytes)
	FileAlign uint32
	// Major operating system version.
	//
	// offset: 0x0028 (2 bytes)
	MajorOSVer uint16
	// Minor operating system version.
	//
	// offset: 0x002A (2 bytes)
	MinorOSVer uint16
	// Major image version.
	//
	// offset: 0x002C (2 bytes)
	MajorImageVer uint16
	// Minor image version.
	//
	// offset: 0x002E (2 bytes)
	MinorImageVer uint16
	// Major subsystem version.
	//
	// offset: 0x0030 (2 bytes)
	MajorSubsystemVer uint16
	// Minor subsystem version.
	//
	// offset: 0x0032 (2 bytes)
	MinorSubsystemVer uint16
	// Reserved.
	//
	// offset: 0x0034 (4 bytes)
	Win32Ver uint32
	// Size of image in bytes.
	//
	// offset: 0x0038 (4 bytes)
	ImageSize uint32
	// Size of headers (rounded up to multiple of FileAlign).
	//
	// offset: 0x003C (4 bytes)
	HeadersSize uint32
	// Image checksum.
	//
	// offset: 0x0040 (4 bytes)
	Checksum uint32
	// Subsystem required to run image.
	//
	// offset: 0x0044 (2 bytes)
	Subsystem enum.Subsystem
	// DLL characteristics.
	//
	// offset: 0x0046 (2 bytes)
	DLLCharacteristics enum.DLLCharacteristics
	// Reserved stack space in bytes.
	//
	// offset: 0x0048 (4 bytes)
	ReservedStackSize uint32
	// Initial stack size in bytes.
	//
	// offset: 0x004C (4 bytes)
	InitialStackSize uint32
	// Reserved heap space in bytes.
	//
	// offset: 0x0050 (4 bytes)
	ReservedHeapSize uint32
	// Initial heap space in bytes.
	//
	// offset: 0x0054 (4 bytes)
	InitialHeapSize uint32
	// Reserved.
	//
	// offset: 0x0058 (4 bytes)
	LoaderFlags uint32
	// Number of data directories.
	//
	// offset: 0x005C (4 bytes)
	NDataDirs uint32
}

// RawOptHeader64 is an optional header of a 64-bit PE file (in raw format).
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#optional-header-image-only
type RawOptHeader64 struct {
	// Magic number (PE32 or PE32+).
	//
	// offset: 0x0000 (2 bytes)
	//Magic uint16
	// Major linker version.
	//
	// offset: 0x0002 (1 bytes)
	MajorLinkerVer uint8
	// Minor linker version.
	//
	// offset: 0x0003 (1 bytes)
	MinorLinkerVer uint8
	// Size of code section(s) in number of bytes.
	//
	// offset: 0x0004 (4 bytes)
	CodeSize uint32
	// Size of initialized data section(s) in number of bytes.
	//
	// offset: 0x0008 (4 bytes)
	InitializedDataSize uint32
	// Size of uninitialized data section(s) in number of bytes.
	//
	// offset: 0x000C (4 bytes)
	UninitializedDataSize uint32
	// Relative address to entry point of the executable (relative to image
	// base).
	//
	// offset: 0x0010 (4 bytes)
	EntryRelAddr uint32
	// Relative address of code section (relative to image base).
	//
	// offset: 0x0014 (4 bytes)
	CodeBase uint32
	// Base address of image when loaded into memory (preferred address).
	//
	// offset: 0x0018 (8 bytes)
	ImageBase uint64
	// Section alignment in bytes.
	//
	// offset: 0x0020 (4 bytes)
	SectionAlign uint32
	// File alignment in bytes.
	//
	// offset: 0x0024 (4 bytes)
	FileAlign uint32
	// Major operating system version.
	//
	// offset: 0x0028 (2 bytes)
	MajorOSVer uint16
	// Minor operating system version.
	//
	// offset: 0x002A (2 bytes)
	MinorOSVer uint16
	// Major image version.
	//
	// offset: 0x002C (2 bytes)
	MajorImageVer uint16
	// Minor image version.
	//
	// offset: 0x002E (2 bytes)
	MinorImageVer uint16
	// Major subsystem version.
	//
	// offset: 0x0030 (2 bytes)
	MajorSubsystemVer uint16
	// Minor subsystem version.
	//
	// offset: 0x0032 (2 bytes)
	MinorSubsystemVer uint16
	// Reserved.
	//
	// offset: 0x0034 (4 bytes)
	Win32Ver uint32
	// Size of image in bytes.
	//
	// offset: 0x0038 (4 bytes)
	ImageSize uint32
	// Size of headers (rounded up to multiple of FileAlign).
	//
	// offset: 0x003C (4 bytes)
	HeadersSize uint32
	// Image checksum.
	//
	// offset: 0x0040 (4 bytes)
	Checksum uint32
	// Subsystem required to run image.
	//
	// offset: 0x0044 (2 bytes)
	Subsystem enum.Subsystem
	// DLL characteristics.
	//
	// offset: 0x0046 (2 bytes)
	DLLCharacteristics enum.DLLCharacteristics
	// Reserved stack space in bytes.
	//
	// offset: 0x0048 (8 bytes)
	ReservedStackSize uint64
	// Initial stack size in bytes.
	//
	// offset: 0x0050 (8 bytes)
	InitialStackSize uint64
	// Reserved heap space in bytes.
	//
	// offset: 0x0058 (8 bytes)
	ReservedHeapSize uint64
	// Initial heap space in bytes.
	//
	// offset: 0x0060 (8 bytes)
	InitialHeapSize uint64
	// Reserved.
	//
	// offset: 0x0068 (4 bytes)
	LoaderFlags uint32
	// Number of data directories.
	//
	// offset: 0x006C (4 bytes)
	NDataDirs uint32
}

// RawSectionHeader is a section header (in raw format).
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#section-table-section-headers
type RawSectionHeader struct {
	// Section name (UTF-8 encoded, NULL-terminated).
	//
	// offset: 0x0000 (8 bytes)
	Name [8]byte
	// Size of section when loaded into memory.
	//
	// offset: 0x0008 (4 bytes)
	VirtualSize uint32
	// Relative address of section (relative to image base).
	//
	// offset: 0x000C (4 bytes)
	RelAddr uint32
	// Size of data on disk.
	//
	// offset: 0x0010 (4 bytes)
	DataSize uint32
	// File offset of section contents.
	//
	// offset: 0x0014 (4 bytes)
	DataOffset uint32
	// File offset of relaction entries.
	//
	// offset: 0x0018 (4 bytes)
	RelocsOffset uint32
	// File offset of line number entries.
	//
	// offset: 0x001C (4 bytes)
	LineNumsOffset uint32
	// Number of relocation entires.
	//
	// offset: 0x0020 (2 bytes)
	NRelocs uint16
	// Number of line number entires.
	//
	// offset: 0x0022 (2 bytes)
	NLineNums uint16
	// Section flags.
	//
	// offset: 0x0024 (8 bytes)
	Flags enum.SectionFlags
}

// --- [ Data directories ] ----------------------------------------------------

// ~~~ [ Base Relocation Table ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// RawBaseRelocBlock is a base relocation block descriptor (in raw format).
type RawBaseRelocBlock struct {
	// Relative address of page. The address of a relocation is computed by
	// adding the image base, the page relative address and the offset of the
	// reloc.
	//
	// offset: 0x0000 (4 bytes)
	PageRelAddr uint32
	// Size of block in number of bytes.
	//
	// offset: 0x0004 (4 bytes)
	BlockSize uint32
}

// RawBaseRelocEntry is a base relocation entry (in raw format).
type RawBaseRelocEntry struct {
	// Bitfield of data.
	//
	//    // Base relocation type.
	//    Type   : 4
	//    // Offset of base relocation from the start of the base relocation
	//    // block.
	//    Offset : 12
	//
	// offset: 0x0000 (2 bytes)
	Bitfield uint16
}

// ~~~ [ Debug ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// RawDebugDirectory is a debug data directory (in raw format).
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#debug-directory-image-only
type RawDebugDirectory struct {
	// Reserved.
	//
	// offset: 0x0000 (4 bytes)
	Characteristics uint32
	// Debug data creation time, measured in number of seconds since Epoch.
	//
	// offset: 0x0004 (4 bytes)
	Date uint32
	// Major debug data format version.
	//
	// offset: 0x0008 (2 bytes)
	MajorVer uint16
	// Minor debug data format version.
	//
	// offset: 0x000A (2 bytes)
	MinorVer uint16
	// Debug data format type.
	//
	// offset: 0x000C (4 bytes)
	Type enum.DebugType
	// Debug data size in bytes.
	//
	// offset: 0x0010 (4 bytes)
	Size uint32
	// Relative address to debug data (relative to image base).
	//
	// offset: 0x0014 (4 bytes)
	RelAddr uint32
	// File offset of debug data.
	//
	// offset: 0x0018 (4 bytes)
	Offset uint32
}

// RawCodeViewInfo contains CodeView debug information (in raw format).
//
// ref: Visual C++ 5.0 Symbolic Debug Information Specification
// ref: https://github.com/Microsoft/microsoft-pdb/blob/master/include/cvinfo.h
type RawCodeViewInfo struct {
	// CodeView signature ("NB10").
	//
	// offset: 0x0000 (4 bytes)
	Signature uint32
	// CodeView offset.
	//
	// offset: 0x0004 (4 bytes)
	Offset uint32
	// CodeView debug data creation time, measured in number of seconds since
	// Epoch.
	//
	// offset: 0x0008 (4 bytes)
	Date uint32
	// Incremental number, initially set to 1 and incremented for each partial
	// write to the PDB file.
	//
	// offset: 0x000C (4 bytes)
	Age uint32
}

// RawFPOData represents the stack frame layout for a function on an x86
// computer when frame pointer omission (FPO) optimization is used. The
// structure is used to locate the base of the call frame.
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#debug-type
// ref: https://docs.microsoft.com/en-us/windows/desktop/api/winnt/ns-winnt-fpo_data
type RawFPOData struct {
	// Offset to first byte of function code.
	//
	// offset: 0x0000 (4 bytes)
	StartOffset uint32
	// Function size in number of bytes.
	//
	// offset: 0x0004 (4 bytes)
	FuncSize uint32
	// Number of local variables / 4.
	//
	// offset: 0x0008 (4 bytes)
	NLocals uint32
	// Size of parameters in number of bytes / 4.
	//
	// offset: 0x000C (2 bytes)
	Params uint16
	// Size of function prolog code in number of bytes.
	//
	// offset: 0x000E (1 bytes)
	Prolog uint8
	// Bitfield of data.
	//
	//    // Number of registers saved.
	//    Regs     : 3 bits
	//    // Specifies whether the function uses structured exception handling.
	//    HasSEH   : 1 bit
	//    // Specifies whether the EBP register has been allocated.
	//    UseBP    : 1 bit
	//    // Reserved.
	//    Reserved : 1 bit
	//    // Frame type of function.
	//    Frame    : 2 bits
	//
	// offset: 0x000F (1 bytes)
	Bitfield uint8
}
