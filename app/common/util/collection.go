// Package util provides common utility functions
package util

// Map generic function to convert from slice of one type to slice of another type
func Map[SourceType any, DestinationType any](
	input []SourceType,
	mapperFn func(index int, value SourceType) (DestinationType, error),
) ([]DestinationType, error) {
	result := make([]DestinationType, len(input))
	for index, value := range input {
		mappedVal, err := mapperFn(index, value)
		if err != nil {
			return nil, err
		}
		result[index] = mappedVal
	}
	return result, nil
}

// SlimMap generic function to convert from slice of one type to slice of another type
func SlimMap[SourceType any, DestinationType any](
	input []SourceType,
	mapperFn func(value SourceType) DestinationType,
) []DestinationType {
	result, _ := Map(input, func(_ int, value SourceType) (DestinationType, error) {
		return mapperFn(value), nil
	})
	return result
}
