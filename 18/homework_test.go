package main

import (
	"reflect"
	"testing"
)

func TestParseSumBracketStart(t *testing.T) {
	testString := `(5+6+8*(5*2*7*2)+2)+3`
	want := Sum{
		n: 2,
		s: []NestedExpression{
			NestedExpression{n: 1, e: []string{"5", "+", "6", "+", "8", "*"}},
			NestedExpression{n: 2, e: []string{"5", "*", "2", "*", "7", "*", "2"}},
			NestedExpression{n: 1, e: []string{"+", "2"}},
			NestedExpression{n: 0, e: []string{"+", "3"}},
		},
	}

	result := parseSum(testString)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestParseSumNumberStart(t *testing.T) {
	testString := `5+6+8*(5*2*7*2)+2+3`
	want := Sum{
		n: 1,
		s: []NestedExpression{
			NestedExpression{n: 0, e: []string{"5", "+", "6", "+", "8", "*"}},
			NestedExpression{n: 1, e: []string{"5", "*", "2", "*", "7", "*", "2"}},
			NestedExpression{n: 0, e: []string{"+", "2", "+", "3"}},
		},
	}

	result := parseSum(testString)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestParseSumNoBrackets(t *testing.T) {
	testString := `5+6+8+2`
	want := Sum{
		n: 0,
		s: []NestedExpression{
			NestedExpression{n: 0, e: []string{"5", "+", "6", "+", "8", "+", "2"}},
		},
	}

	result := parseSum(testString)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestParseSumManyBrackets(t *testing.T) {
	testString := `(((5+6)+(8+2)+1*2))`
	want := Sum{
		n: 3,
		s: []NestedExpression{
			NestedExpression{n: 3, e: []string{"5", "+", "6"}},
			NestedExpression{n: 2, e: []string{"+"}},
			NestedExpression{n: 3, e: []string{"8", "+", "2"}},
			NestedExpression{n: 2, e: []string{"+", "1", "*", "2"}},
		},
	}

	result := parseSum(testString)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

// Evaluate Sum

func TestEvaluateSum1(t *testing.T) {
	testString := "1+2*3+4*5+6"
	want := 71

	result := evaluateSum(parseSum(testString), false)

	if result != want {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestEvaluateSum2(t *testing.T) {
	testString := "1 + (2 * 3) + (4 * (5 + 6))"
	want := 51

	result := evaluateSum(parseSum(testString), false)

	if result != want {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestEvaluateSum3(t *testing.T) {
	testString := "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	want := 12240

	result := evaluateSum(parseSum(testString), false)

	if result != want {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestEvaluateSum4(t *testing.T) {
	testString := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	want := 13632

	result := evaluateSum(parseSum(testString), false)

	if result != want {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestEvaluateSum5(t *testing.T) {
	testString := "6 * (2 + 4 * (3 + 7 * 8 + 4) * 8 + 3 + (9 + 8 + 4 + 9)) * 8 + ((2 * 5 * 9 + 7 + 4 * 3) + (8 + 3 + 7) + (7 + 2) * 8 * 8 + 9)"
	want := 216249

	result := evaluateSum(parseSum(testString), false)

	if result != want {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

// Collapse terms

func TestCollapseTerms1(t *testing.T) {
	testSum := []NestedExpression{
		NestedExpression{n: 1, e: []string{"5", "+", "6", "+", "8", "*"}},
		NestedExpression{n: 1, e: []string{"52"}},
		NestedExpression{n: 1, e: []string{"+", "2"}},
		NestedExpression{n: 0, e: []string{"+", "3"}},
	}

	want := []NestedExpression{
		NestedExpression{n: 1, e: []string{"5", "+", "6", "+", "8", "*", "52", "+", "2"}},
		NestedExpression{n: 0, e: []string{"+", "3"}},
	}

	result := collapseTerms(testSum)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestCollapseTerms2Dont(t *testing.T) {
	testSum := []NestedExpression{
		NestedExpression{n: 1, e: []string{"5", "+", "6", "+", "8", "*"}},
		NestedExpression{n: 2, e: []string{"5", "*", "2", "*", "7", "*", "2"}},
		NestedExpression{n: 1, e: []string{"+", "2"}},
		NestedExpression{n: 0, e: []string{"+", "3"}},
	}

	want := testSum

	result := collapseTerms(testSum)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestCollapseTerms3All(t *testing.T) {
	testSum := []NestedExpression{
		NestedExpression{n: 1, e: []string{"5", "+", "6", "+", "8", "*"}},
		NestedExpression{n: 1, e: []string{"167"}},
		NestedExpression{n: 1, e: []string{"+", "52"}},
	}

	want := []NestedExpression{
		NestedExpression{n: 1, e: []string{"5", "+", "6", "+", "8", "*", "167", "+", "52"}},
	}

	result := collapseTerms(testSum)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestCollapseTerms3Variant(t *testing.T) {
	testSum := []NestedExpression{
		NestedExpression{n: 0, e: []string{"5", "+", "6", "+", "8", "*"}},
		NestedExpression{n: 1, e: []string{"167"}},
		NestedExpression{n: 1, e: []string{"+", "52"}},
	}

	want := []NestedExpression{
		NestedExpression{n: 0, e: []string{"5", "+", "6", "+", "8", "*"}},
		NestedExpression{n: 1, e: []string{"167", "+", "52"}},
	}

	result := collapseTerms(testSum)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

// [{0 [1 +]} {1 [2 * 3]} {0 [+]} {1 [4 *]} {1 [11]}]

func TestCollapseTerms4(t *testing.T) {
	testSum := []NestedExpression{
		NestedExpression{n: 0, e: []string{"1", "+"}},
		NestedExpression{n: 1, e: []string{"2", "*", "3"}},
		NestedExpression{n: 0, e: []string{"+"}},
		NestedExpression{n: 1, e: []string{"4", "*"}},
		NestedExpression{n: 1, e: []string{"11"}},
	}

	want := []NestedExpression{
		NestedExpression{n: 0, e: []string{"1", "+"}},
		NestedExpression{n: 1, e: []string{"2", "*", "3"}},
		NestedExpression{n: 0, e: []string{"+"}},
		NestedExpression{n: 1, e: []string{"4", "*", "11"}},
	}

	result := collapseTerms(testSum)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

// [{0 [6 *]} {1 [2 + 4 *]} {1 [84]} {1 [* 8 + 3 +]} {1 [30]} {0 [* 8]}

func TestCollapseTerms5(t *testing.T) {
	testSum := []NestedExpression{
		NestedExpression{n: 0, e: []string{"6"}},
		NestedExpression{n: 1, e: []string{"2", "+", "4", "*"}},
		NestedExpression{n: 1, e: []string{"84"}},
		NestedExpression{n: 1, e: []string{"*", "8", "+", "3", "+"}},
		NestedExpression{n: 1, e: []string{"30"}},
		NestedExpression{n: 0, e: []string{"*", "8"}},
	}

	want := []NestedExpression{
		NestedExpression{n: 0, e: []string{"6"}},
		NestedExpression{n: 1, e: []string{"2", "+", "4", "*", "84", "*", "8", "+", "3", "+", "30"}},
		NestedExpression{n: 0, e: []string{"*", "8"}},
	}

	result := collapseTerms(testSum)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

// Evaluate expression

func TestEvaluateExpression(t *testing.T) {
	testExpression := []string{"5", "+", "6", "*", "8", "+", "2"}
	want := []string{"90"}

	result := evaluateExpression(testExpression)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

// Evaluate with op precedence

func TestEvaluateExpressionWithOpPrecedence(t *testing.T) {
	testExpression := []string{"5", "+", "6", "*", "8", "+", "2"}
	want := []string{"110"}

	result := evaluateExpressionWithOpPrecedence(testExpression)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestEvaluateSumWithOpPrecedence(t *testing.T) {
	testString := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"

	want := 23340

	result := evaluateSum(parseSum(testString), true)

	if result != want {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestEvaluateSumWithOpPrecedence2(t *testing.T) {
	testString := "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"

	want := 669060

	result := evaluateSum(parseSum(testString), true)

	if result != want {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestEvaluateSumWithOpPrecedence3(t *testing.T) {
	testString := "5 + (8 * 3 + 9 + 3 * 4 * 3)"

	want := 1445

	result := evaluateSum(parseSum(testString), true)

	if result != want {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestEvaluateSumWithOpPrecedence4(t *testing.T) {
	testString := "4 * 7"

	want := 28

	result := evaluateSum(parseSum(testString), true)

	if result != want {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestEvaluateSumWithOpPrecedence5(t *testing.T) {
	testString := "8 + 5 * 7 + 5"

	want := 156

	result := evaluateSum(parseSum(testString), true)

	if result != want {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}

func TestEvaluateSumWithOpPrecedence6(t *testing.T) {
	testString := "(8 + 6) * 3 * 8"

	want := 336

	result := evaluateSum(parseSum(testString), true)

	if result != want {
		t.Fatalf("Result:\n%v\nWant:\n%v\n", result, want)
	}
}
