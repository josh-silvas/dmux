package nlog

import "log/slog"

// LevelFatal : Fatal level
const LevelFatal = slog.Level(12)

// LevelNames : Map of slog.Leveler to string
var LevelNames = map[slog.Leveler]string{
	LevelFatal: "FATAL",
}
