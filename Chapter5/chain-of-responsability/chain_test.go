package chainofresponsability

import (
	"strings"
	"testing"
)

type myTestWriter struct {
	receivedMessage string
}

func (m *myTestWriter) Write(p []byte) (int, error) {
	m.receivedMessage += string(p)
	return len(p), nil
}

func (m *myTestWriter) Next(s string) {
	m.Write([]byte(s))
}

func TestCreateDefaultChain(t *testing.T) {
	//Nosso teste ChainLogger
	myWriter := myTestWriter{}

	writerLogger := WriterLogger{Writer: &myWriter}
	second := SecondLogger{NextChain: &writerLogger}
	chain := FirstLogger{NextChain: &second}

	t.Run("3 loggers, 2 of them writes to console, second only if it founds"+"'hello'", func(t *testing.T) {
		chain.Next("message that breaks the chain\n")

		if myWriter.receivedMessage != "" {
			t.Fatal("Last link should not receive any message")
		}

		chain.Next("Hello\n")

		if !strings.Contains(myWriter.receivedMessage, "Hello") {
			t.Fatal("Last link didn't received expected message")
		}
	})
}
