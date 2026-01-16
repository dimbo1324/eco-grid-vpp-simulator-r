package interfaces

/*
TODO:
! Объединяет все данные, возможно, не понадобится
*/
type Location struct {
	locProf  LocationProfile
	envState EnvironmentState
	weather  WeatherScenario
}
