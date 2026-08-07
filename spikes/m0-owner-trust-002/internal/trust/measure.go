package trust

import (
	"crypto/sha256"
	"runtime"
	"time"

	"golang.org/x/crypto/argon2"
)

type KDFMeasurement struct {
	Algorithm             string `json:"algorithm"`
	MemoryKiB             uint32 `json:"memory_kib"`
	Iterations            uint32 `json:"iterations"`
	Parallelism           uint8  `json:"parallelism"`
	KeyBytes              uint32 `json:"key_bytes"`
	DurationNS            int64  `json:"duration_ns"`
	ConfiguredMemoryBytes int64  `json:"configured_memory_bytes"`
	ObservedPeakHeapDelta int64  `json:"observed_peak_heap_delta_bytes"`
	MeasurementMethod     string `json:"measurement_method"`
}

func MeasureArgon2id(passphrase []byte) KDFMeasurement {
	seed := sha256.Sum256([]byte("aurora-spk002-kdf-measurement-salt-v1"))
	salt := seed[:16]
	runtime.GC()
	var before runtime.MemStats
	runtime.ReadMemStats(&before)
	baseline := before.HeapInuse
	peak := baseline
	done := make(chan struct{})
	start := time.Now()
	go func() {
		key := argon2.IDKey(passphrase, salt, kdfIterations, kdfMemoryKiB, kdfParallelism, kdfKeyBytes)
		clear(key)
		close(done)
	}()
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			var end runtime.MemStats
			runtime.ReadMemStats(&end)
			if end.HeapInuse > peak {
				peak = end.HeapInuse
			}
			delta := int64(0)
			if peak > baseline {
				delta = int64(peak - baseline)
			}
			return KDFMeasurement{
				Algorithm:             "Argon2id",
				MemoryKiB:             kdfMemoryKiB,
				Iterations:            kdfIterations,
				Parallelism:           kdfParallelism,
				KeyBytes:              kdfKeyBytes,
				DurationNS:            time.Since(start).Nanoseconds(),
				ConfiguredMemoryBytes: int64(kdfMemoryKiB) * 1024,
				ObservedPeakHeapDelta: delta,
				MeasurementMethod:     "1ms polling of Go runtime HeapInuse; approximate process heap peak, not RSS",
			}
		case <-ticker.C:
			var current runtime.MemStats
			runtime.ReadMemStats(&current)
			if current.HeapInuse > peak {
				peak = current.HeapInuse
			}
		}
	}
}

func DiagnosticFor(c Classification) string {
	return diagnostic(c)
}

func SameOwnerRoot(a, b *OwnerSession) bool {
	if a == nil || b == nil {
		return false
	}
	fa := rootFingerprint(a)
	fb := rootFingerprint(b)
	return len(fa) == len(fb) && string(fa) == string(fb)
}
