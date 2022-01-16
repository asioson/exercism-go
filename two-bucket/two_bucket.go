// package twobucket implements the routines needed to solve the two bucket problem
package twobucket

import "fmt"

// Gcd computes the greatest common divisor of passed int values
func GCD(x, y int) int {
	if y == 0 {
		return x
	}
	return GCD(y, x%y)
}

// Min returns the minimum among the passed int values
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max returns the maximum among the passed int values
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Solve determines how many steps are needed to reach the goalAmount by strategically
// transfer amounts between two buckets
func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	if sizeBucketOne <= 0 {
		return "", 0, 0, fmt.Errorf("invalid first bucket size")
	}
	if sizeBucketTwo <= 0 {
		return "", 0, 0, fmt.Errorf("invalid second bucket size")
	}
	if goalAmount == 0 {
		return "", 0, 0, fmt.Errorf("invalid goal amount")
	}
	if startBucket != "one" && startBucket != "two" {
		return "", 0, 0, fmt.Errorf("invalid start bucket name")
	}
	if goalAmount%GCD(sizeBucketOne, sizeBucketTwo) == 1 {
		errMsg := fmt.Sprintf("goal %v must be divided evenly by gcd(%v,%v)", goalAmount, sizeBucketOne, sizeBucketTwo)
		return "", 0, 0, fmt.Errorf(errMsg)
	}
	count := 1
	state := []int{0, sizeBucketTwo}
	limits := []int{sizeBucketOne, sizeBucketTwo}
	bucketName := []string{"one", "two"}
	if startBucket == "one" {
		state[1] = sizeBucketOne
		limits[0], limits[1] = limits[1], limits[0]
		bucketName[0], bucketName[1] = bucketName[1], bucketName[0]
	}
	for goalAmount != state[0] && goalAmount != state[1] {
		if state[0] == 0 && limits[0] == goalAmount {
			state[0] = goalAmount
		} else if state[1] == 0 && limits[1] == goalAmount {
			state[1] = goalAmount
		} else if state[0] == limits[0] {
			state[0] = 0
		} else if state[1] == 0 {
			state[1] = limits[1]
		} else {
			state = []int{
				Min(limits[0], state[0]+state[1]),
				Max(0, state[1]-(limits[0]-state[0]))}
		}
		count++
	}
	if state[0] == goalAmount {
		return bucketName[0], count, state[1], nil
	}
	return bucketName[1], count, state[0], nil
}
