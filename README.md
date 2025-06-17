# ✨ media.vector

[![Work in Progress](https://img.shields.io/badge/Status-Work%20in%20Progress-yellow)](https://shields.io)
[![Go Report Card](https://goreportcard.com/badge/github.com/SmartMediaFiles/media.vector)](https://goreportcard.com/report/github.com/SmartMediaFiles/media.vector)
[![GoDoc](https://pkg.go.dev/badge/github.com/SmartMediaFiles/media.vector)](https://pkg.go.dev/github.com/SmartMediaFiles/media.vector)
[![Release](https://img.shields.io/github/release/SmartMediaFiles/media.vector.svg?style=flat)](https://github.com/SmartMediaFiles/media.vector/releases)

## Overview

`media.vector` is a specialized library within the **SmartMediaFiles ecosystem**. Its purpose is to define and categorize all file types and extensions related to vector graphic files. This library provides the necessary definitions for the main `@/media` library to recognize and classify vector files like `SVG`, `AI`, etc.

## Features

- **File Type Definitions**: Provides a list of `FileType` constants for each supported vector format.
- **File Extension Definitions**: Provides a list of all corresponding `FileExtension` constants.
- **Type-to-Extension Mapping**: Contains the `VectorFileTypesExtensions` map that links each file type to its extensions, which is then consumed by the central `@/media` registry.
- **Future-proofing**: This library will be expanded to include functions for parsing vector graphics and extracting relevant metadata.

## Installation

```bash
go get -u github.com/smartmediafiles/media.vector
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---

⚠️ **Note:** This README will be updated regularly as the project progresses. Check back often for the latest information!
