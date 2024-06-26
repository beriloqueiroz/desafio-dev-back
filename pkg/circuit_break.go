package pkg

import "github.com/sony/gobreaker/v2"

func NewCircuitBreak[T any](name string, countRequests uint32, failureRatioIn float64) *gobreaker.CircuitBreaker[T] {
	var st gobreaker.Settings
	st.Name = name
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= countRequests && failureRatio >= failureRatioIn // todo pode ser variável de ambiente
	}

	return gobreaker.NewCircuitBreaker[T](st)
}
