//go:build !aurora_testhooks

package application

func pauseAfterStateCommitForTest() error { return nil }
