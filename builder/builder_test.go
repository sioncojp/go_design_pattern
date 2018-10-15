package builder

import "testing"

func TestBuilderMethod(t *testing.T) {
	familyCar := New().Wheel(NormalWheel).TopSpeed(100).Paint(BLUE).Build()
	familyCar.Drive()

	sportsCar := New().Wheel(SportsWheel).TopSpeed(300).Paint(RED).Build()
	sportsCar.Drive()
}