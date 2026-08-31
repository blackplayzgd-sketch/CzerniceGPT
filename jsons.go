package main

import (
	"encoding/json"
	"os"
)

func loadJokeConfig() []map[string]string {
	configPath := getNicePath("jokes.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config []map[string]string

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func loadInfoConfig() map[string]map[string]string {
	configPath := getNicePath("infos.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config map[string]map[string]string

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func loadStatsData() map[string][]float64 {
	listsPath := getNicePath("lists.json")

	data, err := os.ReadFile(listsPath)
	if err != nil {
		panic(err)
	}

	var statsData map[string][]float64

	err = json.Unmarshal(data, &statsData)
	if err != nil {
		panic(err)
	}
	return statsData
}

func saveStatsData(statsData map[string][]float64) {
	newData, err := json.MarshalIndent(statsData, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(getNicePath("lists.json"), newData, 0644)
	if err != nil {
		panic(err)
	}
}
