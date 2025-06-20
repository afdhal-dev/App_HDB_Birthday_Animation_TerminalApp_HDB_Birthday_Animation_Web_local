package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	// Handle main page
	http.HandleFunc("/", handleHome)
	
	// Handle API endpoints for animations
	http.HandleFunc("/api/animation", handleAnimation)
	
	fmt.Println("🎉 HDB Birthday Web Server Starting...")
	fmt.Println("🌐 Server running at: http://localhost:8080")
	fmt.Println("📱 Access from any device on your network!")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>🎂 Happy Birthday HDB! 🎂</title>
    <style>
        body {
            margin: 0;
            padding: 0;
            background: linear-gradient(45deg, #ff6b6b, #4ecdc4, #45b7d1, #96ceb4, #feca57);
            background-size: 400% 400%;
            animation: gradientShift 3s ease infinite;
            font-family: 'Arial', sans-serif;
            overflow: hidden;
            height: 100vh;
        }
        
        @keyframes gradientShift {
            0% { background-position: 0% 50%; }
            50% { background-position: 100% 50%; }
            100% { background-position: 0% 50%; }
        }
        
        .container {
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            height: 100vh;
            text-align: center;
        }
        
        .hdb-logo {
            font-family: monospace;
            font-size: 1.5em;
            color: white;
            text-shadow: 2px 2px 4px rgba(0,0,0,0.5);
            margin-bottom: 20px;
            animation: glow 2s ease-in-out infinite alternate;
        }
        
        @keyframes glow {
            from { text-shadow: 0 0 5px #fff, 0 0 10px #fff, 0 0 15px #e60073; }
            to { text-shadow: 0 0 10px #fff, 0 0 20px #ff4da6, 0 0 30px #ff4da6; }
        }
        
        .message {
            font-size: 1.5em;
            color: white;
            text-shadow: 2px 2px 4px rgba(0,0,0,0.5);
            margin: 10px 0;
            opacity: 0;
            animation: fadeInUp 1s ease forwards;
        }
        
        @keyframes fadeInUp {
            from {
                opacity: 0;
                transform: translateY(30px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }
        
        .cake {
            font-size: 3em;
            margin: 20px 0;
            animation: bounce 2s infinite;
        }
        
        @keyframes bounce {
            0%, 20%, 50%, 80%, 100% { transform: translateY(0); }
            40% { transform: translateY(-30px); }
            60% { transform: translateY(-15px); }
        }
        
        .balloons {
            position: absolute;
            font-size: 2em;
            animation: float 3s ease-in-out infinite;
        }
        
        @keyframes float {
            0%, 100% { transform: translateY(0px); }
            50% { transform: translateY(-20px); }
        }
        
        .sparkle {
            position: absolute;
            font-size: 1.5em;
            animation: sparkle 1.5s ease-in-out infinite;
        }
        
        @keyframes sparkle {
            0%, 100% { opacity: 0; transform: scale(0); }
            50% { opacity: 1; transform: scale(1); }
        }
        
        .firework {
            position: absolute;
            font-size: 2em;
            animation: explode 2s ease-out infinite;
        }
        
        @keyframes explode {
            0% { transform: scale(0); opacity: 1; }
            100% { transform: scale(3); opacity: 0; }
        }
        
        .btn {
            background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
            border: none;
            color: white;
            padding: 15px 30px;
            font-size: 1.2em;
            border-radius: 25px;
            cursor: pointer;
            margin: 10px;
            transition: transform 0.3s ease;
            box-shadow: 0 4px 15px rgba(0,0,0,0.2);
        }
        
        .btn:hover {
            transform: translateY(-3px);
            box-shadow: 0 6px 20px rgba(0,0,0,0.3);
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="hdb-logo">
██╗  ██╗██████╗ ██████╗ <br>
██║  ██║██╔══██╗██╔══██╗<br>
███████║██║  ██║██████╔╝<br>
██╔══██║██║  ██║██╔══██╗<br>
██║  ██║██████╔╝██║  ██║<br>
╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝
        </div>
        
        <div class="message" style="animation-delay: 0.5s;">🎉 SELAMAT ULANG TAHUN! 🎉</div>
        <div class="message" style="animation-delay: 1s;">✨ Semoga tahun ini membawa kebahagiaan yang tak terbatas ✨</div>
        <div class="message" style="animation-delay: 1.5s;">🌟 Kesuksesan dalam setiap langkah 🌟</div>
        <div class="message" style="animation-delay: 2s;">🎯 Impian yang terwujud 🎯</div>
        <div class="message" style="animation-delay: 2.5s;">💪 Kesehatan yang prima 💪</div>
        <div class="message" style="animation-delay: 3s;">🎊 Dan semua yang kamu impikan! 🎊</div>
        
        <div class="cake">🎂</div>
        
        <div class="message" style="animation-delay: 3.5s;">🎂 Happy Birthday! 🎂</div>
        
        <button class="btn" onclick="startFireworks()">🎆 Start Fireworks! 🎆</button>
        <button class="btn" onclick="startBalloons()">🎈 Release Balloons! 🎈</button>
        <button class="btn" onclick="startSparkles()">✨ Show Sparkles! ✨</button>
    </div>
    
    <script>
        // Create floating balloons
        function createBalloon() {
            const balloon = document.createElement('div');
            balloon.className = 'balloons';
            balloon.innerHTML = '🎈';
            balloon.style.left = Math.random() * window.innerWidth + 'px';
            balloon.style.top = window.innerHeight + 'px';
            balloon.style.animationDelay = Math.random() * 2 + 's';
            document.body.appendChild(balloon);
            
            setTimeout(() => {
                balloon.remove();
            }, 5000);
        }
        
        function startBalloons() {
            for(let i = 0; i < 10; i++) {
                setTimeout(createBalloon, i * 200);
            }
        }
        
        // Create sparkles
        function createSparkle() {
            const sparkle = document.createElement('div');
            sparkle.className = 'sparkle';
            sparkle.innerHTML = ['✨', '💫', '⭐', '🌟', '💎'][Math.floor(Math.random() * 5)];
            sparkle.style.left = Math.random() * window.innerWidth + 'px';
            sparkle.style.top = Math.random() * window.innerHeight + 'px';
            document.body.appendChild(sparkle);
            
            setTimeout(() => {
                sparkle.remove();
            }, 2000);
        }
        
        function startSparkles() {
            for(let i = 0; i < 20; i++) {
                setTimeout(createSparkle, i * 100);
            }
        }
        
        // Create fireworks
        function createFirework() {
            const firework = document.createElement('div');
            firework.className = 'firework';
            firework.innerHTML = ['💥', '🎆', '🎇', '🔥', '✨'][Math.floor(Math.random() * 5)];
            firework.style.left = Math.random() * window.innerWidth + 'px';
            firework.style.top = Math.random() * window.innerHeight + 'px';
            document.body.appendChild(firework);
            
            setTimeout(() => {
                firework.remove();
            }, 2000);
        }
        
        function startFireworks() {
            for(let i = 0; i < 15; i++) {
                setTimeout(createFirework, i * 150);
            }
        }
        
        // Auto start some animations
        setTimeout(startBalloons, 4000);
        setTimeout(startSparkles, 6000);
        setTimeout(startFireworks, 8000);
        
        // Add some random sparkles periodically
        setInterval(() => {
            if(Math.random() > 0.7) {
                createSparkle();
            }
        }, 1000);
    </script>
</body>
</html>`
	
	t, err := template.New("home").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, nil)
}

func handleAnimation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	response := map[string]interface{}{
		"message": "🎉 Animation triggered! 🎉",
		"timestamp": time.Now().Format(time.RFC3339),
		"status": "success",
	}
	
	fmt.Fprintf(w, `{"message":"%s","timestamp":"%s","status":"%s"}`, 
		response["message"], response["timestamp"], response["status"])
} 
