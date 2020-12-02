package pe

// --- [ Resource ] ------------------------------------------------------------

// ResourceDirectory is a resource directory.
type ResourceDirectory struct {
	// Reserved.
	Characteristics uint32
	// Resource directory creation time.
	Date time.Time
	// Major resource directory format version.
	MajorVer uint16
	// Minor resource directory format version.
	MinorVer uint16
	// Number of named entries.
	NNamedEntries uint16
	// Number of ID entries.
	NIDEntries uint16

	//Children []ResourceDirectoryEntry
	//  IMAGE_RESOURCE_DIRECTORY_ENTRY DirectoryEntries[];
}

/*
type ResourceDirectoryEntry {
	Name string
	ID uint16

    union {
        struct {
            DWORD NameOffset:31;
            DWORD NameIsString:1;
        };
        DWORD   Name;
        WORD    Id;
    };
    union {
        DWORD   OffsetToData;
        struct {
            DWORD   OffsetToDirectory:31;
            DWORD   DataIsDirectory:1;
        };
    };
}
*/
