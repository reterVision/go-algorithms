package jumphashing

// JumpHash implements JumpHash algorithm
func JumpHash(key, buckets int) int {
	bucket, jump := -1.0, 0.0
	for jump < float64(buckets) {
		bucket = jump
		key = key*2862933555777941757 + 1
		jump = (bucket + 1) * (float64(1<<31) / float64((key>>33)+1))
	}
	return int(bucket)
}
