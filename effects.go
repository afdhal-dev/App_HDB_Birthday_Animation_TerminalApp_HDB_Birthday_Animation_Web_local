package main

import (
	"fmt"
	"time"
)

// Efek animasi text berkedip
func showBlinkingText(text string, duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
		clearScreen()
		fmt.Println("\n\n")
		fmt.Printf("\033[10;%dH%s", 40-len(text)/2, text)
		time.Sleep(500 * time.Millisecond)
		
		clearScreen()
		fmt.Println("\n\n")
		time.Sleep(200 * time.Millisecond)
	}
}

// Efek animasi text bergerak
func showMovingText(text string, duration time.Duration) {
	start := time.Now()
	position := 0
	direction := 1
	
	for time.Since(start) < duration {
		clearScreen()
		fmt.Println("\n\n")
		
		// Bergerak dari kiri ke kanan dan sebaliknya
		if position >= 60 || position <= 0 {
			direction *= -1
		}
		position += direction
		
		fmt.Printf("\033[10;%dH%s", position, text)
		time.Sleep(100 * time.Millisecond)
	}
}

// Efek animasi text berputar
func showRotatingText(text string, duration time.Duration) {
	start := time.Now()
	angle := 0.0
	
	for time.Since(start) < duration {
		clearScreen()
		fmt.Println("\n\n")
		
		// Simulasi text berputar dengan mengubah posisi
		x := 40 + int(10*float64(angle))
		if x < 0 {
			x = 0
		} else if x > 70 {
			x = 70
		}
		
		fmt.Printf("\033[10;%dH%s", x, text)
		angle += 0.1
		time.Sleep(100 * time.Millisecond)
	}
}

// Efek animasi text dengan warna berubah
func showColorChangingText(text string, duration time.Duration) {
	colors := []string{
		"\033[31m", // Merah
		"\033[32m", // Hijau
		"\033[33m", // Kuning
		"\033[34m", // Biru
		"\033[35m", // Magenta
		"\033[36m", // Cyan
		"\033[37m", // Putih
	}
	reset := "\033[0m"
	
	start := time.Now()
	colorIndex := 0
	
	for time.Since(start) < duration {
		clearScreen()
		fmt.Println("\n\n")
		
		color := colors[colorIndex%len(colors)]
		fmt.Printf("\033[10;%dH%s%s%s", 40-len(text)/2, color, text, reset)
		
		colorIndex++
		time.Sleep(300 * time.Millisecond)
	}
}

// Efek animasi text dengan efek ketikan yang lebih canggih
func showAdvancedTypewriter(text string, delay time.Duration) {
	clearScreen()
	fmt.Println("\n\n")
	
	for i, char := range text {
		fmt.Printf("\033[10;%dH%s", 40-len(text)/2+i, string(char))
		time.Sleep(delay)
	}
	fmt.Println()
}

// Efek animasi text dengan efek fade in
func showFadeInText(text string, duration time.Duration) {
	clearScreen()
	fmt.Println("\n\n")
	
	steps := 10
	stepDuration := duration / time.Duration(steps)
	
	for i := 1; i <= steps; i++ {
		clearScreen()
		fmt.Println("\n\n")
		
		// Simulasi fade in dengan menampilkan sebagian text
		visibleLength := len(text) * i / steps
		visibleText := text[:visibleLength]
		
		fmt.Printf("\033[10;%dH%s", 40-len(text)/2, visibleText)
		time.Sleep(stepDuration)
	}
}

// Efek animasi text dengan efek bounce
func showBouncingText(text string, duration time.Duration) {
	start := time.Now()
	position := 10
	direction := 1
	
	for time.Since(start) < duration {
		clearScreen()
		fmt.Println("\n\n")
		
		// Bergerak naik turun
		if position >= 15 || position <= 5 {
			direction *= -1
		}
		position += direction
		
		fmt.Printf("\033[%d;%dH%s", position, 40-len(text)/2, text)
		time.Sleep(200 * time.Millisecond)
	}
}

// Efek animasi text dengan efek wave
func showWaveText(text string, duration time.Duration) {
	start := time.Now()
	wave := 0.0
	
	for time.Since(start) < duration {
		clearScreen()
		fmt.Println("\n\n")
		
		// Simulasi efek gelombang
		for i, char := range text {
			y := 10 + int(2*float64(wave+float64(i)*0.5))
			if y < 5 {
				y = 5
			} else if y > 15 {
				y = 15
			}
			
			fmt.Printf("\033[%d;%dH%s", y, 40-len(text)/2+i, string(char))
		}
		
		wave += 0.2
		time.Sleep(100 * time.Millisecond)
	}
} 