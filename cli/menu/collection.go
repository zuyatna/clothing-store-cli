package menu

import (
	"bufio"
	"clothing-pair-project/entity"
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
)

func ManageCollectionMenu(db *sqlx.DB) {
	for {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Manage Collection Menu")
		fmt.Println("1. Add Collection")
		fmt.Println("2. Find All Collections")
		fmt.Println("3. Update Collection")
		fmt.Println("4. Delete Collection")
		fmt.Println("0. Back")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		collectionHandler := handler.NewCollectionHandler(db)
		collectionService := service.NewCollectionService(collectionHandler)

		switch input {
		case 1:
			addCollectionMenu(collectionService)
		case 2:
			findAllCollectionsMenu(collectionService)
		case 3:
			updateCollectionMenu(collectionService)
		case 4:
			deleteCollectionMenu(collectionService)
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func addCollectionMenu(collectionService *service.CollectionService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Add Collection")
	fmt.Println("=====================================")

	var name string
	fmt.Print("Name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	collection := entity.Collection{
		Name: name,
	}

	err := collectionService.Add(collection)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Collection added successfully!")
}

func findAllCollectionsMenu(collectionService *service.CollectionService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("All Collections")
	fmt.Println("=====================================")

	allCollection(collectionService)

	var input string
	fmt.Print("0. Back: ")
	fmt.Scanln(&input)
	if input == "0" {
		return
	} else {
		fmt.Println("Invalid input")
	}
}

func allCollection(collectionService *service.CollectionService) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name"})

	collections, err := collectionService.FindAll()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collections) == 0 {
		fmt.Println("No collections found.")
		return
	}

	for _, collection := range collections {
		table.Append([]string{
			strconv.Itoa(collection.CollectionID),
			collection.Name,
		})
	}
	table.Render()

	fmt.Println()
}

func updateCollectionMenu(collectionService *service.CollectionService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Update Collection")
	fmt.Println("=====================================")

	allCollection(collectionService)

	var collectionID int
	fmt.Print("Collection ID: ")
	fmt.Scanln(&collectionID)

	collection, err := collectionService.FindByID(collectionID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("Name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	collection.Name = name

	err = collectionService.Update(collection)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Collection updated successfully!")
}

func deleteCollectionMenu(collectionService *service.CollectionService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Delete Collection")
	fmt.Println("=====================================")

	allCollection(collectionService)

	var collectionID int
	fmt.Print("Collection ID: ")
	fmt.Scanln(&collectionID)

	err := collectionService.Delete(collectionID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Collection deleted successfully!")
}
