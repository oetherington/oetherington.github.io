# Introduction

<span class="todo">NOTE:</span> This article is still a work in progress.

This is a short guide describing the ELF file format for those wishing to build
compilers or assemblers that generate ELF files to be fed into a linker. It is
not a comprehensive guide about every last detail of ELF, but rather the
minimum information needed to generate valid object files on Linux. Having now
written a couple of compilers that generate ELF files, this has proven to be
one of the trickiest (and least interesting) steps, and I have written this
article primarily as a reference for myself in the future should I need to
output ELF files again.

In particular, a number of simplifications and generalisations are made below
about ELF files which, while true for this use case, aren't necessarily true in
general. I also present a much more rigid structure below than is actually
required by the ELF specification; this is to allow the precalculation of
several sizes and offsets. For more accurate and more in depth information, see
the resources section at the bottom of the page.

Many of the sections below make reference to the special precalculated offset
`ELF_BASE`; this is 0x17d for 32 bit code or ??? for 64 bit code.

It's also possible to learn more about the file format by inspecting the output
from more mature compilers, assemblers and linkers with
<code lang="sh">objdump -S example.o</code>,
<code lang="sh">cat example.o | xxd</code> or
<code lang="sh">readelf -a example.o</code>.

An ELF file is split into three parts:
 - The ELF header which contains general information (including the endianness)
 - The section header table
 - The section headers (each corresponding to an entry in the section header
   table)

An executable would also contain a program header table and program headers,
but these are not present in object files to be passed into a linker.

In the tables below, fields with the size <code lang="plaintext">4/8</code> are
4 bytes for 32 bit code and 8 bytes for 64 bit code.

<!--
A reasonably concise implementation of the layout described on this page can be
found in my B compiler, [obc](https://github.com/oetherington/obc).
-->

# ELF Header

First is the ELF header; it consists of the following bytes:

| Size | Value                                                                 |
| ---- | ---------------------------------------------------------------       |
| 4    | 0x464c457f (magic)                                                    |
| 1    | 1 for 32-bit, 2 for 64-bit                                            |
| 1    | 1 for little endian, 2 for big endian                                 |
| 1    | 1 (ELF version)                                                       |
| 1    | OS ABI - Linux generally uses 0x00 which is System V                  |
| 1    | 0 - more info about the ABI (generally not used)                      |
| 7    | 0 (padding)                                                           |
| 2    | File type - 0x01 (`ET_REL`) for object files                          |
| 2    | Architecture (see table below)                                        |
| 4    | 1 (ELF version)                                                       |
| 4/8  | Start of entry point (0x0 - executables are not discussed here)       |
| 4/8  | Start of program header table (0x0)                                   |
| 4/8  | Start of section header table (0x34/0x40 for 32/64 bit)               |
| 4    | 0 (flags)                                                             |
| 2    | Size of this header (0x34/0x40 for 32/64 bit)                         |
| 2    | Size of a single program header table entry (0x38)                    |
| 2    | Number of program header table entries (0x0)                          |
| 2    | Size of a single section header table entry (0x28/0x40 for 32/64 bit) |
| 2    | Number of section header table entries (0x07)                         |
| 2    | Index of the `.shstrtab` section (0x03)                               |

The most common architecture values are as follows:

| Value | Architecture   |
| ----- | -------------- |
| 0x00  | None specified |
| 0x03  | x86            |
| 0x3e  | x86-64         |
| 0x28  | ARM            |
| 0xb7  | ARM (64-bits)  |
| 0xf3  | RISC-V         |

More can be found in the references at the end of the page.

# Section Header Table

As seen in the main ELF header, we will use 7 sections, each with a 0x28 byte
header for 32-bit code or 0x40 byte header for 64-bit code. Each section header
consists of the following bytes:

| Size | Value                                                          |
| ---- | -------------------------------------------------------------- |
| 4    | Offset of the header's name in the `.shstrtab` section         |
| 4    | Header type (detailed for each section below)                  |
| 4/8  | Flags                                                          |
| 4/8  | Virtual address of the section in memory                       |
| 4/8  | Offset of the section in the file image                        |
| 4/8  | Size of the section in bytes (may be 0)                        |
| 4    | Section index of related section (use depends on section type) |
| 4    | Extra section info (use depends on desction type)              |
| 4/8  | Section alignment (must be a power of 2)                       |
| 4/8  | Size in bytes of each entry, or 0 where not applicable         |

## Index 0 Section

The first section is the special index 0 section. Every byte is set to 0 for
our use case.

## Data Section

The second section is the `.data` section which holds static data. The header
is as follows:
 - The shstrtab offset is 1
 - The type is 0x1 (`SHT_PROGBITS`)
 - The flags are 0x3 (`SHF_WRITE | SHF_ALLOC`)
 - The virtual address is 0
 - The offset in the file image is `ELF_BASE`
 - The size of the section must be calculated
 - The index of the related section is 0
 - The extra section info is 0
 - The alignment is 4 for 32 bit code or 8 for 64 bit code
 - The entry size is 0

## Text Section

The third section is the `.text` section which holds the compiled machine code
itself. The header is as follows:
 - The shstrtab offset is 7
 - The type is 1 (`SHT_PROGBITS`)
 - The flags are 0x6 (`SHF_EXECINSTR | SHF_ALLOC`)
 - The virtual address is 0
 - The offset in the file image is `ELF_BASE` plus the size of the data section
 - The size of the section must be calculated
 - The index of the related section is 0
 - The extra section info is 0
 - The alignment is 0x10
 - The entry size is 0

## Shstrtab Section

The fourth section is `.shstrtab` which contains the names of the sections as
strings. Note  that the strings must be NULL-terminated and first byte must
also be 0. The header is as follows:
 - The shstrtab offset is 0xd
 - The type is 0x3 (`SHT_STRTAB`)
 - The flags are 0
 - The virtual address is 0
 - The offset in the file image is 0x14c for 32 bit or ??? for 64 bit
 - The size of the section is 0x31
 - The index of the related section is 0
 - The extra section info is 0
 - The alignment is 0x1
 - The entry size is 0

The layout described in this article has the following data:

```c
const char shstrtab[] = {
	0, '.', 'd', 'a', 't', 'a',
	0, '.', 't', 'e' ,'x', 't',
	0, '.', 's', 'h', 's', 't', 'r', 't', 'a', 'b',
	0, '.', 's', 'y', 'm', 't', 'a', 'b',
	0, '.', 's', 't', 'r', 't', 'a', 'b',
	0, '.', 'r', 'e', 'l', '.', 't', 'e', 'x', 't',
	0,
};
```

## Symtab Section

The fifth section is the symbol table. The header is as follows:
 - The shstrtab offset is 0x17
 - The type is 0x2 (`SHT_SYMTAB`)
 - The flags are 0
 - The virtual address is 0
 - The offset in the file image is `ELF_BASE` plus the size of the data
   section, the text section and the strtab section
 - The size of the section must be calculated
 - The index of the related section is 0x5 (the `.strtab` section which gives
   the symbol names)
 - The extra section info is 0x3 (pointing to this section itself, stating the
   location of the relocation data - in more complex senarios this may be in a
   separate relocation section)
 - The alignment is 0x4
 - The entry size is 0x10

The format of the symbol table entries is described
[below](#symtab-section-entries).

Each entry is 0x10 bytes for 32-bit code or 0x18 bytes for 64-bit code, so the
total size will be this times the number of entries.

## Strtab Section

The sixth section is the string table. It is pointed to by `.symtab` as it's
related section. The header is as follows:
 - The shstrtab offset is 0x1f
 - The type is 0x3 (`SHT_STRTAB`)
 - The flags are 0
 - The virtual address is 0
 - The offset in the file image is `ELF_BASE` plus the size of the data section
   and the text section
 - The size of the section must be calculated
 - The index of the related section is 0
 - The extra section info is 0
 - The alignment is 0x1
 - The entry size is 0

## Relocation Section

The seventh and final section is the relocation table.
 - The shstrtab offset is 0x27
 - The type is 0x9
 - The flags are 0
 - The virtual address is 0
 - The offset in the file image is `ELF_BASE` plus the size of the data
   section, the text section, the strtab section and the symtab section
 - The size of the section is the number of entries times the entry size
 - The index of the related section is 0x4 (the symtab section)
 - The extra section info is 0x2 (the text section)
 - The alignment is 0x4
 - The entry size is 0x8 for 32 bit code or ??? for 64 bit

# Program Header Table

The program header table and program headers are only required for executable
files, so they are not needed for the use case discussed here. Information about
their use can be found in the resources at the bottom of the page.

# Symtab Section Entries

The entries in the symtab have the following structure in a 32-bit ELF file:

| Size | Name  | Value                                                       |
| ---- | ----- | ----------------------------------------------------------- |
| 4    | name  | An index into the strtab section for the section name       |
| 4    | value | The offset from the beginning of the symbol's section       |
| 4    | size  | The size of the symbol                                      |
| 1    | info  | The binding and type in the format given below              |
| 1    | other | For our purposes here, always 0                             |
| 2    | shndx | The index of the section header where the symbol is located |

In a 64-bit ELF file, each of the 4 byte entries becomes 8 bytes and they are
reordered as follows: name, info, other, shndx, value, size.

The first symbol table entry is reserved and must be completely filled with
zeros.

The second symbol is the source file name: the name points into the strtab,
the value is 0, the size is 0, the info (see below) has type 0x4 and binding 0,
and the shndx has the special value `SHN_ABS` which is 0xfff1.

The third symbol is the text section: all fields are 0 except the type which is
0x3 and the shndx which is 0x2.

After this are the symbols defined in the program.

## Symtab Entry Info

The info field holds the symbol's binding and type as
<code lang="c">(binding << 4) + (type & 0xf)</code>. The binding is 0x0 for
static symbols or 0x1 for global symbols. Common symbol types are the following:

| Value | Name                                      |
| ----- | ----------------------------------------- |
| 0x1   | A data object such as a variable or array |
| 0x2   | Executable code such as a function        |
| 0x3   | An ELF section                            |
| 0x4   | The source file name                      |

# Relocation Section Entries

<span class="todo">TODO</span>

# Further Resources

https://en.wikipedia.org/wiki/Executable_and_Linkable_Format

https://cirosantilli.com/elf-hello-world

http://www.sco.com/developers/gabi/2003-12-17/contents.html

https://www.conradk.com/codebase/2017/05/28/elf-from-scratch/

http://www.sunshine2k.de/coding/javascript/onlineelfviewer/onlineelfviewer.html

https://docs.oracle.com/cd/E19683-01/817-3677/chapter6-62988/index.html

https://github.com/torvalds/linux/blob/master/include/uapi/linux/elf.h
