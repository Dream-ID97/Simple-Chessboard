package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Piece merepresentasikan bidak catur
type Piece struct {
	Type  string // Tipe bidak (P, R, N, B, Q, K)
	Color string // Warna bidak (W untuk putih, B untuk hitam)
}

// Board merepresentasikan papan catur
type Board [8][8]*Piece

// Inisialisasi papan catur
func initializeBoard() Board {
	var board Board

	// Bidak putih
	board[0] = [8]*Piece{
		{"R", "W"}, {"N", "W"}, {"B", "W"}, {"Q", "W"},
		{"K", "W"}, {"B", "W"}, {"N", "W"}, {"R", "W"},
	}
	for i := 0; i < 8; i++ {
		board[1][i] = &Piece{"P", "W"}
	}

	// Bidak hitam
	board[7] = [8]*Piece{
		{"R", "B"}, {"N", "B"}, {"B", "B"}, {"Q", "B"},
		{"K", "B"}, {"B", "B"}, {"N", "B"}, {"R", "B"},
	}
	for i := 0; i < 8; i++ {
		board[6][i] = &Piece{"P", "B"}
	}

	return board
}

// Clear terminal
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Tampilkan papan catur dalam ASCII
func (b Board) display() {
	clearScreen()
	fmt.Println("   a b c d e f g h")
	for i, row := range b {
		fmt.Printf("%d ", 8-i)
		for _, piece := range row {
			if piece == nil {
				fmt.Print(". ")
			} else {
				fmt.Printf("%s ", piece.Type)
			}
		}
		fmt.Printf("%d\n", 8-i)
	}
	fmt.Println("   a b c d e f g h")
}

// Validasi langkah bidak
func isValidMove(piece *Piece, startRow, startCol, endRow, endCol int, board Board) bool {
	if endRow < 0 || endRow >= 8 || endCol < 0 || endCol >= 8 {
		return false
	}
	if board[endRow][endCol] != nil && board[endRow][endCol].Color == piece.Color {
		return false
	}
	switch piece.Type {
	case "P":
		if piece.Color == "W" {
			return startCol == endCol && endRow == startRow-1 && board[endRow][endCol] == nil
		}
		return startCol == endCol && endRow == startRow+1 && board[endRow][endCol] == nil
	case "R":
		return startRow == endRow || startCol == endCol
	case "N":
		return (abs(startRow-endRow) == 2 && abs(startCol-endCol) == 1) ||
			(abs(startRow-endRow) == 1 && abs(startCol-endCol) == 2)
	case "B":
		return abs(startRow-endRow) == abs(startCol-endCol)
	case "Q":
		return startRow == endRow || startCol == endCol || abs(startRow-endRow) == abs(startCol-endCol)
	case "K":
		return abs(startRow-endRow) <= 1 && abs(startCol-endCol) <= 1
	}
	return false
}

// Gerakkan bidak
func (b *Board) movePiece(startRow, startCol, endRow, endCol int) {
	b[endRow][endCol] = b[startRow][startCol]
	b[startRow][startCol] = nil
}

// Konversi koordinat dari input pengguna
func parseInput(input string) (int, int, int, int, error) {
	if len(input) != 5 || input[2] != '-' {
		return 0, 0, 0, 0, fmt.Errorf("format salah")
	}
	startCol := int(input[0] - 'a')
	startRow := 8 - int(input[1]-'0')
	endCol := int(input[3] - 'a')
	endRow := 8 - int(input[4]-'0')
	if startRow < 0 || startRow >= 8 || startCol < 0 || startCol >= 8 || endRow < 0 || endRow >= 8 || endCol < 0 || endCol >= 8 {
		return 0, 0, 0, 0, fmt.Errorf("koordinat di luar batas")
	}
	return startRow, startCol, endRow, endCol, nil
}

// Fungsi utama permainan
func main() {
	board := initializeBoard()
	reader := bufio.NewReader(os.Stdin)

	for {
		board.display()
		fmt.Print("Masukkan langkah (contoh: e2-e4): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		startRow, startCol, endRow, endCol, err := parseInput(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		piece := board[startRow][startCol]
		if piece == nil {
			fmt.Println("Tidak ada bidak di posisi awal.")
			continue
		}

		if !isValidMove(piece, startRow, startCol, endRow, endCol, board) {
			fmt.Println("Langkah tidak valid.")
			continue
		}

		board.movePiece(startRow, startCol, endRow, endCol)
	}
}

// Fungsi pembantu
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}