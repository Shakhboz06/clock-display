package main

type placeholder [5]string

func digits() ([]placeholder, []placeholder) {

	zero := placeholder{
		"████",
		"█  █",
		"█  █",
		"█  █",
		"████",
	}

	one := placeholder{
		" ██ ",
		"  █ ",
		"  █ ",
		"  █ ",
		" ███",
	}

	two := placeholder{
		"████",
		"   █",
		"████",
		"█   ",
		"████",
	}

	three := placeholder{
		"████",
		"   █",
		"████",
		"   █",
		"████",
	}

	four := placeholder{

		"█  █",
		"█  █",
		"████",
		"   █",
		"   █",
	}

	five := placeholder{
		"████",
		"█   ",
		"████",
		"   █",
		"████",
	}

	six := placeholder{
		"████",
		"█   ",
		"████",
		"█  █",
		"████",
	}

	seven := placeholder{
		"████",
		"   █",
		"   █",
		"   █",
		"   █",
	}

	eight := placeholder{
		"████",
		"█  █",
		"████",
		"█  █",
		"████",
	}

	nine := placeholder{
		"████",
		"█  █",
		"████",
		"   █",
		"████",
	}


	seperator := placeholder{
		"   ",
		" ▓ ",
		"   ",
		" ▓ ",
		"   ",
	}
	dot := placeholder{
		"   ",
		"   ",
		"   ",
		"   ",
		" █ ",
	}
	
	alarm := placeholder{
		"  ██  	     █             ██        █████      ██    ██ ",
		" █  █ 	     █            █  █       █   █      █ █  █ █ ",
		"█    █	     █           █    █      █████      █  ██  █ ",
		"██████	     █           ██████      █  █       █      █ ",
		"█    █	     ██████      █    █      █   █      █      █ ",
	}

	features := []placeholder{
		seperator, alarm, dot,
	}
	digits := []placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	return features, digits
}
