package logger

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestDefault(t *testing.T) {
	output := captureOutput(func() {
		Default("Test message")
	})

	assertEqual(t, "{\"message\":\"Test message\",\"severity\":\"DEFAULT\"}\n", output)
}

func TestDebug(t *testing.T) {
	output := captureOutput(func() {
		Debug("Test message")
	})

	assertEqual(t, "{\"message\":\"Test message\",\"severity\":\"DEBUG\"}\n", output)
}

func TestInfo(t *testing.T) {
	output := captureOutput(func() {
		Info("Test message")
	})

	assertEqual(t, "{\"message\":\"Test message\",\"severity\":\"INFO\"}\n", output)
}

func TestNotice(t *testing.T) {
	output := captureOutput(func() {
		Notice("Test message")
	})

	assertEqual(t, "{\"message\":\"Test message\",\"severity\":\"NOTICE\"}\n", output)
}

func TestWarning(t *testing.T) {
	output := captureOutput(func() {
		Warning("Test message")
	})

	assertEqual(t, "{\"message\":\"Test message\",\"severity\":\"WARNING\"}\n", output)
}

func TestError(t *testing.T) {
	output := captureOutput(func() {
		Error("Test message")
	})

	assertEqual(t, "{\"message\":\"Test message\",\"severity\":\"ERROR\"}\n", output)
}

func TestCritical(t *testing.T) {
	output := captureOutput(func() {
		Critical("Test message")
	})

	assertEqual(t, "{\"message\":\"Test message\",\"severity\":\"CRITICAL\"}\n", output)
}

func TestAlert(t *testing.T) {
	output := captureOutput(func() {
		Alert("Test message")
	})
	assertEqual(t, "{\"message\":\"Test message\",\"severity\":\"ALERT\"}\n", output)
}

func TestEmergency(t *testing.T) {
	output := captureOutput(func() {
		Emergency("Test message")
	})

	assertEqual(t, "{\"message\":\"Test message\",\"severity\":\"EMERGENCY\"}\n", output)
}
