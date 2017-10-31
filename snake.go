package main

import (
	log "github.com/sirupsen/logrus"
)

type Snake struct {
	Name         string
	GameSettings GameSettings
}

func (snake Snake) GetNextMove(m Map) string {
	return "DOWN"
}

func (snake Snake) OnPlayerRegistered(s GameSettings) {
	snake.GameSettings = s
	log.Debug("Player registered.")
}

func (snake Snake) OnSnakeDead(reason string) {
	log.WithField("reason", reason).Debug("The snake died")
}

func (snake Snake) OnGameStarting() {
	log.Debug("All snakes are ready to rock. Game is starting.")
}

func (snake Snake) OnInvalidPlayername(reasonCode int) {
	log.Debug("Player name is invalid.")
}

func GetSnake() Snake {
	return Snake{
		Name: "golang",
	}
}
