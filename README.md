# Go Slice Modification Unexpectedly Affects Underlying Array

This example demonstrates a common issue in Go related to how slices interact with their underlying arrays.  Modifying a slice that shares the same underlying array as another slice (or the original array) can lead to unexpected changes in the original array.

## Bug Description
The code appends an element to a slice `b`, which is a slice of `a`.  This unexpectedly modifies the original array `a` because both slices share the same underlying array.  Observe the behavior of `a` after the modification of `b`.

## How to Reproduce
1. Save the code in `bug.go`.
2. Run the code using `go run bug.go`.

## Solution
The solution involves creating a copy of the slice before modification to avoid affecting the original array.