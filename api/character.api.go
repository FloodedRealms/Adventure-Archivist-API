package api

import (
	"net/http"
	"strconv"

	"github.com/floodedrealms/adventure-archivist/services"
	"github.com/floodedrealms/adventure-archivist/types"
	"github.com/floodedrealms/adventure-archivist/util"
)

type CharacterApi struct {
	characterService services.CharacterService
}

func NewCharacterApi(as services.CharacterService) *CharacterApi {
	return &CharacterApi{characterService: as}
}

func (c CharacterApi) CreateCharacterForCampaign(w http.ResponseWriter, r *http.Request) {
	campaignId, err := strconv.Atoi(r.PathValue("campaignId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	campaign := types.NewCampaign(campaignId)

	var characterToInsert types.CreateCharacterRecordRequest
	decodeJSONBody(w, r, &characterToInsert)

	created, err := c.characterService.CreateCharacterForCampaign(campaign, &characterToInsert)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sendGoodResponseWithObject(w, created)
}

func (c CharacterApi) UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("characterId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var characterToUpdate *types.UpdateCharacterRecordRequest
	created, err := c.characterService.UpdateCharacter(id, characterToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sendGoodResponseWithObject(w, created)
}
func (c CharacterApi) ManageCharactersForAdventure(w http.ResponseWriter, r *http.Request) {
	adventureId, err := strconv.Atoi(r.PathValue("advnetureId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	characterId, err := strconv.Atoi(r.PathValue("characterId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	params := r.URL.Query()

	operation := params.Get("operation")
	halfshare := params.Get("halfshare")

	adventure := types.NewAdventureRecordById(adventureId)
	character := types.NewCharacterById(characterId)
	status, err := c.characterService.ManageCharactersForAdventure(*adventure, character, operation, halfshare)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sendGoodResponseWithObject(w, status)
}

func (c CharacterApi) GetCharacterById(w http.ResponseWriter, r *http.Request) {
	http.Error(w, util.NotYetImplmented().Error(), http.StatusNotImplemented)
}
