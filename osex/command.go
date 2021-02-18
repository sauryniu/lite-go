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
	cmd     *exec.Cmd
	dirPath string
	expires time.Duration
}

func (m osCommand) Exec(name string, args ...string) (string, string, error) {
	m.cmd = exec.Command(name, args...)
	if m.dirPath != "" {
		m.cmd.Dir = m.dirPath
	}
	errReader, err := m.cmd.StderrPipe()
	if err != nil {
		return "", "", err
	}
	defer errReader.Close()

	outReader, err := m.cmd.StdoutPipe()
	if err != nil {
		return "", "", err
	}
	defer outReader.Close()

	if err = m.cmd.Start(); err != nil {
		return "", "", err
	}

	return m.exec(errReader, outReader)
}

func (m *osCommand) SetDir(format string, args ...interface{}) ICommand {
	m.dirPath = fmt.Sprintf(format, args...)
	return m
}

func (m *osCommand) SetExpires(expires time.Duration) ICommand {
	m.expires = expires
	return m
}

func (m osCommand) exec(errReader, outReader io.ReadCloser) (string, string, error) {
	errChan := make(chan error)
	go func() {
		errChan <- m.cmd.Wait()
	}()

	stderrChan := make(chan string)
	m.scan(errReader, stderrChan)

	stdoutChan := make(chan string)
	m.scan(outReader, stdoutChan)

	select {
	case err := <-errChan:
		close(errChan)
		return "", "", err
	case <-time.After(m.expires):
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

// NewCommand is 创建ICommand
func NewCommand() ICommand {
	return &osCommand{
		expires: 4 * time.Second,
	}
}
