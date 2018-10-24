package findMedianSortedArrays

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	// ensure n >= m
	if m > n {
		nums1, m, nums2, n = nums2, n, nums1, m
	}
	if n == 0 {
		return 0
	}
	imin, imax, half_len := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := half_len - i
		if i > 0 && nums1[i-1] > nums2[j] {
			// i is too big, must decrease it
			imax = i - 1
		} else if i < m && nums2[j-1] > nums1[i] {
			// i is too small, must increase it
			imin = i + 1
		} else {
			// i is perfect
			var max_of_left, min_of_right int
			if i == 0 {
				max_of_left = nums2[j-1]
			} else if j == 0 {
				max_of_left = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					max_of_left = nums1[i-1]
				} else {
					max_of_left = nums1[j-1]
				}
			}

			// even
			if (m+n)%2 == 1 {
				return float64(max_of_left)
			}

			if i == m {
				min_of_right = nums2[j]
			} else if j == n {
				min_of_right = nums1[i]
			} else {
				if nums1[i] < nums2[j] {
					min_of_right = nums1[i]
				} else {
					min_of_right = nums2[j]
				}
			}

			// odd
			return float64(min_of_right+max_of_left) / 2.0
		}
	}
	return 0
}
