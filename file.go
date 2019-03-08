package pe

import (
	"fmt"
	"time"

	"github.com/mewmew/pe/enum"
)

// File is a Portable Executable (PE) file.
type File struct {
	// File contents.
	Content []byte
	// COFF file header.
	FileHdr *FileHeader
	// Optional header.
	OptHdr *OptHeader
	// Data directories.
	DataDirs []DataDirectory
	// Section headers.
	SectHdrs []SectionHeader
	// Data directory contents.
	//
	// 0 - Export Table
	// 1 - Import Table
	Imps []ImportEntry
	// 2 - Resource Table
	// 3 - Exception Table
	// 4 - Certificate Table
	// 5 - Base Relocation Table
	BaseRelocBlocks []BaseRelocBlock
	// 6 - Debug data
	DbgData []DebugData
	// 7 - Architecture
	// 8 - Global Pointer Register
	// 9 - TLS Table
	// 10 - Load Config Table
	// 11 - Bound Import Table
	// 12 - Import Address Table
	// 13 - Delay Import Descriptor
	// 14 - CLR Header
	// 15 - Reserved
}

// ReadData reads the data with the specified address and length from the
// section containing the memory range. It panics if no such section is located.
func (file *File) ReadData(addr uint64, n int64) []byte {
	for _, sectHdr := range file.SectHdrs {
		sectStartAddr := file.OptHdr.ImageBase + uint64(sectHdr.RelAddr)
		sectEndAddr := sectStartAddr + uint64(sectHdr.DataSize)
		if !(sectStartAddr <= addr && addr+uint64(n) < sectEndAddr) {
			continue
		}
		offset := addr - sectStartAddr
		start := uint64(sectHdr.DataOffset) + offset
		end := start + uint64(n)
		return file.Content[start:end]
	}
	panic(fmt.Errorf("unable to locate data at address 0x%08X (%d bytes)", addr, n))
}

// FileHeader is a COFF file header.
type FileHeader struct {
	// Target CPU type.
	Machine enum.MachineType
	// Number of sections.
	NSections uint16
	// File creation time.
	Date time.Time
	// File offset of COFF symbol table.
	SymbolTableOffset uint32
	// Number of entries in symbol table.
	NSymbols uint32
	// Size in bytes of optional header.
	OptHdrSize uint16
	// Image characteristics.
	Characteristics enum.Characteristic
}

// OptHeader is an optional header of a PE file.
type OptHeader struct {
	// Magic number (PE32 or PE32+).
	Magic uint16
	// Major linker version.
	MajorLinkerVer uint8
	// Minor linker version.
	MinorLinkerVer uint8
	// Size of code section(s) in number of bytes.
	CodeSize uint32
	// Size of initialized data section(s) in number of bytes.
	InitializedDataSize uint32
	// Size of uninitialized data section(s) in number of bytes.
	UninitializedDataSize uint32
	// Relative address to entry point of the executable (relative to image
	// base).
	EntryRelAddr uint32
	// Relative address of code section (relative to image base).
	CodeBase uint32
	// Relative address of data section (relative to image base); field only used
	// by 32-bit optional header.
	DataBase uint32
	// Base address of image when loaded into memory (preferred address).
	ImageBase uint64
	// Section alignment in bytes.
	SectionAlign uint32
	// File alignment in bytes.
	FileAlign uint32
	// Major operating system version.
	MajorOSVer uint16
	// Minor operating system version.
	MinorOSVer uint16
	// Major image version.
	MajorImageVer uint16
	// Minor image version.
	MinorImageVer uint16
	// Major subsystem version.
	MajorSubsystemVer uint16
	// Minor subsystem version.
	MinorSubsystemVer uint16
	// Reserved.
	Win32Ver uint32
	// Size of image in bytes.
	ImageSize uint32
	// Size of headers (rounded up to multiple of FileAlign).
	HeadersSize uint32
	// Image checksum.
	Checksum uint32
	// Subsystem required to run image.
	Subsystem enum.Subsystem
	// DLL characteristics.
	DLLCharacteristics enum.DLLCharacteristic
	// Reserved stack space in bytes.
	ReservedStackSize uint64
	// Initial stack size in bytes.
	InitialStackSize uint64
	// Reserved heap space in bytes.
	ReservedHeapSize uint64
	// Initial heap space in bytes.
	InitialHeapSize uint64
	// Reserved.
	LoaderFlags uint32
	// Number of data directories.
	NDataDirs uint32
}

// DataDirectory is a data directory of a PE file.
type DataDirectory struct {
	// Relative address to table.
	RelAddr uint32
	// Size of table in bytes.
	Size uint32
}

// SectionHeader is a section header.
type SectionHeader struct {
	// Section name.
	Name string
	// Size of section when loaded into memory.
	VirtualSize uint32
	// Relative address of section (relative to image base).
	RelAddr uint32
	// Size of data on disk.
	DataSize uint32
	// File offset of section contents.
	DataOffset uint32
	// File offset of relaction entries.
	RelocsOffset uint32
	// File offset of line number entries.
	LineNumsOffset uint32
	// Number of relocation entires.
	NRelocs uint16
	// Number of line number entires.
	NLineNums uint16
	// Section flags.
	Flags enum.SectionFlag
}
