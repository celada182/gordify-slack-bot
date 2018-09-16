package main

import (
	"testing"
	"fmt"
)

func TestGetNumberOfGroups(t *testing.T) {
	TestingGetGroupsNumber(t, 22, 3, 7)
	TestingGetGroupsNumber(t, 23, 4, 6)
	TestingGetGroupsNumber(t, 10, 2, 5)
}

func TestingGetGroupsNumber(t *testing.T, 
							testedUsersNumber int, 
							expectedGroupsNumber int, 
							expectedGroupSize int) {
    groupsNumber, groupSize := getNumberOfGroups(testedUsersNumber)
    if groupsNumber != expectedGroupsNumber || groupSize != expectedGroupSize {
       	t.Errorf("Error getting groups number, got: %d of %d users, want: %d of %d users.", groupsNumber, groupSize, expectedGroupsNumber, expectedGroupSize)
    } else {
    	message:=fmt.Sprintf("Given %d users then return %d groups of %d users", testedUsersNumber, groupsNumber, groupSize)
    	fmt.Println(message)
    }
}

func TestGetGroups(t *testing.T) {
	users23 := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w"}
	users22 := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v"}
	users10 := []string{"a","b","c","d","e","f","g","h","i","j","k"}

	TestingGetGroups(t, users22, 3, 7)
	TestingGetGroups(t, users23, 4, 6)
	TestingGetGroups(t, users10, 2, 5)
}

func TestingGetGroups(t *testing.T, 
						testedUsers []string, 
						expectedGroupsNumber int, 
						expectedGroupSize int) {
	groups := getGroups(testedUsers, expectedGroupsNumber, expectedGroupSize)
	if (len(groups) != expectedGroupsNumber) {
       	t.Errorf("Error getting groups, got: %d groups, want: %d groups", len(groups), expectedGroupsNumber)
	}
	if (len(groups[0]) != expectedGroupSize) {
       	t.Errorf("Error getting group size, got: %d users, want: %d users", len(groups[0]), expectedGroupSize)
	}
}