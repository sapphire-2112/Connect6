package commands

func CalculateTrustRatio(trustedBy, totalContacts int) float64 {
	if totalContacts == 0 {
		return 0
	}

	return (float64(trustedBy) / float64(totalContacts)) * 100
}