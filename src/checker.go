package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/gonfig"
)

func checkAddress(config gonfig.Configuration) {
	targetAddress, err := config.GetString("TARGET_ADDR")
	if err != nil {
		log.Fatal().Msgf("No target address")
	}
	log.Info().Msgf("Starting to check %s", targetAddress)
	var checkFrequency = config.GetIntOrDefault("CHECK_FREQUENCY", 1)

	buckets := config.GetFloatArrayOrDefault("BUCKETS", []float64{0.1, 0.5, 1})

	responseTimes := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "response_times",
		Help:    "The response times of the tracked system",
		Buckets: buckets,
	}, []string{"ResponseCode", "Error"})

	errorMessages := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "error_messages",
		Help: "The error messages returned by the tracked system",
	}, []string{"ErrorMessage"})
	for {
		timeStart := time.Now()
		resp, err := http.Get(targetAddress)
		timeEnd := time.Now()
		duration := timeEnd.Sub(timeStart).Milliseconds()
		if err != nil {
			responseTimes.WithLabelValues("", "true").Observe(float64(duration))
			errorMessages.WithLabelValues(err.Error()).Inc()
		} else {
			responseTimes.WithLabelValues(fmt.Sprintf("%d", resp.StatusCode), "false").Observe(float64(duration))
		}
		fmt.Println(fmt.Sprintf("Duration %d milliseconds", duration))
		time.Sleep(time.Duration(checkFrequency) * time.Second)
	}
}
