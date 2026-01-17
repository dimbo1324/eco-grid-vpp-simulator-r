package interfaces

var Locations = map[int]LocationProfile{
	1: {
		ID:         1,
		City:       "New York",
		Region:     "New York State",
		Country:    "USA",
		Latitude:   40.7128,
		Longitude:  -74.0060,
		Timezone:   "America/New_York",
		BaseHeight: 10.0,
		/*
			TODO: подумать над данным параметром HubHeight, поскольку это параметр конструкционного характера, его нужно будет передавать через цифрового двойника!!!
		*/
		HubHeight: 100.0,
		Roughness: 0.3,
	},
}
