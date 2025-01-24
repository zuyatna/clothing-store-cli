package feat

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/service"
	"log"
)

func AddCollection(collectionService service.CollectionService, name string) {
	addCollection := entity.Collection{
		Name: name,
	}

	err := collectionService.Add(addCollection)
	if err != nil {
		log.Fatal("Failed to add collection:", err.Error())
	}
	log.Println("Successfully added collection:", addCollection)
}

func FindAllCollections(collectionService service.CollectionService) {
	collections, err := collectionService.FindAll()
	if err != nil {
		log.Fatal("Failed to find all collections:", err.Error())
	}
	log.Println("Successfully found collections:", collections)
}

func FindCollectionByID(collectionService service.CollectionService, id int) {
	collection, err := collectionService.FindByID(id)
	if err != nil {
		log.Fatal("Failed to find collection by ID:", err.Error())
	}
	log.Println("Successfully found collection by ID:", collection)
}

func UpdateCollection(collectionService service.CollectionService, id int, name string) {
	updateCollection := entity.Collection{
		CollectionID: id,
		Name:         name,
	}
	err := collectionService.Update(updateCollection)
	if err != nil {
		log.Fatal("Failed to update collection:", err.Error())
	}
	log.Println("Successfully updated collection:", updateCollection)
}

func DeleteCollection(collectionService service.CollectionService, id int) {
	err := collectionService.Delete(id)
	if err != nil {
		log.Fatal("Failed to delete collection:", err.Error())
	}
	log.Println("Successfully deleted collection with ID:", id)
}
