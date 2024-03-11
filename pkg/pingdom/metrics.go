package pingdom

import "github.com/prometheus/client_golang/prometheus"

type clientMetrics struct {
	requestDuration                      *prometheus.HistogramVec
	pingdomApiRateLimitRemainingShort    *prometheus.GaugeVec
	pingdomApiRateLimitRemainingLong     *prometheus.GaugeVec
	pingdomApiRateLimitResetShortSeconds *prometheus.GaugeVec
	pingdomApiRateLimitResetLongSeconds  *prometheus.GaugeVec
}

func newMetrics(reg prometheus.Registerer) *clientMetrics {
	m := &clientMetrics{
		requestDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: "pingdom_request_duration_seconds",
			Help: "Duration of requests to the Pingdom API",
		}, []string{"method", "path"}),
		pingdomApiRateLimitRemainingShort: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "pingdom_api_rate_limit_remaining_short",
			Help: "Pingdom API rate limit remaining for short term",
		}, []string{"method", "path"}),
		pingdomApiRateLimitRemainingLong: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "pingdom_api_rate_limit_remaining_long",
			Help: "Pingdom API rate limit remaining for long term",
		}, []string{"method", "path"}),
		pingdomApiRateLimitResetShortSeconds: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "pingdom_api_rate_limit_reset_short_seconds",
			Help: "Pingdom API rate limit reset for short term",
		}, []string{"method", "path"}),
		pingdomApiRateLimitResetLongSeconds: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "pingdom_api_rate_limit_reset_long_seconds",
			Help: "Pingdom API rate limit reset for long term",
		}, []string{"method", "path"}),
	}
	reg.MustRegister(
		m.requestDuration,
		m.pingdomApiRateLimitRemainingShort,
		m.pingdomApiRateLimitRemainingLong,
		m.pingdomApiRateLimitResetShortSeconds,
		m.pingdomApiRateLimitResetLongSeconds,
	)

	return m
}
