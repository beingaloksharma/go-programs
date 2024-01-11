// package main

// import "fmt"

// type Data interface {
// 	Print()
// }

// type Employee struct {
// 	Name        string
// 	Designation string
// 	Address     string
// 	Asset       Asset
// }

// type Asset struct {
// 	Laptop string
// 	System Config
// }

// type Config struct {
// 	RAM       int
// 	Processor string
// }

// func (e *Employee) Print() {
// 	fmt.Printf("Name :: %s \nDesignation :: %s \nAddress :: %s \n", e.Name, e.Designation, e.Address)
// 	fmt.Printf("Assest Details \nLaptop Name :: %s \n", e.Asset.Laptop)
// 	fmt.Printf("System Details \nRAM :: %d \nProcessor :: %s \n", e.Asset.System.RAM, e.Asset.System.Processor)
// }
// func (a *Asset) Print() {
// 	fmt.Printf("Assest Details \nLaptop Name :: %s \nRAM :: %d \nProcessor :: %s\n ", a.Laptop, a.System.RAM, a.System.Processor)
// }

// func main() {
// 	fmt.Println("........ Start ........")
// 	data := Employee{
// 		Name:        "Alok Kumar",
// 		Designation: "Software Engineer",
// 		Address:     "New Delhi",
// 		Asset: Asset{
// 			Laptop: "Hp-Notebook",
// 			System: Config{
// 				RAM:       4,
// 				Processor: "inter 11th Generation",
// 			},
// 		},
// 	}
// 	data.Print()
// 	asset := Asset{
// 		Laptop: "Hp-Notebook",
// 		System: Config{
// 			RAM:       4,
// 			Processor: "inter 11th Generation",
// 		},
// 	}
// 	fmt.Print("************************************************************************")
// 	fmt.Println()
// 	asset.Print()
// 	fmt.Println("........ End ........")
// }
