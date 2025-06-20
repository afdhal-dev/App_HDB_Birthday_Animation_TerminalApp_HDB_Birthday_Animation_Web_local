# 📦 Instalasi dan Cara Menjalankan HDB Birthday Animation

## 🔧 Prasyarat

Program ini membutuhkan Go (Golang) versi 1.21 atau lebih baru.

### 📥 Instalasi Go di Windows

1. **Download Go:**

   - Kunjungi https://golang.org/dl/
   - Download versi terbaru untuk Windows (misalnya: `go1.21.0.windows-amd64.msi`)

2. **Install Go:**

   - Double click file `.msi` yang sudah didownload
   - Ikuti wizard instalasi (gunakan default settings)
   - Klik "Install" dan tunggu proses selesai

3. **Verifikasi Instalasi:**
   - Restart Command Prompt
   - Ketik: `go version`
   - Jika berhasil, akan muncul versi Go yang terinstall

### 📥 Instalasi Go di Linux/macOS

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# macOS (dengan Homebrew)
brew install go

# Atau download dari https://golang.org/dl/
```

## 🚀 Cara Menjalankan Program

### Windows:

1. Buka Command Prompt sebagai Administrator
2. Navigasi ke folder HBD:
   ```cmd
   cd "D:\Source Code\HDB\HBD"
   ```
3. Jalankan dengan salah satu cara:

   ```cmd
   # Cara 1: Menggunakan file batch (RECOMMENDED)
   run.bat

   # Cara 2: Langsung dengan Go
   go run main.go effects.go
   ```

### Linux/macOS:

1. Buka Terminal
2. Navigasi ke folder HBD:
   ```bash
   cd /path/to/HBD
   ```
3. Jalankan dengan salah satu cara:

   ```bash
   # Cara 1: Menggunakan shell script
   chmod +x run.sh
   ./run.sh

   # Cara 2: Langsung dengan Go
   go run main.go effects.go
   ```

## 🎯 Fitur Program

Program ini menampilkan animasi ulang tahun yang menarik dengan:

- 🎨 Logo HDB dengan efek ketikan
- 🎂 Kue ulang tahun animasi
- 🎈 Balon terbang
- ✨ Efek sparkle
- 🌧️ Hujan emoji
- 💖 Jantung berdetak
- 🎆 Kembang api
- 🎊 Confetti berwarna-warni
- 💫 Efek teks modern (warna, gelombang, bouncing, blinking)
- 💫 Pesan ucapan yang personal

## 🔍 Troubleshooting

### Error: 'go' is not recognized

- Pastikan Go sudah terinstall dengan benar
- Restart Command Prompt setelah instalasi
- Cek PATH environment variable

### Error: module not found

- Jalankan: `go mod tidy`
- Pastikan berada di folder yang benar

### Animasi tidak muncul dengan baik

- Pastikan terminal mendukung ANSI escape codes
- Gunakan Windows Terminal atau PowerShell untuk hasil terbaik

## 🎊 Selamat Menikmati!

Setelah Go terinstall, program siap untuk dijalankan dan menampilkan animasi ulang tahun yang spektakuler untuk HDB! 🎉

**Tips:** Untuk hasil terbaik, gunakan Windows Terminal atau PowerShell dengan font yang mendukung emoji.
