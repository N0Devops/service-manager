package program

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type ProgramAction struct {
	program Program
}

func NewProgramAction(p Program) ProgramAction {
	return ProgramAction{
		program: p,
	}
}

func (pa *ProgramAction) ReadConfig(conf string) ([]byte, error) {
	config, ok := pa.program.Config[conf]
	if !ok {
		return nil, fmt.Errorf("config undefined")
	}
	return os.ReadFile(config)
}

func (pa *ProgramAction) WriteConfig(conf string, b []byte) error {
	config, ok := pa.program.Config[conf]
	if !ok {
		return fmt.Errorf("config undefined")
	}
	file, err := os.OpenFile(config, os.O_WRONLY|os.O_SYNC|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(b); err != nil {
		return err
	}
	return nil
}

func (pa *ProgramAction) Start() ([]byte, error) {
	args := strings.Split(pa.program.Operation.Start, " ")
	if len(args) == 0 {
		return nil, fmt.Errorf("action undefined")
	}
	cmd := exec.Command(args[0], args[1:]...)
	return cmd.CombinedOutput()
}

func (pa *ProgramAction) Stop() ([]byte, error) {
	args := strings.Split(pa.program.Operation.Stop, " ")
	if len(args) == 0 {
		return nil, fmt.Errorf("action undefined")
	}
	cmd := exec.Command(args[0], args[1:]...)
	return cmd.CombinedOutput()
}

func (pa *ProgramAction) Restart() ([]byte, error) {
	args := strings.Split(pa.program.Operation.Restart, " ")
	if len(args) == 0 {
		return nil, fmt.Errorf("action undefined")
	}
	cmd := exec.Command(args[0], args[1:]...)
	return cmd.CombinedOutput()
}

func (pa *ProgramAction) Status() ([]byte, error) {
	args := strings.Split(pa.program.Operation.Status, " ")
	if len(args) == 0 {
		return nil, fmt.Errorf("action undefined")
	}
	cmd := exec.Command(args[0], args[1:]...)
	return cmd.CombinedOutput()
}
