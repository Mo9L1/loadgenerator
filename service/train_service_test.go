package service

import (
	"math/rand"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestTrainService_FullIntegration(t *testing.T) {
	cli, _ := GetAdminClient()
	var trainSvc TrainService = cli

	// Mock data
	MockedID := faker.UUIDHyphenated()
	MockedName := faker.Name()
	MockedEconomyClass := rand.Intn(3)
	MockedConfortClass := rand.Intn(2)
	MockedAverageSpeed := 250 + rand.Intn(20)
	// input
	trainType := TrainType{
		ID:           MockedID,
		Name:         MockedName,
		EconomyClass: MockedEconomyClass,
		ConfortClass: MockedConfortClass,
		AverageSpeed: MockedAverageSpeed,
	}

	// Create Test
	createResp, err := trainSvc.Create(&trainType)
	if err != nil {
		t.Errorf("Create request failed, err %s", err)
	}
	if createResp.Status != 1 {
		t.Errorf("Create failed: %s", createResp.Msg)
	}

	// Query Test
	resp, err := trainSvc.Query()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	t.Logf("Query returned results: %v", resp)

	// Query all
	allTrainTypes, err := trainSvc.Query()
	if err != nil {
		t.Errorf("Query all request failed, err %s", err)
	}
	if allTrainTypes.Status != 1 {
		t.Errorf("allTrainTypes.Status != 1")
	}
	if len(allTrainTypes.Data) == 0 {
		t.Errorf("Query all returned no results")
	}
	found := false
	for _, trainTypeElement := range allTrainTypes.Data {
		if trainTypeElement.Id == trainType.ID &&
			trainTypeElement.Name == trainType.Name &&
			trainTypeElement.AverageSpeed == trainType.AverageSpeed &&
			trainTypeElement.ConfortClass == trainType.ConfortClass &&
			trainTypeElement.EconomyClass == trainType.ConfortClass {
			found = true
		}
	}
	if !found {
		t.Errorf("Query all not get the corresponsing result, whcih means 'Creation Fails'")
	}

	// Test Update
	UpdatedAverageSpeed := 275 + rand.Intn(10)
	updateTrainType := TrainType{
		ID:           trainType.ID,
		Name:         trainType.Name,
		EconomyClass: trainType.EconomyClass,
		ConfortClass: trainType.ConfortClass,
		AverageSpeed: UpdatedAverageSpeed,
	}
	updateResp, err := trainSvc.Update(&updateTrainType)
	if err != nil {
		t.Errorf("Update request failed, err %s", err)
	}
	if updateResp.Status != 1 {
		t.Errorf("Update failed: %s", updateResp.Msg)
	}

	// Test Retrieve by ID
	retrieveResp, err := trainSvc.Retrieve(trainType.ID)
	if err != nil {
		t.Errorf("Retrieve request failed, err %s", err)
	}
	//if len(retrieveResp.Data) == 0 {
	//	t.Errorf("Retrieve returned no result")
	//}
	if retrieveResp.Data == nil {
		t.Errorf("Retrieve returned no result")
	}

	// Test Retrieve by Name
	retrieveByNameResp, err := trainSvc.RetrieveByName(trainType.Name)
	if err != nil {
		t.Errorf("Retrieve by name request failed, err %s", err)
	}
	if retrieveByNameResp == nil {
		t.Errorf("Retrieve by name returned no result")
	}

	// Test Retrieve by Names
	names := []string{"GaoTieOne", "GaoTieTwo", "DongCheOne"}
	retrieveByNamesResp, err := cli.RetrieveByNames(names)
	if err != nil {
		t.Errorf("Retrieve by names request failed, err %s", err)
	}
	if len(retrieveByNamesResp.Data) == 0 {
		t.Errorf("Retrieve by names returned no results")
	}

	// Test Delete
	//var deleteID string
	//if len(allTrainTypes.Data) > 0 {
	//	deleteID = allTrainTypes.Data[len(allTrainTypes.Data)-1].Id
	//} else {
	//	t.Errorf("Query all returned empty data")
	//}
	deleteResp, err := trainSvc.Delete(trainType.ID)
	if err != nil {
		t.Errorf("Delete request failed, err %s", err)
	}
	if deleteResp.Status != 1 {
		t.Errorf("Delete failed: %s", deleteResp.Msg)
	}

}
