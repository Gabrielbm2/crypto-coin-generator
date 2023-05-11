package mappers

import "desafioKlever/models"

//Esse package "mappers" contém uma função MapVotesToPayload que recebe uma lista de votos e mapeia cada voto para um modelo de dados simplificado, com apenas as informações necessárias para exibição em uma interface de usuário.

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
