package nlog

import "log/slog"

const LevelFatal = slog.Level(12)

var LevelNames = map[slog.Leveler]string{
	LevelFatal: "FATAL",
}
