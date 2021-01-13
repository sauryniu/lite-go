package osex

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"time"
)

const expiredText = "执行超时"

type osCommand struct {
	Cmd     *exec.Cmd
	Expires time.Duration
}

func (m osCommand) Exec() (string, string, error) {
	errReader, err := m.Cmd.StderrPipe()
	if err != nil {
		return "", "", err
	}
	defer errReader.Close()

	outReader, err := m.Cmd.StdoutPipe()
	if err != nil {
		return "", "", err
	}
	defer outReader.Close()

	if err = m.Cmd.Start(); err != nil {
		return "", "", err
	}

	return m.exec(errReader, outReader)
}

func (m *osCommand) SetDir(format string, args ...interface{}) ICommand {
	m.Cmd.Dir = fmt.Sprintf(format, args...)
	return m
}

func (m *osCommand) SetExpires(expires time.Duration) ICommand {
	m.Expires = expires
	return m
}

func (m osCommand) exec(errReader, outReader io.ReadCloser) (string, string, error) {
	errChan := make(chan error)
	go func() {
		errChan <- m.Cmd.Wait()
	}()

	stderrChan := make(chan string)
	m.scan(errReader, stderrChan)

	stdoutChan := make(chan string)
	m.scan(outReader, stdoutChan)

	select {
	case err := <-errChan:
		close(errChan)
		return "", "", err
	case <-time.After(m.Expires):
		return "", expiredText, nil
	case errStr := <-stderrChan:
		close(stderrChan)
		return "", errStr, nil
	case out := <-stdoutChan:
		close(stdoutChan)
		return out, "", nil
	}
}

func (m osCommand) scan(reader io.ReadCloser, msg chan string) {
	go func() {
		var bf bytes.Buffer
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			bf.WriteString(scanner.Text())
		}

		if bf.Len() == 0 {
			return
		}

		msg <- bf.String()
	}()
}

// NewOSCommand is 创建ICommand
func NewOSCommand(name string, args ...string) ICommand {
	return &osCommand{
		Cmd:     exec.Command(name, args...),
		Expires: 4 * time.Second,
	}
}
