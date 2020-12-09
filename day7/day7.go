package day7

type LuggageRules map[string]LuggageBag

type LuggageBag struct {
	Name   string
	Holds  map[string]int
	HeldBy []string
}

func (m LuggageRules) Bag(name string) *LuggageBag {
	bag, ok := m[name]
	if !ok {
		bag = NewLuggageBag(name)
		m[name] = bag
	}
	return &bag
}

func NewLuggageBag(name string) LuggageBag {
	return LuggageBag{Name: name, Holds: make(map[string]int), HeldBy: make([]string, 0, 20)}
}

func (m LuggageRules) bagsWhichCanContainHelper(bagName string) []string {
	bag := m.Bag(bagName)
	resultList := []string{bagName}
	for _, otherBag := range bag.HeldBy {
		resultList = append(resultList, m.bagsWhichCanContainHelper(otherBag)...)
	}
	return resultList
}

func (m LuggageRules) BagsWhichCanContain(bagName string) []string {
	bags := m.bagsWhichCanContainHelper(bagName)[1:]
	bagsCounter := make(map[string]bool)
	for _, bag := range bags {
		bagsCounter[bag] = true
	}
	uniqueBags := []string{}
	for bag := range bagsCounter {
		uniqueBags = append(uniqueBags, bag)
	}
	return uniqueBags
}

func (m LuggageRules) BagsWithin(bagName string) int {
	bag := m.Bag(bagName)
	sum := 0
	for otherBag, count := range bag.Holds {
		sum += count + count*m.BagsWithin(otherBag)
	}
	return sum
}
