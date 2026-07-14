package mobile
import (
	"connect6/commands"
	"connect6/data"
)
type MyProfile struct {
	ID               string
	Name             string
	TrustedBy        int
	TrustedByPeers   []string
	TotalContacts    int
	TrustRatio       float64
	ConfidenceScore  float64
	Confidence       string
}
func (e *Engine) Profile() (MyProfile, error) {
    identity, err := data.LoadIdentity()
    if err != nil {
        return MyProfile{}, err
    }

    peers, err := data.LoadPeers()
    if err != nil {
        return MyProfile{}, err
    }

    totalContacts := len(peers)

    trustRatio := commands.CalculateTrustRatio(
        identity.TrustedBy,
        totalContacts,
    )

    confidenceScore := commands.CalculateConfidenceScore(
        identity.TrustedBy,
        totalContacts,
    )

    confidence := commands.GetConfidence(confidenceScore)

    profile := MyProfile{
        ID:              identity.ID,
        Name:            identity.Name,
        TrustedBy:       identity.TrustedBy,
        TrustedByPeers:  identity.TrustedByPeers,
        TotalContacts:   totalContacts,
        TrustRatio:      trustRatio,
        ConfidenceScore: confidenceScore,
        Confidence:      confidence,
    }

    return profile, nil
}