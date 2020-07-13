package eval

import (
	"fmt"
	"testing"
)

/*
func TestIsFlush(t *testing.T) {
	tables := []struct {
		description string
		c1          Card
		c2          Card
		c3          Card
		c4          Card
		c5          Card
		want        bool
	}{
		{
			"FlushSpades",
			[]byte("9S"), // 8394515
			[]byte("QS"), 	// 67115551
			[]byte("KS"), 	// 134224677
			[]byte("2S"), // 69634
			[]byte("AS"), // 268442665
			true,
		},
		{
			"FlushHearts",
			[]byte("JH"), // 33564957
			[]byte("QH"), // 67119647
			1057803,
			270853,
			[]byte("8H"), // 4204049
			true,
		},
		{
			"FlushDiamonds",
			134236965,
			[]byte("4D"), // 279045
			[]byte("2D"), // 81922
			[]byte("AD"), // 268454953
			[]byte("6D"), // 1065995
			true,
		},
		{
			"FlushClubs",
			[]byte("KC"), // 134253349
			16812055,
			164099,
			[]byte("JC"), // 33589533
			98306,
			true,
		},
		{
			"NotFlush",
			[]byte("KC"), // 134253349
			16812055,
			[]byte("9D"), // 8406803
			[]byte("JC"), // 33589533
			98306,
			false,
		},

		{
			"SixHighStraightFlushSpades",
			[]byte("2S"), // 69634
			[]byte("3S"), // 135427
			[]byte("4S"), // 266757
			[]byte("5S"), // 529159
			[]byte("6S"), // 1053707
			true,
		},
	}

	for _, table := range tables {
		testname := fmt.Sprintf("%s", table.description)
		t.Run(testname, func(t *testing.T) {
			result := isFlush(table.c1, table.c2, table.c3, table.c4, table.c5)
			if result != table.want {
				t.Errorf("\nFAIL:\nIn: %d, %d, %d, %d, %d \nWant: %t \nGot: %t \n", table.c1, table.c2, table.c3, table.c4, table.c5, table.want, result)
			}
		})
	}
}

*/

/*
func TestFlushValue(t *testing.T) {
	tables := []struct {
		description string
		c1          Card
		c2          Card
		c3          Card
		c4          Card
		c5          Card
		want        int16
	}{
		{
			"SixHighStraightFlushSpades",
			[]byte("2S"), // 69634
			[]byte("3S"), // 135427
			[]byte("4S"), // 266757
			[]byte("5S"), // 529159
			[]byte("6S"), // 1053707
			9,
		},
		{
			"AKQJ9FlushHearts",
			[]byte("9H"), // 8398611
			[]byte("JH"), // 33564957
			[]byte("QH"), // 67119647
			[]byte("KH"), // 134228773
			[]byte("AH"), // 268446761
			323,
		},
		{
			"AKQJ9FlushClubs",
			[]byte("9C"), // 8423187
			[]byte("JC"), // 33589533
			[]byte("QC"), // 67144223,
			[]byte("KC"), // 134253349
			[]byte("AC"), // 268471337
			323,
		},
	}

	for _, table := range tables {
		testname := fmt.Sprintf("%s", table.description)
		t.Run(testname, func(t *testing.T) {
			result := flushValue(table.c1, table.c2, table.c3, table.c4, table.c5)
			if result != table.want {
				t.Errorf("\nFAIL:\nIn: %d, %d, %d, %d, %d \nWant: %d \nGot: %d \n", table.c1, table.c2, table.c3, table.c4, table.c5, table.want, result)
			}
		})
	}
}

*/

/*
func TestUniqueValue(t *testing.T) {
	tables := []struct {
		description string
		c1          Card
		c2          Card
		c3          Card
		c4          Card
		c5          Card
		want        int16
	}{
		{
			"NotUniqueSameSuit",
			[]byte("2S"), // 69634
			[]byte("3S"), // 135427
			[]byte("4S"), // 266757
			[]byte("6S"), // 1053707
			[]byte("6S"), // 1053707
			0,
		},
		{
			"NotUniqueDifferentSuits",
			[]byte("9H"), // 8398611
			[]byte("JH"), // 33564957
			[]byte("QH"), // 67119647
			[]byte("AC"), // 268471337
			[]byte("AH"), // 268446761
			0,
		},
		{
			"75432Unsuited",
			[]byte("2D"), // 81922
			[]byte("3D"), // 147715
			[]byte("4D"), // 279045
			[]byte("5H"), // 533255
			[]byte("7H"), // 2106637
			7462,
		},
		{
			"KQJ109Unsuited",
			[]byte("9S"), // 8394515
			[]byte("10S"), // 16783383
			[]byte("JH"), // 33564957
			[]byte("QS"), 	// 67115551
			[]byte("KS"), 	// 134224677
			1601,
		},
	}

	for _, table := range tables {
		testname := fmt.Sprintf("%s", table.description)
		t.Run(testname, func(t *testing.T) {
			result := uniqueValue(table.c1, table.c2, table.c3, table.c4, table.c5)
			if result != table.want {
				t.Errorf("\nFAIL:\nIn: %d, %d, %d, %d, %d \nWant: %d \nGot: %d \n", table.c1, table.c2, table.c3, table.c4, table.c5, table.want, result)
			}
		})
	}
}

*/

func TestHashTableLookup(t *testing.T) {
	tables := []struct {
		description string
		key         uint32
		want        uint16
	}{
		{
			"InTable",
			104553157,
			11,
		},
		{
			"InTable",
			48,
			166,
		},
		// {
		// 	"NotInTable",
		// 	104553156,
		// 	0,
		// },
	}

	for _, table := range tables {
		testname := table.description
		t.Run(testname, func(t *testing.T) {
			result := uint16(0xFFFF)
			result = hashes.get(table.key)
			if result != table.want {
				t.Errorf("\nFAIL:\nIn: %d \nWant: %d \nGot: %d \n", table.key, table.want, result)
			} else {
				t.Logf("\nPASS:\nIn: %d \nWant: %d \nGot: %d \n", table.key, table.want, result)
			}
		})
	}
}

func TestHandValue(t *testing.T) {
	tables := []struct {
		description string
		c1          []byte
		c2          []byte
		c3          []byte
		c4          []byte
		c5          []byte
		want        int
	}{
		{
			"75432Unsuited",
			[]byte("2D"), // 81922
			[]byte("3D"), // 147715
			[]byte("4D"), // 279045
			[]byte("5H"), // 533255
			[]byte("7H"), // 2106637
			7462,
		},
		{
			"KQJ109Unsuited",
			[]byte("9S"),  // 8394515
			[]byte("10S"), // 16783383
			[]byte("JH"),  // 33564957
			[]byte("QS"),  // 67115551
			[]byte("KS"),  // 134224677
			1601,
		},
		{
			"SixHighStraightFlushSpades",
			[]byte("2S"), // 69634
			[]byte("3S"), // 135427
			[]byte("4S"), // 266757
			[]byte("5S"), // 529159
			[]byte("6S"), // 1053707
			9,
		},
		{
			"SixHighStraightFlushSpadesDifferentOrder",
			[]byte("6S"), // 1053707
			[]byte("2S"), // 69634
			[]byte("3S"), // 135427
			[]byte("4S"), // 266757
			[]byte("5S"), // 529159
			9,
		},
		{
			"AKQJ9FlushHearts",
			[]byte("9H"), // 8398611
			[]byte("JH"), // 33564957
			[]byte("QH"), // 67119647
			[]byte("KH"), // 134228773
			[]byte("AH"), // 268446761
			323,
		},
		{
			"AKQJ9FlushClubs",
			[]byte("9C"), // 8423187
			[]byte("JC"), // 33589533
			[]byte("QC"), // 67144223,
			[]byte("KC"), // 134253349
			[]byte("AC"), // 268471337
			323,
		},
		{
			"AAAAK",
			[]byte("AC"), // 268471337
			[]byte("AD"), // 268454953
			[]byte("AH"), // 268446761
			[]byte("KS"), // 134224677
			[]byte("AS"), // 268442665
			11,
		},
		{
			"AKAAA",
			[]byte("AC"), // 268471337
			[]byte("KS"), // 134224677
			[]byte("AD"), // 268454953
			[]byte("AH"), // 268446761
			[]byte("AS"), // 268442665
			11,
		},
		{
			"22223",
			[]byte("2S"), // 69634
			[]byte("3S"), // 135427
			[]byte("2H"), // 73730
			[]byte("2D"), // 81922
			[]byte("2C"), // 98306
			166,
		},
		{
			"Problem Hand",
			[]byte("6S"), // 1053707
			[]byte("3S"), // 135427
			[]byte("4S"), // 266757
			[]byte("5S"), // 529159
			[]byte("6D"), // 1065995
			5302,
		},
	}

	for _, table := range tables {
		testname := fmt.Sprintf("%s", table.description)
		t.Run(testname, func(t *testing.T) {
			result := HandValue(
				ParseCardStr(table.c1),
				ParseCardStr(table.c2),
				ParseCardStr(table.c3),
				ParseCardStr(table.c4),
				ParseCardStr(table.c5),
			)
			if result != table.want {
				t.Errorf("\nFAIL:\nIn: %s, %s, %s, %s, %s \nWant: %d \nGot: %d \n", table.c1, table.c2, table.c3, table.c4, table.c5, table.want, result)
			}
		})
	}
}

func TestBestFiveOfSix(t *testing.T) {
	tables := []struct {
		description string
		c1          []byte
		c2          []byte
		c3          []byte
		c4          []byte
		c5          []byte
		c6          []byte
		want        int
	}{
		{
			"KQJ10987Unsuited",
			[]byte("9S"),  // 8394515
			[]byte("10S"), // 16783383
			[]byte("JH"),  // 33564957
			[]byte("QS"),  // 67115551
			[]byte("KS"),  // 134224677
			[]byte("7H"),  // 2106637
			1601,
		},
		{
			"SixHighStraightFlushSpadesPlusSevenHighStraight",
			[]byte("2S"), // 69634
			[]byte("3S"), // 135427
			[]byte("4S"), // 266757
			[]byte("5S"), // 529159
			[]byte("6S"), // 1053707
			[]byte("7D"), // 2114829
			9,
		},
		{
			"SixHighStraightFlushSpadesPlusSixHighStraight",
			[]byte("6S"), // 1053707
			[]byte("2S"), // 69634
			[]byte("3S"), // 135427
			[]byte("4S"), // 266757
			[]byte("5S"), // 529159
			[]byte("6D"), // 1065995
			9,
		},
		{
			"AKQJ9FlushHearts",
			[]byte("9H"), // 8398611
			[]byte("JH"), // 33564957
			[]byte("QH"), // 67119647
			[]byte("KH"), // 134228773
			[]byte("AH"), // 268446761
			[]byte("9D"), // 8406803
			323,
		},
		{
			"AKQJ9FlushClubs",
			[]byte("9C"), // 8423187
			[]byte("JC"), // 33589533
			[]byte("QC"), // 67144223
			[]byte("KC"), // 134253349
			[]byte("AC"), // 268471337
			[]byte("8D"), // 4212241
			323,
		},
		{
			"AAAAKK",
			[]byte("AC"), // 268471337
			[]byte("AD"), // 268454953
			[]byte("AH"), // 268446761
			[]byte("KS"), // 134224677
			[]byte("AS"), // 268442665
			[]byte("KH"), // 134228773
			11,
		},
		{
			"AKAAAK",
			[]byte("AC"), // 268471337
			[]byte("KS"), // 134224677
			[]byte("AD"), // 268454953
			[]byte("AH"), // 268446761
			[]byte("AS"), // 268442665
			[]byte("KH"), // 134228773
			11,
		},
	}

	for _, table := range tables {
		testname := fmt.Sprintf("%s", table.description)
		t.Run(testname, func(t *testing.T) {

			_, result := BestFiveOfSix(
				ParseCardStr(table.c1),
				ParseCardStr(table.c2),
				ParseCardStr(table.c3),
				ParseCardStr(table.c4),
				ParseCardStr(table.c5),
				ParseCardStr(table.c6),
			)
			if result != table.want {
				t.Errorf("\nFAIL:\nIn: %s, %s, %s, %s, %s, %s \nWant: %d \nGot: %d \n", table.c1, table.c2, table.c3, table.c4, table.c5, table.c6, table.want, result)
			}
		})
	}
}

func TestBestFiveOfSeven(t *testing.T) {
	tables := []struct {
		description string
		c1          []byte
		c2          []byte
		c3          []byte
		c4          []byte
		c5          []byte
		c6          []byte
		c7          []byte
		want        int
	}{
		{
			"KQJ10987Unsuited",
			[]byte("9S"),  // 8394515
			[]byte("10S"), // 16783383
			[]byte("JH"),  // 33564957
			[]byte("QS"),  // 67115551
			[]byte("KS"),  // 134224677
			[]byte("7H"),  // 2106637
			[]byte("8H"),  // 4204049
			1601,
		},
		{
			"SixHighStraightFlushSpadesPlusEightHighStraight",
			[]byte("2S"), // 69634
			[]byte("3S"), // 135427
			[]byte("4S"), // 266757
			[]byte("5S"), // 529159
			[]byte("6S"), // 1053707
			[]byte("7D"), // 2114829
			[]byte("8D"), // 4212241
			9,
		},
		{
			"SixHighStraightFlushSpadesPlusSixHighStraight",
			[]byte("6S"), // 1053707
			[]byte("2S"), // 69634
			[]byte("3S"), // 135427
			[]byte("4S"), // 266757
			[]byte("5S"), // 529159
			[]byte("5D"), // 541447
			[]byte("6D"), // 1065995
			9,
		},
		{
			"AKQJ9FlushHearts",
			[]byte("9H"), // 8398611
			[]byte("JH"), // 33564957
			[]byte("QH"), // 67119647
			[]byte("KH"), // 134228773
			[]byte("AH"), // 268446761
			[]byte("8D"), // 4212241
			[]byte("9D"), // 8406803
			323,
		},
		{
			"AKQJ9FlushClubs",
			[]byte("9C"), // 8423187
			[]byte("JC"), // 33589533
			[]byte("QC"), // 67144223,
			[]byte("KC"), // 134253349
			[]byte("AC"), // 268471337
			[]byte("8D"), // 4212241
			[]byte("9D"), // 8406803
			323,
		},
		{
			"AAAAKQK",
			[]byte("AC"), // 268471337
			[]byte("AD"), // 268454953
			[]byte("AH"), // 268446761
			[]byte("KS"), // 134224677
			[]byte("AS"), // 268442665
			[]byte("QH"), // 67119647
			[]byte("KH"), // 134228773
			11,
		},
		{
			"AKAAAQK",
			[]byte("AC"), // 268471337
			[]byte("KS"), // 134224677
			[]byte("AD"), // 268454953
			[]byte("AH"), // 268446761
			[]byte("AS"), // 268442665
			[]byte("QH"), // 67119647
			[]byte("KH"), // 134228773
			11,
		},
	}

	for _, table := range tables {
		testname := fmt.Sprintf("%s", table.description)
		t.Run(testname, func(t *testing.T) {

			_, result := BestFiveOfSeven(
				ParseCardStr(table.c1),
				ParseCardStr(table.c2),
				ParseCardStr(table.c3),
				ParseCardStr(table.c4),
				ParseCardStr(table.c5),
				ParseCardStr(table.c6),
				ParseCardStr(table.c7),
			)
			if result != table.want {
				t.Errorf("\nFAIL:\nIn: %s, %s, %s, %s, %s, %s, %s \nWant: %d \nGot: %d \n", table.c1, table.c2, table.c3, table.c4, table.c5, table.c6, table.c7, table.want, result)
			}
		})
	}
}

func TestParseCardStr(t *testing.T) {
	type args struct {
		c []byte
	}
	tests := []struct {
		name string
		args args
		want Card
	}{
		{
			"Two of Clubs (2C)",
			args{c: []byte("2C")},
			98306,
		},
		{
			"Two of Clubs (2c)",
			args{c: []byte("2c")},
			98306,
		},
		{
			"Jack of Hearts (JH)",
			args{c: []byte("JH")},
			33564957,
		},
		{
			"Jack of Hearts (Jh)",
			args{c: []byte("Jh")},
			33564957,
		},
		{
			"Jack of Hearts (jH)",
			args{c: []byte("jH")},
			33564957,
		},
		{
			"Jack of Hearts (jh)",
			args{c: []byte("jh")},
			33564957,
		},
		{
			"Ten of Spades (10S)",
			args{c: []byte("10S")},
			16783383,
		},
		{
			"Ten of Spades (TS)",
			args{c: []byte("TS")},
			16783383,
		},
		{
			"Nonsense (ABC)",
			args{c: []byte("ABC")},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCardStr(tt.args.c); got != tt.want {
				t.Errorf("ParseCardStr() = %v, want %v", got, tt.want)
			}

		})
	}
}
