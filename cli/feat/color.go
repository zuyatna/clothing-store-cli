package feat

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/service"
	"log"
)

func AddColor(colorService service.ColorService, name string) {
	addColor := entity.Color{
		Name: name,
	}

	err := colorService.Add(addColor)
	if err != nil {
		log.Fatal("Failed to add color:", err.Error())
	}
	log.Println("Successfully added color:", addColor)
}

func FindAllColors(colorService service.ColorService) {
	colors, err := colorService.FindAll()
	if err != nil {
		log.Fatal("Failed to find all colors:", err.Error())
	}
	log.Println("Successfully found colors:", colors)
}

func FindColorByID(colorService service.ColorService, id int) {
	color, err := colorService.FindByID(id)
	if err != nil {
		log.Fatal("Failed to find color by ID:", err.Error())
	}
	log.Println("Successfully found color by ID:", color)
}

func UpdateColor(colorService service.ColorService, id int, name string) {
	updateColor := entity.Color{
		ColorID: id,
		Name:    name,
	}
	err := colorService.Update(updateColor)
	if err != nil {
		log.Fatal("Failed to update color:", err.Error())
	}
	log.Println("Successfully updated color:", updateColor)
}

func DeleteColor(colorService service.ColorService, id int) {
	err := colorService.Delete(id)
	if err != nil {
		log.Fatal("Failed to delete color:", err.Error())
	}
	log.Println("Successfully deleted color with ID:", id)
}
