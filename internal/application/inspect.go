package application

import "context"

type InspectResult struct {
	AuroraID            string     `json:"aurora_id"`
	OwnerOperatorID     string     `json:"owner_operator_id"`
	GoverningGeneration uint64     `json:"governing_generation"`
	TrustState          TrustState `json:"trust_state"`
}

func resultFromAssessment(a TrustAssessment) InspectResult {
	return InspectResult{AuroraID: a.Current.AuroraID, OwnerOperatorID: a.Current.OwnerOperatorID, GoverningGeneration: a.Current.GoverningGeneration, TrustState: a.State}
}
func (s *Service) Status(ctx context.Context, pass []byte) (InspectResult, error) {
	a, err := s.AssessTrust(ctx, pass)
	if err != nil {
		return InspectResult{}, err
	}
	return resultFromAssessment(a), nil
}
func (s *Service) Inspect(ctx context.Context, pass []byte) (InspectResult, error) {
	a, err := s.requireNormalTrust(ctx, pass)
	if err != nil {
		return resultFromAssessment(a), err
	}
	return resultFromAssessment(a), nil
}
