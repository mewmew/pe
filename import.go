package pe

import "time"

// ImportDirectory is an import data directory.
type ImportDirectory struct {
	// (optional) Relative address of import name table (INT); zero if not
	// present.
	INTRelAddr uint32
	// DLL creation time (set by the dynmaic linker).
	Date time.Time
	// Index of the first forwarder reference.
	ForwardChain uint32
	// DLL name.
	Name string
	// Relative address of import address table (IAT).
	IATRelAddr uint32
}

// ImportEntry contains the contents of an import entry.
type ImportEntry struct {
	// Import data directory.
	ImpDir ImportDirectory
	// Import name table entries.
	INTs []INTEntry
	// Import address table entries.
	IATs []INTEntry
}

// INTEntry is an import name table entry of a PE file.
//
// Import address table (IAT) entries are identical to INT entries prior to
// dynamic linking. After dynamic linking, IAT entries contain the address of
// the symbol being imported.
type INTEntry struct {
	// Specifies whether to import by ordinal or name.
	IsOrdinal bool
	// Ordinal number (used if IsOrdinal is set).
	Ordinal uint16
	// Name entry (used if IsOrdinal is clear).
	NameEntry NameEntry
}

// NameEntry is a name table entry.
type NameEntry struct {
	// Index into the export name table. Used as a hint to try and locate the
	// entry; if not successful, binary search of the DLL's export name table.
	Hint uint16
	// Name of the entry.
	Name string
}
