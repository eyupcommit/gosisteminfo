package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kardianos/service"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	// #Start
	go p.run()
	return nil
}

func (p *program) run() {
	path := "C:/os.info"
	cf := "systeminfo"

	command := exec.Command(cf)

	stdout, err := command.Output()

	if err != nil {
		fmt.Println("Çıktı hata")
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(string(stdout))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "byte başarıyla yazıldı.")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
func (p *program) Stop(s service.Service) error {
	// #Stop
	return nil
}
func main() {
	// Main fonksiyonu ile uygulamamız çalıştırılır.
	svcConfig := &service.Config{
		Name: "Go Servis",
	}
	prg := &program{}
	s, _ := service.New(prg, svcConfig)
	_ = s.Run()
	s.Install()
}
