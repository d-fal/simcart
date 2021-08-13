package scaffold

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

func (s *skeleton) Close() {
	fmt.Println(aurora.Yellow("\n\tShutting Down .................\n").BgBlue())
	s.cancel()
	defer s.closer.Close()
	defer s.cancel()

	os.Exit(1)
}
