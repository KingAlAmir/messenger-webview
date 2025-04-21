# Messenger Lite (Webview Version)

A minimal desktop wrapper for Messenger.com using Go and [webview_go](https://github.com/webview/webview_go), which provides a lightweight native WebView window.

This project embeds the Messenger web interface (`https://www.messenger.com/t/`) inside a minimal native window, optionally with a custom icon

## 🚀 Features

- Loads Messenger in a lightweight webview.
- Native look and feel (no browser UI).
- Small binary size.
- No telemetry or tracking.
- Customizable icon support.

## 🛠 Getting Started

These instructions will get you a copy of the project up and running on your local machine.

### Prerequisites

- [Go](https://go.dev/dl/) 1.18+
- C compiler (MinGW-w64 for Windows)
- Git
- Windows system (64-bit)
- Internet access

### 🧱 Build Instructions

```bash
# 1. Create your project directory
mkdir messenger-webview && cd messenger-webview

# Or clone this repo

git clone https://github.com/KingAlAmir/messenger-webview.git
cd messenger-webview

# 2. Initialize Go module
go mod init example.com/messenger

> Skip this if `go.mod` already exists.

# 3. Download basic example file
curl -sSLo main.go "https://raw.githubusercontent.com/webview/webview_go/master/examples/basic/main.go"

# 4. Get the webview_go dependency
go get github.com/webview/webview_go

# 5. Build the project
go build -tags=windows -ldflags="-H windowsgui" -o dist/messenger-lite.exe

> This will produce `messenger-lite.exe` inside the `dist/` folder with the custom icon if `rsrc` was used.
```

### 💡 Notes

- If you get `mm_malloc.h: No such file or directory`, your MinGW is missing required headers. Use [MSYS2](https://www.msys2.org/) with `mingw-w64-x86_64-toolchain`.
- 32-bit GCC won’t work. Install **64-bit** version of GCC (`x86_64-w64-mingw32-gcc`).

### 🖼 Add Icon (Optional)

Make sure you have `icon.ico` in the `assets/` directory.

Then generate a `.syso` file using [`rsrc`](https://github.com/akavel/rsrc):

```bash
rsrc -ico assets/icon.ico -o dist/rsrc_windows_amd64.syso
```


## 🧪 Running the App

Simply double-click `dist/messenger-lite.exe` or run:

```bash
./dist/messenger-lite.exe
```

## 📦 Project Structure

```
messenger-webview/
├── assets/
│   └── icon.ico                # Your app icon (used during build)
├── dist/
│   ├── messenger-lite.exe      # Final built Windows executable
│   └── rsrc_windows_amd64.syso # Compiled resource file (icon)
├── go.mod
├── go.sum
├── main.go                     # Go code to launch WebView
├── README.md
└── DISCLAIMER.md
```

## Contributions

Contributions are welcome! Please fork the repository and submit a pull request with your enhancements.

# 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.

# 📢 Disclaimer

Please read the [DISCLAIMER](disclaimer.md) before using this app.


