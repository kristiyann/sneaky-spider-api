package backend

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func normalizeWhitespace(input string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(strings.TrimSpace(input), " ")
}

func TestGeneratePaginatedSqlStatement(t *testing.T) {
	selectStatement := "SELECT * FROM products"
	filterStatement := "WHERE category = 'Electronics'"
	orderBy := "product_name ASC"

	tests := []struct {
		name            string
		selectStatement string
		top             int
		skip            int
		orderBy         string
		filterStatement string
		expected        string
	}{
		{
			name:            "BasicPagination",
			selectStatement: selectStatement,
			top:             10,
			skip:            0,
			orderBy:         orderBy,
			filterStatement: filterStatement,
			expected: `
				SELECT *
				FROM (
					SELECT *, COUNT(*) OVER () AS count
					FROM ( SELECT * FROM products WHERE category = 'Electronics' ) AS t1 ) AS t2 
				ORDER BY product_name ASC OFFSET 0 ROWS FETCH NEXT 10 ROWS ONLY`,
		},
		{
			name:            "PaginationWithSkip",
			selectStatement: selectStatement,
			top:             5,
			skip:            15,
			orderBy:         orderBy,
			filterStatement: filterStatement,
			expected: `
				SELECT *
				FROM (
					SELECT *, COUNT(*) OVER () AS count
					FROM ( SELECT * FROM products WHERE category = 'Electronics' ) AS t1 ) AS t2 
				ORDER BY product_name ASC OFFSET 15 ROWS FETCH NEXT 5 ROWS ONLY`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := generatePaginatedSqlStatement(
				test.selectStatement,
				test.top,
				test.skip,
				test.orderBy,
				test.filterStatement,
			)
			normalizedExpected := normalizeWhitespace(test.expected)
			normalizedResult := normalizeWhitespace(result)
			if normalizedResult != normalizedExpected {
				t.Errorf("Expected: %s\nGot: %s", normalizedExpected, normalizedResult)
			}
		})
	}
}

type StringerMock struct {
	Value string
}

func (s StringerMock) String() string {
	return s.Value
}

func TestParseArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []fmt.Stringer
		expected string
	}{
		{
			name:     "EmptyArray",
			input:    []fmt.Stringer{},
			expected: "",
		},
		{
			name: "SingleElementArray",
			input: []fmt.Stringer{
				StringerMock{Value: "element1"},
			},
			expected: "element1",
		},
		{
			name: "MultipleElementArray",
			input: []fmt.Stringer{
				StringerMock{Value: "element1"},
				StringerMock{Value: "element2"},
				StringerMock{Value: "element3"},
			},
			expected: "element1, element2, element3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := parseArray(test.input)
			if result != test.expected {
				t.Errorf("Expected: %s, Got: %s", test.expected, result)
			}
		})
	}
}
