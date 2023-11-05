package noizy

type BackgroundNoiseList struct {
	Category string
	Sounds   []string
}

func GetNoiseList() []BackgroundNoiseList {
	list := []BackgroundNoiseList{
		{
			Category: "Noise Sounds",
			Sounds: []string{
				"Brown Noise",
				"Pink Noise",
				"White Noise",
			},
		},
		{
			Category: "Water Sounds",
			Sounds: []string{
				"Brook",
				"Creek",
				"Stream",
				"Close Waterfall",
				"Distant Waterfall",
			},
		},
		{
			Category: "Coastal & Ocean",
			Sounds: []string{
				"Calm Shore",
				"Shore",
				"Wild Shore",
				"Ocean Waves",
				"Large Waves",
			},
		},
		{
			Category: "Rain Sounds",
			Sounds: []string{
				"Rain Drops",
				"Pouring Rain",
			},
		},
		{
			Category: "Thunderstorms",
			Sounds: []string{
				"Distant Thunder",
				"Closer Thunder",
			},
		},
		{
			Category: "Types of Winds",
			Sounds: []string{
				"Coastal Wind",
				"Forest Wind",
				"Autumn Breeze",
			},
		},
		{
			Category: "Nature Sounds",
			Sounds: []string{
				"Birds",
				"Frogs",
				"Summer Night",
				"Heat Wave",
			},
		},
		{
			Category: "Social Settings",
			Sounds: []string{
				"Coffee House",
				"Cocktail Voices",
			},
		},
		{
			Category: "Relaxation",
			Sounds: []string{
				"Meditation Time",
				"Wind Chimes",
			},
		},
		{
			Category: "Other",
			Sounds: []string{
				"Bonfire",
				"Fan Noise",
			},
		},
	}

	return list
}
