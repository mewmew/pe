// Package enum defines enumerate types of the PE file format.
package enum

import "strings"

//go:generate stringer -trimprefix MachineType -type MachineType

// MachineType specifies the CPU type of a PE image.
type MachineType uint16

// Machine types.
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#machine-types
const (
	MachineTypeUnknown   MachineType = 0x0000 // The contents of this field are assumed to be applicable to any machine type
	MachineTypeAM33      MachineType = 0x01D3 // Matsushita AM33
	MachineTypeAMD64     MachineType = 0x8664 // x64
	MachineTypeARM       MachineType = 0x01C0 // ARM little endian
	MachineTypeARM64     MachineType = 0xAA64 // ARM64 little endian
	MachineTypeARMNT     MachineType = 0x01C4 // ARM Thumb-2 little endian
	MachineTypeEBC       MachineType = 0x0EBC // EFI byte code
	MachineTypeI386      MachineType = 0x014C // Intel 386 or later processors and compatible processors
	MachineTypeIA64      MachineType = 0x0200 // Intel Itanium processor family
	MachineTypeM32R      MachineType = 0x9041 // Mitsubishi M32R little endian
	MachineTypeMIPS16    MachineType = 0x0266 // MIPS16
	MachineTypeMIPSFPU   MachineType = 0x0366 // MIPS with FPU
	MachineTypeMIPSFPU16 MachineType = 0x0466 // MIPS16 with FPU
	MachineTypePowerPC   MachineType = 0x01F0 // Power PC little endian
	MachineTypePowerPCFP MachineType = 0x01F1 // Power PC with floating point support
	MachineTypeR4000     MachineType = 0x0166 // MIPS little endian
	MachineTypeRISCV32   MachineType = 0x5032 // RISC-V 32-bit address space
	MachineTypeRISCV64   MachineType = 0x5064 // RISC-V 64-bit address space
	MachineTypeRISCV128  MachineType = 0x5128 // RISC-V 128-bit address space
	MachineTypeSH3       MachineType = 0x01A2 // Hitachi SH3
	MachineTypeSH3DSP    MachineType = 0x01A3 // Hitachi SH3 DSP
	MachineTypeSH4       MachineType = 0x01A6 // Hitachi SH4
	MachineTypeSH5       MachineType = 0x01A8 // Hitachi SH5
	MachineTypeThumb     MachineType = 0x01C2 // Thumb
	MachineTypeWCEMIPSv2 MachineType = 0x0169 // MIPS little-endian WCE v2
)

//go:generate stringer -trimprefix Characteristic -type Characteristic

// Characteristic is an image characteristic.
type Characteristic uint16

// Image characteristics.
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#characteristics
const (
	CharacteristicRelocsStripped       Characteristic = 0x0001 // Image only, Windows CE, and Microsoft Windows NT and later. This indicates that the file does not contain base relocations and must therefore be loaded at its preferred base address. If the base address is not available, the loader reports an error. The default behavior of the linker is to strip base relocations from executable (EXE) files.
	CharacteristicExecutableImage      Characteristic = 0x0002 // Image only. This indicates that the image file is valid and can be run. If this flag is not set, it indicates a linker error.
	CharacteristicLineNumsStripped     Characteristic = 0x0004 // COFF line numbers have been removed. This flag is deprecated and should be zero.
	CharacteristicLocalSymsStripped    Characteristic = 0x0008 // COFF symbol table entries for local symbols have been removed. This flag is deprecated and should be zero.
	CharacteristicAggressiveWSTrim     Characteristic = 0x0010 // Obsolete. Aggressively trim working set. This flag is deprecated for Windows 2000 and later and must be zero.
	CharacteristicLargeAddressAware    Characteristic = 0x0020 // Application can handle > 2-GB addresses.
	CharacteristicReserved0040         Characteristic = 0x0040 // This flag is reserved for future use.
	CharacteristicBytesReversedLo      Characteristic = 0x0080 // Little endian: the least significant bit (LSB) precedes the most significant bit (MSB) in memory. This flag is deprecated and should be zero.
	Characteristic32bitMachine         Characteristic = 0x0100 // Machine is based on a 32-bit-word architecture.
	CharacteristicDebugStripped        Characteristic = 0x0200 // Debugging information is removed from the image file.
	CharacteristicRemovableRunFromSwap Characteristic = 0x0400 // If the image is on removable media, fully load it and copy it to the swap file.
	CharacteristicNetRunFromSwap       Characteristic = 0x0800 // If the image is on network media, fully load it and copy it to the swap file.
	CharacteristicSystem               Characteristic = 0x1000 // The image file is a system file, not a user program.
	CharacteristicDLL                  Characteristic = 0x2000 // The image file is a dynamic-link library (DLL). Such files are considered executable files for almost all purposes, although they cannot be directly run.
	CharacteristicUpSystemOnly         Characteristic = 0x4000 // The file should be run only on a uniprocessor machine.
	CharacteristicBytesReversedHi      Characteristic = 0x8000 // Big endian: the MSB precedes the LSB in memory. This flag is deprecated and should be zero.
)

// Characteristics is a bitfield of image characteristics.
type Characteristics Characteristic

// String returns the string representation of the image characteristics.
func (c Characteristics) String() string {
	var ss []string
	for mask := uint32(1); mask < 0xFFFF; mask <<= 1 {
		v := Characteristic(c) & Characteristic(mask)
		if v != 0 {
			s := v.String()
			ss = append(ss, s)
		}
	}
	return strings.Join(ss, " | ")
}

//go:generate stringer -trimprefix Subsystem -type Subsystem

// Subsystem is a Windows subsystem.
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#windows-subsystem
type Subsystem uint16

// Windows subsystems.
const (
	SubsystemUnknown                Subsystem = 0  // An unknown subsystem
	SubsystemNative                 Subsystem = 1  // Device drivers and native Windows processes
	SubsystemWindowsGUI             Subsystem = 2  // The Windows graphical user interface (GUI) subsystem
	SubsystemWindowsCUI             Subsystem = 3  // The Windows character subsystem
	SubsystemOS2CUI                 Subsystem = 5  // The OS/2 character subsystem
	SubsystemPosixCUI               Subsystem = 7  // The Posix character subsystem
	SubsystemNativeWindows          Subsystem = 8  // Native Win9x driver
	SubsystemWindowsCEGUI           Subsystem = 9  // Windows CE
	SubsystemEFIApplication         Subsystem = 10 // An Extensible Firmware Interface (EFI) application
	SubsystemEFIBootServiceDriver   Subsystem = 11 // An EFI driver with boot services
	SubsystemEFIRuntimeDriver       Subsystem = 12 // An EFI driver with run-time services
	SubsystemEFIRom                 Subsystem = 13 // An EFI ROM image
	SubsystemXbox                   Subsystem = 14 // XBOX
	SubsystemWindowsBootApplication Subsystem = 16 // Windows boot application.
)

//go:generate stringer -trimprefix DLLCharacteristic -type DLLCharacteristic

// DLLCharacteristic is a DLL characteristic.
type DLLCharacteristic uint16

// DLL characteristics.
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#dll-characteristics
const (
	DLLCharacteristicsReserved0001        DLLCharacteristic = 0x0001 // Reserved, must be zero.
	DLLCharacteristicsReserved0002        DLLCharacteristic = 0x0002 // Reserved, must be zero.
	DLLCharacteristicsReserved0004        DLLCharacteristic = 0x0004 // Reserved, must be zero.
	DLLCharacteristicsReserved0008        DLLCharacteristic = 0x0008 // Reserved, must be zero.
	DLLCharacteristicsHighEntropyVA       DLLCharacteristic = 0x0020 // Image can handle a high entropy 64-bit virtual address space.
	DLLCharacteristicsDynamicBase         DLLCharacteristic = 0x0040 // DLL can be relocated at load time.
	DLLCharacteristicsForceIntegrity      DLLCharacteristic = 0x0080 // Code Integrity checks are enforced.
	DLLCharacteristicsNXCompat            DLLCharacteristic = 0x0100 // Image is NX compatible.
	DLLCharacteristicsNoIsolation         DLLCharacteristic = 0x0200 // Isolation aware, but do not isolate the image.
	DLLCharacteristicsNoSEH               DLLCharacteristic = 0x0400 // Does not use structured exception (SE) handling. No SE handler may be called in this image.
	DLLCharacteristicsNoBind              DLLCharacteristic = 0x0800 // Do not bind the image.
	DLLCharacteristicsAppContainer        DLLCharacteristic = 0x1000 // Image must execute in an AppContainer.
	DLLCharacteristicsWDMDriver           DLLCharacteristic = 0x2000 // A WDM driver.
	DLLCharacteristicsGuardCF             DLLCharacteristic = 0x4000 // Image supports Control Flow Guard.
	DLLCharacteristicsTerminalServerAware DLLCharacteristic = 0x8000 // Terminal Server aware.
)

// DLLCharacteristics is a bitfield of DLL characteristics.
type DLLCharacteristics DLLCharacteristic

// String returns the string representation of the DLL characteristics.
func (c DLLCharacteristics) String() string {
	var ss []string
	for mask := uint32(1); mask < 0xFFFF; mask <<= 1 {
		v := DLLCharacteristic(c) & DLLCharacteristic(mask)
		if v != 0 {
			s := v.String()
			ss = append(ss, s)
		}
	}
	return strings.Join(ss, " | ")
}

//go:generate stringer -trimprefix SectionFlag -type SectionFlag

// SectionFlag is a section flag.
type SectionFlag uint32

// Section flags.
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#section-flags
const (
	SectionFlagReserved00000000          SectionFlag = 0x00000000 // Reserved for future use.
	SectionFlagReserved00000001          SectionFlag = 0x00000001 // Reserved for future use.
	SectionFlagReserved00000002          SectionFlag = 0x00000002 // Reserved for future use.
	SectionFlagReserved00000004          SectionFlag = 0x00000004 // Reserved for future use.
	SectionFlagTypeNoPad                 SectionFlag = 0x00000008 // The section should not be padded to the next boundary. This flag is obsolete and is replaced by SectionFlagAlign1. This is valid only for object files.
	SectionFlagReserved00000010          SectionFlag = 0x00000010 // Reserved for future use.
	SectionFlagContainsCode              SectionFlag = 0x00000020 // The section contains executable code.
	SectionFlagContainsInitializedData   SectionFlag = 0x00000040 // The section contains initialized data.
	SectionFlagContainsUninitializedData SectionFlag = 0x00000080 // The section contains uninitialized data.
	SectionFlagLinkOther                 SectionFlag = 0x00000100 // Reserved for future use.
	SectionFlagLinkInfo                  SectionFlag = 0x00000200 // The section contains comments or other information. The .drectve section has this type. This is valid for object files only.
	SectionFlagReserved00000400          SectionFlag = 0x00000400 // Reserved for future use.
	SectionFlagLinkRemove                SectionFlag = 0x00000800 // The section will not become part of the image. This is valid only for object files.
	SectionFlagLinkComdat                SectionFlag = 0x00001000 // The section contains COMDAT data. For more information, see COMDAT Sections (Object Only). This is valid only for object files.
	SectionFlagGPRel                     SectionFlag = 0x00008000 // The section contains data referenced through the global pointer (GP).
	SectionFlagMemPurgeable              SectionFlag = 0x00020000 // Reserved for future use.
	SectionFlagMem16bit                  SectionFlag = 0x00020000 // Reserved for future use.
	SectionFlagMemLocked                 SectionFlag = 0x00040000 // Reserved for future use.
	SectionFlagMemPreload                SectionFlag = 0x00080000 // Reserved for future use.
	SectionFlagAlign1                    SectionFlag = 0x00100000 // Align data on a 1-byte boundary. Valid only for object files.
	SectionFlagAlign2                    SectionFlag = 0x00200000 // Align data on a 2-byte boundary. Valid only for object files.
	SectionFlagAlign4                    SectionFlag = 0x00300000 // Align data on a 4-byte boundary. Valid only for object files.
	SectionFlagAlign8                    SectionFlag = 0x00400000 // Align data on an 8-byte boundary. Valid only for object files.
	SectionFlagAlign16                   SectionFlag = 0x00500000 // Align data on a 16-byte boundary. Valid only for object files.
	SectionFlagAlign32                   SectionFlag = 0x00600000 // Align data on a 32-byte boundary. Valid only for object files.
	SectionFlagAlign64                   SectionFlag = 0x00700000 // Align data on a 64-byte boundary. Valid only for object files.
	SectionFlagAlign128                  SectionFlag = 0x00800000 // Align data on a 128-byte boundary. Valid only for object files.
	SectionFlagAlign256                  SectionFlag = 0x00900000 // Align data on a 256-byte boundary. Valid only for object files.
	SectionFlagAlign512                  SectionFlag = 0x00A00000 // Align data on a 512-byte boundary. Valid only for object files.
	SectionFlagAlign1024                 SectionFlag = 0x00B00000 // Align data on a 1024-byte boundary. Valid only for object files.
	SectionFlagAlign2048                 SectionFlag = 0x00C00000 // Align data on a 2048-byte boundary. Valid only for object files.
	SectionFlagAlign4096                 SectionFlag = 0x00D00000 // Align data on a 4096-byte boundary. Valid only for object files.
	SectionFlagAlign8192                 SectionFlag = 0x00E00000 // Align data on an 8192-byte boundary. Valid only for object files.
	SectionFlagLinkNRelocOverflow        SectionFlag = 0x01000000 // The section contains extended relocations.
	SectionFlagMemDiscardable            SectionFlag = 0x02000000 // The section can be discarded as needed.
	SectionFlagMemNotCached              SectionFlag = 0x04000000 // The section cannot be cached.
	SectionFlagMemNotPaged               SectionFlag = 0x08000000 // The section is not pageable.
	SectionFlagMemShared                 SectionFlag = 0x10000000 // The section can be shared in memory.
	SectionFlagMemExecute                SectionFlag = 0x20000000 // The section can be executed as code.
	SectionFlagMemRead                   SectionFlag = 0x40000000 // The section can be read.
	SectionFlagMemWrite                  SectionFlag = 0x80000000 // The section can be written to.
)

// SectionFlags is a bitfield of section flags.
type SectionFlags SectionFlag

// String returns the string representation of the section flags.
func (flag SectionFlags) String() string {
	var ss []string
	for mask := uint64(1); mask < 0xFFFF; mask <<= 1 {
		v := SectionFlag(flag) & SectionFlag(mask)
		if v != 0 {
			s := v.String()
			ss = append(ss, s)
		}
	}
	return strings.Join(ss, " | ")
}

// --- [ Data directories ] ----------------------------------------------------

// ~~~ [ Base Relocation Table ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

//go:generate stringer -trimprefix BaseRelocType -type BaseRelocType

// BaseRelocType indicates the type of base relocation to apply by the linker.
type BaseRelocType uint8

// Base relocation types.
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#base-relocation-types
const (
	BaseRelocTypeAbsolute      BaseRelocType = 0  // The base relocation is skipped. This type can be used to pad a block.
	BaseRelocTypeHigh          BaseRelocType = 1  // The base relocation adds the high 16 bits of the difference to the 16-bit field at offset. The 16-bit field represents the high value of a 32-bit word.
	BaseRelocTypeLow           BaseRelocType = 2  // The base relocation adds the low 16 bits of the difference to the 16-bit field at offset. The 16-bit field represents the low half of a 32-bit word.
	BaseRelocTypeHighLow       BaseRelocType = 3  // The base relocation applies all 32 bits of the difference to the 32-bit field at offset.
	BaseRelocTypeHighAdj       BaseRelocType = 4  // The base relocation adds the high 16 bits of the difference to the 16-bit field at offset. The 16-bit field represents the high value of a 32-bit word. The low 16 bits of the 32-bit value are stored in the 16-bit word that follows this base relocation. This means that this base relocation occupies two slots.
	BaseRelocTypeMipsJmpAddr   BaseRelocType = 5  // The relocation interpretation is dependent on the machine type. When the machine type is MIPS, the base relocation applies to a MIPS jump instruction.
	BaseRelocTypeARMMov32      BaseRelocType = 5  // This relocation is meaningful only when the machine type is ARM or Thumb. The base relocation applies the 32-bit address of a symbol across a consecutive MOVW/MOVT instruction pair.
	BaseRelocTypeRISCVHigh20   BaseRelocType = 5  // This relocation is only meaningful when the machine type is RISC-V. The base relocation applies to the high 20 bits of a 32-bit absolute address.
	BaseRelocTypeReserved6     BaseRelocType = 6  // Reserved, must be zero.
	BaseRelocTypeThumbMov32    BaseRelocType = 7  // This relocation is meaningful only when the machine type is Thumb. The base relocation applies the 32-bit address of a symbol to a consecutive MOVW/MOVT instruction pair.
	BaseRelocTypeRISCVLow12i   BaseRelocType = 7  // This relocation is only meaningful when the machine type is RISC-V. The base relocation applies to the low 12 bits of a 32-bit absolute address formed in RISC-V I-type instruction format.
	BaseRelocTypeRISCVLow12s   BaseRelocType = 8  // This relocation is only meaningful when the machine type is RISC-V. The base relocation applies to the low 12 bits of a 32-bit absolute address formed in RISC-V S-type instruction format.
	BaseRelocTypeMIPSJmpAddr16 BaseRelocType = 9  // The relocation is only meaningful when the machine type is MIPS. The base relocation applies to a MIPS16 jump instruction.
	BaseRelocTypeDir64         BaseRelocType = 10 // The base relocation applies the difference to the 64-bit field at offset.
)

// ~~~ [ Debug ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

//go:generate stringer -trimprefix DebugType -type DebugType

// DebugType specifies the format type of the debug data pointed to by the
// debug data directory.
type DebugType uint32

// Debug data format types.
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#debug-type
const (
	DebugTypeUnknown     DebugType = 0  // An unknown value that is ignored by all tools.
	DebugTypeCOFF        DebugType = 1  // The COFF debug information (line numbers, symbol table, and string table). This type of debug information is also pointed to by fields in the file headers.
	DebugTypeCodeView    DebugType = 2  // The Visual C++ debug information. The format of the data block is described by the CodeView 4.0 specification.
	DebugTypeFPO         DebugType = 3  // The frame pointer omission (FPO) information. This information tells the debugger how to interpret nonstandard stack frames, which use the EBP register for a purpose other than as a frame pointer.
	DebugTypeMisc        DebugType = 4  // The location of DBG file.
	DebugTypeException   DebugType = 5  // Exception information. A copy of .pdata section.
	DebugTypeFixup       DebugType = 6  // Reserved. Fixup information.
	DebugTypeOMapToSrc   DebugType = 7  // The mapping from an RVA in image to an RVA in source image.
	DebugTypeOMapFromSrc DebugType = 8  // The mapping from an RVA in source image to an RVA in image.
	DebugTypeBorland     DebugType = 9  // Reserved for Borland.
	DebugTypeReserved10  DebugType = 10 // Reserved.
	DebugTypeCLSID       DebugType = 11 // Reserved.
	DebugTypeRepro       DebugType = 16 // PE determinism or reproducibility.
)

//go:generate stringer -trimprefix FrameType -type FrameType

// FrameType specifies the frame type of a function.
type FrameType uint8

// Frame types.
//
// ref: https://docs.microsoft.com/en-us/windows/desktop/api/winnt/ns-winnt-fpo_data#members
const (
	// FPO frame
	FrameTypeFPO FrameType = 0
	// Non-FPO frame
	FrameTypeNonFPO FrameType = 3
	// Trap frame
	FrameTypeTrap FrameType = 1
	// TSS frame
	FrameTypeTSS FrameType = 2
)
