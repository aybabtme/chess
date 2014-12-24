package board

import (
	"os"
)

func ExampleBoard() {
	b := NewBoard()
	os.Stdout.Write(b.DrawANSI())
	// Output:
	// _ a b c d e f g h
	// 8[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m8
	// 7[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m7
	// 6[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m6
	// 5[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m5
	// 4[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m4
	// 3[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m3
	// 2[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m2
	// 1[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m[38;5;16;48;5;59m  [0;00m[38;5;16;48;5;145m  [0;00m1
	//   a b c d e f g h
}