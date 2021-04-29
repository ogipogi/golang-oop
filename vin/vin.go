package vin

func Manufacturer(vin string) string {

	if string(vin[2]) == "9" {
		return vin[:3] + vin[11:14]
	}
	return vin[:3]

	// Variante 2
	/*
		manufactorer := vin[:3]
		if string(vin[2]) == "9" {
			manufactorer += vin[11:14]
		}
		return manufactorer

	*/

}
