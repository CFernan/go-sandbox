// Sorting integers with Dial's algorithm
package dial

import "fmt"

// Definition of the maximum ordinal
const upperBound = 1000

// Initialization of a bucket with upperBound+1 entries in the
// range [0, upperBound+1)
func initBucket() (bucket [][]int) {
	// allocation of upperBound+1 entries, and initialization of
	// each bucket
	bucket = make([][]int, upperBound+1)
	for i := range upperBound {
		bucket[i] = []int{}
	}

	return
}

// Add number i to the i-th bucket
func insert(i int, bucket [][]int) error {
	if len(bucket) <= i {
		return fmt.Errorf("item out of bounds: %v > %v", i, upperBound)
	}

	bucket[i] = append(bucket[i], i)

	return nil
}

// Return all items in the bucket in ascending order just by
// transversing all buckets to 0 to upperBound
func pop(bucket [][]int) []int {
	// initialize the result
	result := []int{}
	for i := range upperBound {
		// if this bucket is not empty
		if len(bucket[i]) > 0 {
			result = append(result, bucket[i]...)
		}
	}

	// and return the result
	return result
}

// Sort all numbers in the given slice using buckets. It returns nil
// if no error happened
func Sort(items []int) error {
	// first create a bucket as a slice of slices and initialize it
	bucket := initBucket()

	// not insert all intes in the bucket
	for _, v := range items {
		if err := insert(v, bucket); err != nil {
			return err
		}
	}

	// modify the given sclice
	copy(items, pop(bucket))
	return nil
}
