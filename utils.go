package gomathutils

// float32

func MaxFloat32v(vars ...float32) float32 {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func MaxFloat32x2(f1, f2 float32) float32 {
	if f1 > f2 {
		return f1
	}
	return f2
}

func MaxFloat32x3(f1, f2, f3 float32) float32 {
	if f1 > f2 {
		if f1 > f3 {
			return f1
		}
		return f3
	}
	if f2 > f3 {
		return f2
	}
	return f3
}

func MinFloat32v(vars ...float32) float32 {
	min := vars[0]
	for _, v := range vars {
		if v < min {
			min = v
		}
	}
	return min
}

func AbsFloat32(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

func MaxminFloat32(arr []float32) (float32, float32) {
	min := arr[0]
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

func EqualFloat32(a, b []float32, tol float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if AbsFloat32((v - b[i])) > tol {
			return false
		}
	}
	return true
}

// int32

func AbsInt32(n int32) int32 {
	y := n >> 31
	return (n ^ y) - y
}

func EqualInt32(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
