@echo off
echo 🎉 HDB Birthday Web Server
echo =========================
echo.
echo Starting web server...
echo.
echo 🌐 Local URL: http://localhost:8080
echo 📱 Network URL: http://%COMPUTERNAME%:8080
echo.
echo Press Ctrl+C to stop the server
echo.
go run web_main.go
pause 