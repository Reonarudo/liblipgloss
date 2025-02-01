# C Wrapper for Lipgloss

![Lipgloss](https://github.com/charmbracelet/lipgloss/raw/main/docs/lipgloss.png)

## Overview
This is a **C wrapper** around the [Lipgloss](https://github.com/charmbracelet/lipgloss) Go library, which provides beautiful ANSI styling for terminal applications. This wrapper enables **C applications** to leverage Lipgloss's powerful styling capabilities.

## Features
- **Text Styling**: Apply bold, italic, underline, and other styles.
- **Color Management**: Supports foreground, background, and ANSI color profiles.
- **Border Support**: Add stylish borders around text elements.
- **Alignment & Layout**: Position elements using horizontal/vertical alignment.
- **Memory Management**: Ensures proper allocation and cleanup for C integration.

## Getting Started

### 1. Clone the Repository
```sh
git clone https://github.com/Reonarudo/liblipgloss.git
cd liblipgloss
```

### 2. Install Dependencies
Ensure Go is installed and fetch the required modules:
```sh
go mod tidy
```

### 3. Build the Shared Library
```sh
make
```
This compiles the Go files and generates `liblipgloss.dylib`, which can be linked from C applications.

### 4. Run Tests
```sh
make test
```
This compiles and executes test cases to verify the wrapper's functionality.

## Usage
### Linking with C Applications
To use the library in a C project, include `CLipgloss.h` and link against `liblipgloss.dylib`:
```c
#include "CLipgloss.h"

int main() {
    DefaultRenderer();
    uint64_t style = NewStyle();
    style = StyleBold(style, 1);
    char* rendered = StyleRender(style, "Hello, Lipgloss!");
    printf("%s\n", rendered);
    FreeString(rendered);
    FreeStyle(style);
    return 0;
}
```

## Credits

LipglossSwift is a Swift wrapper around [Lipgloss](https://github.com/charmbracelet/lipgloss), created by [Charm](https://charm.sh). All credit for the underlying styling engine goes to the Lipgloss team:

- [All Lipgloss contributors](https://github.com/charmbracelet/lipgloss/graphs/contributors)

## License

This project is licensed under MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## Support

If you encounter any issues or have questions:

1. Check the [Issues](YOUR_REPOSITORY_URL/issues) section
2. For Lipgloss-specific questions, refer to the [Lipgloss documentation](https://github.com/charmbracelet/lipgloss)
3. Open a new issue if needed

---

Happy styling! 🎨✨