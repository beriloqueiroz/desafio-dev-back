package pkg

import "github.com/sony/gobreaker/v2"

func NewCircuitBreak(name string, countRequests uint32, failureRatio float64) *gobreaker.CircuitBreaker[[]byte] {
	var st gobreaker.Settings
	st.Name = name
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= countRequests && failureRatio >= failureRatio // todo pode ser variável de ambiente
	}

	return gobreaker.NewCircuitBreaker[[]byte](st)
}
