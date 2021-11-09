package mediator

type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}

func main() {
	stationManager := newStationManger()

	// 乗客列車
	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}

	// 貨物列車
	freightTrain := &freightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
