## General Rules
- Short and Clear: Package names should be short and convey the purpose of the package.
- Lowercase Only: Use lowercase letters only; avoid mixedCaps or underscores.
- No Plurals: Stick to singular forms to avoid ambiguity.
- Reflect Purpose: The name should give a hint about the package's functionality.


## Single vs. Multiple Words
- Prefer Single Words: Use single-word names whenever possible.
- Multiple Words: If a package name consists of multiple words, use underscores to separate them, though this is less common and generally discouraged.

## File Naming
- Naming Convention: Use lowercase letters with underscores to separate words. Avoid starting filenames with "." or "_".
- Test Files: Append _test.go to test files.
- Platform-Specific Files: Include the platform name in the filename for OS/architecture-specific files, e.g., filename_linux.go.

## Folders Naming
- Naming Convention: Use lowercase letters with underscores to separate words. Avoid starting filenames with "." or "_".
- Test Files: Append _test.go to test files.
- Platform-Specific Files: Include the platform name in the filename for OS/architecture-specific files, e.g., filename_linux.go.

## Example good pkg name
strconv (string conversion)
syscall (system call)
fmt (formatted I/O)

## Example bad pkg name
string_conversion
stringConversion