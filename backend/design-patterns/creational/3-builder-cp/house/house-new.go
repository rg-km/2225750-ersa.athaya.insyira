package house

// Nah untuk builder sendiri kita harus menentukan builder nya itu yang mana. Anggap saja seperti Tukang yang ahli dalam membuat sesuatu
// Disini masih sama kita masih mengambil contoh untuk membuat sebuah Rumah
// Katakan terdapat 2 jenis builder (Tukang). Builder Zimbabwe dan Builder Indonesia

// Kita definisikan dulu interface(kontrak) general untuk builder

type House struct {
	NumOfWindows    int
	NumOfDoors      int
	HasGarage       bool
	HasSwimmingPool bool
}

type HouseBuilder interface {
	BuildHouseWithoutSwimmingPool() House
	getHouse() House
}

type ZimbabweHouseBuilder struct{}

func (z *ZimbabweHouseBuilder) getHouse() House {
	return House{
		NumOfWindows:    2,
		NumOfDoors:      1,
		HasGarage:       true,
		HasSwimmingPool: true,
	}
}

func (i *ZimbabweHouseBuilder) BuildHouseWithoutSwimmingPool() House {
	return House{
		NumOfWindows:    1,
		NumOfDoors:      1,
		HasGarage:       true,
		HasSwimmingPool: false,
	}
}

func NewBuilder(builderType string) HouseBuilder {
	return &ZimbabweHouseBuilder{}
}
