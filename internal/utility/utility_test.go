package utility

import (
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
