package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func typeWriter(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
	fmt.Println()
}

func showLoading(duration time.Duration) {
	chars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	start := time.Now()
	
	for time.Since(start) < duration {
		for _, char := range chars {
			fmt.Printf("\r%s Loading... ", char)
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println()
}

func showConfetti() {
	colors := []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m"}
	symbols := []string{"*", "✨", "🎉", "🎊", "💫", "⭐"}
	reset := "\033[0m"
	
	for i := 0; i < 20; i++ {
		clearScreen()
		fmt.Println("\n\n")
		
		for j := 0; j < 10; j++ {
			x := rand.Intn(80)
			y := rand.Intn(20)
			color := colors[rand.Intn(len(colors))]
			symbol := symbols[rand.Intn(len(symbols))]
			
			fmt.Printf("\033[%d;%dH%s%s%s", y, x, color, symbol, reset)
		}
		
		time.Sleep(200 * time.Millisecond)
	}
}

// Animasi efek sparkle
func showSparkleEffect() {
	clearScreen()
	
	sparkles := []string{"✨", "💫", "⭐", "🌟", "💎", "🔥"}
	colors := []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m", "\033[37m"}
	reset := "\033[0m"
	
	for i := 0; i < 15; i++ {
		clearScreen()
		fmt.Println("\n\n")
		
		// Menampilkan sparkle di posisi acak
		for j := 0; j < 15; j++ {
			x := rand.Intn(80)
			y := rand.Intn(20)
			color := colors[rand.Intn(len(colors))]
			sparkle := sparkles[rand.Intn(len(sparkles))]
			
			fmt.Printf("\033[%d;%dH%s%s%s", y, x, color, sparkle, reset)
		}
		
		time.Sleep(150 * time.Millisecond)
	}
}

// Animasi efek hujan emoji
func showEmojiRain() {
	clearScreen()
	
	emojis := []string{"🎉", "🎊", "🎈", "🎂", "🎁", "💝", "💖", "💕", "💗", "💓"}
	colors := []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m"}
	reset := "\033[0m"
	
	for i := 0; i < 25; i++ {
		clearScreen()
		fmt.Println("\n\n")
		
		// Menampilkan emoji yang jatuh
		for j := 0; j < 8; j++ {
			y := i - j*3
			if y > 0 && y < 20 {
				x := rand.Intn(80)
				color := colors[rand.Intn(len(colors))]
				emoji := emojis[rand.Intn(len(emojis))]
				
				fmt.Printf("\033[%d;%dH%s%s%s", y, x, color, emoji, reset)
			}
		}
		
		time.Sleep(200 * time.Millisecond)
	}
}

// Animasi efek heartbeat
func showHeartbeat() {
	clearScreen()
	
	hearts := []string{
		"  💖💖  💖💖  ",
		"💖💖💖💖💖💖💖💖",
		"💖💖💖💖💖💖💖💖",
		" 💖💖💖💖💖💖 ",
		"  💖💖💖💖  ",
		"   💖💖   ",
		"    💖    ",
	}
	
	for i := 0; i < 5; i++ {
		clearScreen()
		fmt.Println("\n\n")
		
		// Menampilkan jantung yang berdetak
		for j, line := range hearts {
			fmt.Printf("\033[%d;%dH%s", 10+j, 30, line)
		}
		
		time.Sleep(500 * time.Millisecond)
		
		clearScreen()
		fmt.Println("\n\n")
		
		// Jantung yang lebih besar
		for j, line := range hearts {
			fmt.Printf("\033[%d;%dH%s", 9+j, 28, line)
		}
		
		time.Sleep(300 * time.Millisecond)
	}
}

// Animasi efek fireworks
func showFireworks() {
	clearScreen()
	
	colors := []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m"}
	symbols := []string{"*", "✨", "💥", "🎆", "🎇", "🔥"}
	reset := "\033[0m"
	
	for i := 0; i < 8; i++ {
		clearScreen()
		fmt.Println("\n\n")
		
		// Menampilkan kembang api di posisi acak
		centerX := rand.Intn(60) + 10
		centerY := rand.Intn(15) + 5
		
		for j := 0; j < 20; j++ {
			angle := float64(j) * 18 * 3.14159 / 180
			radius := 3 + rand.Intn(5)
			
			x := centerX + int(float64(radius)*math.Cos(angle))
			y := centerY + int(float64(radius)*math.Sin(angle))
			
			if x >= 0 && x < 80 && y >= 0 && y < 20 {
				color := colors[rand.Intn(len(colors))]
				symbol := symbols[rand.Intn(len(symbols))]
				
				fmt.Printf("\033[%d;%dH%s%s%s", y, x, color, symbol, reset)
			}
		}
		
		time.Sleep(400 * time.Millisecond)
	}
}

func showHDBAnimation() {
	clearScreen()
	
	hdbArt := []string{
		"██╗  ██╗██████╗ ██████╗ ",
		"██║  ██║██╔══██╗██╔══██╗",
		"███████║██║  ██║██████╔╝",
		"██╔══██║██║  ██║██╔══██╗",
		"██║  ██║██████╔╝██║  ██║",
		"╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝",
	}
	
	for _, line := range hdbArt {
		typeWriter(line, 50*time.Millisecond)
	}
	
	time.Sleep(1 * time.Second)
	
	messages := []string{
		"🎉 SELAMAT ULANG TAHUN! 🎉",
		"",
		"✨ Semoga tahun ini membawa:",
		"   💖 Kebahagiaan yang tak terbatas",
		"   🌟 Kesuksesan dalam setiap langkah",
		"   🎯 Impian yang terwujud",
		"   💪 Kesehatan yang prima",
		"   🎊 Dan semua yang kamu impikan!",
		"",
		"🎂 Happy Birthday! 🎂",
	}
	
	for _, msg := range messages {
		if msg == "" {
			fmt.Println()
		} else {
			typeWriter(msg, 100*time.Millisecond)
		}
	}
}

func showCakeAnimation() {
	clearScreen()
	
	cake := []string{
		"    🎂    ",
		"   ╱│╲   ",
		"  ╱ │ ╲  ",
		" ╱  │  ╲ ",
		"╱   │   ╲",
		"███████████",
		"│         │",
		"│  HAPPY  │",
		"│ BIRTHDAY│",
		"│   HDB   │",
		"│         │",
		"███████████",
	}
	
	for _, line := range cake {
		fmt.Println(line)
		time.Sleep(200 * time.Millisecond)
	}
	
	for i := 0; i < 5; i++ {
		fmt.Printf("\033[6;%dH🔥", 5+i*2)
		time.Sleep(300 * time.Millisecond)
	}
	
	time.Sleep(2 * time.Second)
}

func showBalloonAnimation() {
	clearScreen()
	
	balloons := []string{"🎈", "🎈", "🎈", "🎈", "🎈"}
	colors := []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m"}
	reset := "\033[0m"
	
	for i := 0; i < 10; i++ {
		clearScreen()
		fmt.Println("\n\n")
		
		for j, balloon := range balloons {
			y := 10 - i + j
			if y > 0 {
				fmt.Printf("\033[%d;%dH%s%s%s", y, 10+j*3, colors[j], balloon, reset)
			}
		}
		
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	clearScreen()
	fmt.Println("🎉 HDB Birthday Animation 🎉")
	fmt.Println("=============================")
	time.Sleep(1 * time.Second)
	
	showLoading(2 * time.Second)
	
	showHDBAnimation()
	time.Sleep(2 * time.Second)
	
	showCakeAnimation()
	time.Sleep(1 * time.Second)
	
	showBalloonAnimation()
	time.Sleep(1 * time.Second)
	
	showSparkleEffect()
	time.Sleep(1 * time.Second)
	
	showEmojiRain()
	time.Sleep(1 * time.Second)
	
	showHeartbeat()
	time.Sleep(1 * time.Second)
	
	showFireworks()
	time.Sleep(1 * time.Second)
	
	// Menambahkan efek-efek text yang menarik
	showColorChangingText("🎂 HAPPY BIRTHDAY HDB! 🎂", 3*time.Second)
	time.Sleep(1 * time.Second)
	
	showBouncingText("🎉 SELAMAT ULANG TAHUN! 🎉", 3*time.Second)
	time.Sleep(1 * time.Second)
	
	showWaveText("✨ SEMOGA SUKSES SELALU! ✨", 3*time.Second)
	time.Sleep(1 * time.Second)
	
	showBlinkingText("💖 WISHING YOU ALL THE BEST! 💖", 3*time.Second)
	time.Sleep(1 * time.Second)
	
	showConfetti()
	
	clearScreen()
	fmt.Println("\n\n")
	typeWriter("🎊 TERIMA KASIH TELAH MENONTON! 🎊", 100*time.Millisecond)
	fmt.Println()
	typeWriter("✨ Semoga hari ini dan seterusnya penuh dengan kebahagiaan! ✨", 80*time.Millisecond)
	fmt.Println()
	typeWriter("🎂 Happy Birthday HDB! 🎂", 100*time.Millisecond)
	fmt.Println()
	
	fmt.Print("\nTekan Enter untuk keluar...")
	fmt.Scanln()
} 
