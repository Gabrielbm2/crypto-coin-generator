package mappers

import "desafioKlever/models"

func MapVotesToPayload(votes []*models.Votes) []*models.VotesPayload {
	votesPayload := make([]*models.VotesPayload, len(votes))
	for i, vote := range votes {
		votesPayload[i] = &models.VotesPayload{
			Likes:    vote.Likes,
			Dislikes: vote.Dislikes,
		}
	}
	return votesPayload
}
