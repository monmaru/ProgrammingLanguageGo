package lengthconv

func MToF(m Meter) Feet { return Feet(m * 3.2808) }

func FToM(f Feet) Meter { return Meter(f * 0.3048) }
