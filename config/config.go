package config

type Configurations struct {
	Wled WledConfigurations
}

type WledConfigurations struct {
	Url            string
	LevelZeroBody  string
	LevelOneBody   string
	LevelTwoBody   string
	LevelThreeBody string
	LevelFourBody  string
	LevelFiveBody  string
}
