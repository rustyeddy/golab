/*
logerr is a convenience wrapper around the excellent logrus and
Go error packages.  With some other useful stuff thrown in.

The code in here is mostly just boiler plate and glue wrapping
the above packages with configuration variables and logging patterns
that repeated app after app.

The goal is to keep these convenience function small and out of
the way, all while fully supporting a modern-microservice-12factor(ish) app.

*/
package logerr

import (
	"os"
	"strings"
	"time"

	"github.com/rustyeddy/logrus"
)

// Logerr wraps common error handling tasks
type Logerr struct {
	Name           string // May have multiple loggers ..
	*logrus.Logger        // Embedded logger gives us easy access to methods
}

const (
	None       = "none"
	Debug      = "debug"
	Test       = "test"
	Staging    = "staging"
	Production = "prod"
)

var (
	log *Logerr
)

func init() {
	log = NewLogger("default")
	log.RegisterExitHandler(GraceFullShutdown)
	log.Debugln("completed logerr init")
}

// LoggerDefaults sets the global
func NewLogerr(name string) *Logerr {
	ll := Logerr{
		Name:   name,
		Logger: log.New(),
	}
	return &ll
}

func LogDebug() *Logerr {
	ll := NewLogerr("debug")
	ll.SetFormatter(&L.TextFormatter)
	ll.SetLevel(L.LevelDebug)
	ll.SetOutput(os.Stdout)
	return ll
}

// LoggerDefaults sets the global
func LogTest() *Logerr {
	ll := NewLogerr("test")
	//ll.SetFormatter(&L.JSONFormatter)
	//ll.SetLevel(L.LevelWarn)
	//ll.SetOutput(os.Stdout) // TODO: change to a test log file
	return ll
}

// LoggerDefaults sets the global
func LogProduction() *Logerr {
	ll := NewLogerr("production")
	//L.SetFormatter(&L.JSONFormatter)
	//L.SetLevel(L.LevelWarn)
	//L.SetOutput(os.Stdout) // TODO: change to a test log file
	return ll
}

// OutputJSON it makes sense, so use it, really!
func (l *Logerr) SetFormatText() {
	//l.SetFormatter(&log.TextFormatter)
}

// OutputJSON it makes sense, so use it, really!
func (l *Logerr) SetFormatJSON() {
	//l.SetFormatter(&log.TextFormatter)
}

// #################### Fatal Handler ####################
func GraceFullShutdown() {
	L.Debug("We are gracefully shutting down")
}

// #################### Utility Functions ####################

// timeStamp returns a filename friendly timestamp.
func timeStamp() string {
	ts := time.Now().UTC().Format(time.RFC3339)
	return strings.Replace(ts, ":", "", -1) // get rid of offesnive colons
}
