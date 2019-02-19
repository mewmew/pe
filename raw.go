package pe

import (
	"github.com/mewmew/pe/enum"
	"github.com/mewmew/pe/internal/pe"
)

// goFileHeader converts the raw file header into a corresponding Go version.
func goFileHeader(raw *pe.RawFileHeader) *FileHeader {
	return &FileHeader{
		Machine:           raw.Machine,
		NSections:         raw.NSections,
		Date:              parseDateFromEpoch(raw.Date),
		SymbolTableOffset: raw.SymbolTableOffset,
		NSymbols:          raw.NSymbols,
		OptHdrSize:        raw.OptHdrSize,
		Characteristics:   raw.Characteristics,
	}
}

// goOptHeader32 converts the raw optional header into a corresponding Go
// version.
func goOptHeader32(raw *pe.RawOptHeader32, magic uint16) *OptHeader {
	return &OptHeader{
		Magic:                 magic,
		MajorLinkerVer:        raw.MajorLinkerVer,
		MinorLinkerVer:        raw.MinorLinkerVer,
		CodeSize:              raw.CodeSize,
		InitializedDataSize:   raw.InitializedDataSize,
		UninitializedDataSize: raw.UninitializedDataSize,
		EntryRelAddr:          raw.EntryRelAddr,
		CodeBase:              raw.CodeBase,
		DataBase:              raw.DataBase,
		ImageBase:             uint64(raw.ImageBase),
		SectionAlign:          raw.SectionAlign,
		FileAlign:             raw.FileAlign,
		MajorOSVer:            raw.MajorOSVer,
		MinorOSVer:            raw.MinorOSVer,
		MajorImageVer:         raw.MajorImageVer,
		MinorImageVer:         raw.MinorImageVer,
		MajorSubsystemVer:     raw.MajorSubsystemVer,
		MinorSubsystemVer:     raw.MinorSubsystemVer,
		Win32Ver:              raw.Win32Ver,
		ImageSize:             raw.ImageSize,
		HeadersSize:           raw.HeadersSize,
		Checksum:              raw.Checksum,
		Subsystem:             raw.Subsystem,
		DLLCharacteristics:    raw.DLLCharacteristics,
		ReservedStackSize:     uint64(raw.ReservedStackSize),
		InitialStackSize:      uint64(raw.InitialStackSize),
		ReservedHeapSize:      uint64(raw.ReservedHeapSize),
		InitialHeapSize:       uint64(raw.InitialHeapSize),
		LoaderFlags:           raw.LoaderFlags,
		NDataDirs:             raw.NDataDirs,
	}
}

// goOptHeader64 converts the raw optional header into a corresponding Go
// version.
func goOptHeader64(raw *pe.RawOptHeader64, magic uint16) *OptHeader {
	return &OptHeader{
		Magic:                 magic,
		MajorLinkerVer:        raw.MajorLinkerVer,
		MinorLinkerVer:        raw.MinorLinkerVer,
		CodeSize:              raw.CodeSize,
		InitializedDataSize:   raw.InitializedDataSize,
		UninitializedDataSize: raw.UninitializedDataSize,
		EntryRelAddr:          raw.EntryRelAddr,
		CodeBase:              raw.CodeBase,
		ImageBase:             raw.ImageBase,
		SectionAlign:          raw.SectionAlign,
		FileAlign:             raw.FileAlign,
		MajorOSVer:            raw.MajorOSVer,
		MinorOSVer:            raw.MinorOSVer,
		MajorImageVer:         raw.MajorImageVer,
		MinorImageVer:         raw.MinorImageVer,
		MajorSubsystemVer:     raw.MajorSubsystemVer,
		MinorSubsystemVer:     raw.MinorSubsystemVer,
		Win32Ver:              raw.Win32Ver,
		ImageSize:             raw.ImageSize,
		HeadersSize:           raw.HeadersSize,
		Checksum:              raw.Checksum,
		Subsystem:             raw.Subsystem,
		DLLCharacteristics:    raw.DLLCharacteristics,
		ReservedStackSize:     raw.ReservedStackSize,
		InitialStackSize:      raw.InitialStackSize,
		ReservedHeapSize:      raw.ReservedHeapSize,
		InitialHeapSize:       raw.InitialHeapSize,
		LoaderFlags:           raw.LoaderFlags,
		NDataDirs:             raw.NDataDirs,
	}
}

// goSectionHeader converts the raw section header into a corresponding Go
// version.
func goSectionHeader(raw pe.RawSectionHeader) SectionHeader {
	return SectionHeader{
		Name:           parseCString(raw.Name[:]),
		VirtualSize:    raw.VirtualSize,
		RelAddr:        raw.RelAddr,
		DataSize:       raw.DataSize,
		DataOffset:     raw.DataOffset,
		RelocsOffset:   raw.RelocsOffset,
		LineNumsOffset: raw.LineNumsOffset,
		NRelocs:        raw.NRelocs,
		NLineNums:      raw.NLineNums,
		Flags:          raw.Flags,
	}
}

// --- [ Data directories ] ----------------------------------------------------

// ~~~ [ Base Relocation Table ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// goBaseRelocEntry converts the raw base relocation entry into a corresponding
// Go version.
func goBaseRelocEntry(raw pe.RawBaseRelocEntry) BaseRelocEntry {
	return BaseRelocEntry{
		Type:   enum.BaseRelocType(raw.Bitfield & 0xF000 >> 12), // 0b1111000000000000
		Offset: raw.Bitfield & 0x0FFF,                           // 0b111111111111
	}
}

// ~~~ [ Debug ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// goDebugDirectory converts the raw debug data directory into a corresponding Go
// version.
func goDebugDirectory(raw pe.RawDebugDirectory) DebugDirectory {
	return DebugDirectory{
		Characteristics: raw.Characteristics,
		Date:            parseDateFromEpoch(raw.Date),
		MajorVer:        raw.MajorVer,
		MinorVer:        raw.MinorVer,
		Type:            raw.Type,
		Size:            raw.Size,
		RelAddr:         raw.RelAddr,
		Offset:          raw.Offset,
	}
}

// goCodeViewInfo converts the raw CodeView debug info into a corresponding Go
// version.
func goCodeViewInfo(raw pe.RawCodeViewInfo, pdbPath string) CodeViewInfo {
	return CodeViewInfo{
		Signature: raw.Signature,
		Offset:    raw.Offset,
		Date:      parseDateFromEpoch(raw.Date),
		Age:       raw.Age,
		PDBPath:   pdbPath,
	}
}

// goFPOData converts the raw FPO data into a corresponding Go version.
func goFPOData(raw pe.RawFPOData) FPOData {
	// TODO: use binary literal
	// Regs     : 3 bits
	regs := uint8(raw.Bitfield & 0x07) // 0b00000111
	// HasSEH   : 1 bit
	hasSEH := (raw.Bitfield & 0x08) != 0 // 0b00001000
	// UseBP    : 1 bit
	useBP := (raw.Bitfield & 0x10) != 0 // 0b00010000
	// Reserved : 1 bit
	reserved := uint8(raw.Bitfield & 0x20 >> 5) // 0b00100000
	// Frame    : 2 bits
	frame := enum.FrameType(raw.Bitfield & 0xC0 >> 6) // 0b11000000
	fpo := FPOData{
		StartOffset: raw.StartOffset,
		FuncSize:    raw.FuncSize,
		NLocals:     uint64(raw.NLocals) * 4, // raw format was stored as / 4.
		Params:      uint32(raw.Params) * 4,  // raw format was stored as / 4.
		Prolog:      raw.Prolog,
		Regs:        regs,
		HasSEH:      hasSEH,
		UseBP:       useBP,
		Reserved:    reserved,
		Frame:       frame,
	}
	return fpo
}
