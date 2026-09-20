package main

import "fmt"

type BloodBag struct {
	UnitID         string
	BloodGroup     string
	VolumeML       int
	DonorAge       int
	CollectionDate string
}

var bloodBags []BloodBag

func main() {

	for {

		fmt.Println("\nBlood Bank")
		fmt.Println("1. Collect blood unit")
		fmt.Println("2. Search blood unit")
		fmt.Println("3. Update blood volume")
		fmt.Println("4. Delete blood unit")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter option: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			collectFromUser()
		case 2:
			var unitID string
			fmt.Print("Enter unit ID: ")
			fmt.Scanln(&unitID)

			unit := SearchUnit(unitID)

			if unit == nil {
				fmt.Println("Blood unit not found")
			} else {
				fmt.Printf("Blood unit is found: %+v\n", *unit)
			}
		case 3:
			var unitID string
			var volume int
			fmt.Print("Enter unit ID: ")
			fmt.Scanln(&unitID)

			fmt.Print("Enter volume in  ML: ")
			fmt.Scanln(&volume)
			fmt.Println("Volume updated successfully")
			fmt.Println(UpdateVolume(unitID, volume))
		case 4:
			var unitID string

			fmt.Print("Enter unit ID: ")
			fmt.Scanln(&unitID)

			fmt.Println(DiscardUnit(unitID))
		case 5:
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func collectFromUser() {

	var b BloodBag

	fmt.Print("Enter unit ID: ")
	fmt.Scanln(&b.UnitID)

	fmt.Print("Enter blood group: ")
	fmt.Scanln(&b.BloodGroup)

	fmt.Print("Enter collection date: ")
	fmt.Scanln(&b.CollectionDate)

	fmt.Print("Enter volume in ML: ")
	fmt.Scanln(&b.VolumeML)

	fmt.Print("Enter donor age: ")
	fmt.Scanln(&b.DonorAge)

	fmt.Println(CollectUnit(b))
}

func CollectUnit(b BloodBag) string {

	if b.DonorAge < 18 {
		return "Invalid age"
	}

	if b.BloodGroup != "A+" &&
		b.BloodGroup != "A-" &&
		b.BloodGroup != "B+" &&
		b.BloodGroup != "B-" &&
		b.BloodGroup != "AB+" &&
		b.BloodGroup != "AB-" &&
		b.BloodGroup != "O+" &&
		b.BloodGroup != "O-" {
		return "Invalid blood group"
	}

	bloodBags = append(bloodBags, b)

	return "Blood bag collected"
}

func SearchUnit(unitID string) *BloodBag {

	for i := 0; i < len(bloodBags); i++ {

		if bloodBags[i].UnitID == unitID {
			return &bloodBags[i]
		}
	}

	return nil
}

func UpdateVolume(unitID string, volume int) string {

	for i := 0; i < len(bloodBags); i++ {

		if bloodBags[i].UnitID == unitID {
			bloodBags[i].VolumeML = volume
			return "Volume updated"
		}
	}

	return "Blood unit not found"
}

func DiscardUnit(unitID string) string {

    newList := []BloodBag{}

    for i := 0; i < len(bloodBags); i++ {

        if bloodBags[i].UnitID == unitID {
            continue
        }

        newList = append(newList, bloodBags[i])
    }

    bloodBags = newList

    return "Blood bag discarded"
}