package logger_test

import (
	"testing"

	"github.com/coneno/logger"
)

func TestErrorLevel(t *testing.T) {
	logger.SetLevel(logger.LEVEL_ERROR)

	t.Run("with error output", func(t *testing.T) {
		n, err := logger.Error.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 8 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with warning output", func(t *testing.T) {
		n, err := logger.Warning.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 0 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with info output", func(t *testing.T) {
		n, err := logger.Info.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 0 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with debug output", func(t *testing.T) {
		n, err := logger.Debug.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 0 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})
}

func TestWarningLevel(t *testing.T) {
	logger.SetLevel(logger.LEVEL_WARNING)

	t.Run("with error output", func(t *testing.T) {
		n, err := logger.Error.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 8 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with warning output", func(t *testing.T) {
		n, err := logger.Warning.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 8 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with info output", func(t *testing.T) {
		n, err := logger.Info.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 0 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with debug output", func(t *testing.T) {
		n, err := logger.Debug.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 0 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})
}

func TestInfoLevel(t *testing.T) {
	logger.SetLevel(logger.LEVEL_INFO)

	t.Run("with error output", func(t *testing.T) {
		n, err := logger.Error.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 8 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with warning output", func(t *testing.T) {
		n, err := logger.Warning.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 8 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with info output", func(t *testing.T) {
		n, err := logger.Info.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 8 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with debug output", func(t *testing.T) {
		n, err := logger.Debug.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 0 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})
}

func TestDebugLevel(t *testing.T) {
	logger.SetLevel(logger.LEVEL_DEBUG)

	t.Run("with error output", func(t *testing.T) {
		n, err := logger.Error.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 8 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with warning output", func(t *testing.T) {
		n, err := logger.Warning.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 8 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with info output", func(t *testing.T) {
		n, err := logger.Info.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 8 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})

	t.Run("with debug output", func(t *testing.T) {
		n, err := logger.Info.Writer().Write([]byte("testmsg\n"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != 8 {
			t.Errorf("unexpected number of bytes written: %d", n)
			return
		}
	})
}

func ExampleSetLevel() {
	logger.SetLevel(logger.LEVEL_DEBUG)
	logger.Error.Printf("error")
	logger.Warning.Printf("warning")
	logger.Debug.Printf("debug")
	logger.Info.Printf("info")
}
