package order

type Raw string

func (s Raw) String() string {
	return string(s)
}

const SetNotNull = Raw(`SET NOT NULL`)

var _ Change = SetNotNull

const DropNotNull = Raw(`DROP NOT NULL`)

var _ Change = DropNotNull
