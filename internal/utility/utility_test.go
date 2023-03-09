package utility

import (
	"fmt"
	"sort"
	"testing"
)

func TestGetMentionsFromNotification(t *testing.T) {
	notification := "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com"
	expected := []string{
		"studentagnes@gmail.com",
		"studentmiche@gmail.com",
	}
	result := GetMentionsFromNotification(notification)
	if len(result) != len(expected) {
		t.Errorf("Mismatch: too many/few results")
	}
	sort.Strings(result)
	for i := 0; i < len(result); i++ {
		if expected[i] != result[i] {
			t.Errorf("Mismatch: expected %s, got %s instead", expected[i], result[i])
		}
	}
}

func TestGetMentionsFromNotificationFakeMentions(t *testing.T) {
	notification := "Hello students! @studentagnes @studentmiche"
	result := GetMentionsFromNotification(notification)
	if len(result) != 0 {
		t.Errorf("Expected no matches")
	}
}

func TestGetMentionsFromNotificationMixed(t *testing.T) {
	notification := "Hello students! @studentagnes@gmail.com @fake @studentmiche@gmail.com @fake @fake"
	expected := []string{
		"studentagnes@gmail.com",
		"studentmiche@gmail.com",
	}
	result := GetMentionsFromNotification(notification)
	if len(result) != len(expected) {
		t.Errorf("Mismatch: too many/few results")
	}
	sort.Strings(result)
	for i := 0; i < len(result); i++ {
		if expected[i] != result[i] {
			t.Errorf("Mismatch: expected %s, got %s instead", expected[i], result[i])
		}
	}
}

func TestRepeat(t *testing.T) {
	expected := "?, ?, ?"
	actual := Repeat("?", 3, ", ")
	if actual != expected {
		t.Errorf("Mismach: expected %s, got %s instead", expected, actual)
	}
}

func TestMap(t *testing.T) {
	expected := []string{
		"str1",
		"str2",
		"str3",
	}
	actual := Map([]int{1, 2, 3}, func(el int) string {
		return fmt.Sprintf("str%d", el)
	})
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("Mismatch: expected %v, got %v instead", expected, actual)
		}
	}
}
