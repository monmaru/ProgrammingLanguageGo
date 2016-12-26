package weightconv

func PToKG(m Pound) Kilogram { return Kilogram(m * 0.45359237) }

func KGToP(f Kilogram) Pound { return Pound(f * 2.2046) }
