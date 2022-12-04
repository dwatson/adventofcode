package day03

type ElfGroup struct {
	Rucksacks []Rucksack
}

func (eg ElfGroup) Find() (rune, bool) {
	for _, item := range eg.Rucksacks[0].Data {
		if Contains(eg.Rucksacks[1].Data, item) && Contains(eg.Rucksacks[2].Data, item) {
			return item, true
		}
	}
	return 0, false
}

type Rucksack struct {
	Data        []rune
	Match       rune
	MatchValue  int
	Compartment []Compartment
}

func (r *Rucksack) Init(value string) {
	r.Data = []rune(value)
	length := len(r.Data)
	r.Compartment = append(r.Compartment, Compartment{Content: r.Data[0 : length/2]})
	r.Compartment = append(r.Compartment, Compartment{Content: r.Data[length/2 : length]})
}

func (r *Rucksack) Find() {
	for _, i := range r.Compartment[0].Content {
		for _, j := range r.Compartment[1].Content {
			if i == j {
				r.Match = i
				r.MatchValue = ItemValue(i)
			}
		}
	}
}

type Compartment struct {
	Content []rune
}

func ItemValue(item rune) int {
	var itemValue int
	if item >= 'a' && item <= 'z' {
		itemValue = int(item) - 96
	}
	if item >= 'A' && item <= 'Z' {
		itemValue = int(item) - 38
	}
	return itemValue
}

func Contains(content []rune, item rune) bool {
	for _, i := range content {
		if i == item {
			return true
		}
	}
	return false
}

func Day3_1(values []string) int {
	rucksacks := []Rucksack{}
	total := 0
	for _, v := range values {
		rs := Rucksack{}
		rs.Init(v)
		rs.Find()
		total += rs.MatchValue
		rucksacks = append(rucksacks, rs)
	}
	return total
}

func Day3_2(values []string) int {
	elves := []ElfGroup{}
	for i, v := range values {
		if i%3 == 0 {
			elves = append(elves, ElfGroup{})
		}
		rs := Rucksack{}
		rs.Init(v)
		elves[len(elves)-1].Rucksacks = append(elves[len(elves)-1].Rucksacks, rs)
	}

	total := 0
	for _, e := range elves {
		found, _ := e.Find()
		total += ItemValue(found)
	}

	return total
}
