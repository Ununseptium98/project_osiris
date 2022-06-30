package procexplorer

import (
	"fmt"
	"log"

	//"project_osiris/hashing"

	"github.com/shirou/gopsutil/process"
)

func Hello() {
	fmt.Printf("Hello!")

}
func Testing() {

	pid_table, err := process.Pids()
	if err != nil {
		log.Fatal(err)
	}

	for _, pid := range pid_table {

		fmt.Printf("PID : %d  => PATH : %s\n", pid, GetProcessExePath(pid))
		//fmt.Printf("\n HASH => %s", hashing.Testing(GetProcessExePath(pid)))
	}

}

func GetProcessExePath(PID int32) string {
	/*
		Takes a PID as input, returns the path to its exe
	*/

	proc, err := process.NewProcess(PID)
	if err != nil {
		log.Fatal("Couille avec le pid")
	}

	path, err := proc.Exe()

	if err != nil {
		//log.Fatal("Couille avec le path du proc")
		return "couille"

	}
	return path

}
