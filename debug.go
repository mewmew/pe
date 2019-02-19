package pe

import (
	"time"

	"github.com/mewmew/pe/enum"
)

// --- [ Debug ] ---------------------------------------------------------------

// DebugData contains the contents of a debug data directory.
//
// DebugData is one of the following types.
//
//    *DebugCodeView
//    *DebugFPO
//    *DebugMisc
type DebugData interface {
	// dbgDir returns the debug data directory of the debug data.
	dbgDir() DebugDirectory
}

// DebugDirectory is a debug data directory.
type DebugDirectory struct {
	// Reserved.
	Characteristics uint32
	// Debug data creation time.
	Date time.Time
	// Major debug data format version.
	MajorVer uint16
	// Minor debug data format version.
	MinorVer uint16
	// Debug data format type.
	Type enum.DebugType
	// Debug data size in bytes.
	Size uint32
	// Relative address to debug data (relative to image base).
	RelAddr uint32
	// File offset of debug data.
	Offset uint32
}

// ~~~ [ CodeView ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DebugCodeView contains the contents of a CodeView debug data directory.
type DebugCodeView struct {
	// Debug data directory.
	DbgDir DebugDirectory
	// CodeView debug information.
	CodeViewInfo
}

// dbgDir returns the debug data directory of the CodeView debug data.
func (dbg *DebugCodeView) dbgDir() DebugDirectory {
	return dbg.DbgDir
}

// CodeViewInfo contains CodeView debug information.
type CodeViewInfo struct {
	// CodeView signature ("NB10").
	Signature uint32
	// CodeView offset (set to 0 since debug info is stored in a separate file).
	Offset uint32
	// CodeView debug data creation time.
	Date time.Time
	// Incremental number, initially set to 1 and incremented for each partial
	// write to the PDB file.
	Age uint32
	// Path to PDB file.
	PDBPath string
}

// ~~~ [ FPO ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DebugFPO contains the contents of a FPO debug data directory.
type DebugFPO struct {
	// Debug data directory.
	DbgDir DebugDirectory
	// FPO data.
	FPOData []FPOData
}

// dbgDir returns the debug data directory of the CodeView debug data.
func (dbg *DebugFPO) dbgDir() DebugDirectory {
	return dbg.DbgDir
}

// FPOData represents the stack frame layout for a function on an x86 computer
// when frame pointer omission (FPO) optimization is used. The structure is used
// to locate the base of the call frame.
type FPOData struct {
	// Offset to first byte of function code.
	StartOffset uint32
	// Function size in number of bytes.
	FuncSize uint32
	// Number of local variables.
	NLocals uint64
	// Size of parameters in number of bytes.
	Params uint32
	// Size of function prolog code in number of bytes.
	Prolog uint8
	// Number of registers saved.
	Regs uint8
	// Specifies whether the function uses structured exception handling.
	HasSEH bool
	// Specifies whether the EBP register has been allocated.
	UseBP bool
	// Reserved.
	Reserved uint8
	// Frame type of function.
	Frame enum.FrameType
}

// ~~~ [ Misc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// DebugMisc contains the contents of a miscellaneous debug data directory.
type DebugMisc struct {
	// Debug data directory.
	DbgDir DebugDirectory
	// Contents of miscellaneous debug data directory.
	Content []byte
}

// dbgDir returns the debug data directory of the CodeView debug data.
func (dbg *DebugMisc) dbgDir() DebugDirectory {
	return dbg.DbgDir
}
