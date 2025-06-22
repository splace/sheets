package sheets

import "fmt"
import . "golang.org/x/exp/constraints"
import "math"
import "testing"
import "iter"


func Primes[T Unsigned]() iter.Seq[T] {
	return PrimesAbove[T](0)
}

func PrimesAbove[T Unsigned](p T) iter.Seq[T] {
	return Filter(func(i T) bool { return IsPrime(i) }, Geometric(p/2*2+1, 2))
}

func IsPrime[T Unsigned](n T) bool {
	return !IsNotPrime(n)
}

func IsNotPrime[T Unsigned](n T) bool {
	for i := T(2); i <= Sqrt(n); i += 1 { // cann't be divisible above the sqrt, if wasn't below.
		if n%i == 0 {
			return true
		}
	}
	return false
}

func Sqrt[T Unsigned](x T) (r T) {
	var b T = 1
	t := x
	for t > 4 {
		t >>= 2
		b <<= 2
	}
	for ; b != 0; b >>= 2 {
		bb := r | b
		r >>= 1
		if x >= bb {
			x -= bb
			r |= b
		}
	}
	return
}




func ExampleLimit() {
	fmt.Print( List[int](Limit(Odds[int](),10)))
	// Output:
	// 1 3 5 7 9 11 13 15 17 19
}

func ExampleRepeat() {
	fmt.Print( List[float32](Limit(Repeat[float32](9),10)))
	// Output:
	// 9 9 9 9 9 9 9 9 9 9
}

func ExampleRepeatSequence() {
	fmt.Print(List[rune](Limit(RepeatSequence(Runes("Ab")),10)))
	// Output:
	// 65 98 65 98 65 98 65 98 65 98
}

func ExampleOdds() {
	fmt.Print( List[int](Limit(Odds[int](),10)))
	// Output:
	// 1 3 5 7 9 11 13 15 17 19
}

func ExampleFilter() {
	fmt.Print( List[int](Limit(Filter(Between(10, 100), Odds[int]()),10)))
	// Output:
	// 11 13 15 17 19 21 23 25 27 29
}

func ExampleMake() {
	fmt.Print( List[[2]int](Limit(Make(func(i int) [2]int { return [2]int{i, 2 * i} }),10)))
	// Output:
	// [0 0] [1 2] [2 4] [3 6] [4 8] [5 10] [6 12] [7 14] [8 16] [9 18]
}

func ExampleUntil() {
	fmt.Print( List[uint](Until(func(n uint) bool { return n > 10 }, Odds[uint]())))
	// Output:
	// 1 3 5 7 9 11
}

func ExampleMultiply() {
	fmt.Print( List[uint](Limit(Multiply(PrimesAbove[uint](0), 2),10)))
	// Output:
	// 2 6 10 14 22 26 34 38 46 58
}

func ExampleApply() {
	fmt.Print( List[float32](Limit(Apply(PrimesAbove[uint](0), func(i uint) float32 { return float32(math.Sqrt(float64(i))) }),5)))
	// Output:
	// 1 1.7320508 2.236068 2.6457512 3.3166249
}

func ExampleApply_typeChanging() {
	fmt.Print( List[float32](Limit(Apply(PrimesAbove[uint](0), func(i uint) float32 { return float32(i) / 2 }),10)))
	// Output:
	// 0.5 1.5 2.5 3.5 5.5 6.5 8.5 9.5 11.5 14.5
}

func UntilAtLeastDelta[T Integer | Float](diff T) func(...T) bool {
	return func(cs ...T) bool {
		return len(cs) > 2 && (cs[len(cs)-1]-cs[len(cs)-2]) > diff
	}
}

func ExampleUntilHistory() {
	fmt.Print( List[uint8](UntilHistory(UntilAtLeastDelta[uint8](10), Primes[uint8]())))
	// Output:
	// 1 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97 101 103 107 109 113 127
}

func ExampleGeometric_float() {
	fmt.Print( List[float32](Limit(Geometric[float32](10.0, 1.0),10)))
	// Output:
	// 10 11 12 13 14 15 16 17 18 19
}

func ExampleGeometric_complex() {
	fmt.Print(List[complex128](Limit(Geometric(10.0+1i, 1.0+2i),10)))
	// Output:
	// (10+1i) (11+3i) (12+5i) (13+7i) (14+9i) (15+11i) (16+13i) (17+15i) (18+17i) (19+19i)
}

func ExampleFibonacci_complex() {
	var count uint = 10
	fmt.Print( List[complex64](Limit(Fibonacci[complex64](),count)))
	// Output:
	// (1+0i) (1+0i) (2+0i) (3+0i) (5+0i) (8+0i) (13+0i) (21+0i) (34+0i) (55+0i)

}

func ExampleConcat() {
	fmt.Print( List[uint](Limit(Concat(Limit(Odds[uint](),4), Evens[uint]()),8)))
	// Output:
	// 1 3 5 7 0 2 4 6
}

func ExampleTotalise() {
	fmt.Print( List[uint8](Limit(Totalise(Odds[uint8]()),9)))
	// Output:
	// 1 4 9 16 25 36 49 64 81
}


func ExampleRunes() {
	fmt.Print(List[rune](Runes("hello")))
	// Output:
	// 104 101 108 108 111
}

func ExamplePermutations() {
	fmt.Print( List[rune](Permutations(Runes("ab"))))
	// Output:
	// 'h' 'e' 'l' 'l' 'o' 'h' 'e' 'l' 'l' 'o'
}

func ExampleAmalgomate() {
	fmt.Print( List[string](Amalgomate[rune,string](
		func(rs ...rune)string{return string(rs)},
		Runes("hello"),
		Runes("hello"),
	)))
	// Output:
	// hh ee ll ll oo
}

////func ExampleSequenceCycle(){
////	ptc:=make(chan uint)
////	go Count(9).On(Cycle{Odds(),Evens()},ptc)
////	for p:= range ptc{
////		fmt.Print("\t",p)
////	}
////	// Output:
////	// 1	2	3	4	5	6	7	8	9
////}


func ExampleMatched() {
	fmt.Println(Matched(Runes("hello"),Runes("hello")))
	fmt.Println(Matched(Runes("hello"),Runes("hellu")))
	fmt.Println(Matched(Runes("hello"),Runes("hello"),Runes("hello"),Runes("hello")))
	fmt.Println(Matched(Runes("hello"),Runes("hello"),Runes("hellu"),Runes("hello")))
	fmt.Println(Matched(Runes("hello"),Runes("hello there"))) // ignores " there"
	fmt.Println(Matched(Runes("hello there"),Runes("hello")))
	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
}

func ExampleSame() {
//	fmt.Println(Compare(iter.Seq[rune](NewRow([]rune("hello")...)),iter.Seq[rune](NewRow([]rune("hello")...))))
	fmt.Println(Same(Runes("hello"),Runes("hello")))
	fmt.Println(Same(Runes("hello"),Runes("hellu")))
	// Output:
	// true
	// false

}

//func ExampleCompare_c() {
////	fmt.Println(Compare(iter.Seq[rune](NewRow([]rune("hello")...)),iter.Seq[rune](NewRow([]rune("hello")...))))
//	fmt.Println(compare2(Runes("hello"),Runes("hello")))
//	// Output:
//	// true
//	// false
//	// true
//	// false
//}


func TestSequence(t *testing.T) {

	//	t.Error(args ...any)
	//// equivalent to Log followed by Fail.

	//	t.Errorf(format string, args ...any)
	//// equivalent to Logf followed by Fail.

	//	Fail()
	//// marks the function as having failed but continues execution.

	//	t.FailNow()
	//// marks the function as having failed and stops its execution by calling runtime.Goexit (which then runs all deferred calls in the current goroutine). Execution will continue at the next test or benchmark. FailNow must be called from the goroutine running the test or benchmark function,not from other goroutines created during the test. Calling FailNow does not stop those other goroutines.

	//	t.Failed() bool
	//// reports whether the function has failed.

	//	t.Fatal(args ...any)
	//// equivalent to Log followed by FailNow.

	//	t.Fatalf(format string, args ...any)
	//// equivalent to Logf followed by FailNow.

}

func BenchmarkSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Limit(Apply(PrimesAbove[uint](10000), func(i uint) float32 { return float32(i) / 2 }),1000)
	}
}

