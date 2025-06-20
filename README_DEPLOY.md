# 🎂 HDB Birthday Web App - Deployment Guide

## 🌐 Cara Deploy ke Cloud (Dapat URL Publik)

### **Opsi 1: Railway (Paling Mudah & Gratis)**

1. **Daftar di Railway**

   - Kunjungi [railway.app](https://railway.app)
   - Sign up dengan GitHub

2. **Deploy Project**

   - Klik "New Project"
   - Pilih "Deploy from GitHub repo"
   - Pilih repository kamu
   - Railway akan otomatis detect Go project

3. **Dapat URL**
   - Setelah deploy selesai, kamu akan dapat URL seperti:
   - `https://hdb-birthday-production.up.railway.app`

### **Opsi 2: Vercel (Gratis)**

1. **Daftar di Vercel**

   - Kunjungi [vercel.com](https://vercel.com)
   - Sign up dengan GitHub

2. **Deploy**

   - Import repository dari GitHub
   - Vercel akan otomatis build dan deploy

3. **Dapat URL**
   - URL akan seperti: `https://hdb-birthday.vercel.app`

### **Opsi 3: Render (Gratis)**

1. **Daftar di Render**

   - Kunjungi [render.com](https://render.com)
   - Sign up dengan GitHub

2. **Deploy**
   - Create "Web Service"
   - Connect GitHub repository
   - Build Command: `go build -o main web_main.go`
   - Start Command: `./main`

### **Opsi 4: Heroku (Gratis dengan batasan)**

1. **Install Heroku CLI**
2. **Login dan deploy:**
   ```bash
   heroku login
   heroku create hdb-birthday-app
   git push heroku main
   ```

## 🚀 Cara Test Lokal

```bash
# Jalankan web server lokal
go run web_main.go

# Buka browser ke: http://localhost:8080
```

## 📱 Akses dari Device Lain

Jika ingin akses dari HP/device lain di jaringan yang sama:

```bash
# Cari IP address komputer kamu
ipconfig  # Windows
ifconfig  # Mac/Linux

# Akses dari device lain: http://IP_ADDRESS:8080
# Contoh: http://192.168.1.100:8080
```

## 🔧 Troubleshooting

### Port 8080 sudah digunakan?

```bash
# Ganti port di web_main.go
log.Fatal(http.ListenAndServe(":3000", nil))
```

### Error build?

```bash
# Pastikan Go terinstall
go version

# Update dependencies
go mod tidy
```

## 🎯 Tips

1. **Railway** adalah pilihan terbaik untuk pemula
2. **Vercel** bagus untuk performance
3. **Render** bagus untuk reliability
4. Semua platform di atas punya **free tier** yang cukup untuk project kecil

## 📞 Butuh Bantuan?

Jika ada error atau butuh bantuan lebih lanjut, share error message-nya!
