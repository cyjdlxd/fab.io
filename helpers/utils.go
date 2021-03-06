package helpers

import (
	"fmt"
	"math/rand"
	"reflect"
)

// H is alias for map[string]interface{}
type H map[string]interface{}

// Range used to create slice
func Range(min int, max int) []int {
	ls := []int{}

	for i := min; i < max; i++ {
		ls = append(ls, i)
	}

	return ls
}

// Times used to iterate n times with handler
func Times(n int, handler func(int) bool) {
	for i := range Range(0, n) {
		shouldContinue := handler(i)

		if !shouldContinue {
			break
		}
	}
}

// Merge used to merge to map
func Merge(h1 H, h2 H) H {
	for k, v := range h2 {
		h1[k] = v
	}

	return h1
}

// NotFoundError used to generate error for not round record
func NotFoundError(itemName string) error {
	return fmt.Errorf("%v not found", itemName)
}

func MapValues(ss interface{}, field string) []interface{} {
	values := []interface{}{}
	s := reflect.ValueOf(ss)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	slice := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		slice[i] = s.Index(i).Interface()
	}

	for _, s := range slice {
		values = append(values, GetFieldValueByName(s, field))
	}

	return values
}

// RandRange used to generate a random number from min to max inclusively
func RandRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

// Rand used to generate a random float from 0.0 - 1.0
func Rand() float32 {
	return rand.Float32()
}
