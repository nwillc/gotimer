package typeface

var Medium = make(Font)

func init() {
	Medium['0'] = FontRune{
		{true, true, true, false},
		{true, false, false, true},
		{true, false, false, true},
		{true, false, false, true},
		{true, false, false, true},
		{false, true, true, true},
	}
	Medium['1'] = FontRune{
		{false, false, true, false},
		{false, true, true, false},
		{false, false, true, false},
		{false, false, true, false},
		{false, false, true, false},
		{false, false, true, false},
	}
	Medium['2'] = FontRune{
		{true, true, true, false},
		{false, false, false, true},
		{false, false, true, false},
		{false, true, false, false},
		{true, false, false, false},
		{true, true, true, true},
	}
	Medium['3'] = FontRune{
		{true, true, true, true},
		{false, false, false, true},
		{false, true, true, true},
		{false, false, false, true},
		{false, false, false, true},
		{true, true, true, true},
	}
	Medium['4'] = FontRune{
		{false, false, true, true},
		{false, true, false, true},
		{true, false, false, true},
		{true, true, true, true},
		{false, false, false, true},
		{false, false, false, true},
	}
	Medium['5'] = FontRune{
		{true, true, true, true},
		{true, false, false, false},
		{true, true, true, false},
		{false, false, false, true},
		{false, false, false, true},
		{true, true, true, true},
	}
	Medium['6'] = FontRune{
		{true, true, true, true},
		{true, false, false, false},
		{true, true, true, false},
		{true, false, false, true},
		{true, false, false, true},
		{true, true, true, true},
	}
	Medium['7'] = FontRune{
		{true, true, true, true},
		{false, false, false, true},
		{false, false, true, false},
		{false, true, false, false},
		{true, false, false, false},
		{true, false, false, false},
	}
	Medium['8'] = FontRune{
		{false, true, true, false},
		{true, false, false, true},
		{true, false, false, true},
		{false, true, true, false},
		{true, false, false, true},
		{true, true, true, true},
	}
	Medium['9'] = FontRune{
		{false, true, true, false},
		{true, false, false, true},
		{true, true, true, true},
		{false, false, false, true},
		{false, false, true, false},
		{true, true, false, false},
	}
	Medium[':'] = FontRune{
		{false, false, false},
		{false, true, false},
		{false, true, false},
		{false, false, false},
		{false, true, false},
		{false, true, false},
	}
	Medium['.'] = FontRune{
		{false, false, false},
		{false, false, false},
		{false, false, false},
		{false, false, false},
		{false, true, false},
		{false, true, false},
	}
}
