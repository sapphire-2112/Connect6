package commands
import (
	"math"
	
)
func CalculateTrustRatio(trustedBy, totalContacts int) float64 {
	if totalContacts == 0 {
		return 0
	}

	return (float64(trustedBy) / float64(totalContacts)) * 100
}
func CalculateConfidenceScore(trustedBy, totalContacts int) float64 {

	if totalContacts == 0 {
		return 0
	}

	p := float64(trustedBy) / float64(totalContacts)
	n := float64(totalContacts)

	z := 1.96 // 95% confidence

	score :=
		(p + (z*z)/(2*n) -
			z*math.Sqrt((p*(1-p))/n+(z*z)/(4*n*n))) /
			(1 + (z*z)/n)

	return score * 100
}
func GetConfidence(score float64) string {

	switch {

	case score >= 80:
		return "Very High"

	case score >= 60:
		return "High"

	case score >= 40:
		return "Medium"

	case score >= 20:
		return "Low"

	default:
		return "Very Low"
	}
}