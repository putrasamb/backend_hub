package naming_test

import (
	"fmt"
	"service-collection/pkg/common/naming"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFullYear(t *testing.T) {
	s := naming.NamingSeries{
		Format: "CPO.YYYY.MM.DD.##",
		Number: 1,
	}
	expectedDate := time.Now().Format("20060102")
	expected := fmt.Sprintf("CPO%s01", expectedDate)
	result := s.Parse()
	require.Equal(t, expected, *result)

	t.Log("Expected: ", expected)
	t.Log("Result: ", *result)
}

func TestShortYear(t *testing.T) {
	s := naming.NamingSeries{
		Format: "CPO.YY.MM.DD.##",
		Number: 1,
	}
	expectedDate := time.Now().Format("060102")
	expected := fmt.Sprintf("CPO%s01", expectedDate)
	result := s.Parse()
	require.Equal(t, expected, *result)

	t.Log("Expected: ", expected)
	t.Log("Result: ", *result)
}

func TestWeek(t *testing.T) {
	s := naming.NamingSeries{
		Format: "CPO.YY.MM.WW.##",
		Number: 1,
	}
	now := time.Now()
	expectedDate := now.Format("0601")
	_, weekNumber := now.ISOWeek()
	expected := fmt.Sprintf("CPO%s%02d01", expectedDate, weekNumber)
	result := s.Parse()
	require.Equal(t, expected, *result)

	t.Log("Expected: ", expected)
	t.Log("Result: ", *result)
}
