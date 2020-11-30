package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signal.Ignore(syscall.SIGINT)  // SIGINTER = interrupt | ctrl - c
	signal.Ignore(syscall.SIGKILL) // kill'e ignore | killall
	signal.Ignore(syscall.SIGTERM) // terminated    | killall

	signal.Ignore(syscall.SIGTSTP) // Bu eleman , ctrl-z atılsa dahi arkaplanda çalışmaya devam ediyor . Müthiş birşey :)
	/*
		//signal.Ignore(syscall.SIGABRT)
		//signal.Ignore(syscall.SIGALRM)
		//signal.Ignore(syscall.SIGBUS)
		//signal.Ignore(syscall.SIGFPE)
		//signal.Ignore(syscall.SIGHUP)
		//signal.Ignore(syscall.SIGILL)
		//signal.Ignore(syscall.SIGINT)
		//signal.Ignore(syscall.SIGPIPE)
		//signal.Ignore(syscall.SIGPOLL)
		//signal.Ignore(syscall.SIGPROF) /*
		//signal.Ignore(syscall.SIGQUIT) /*
		//signal.Ignore(syscall.SIGSEGV) /*
		//signal.Ignore(syscall.SIGSTOP) /*
		//signal.Ignore(syscall.SIGSEGV) /*
		//signal.Ignore(syscall.SIGSYS) /*
		//signal.Ignore(syscall.SIGTRAP) /*
		signal.Ignore(syscall.SIGTSTP)
			signal.Ignore(syscall.SIGTTIN)
			signal.Ignore(syscall.SIGTTOU)
			signal.Ignore(syscall.SIGUSR1)
			signal.Ignore(syscall.SIGUSR2)
			signal.Ignore(syscall.SIGVTALRM)
			signal.Ignore(syscall.SIGXCPU)
			signal.Ignore(syscall.SIGXFSZ)*/

	for {
		fmt.Println("Merhaba Dünya")
		time.Sleep(time.Second)
	}
}

// Böylece Ctrl C 'yi (SIGINT - interrupt ) 'u ignore etmiş (görmezden gelmiş) bulunduk .

// En üstteki 3 eleman kill atmayı engelleyecektir , gerei kalanlar şov :)
