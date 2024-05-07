package sample

import (
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().Unix())
}
func randomProcessorBrand() string {
	return randomStringfromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	switch brand {
	case "Intel":
		return randomStringfromSet(
			"pentium 3500",
			"i3 3500h",
			"i5 4500p",
			"i7 4567e",
			"i9 6500p",
		)
	case "AMD":
		return randomStringfromSet(
			"ryzen 5 PRO 2500U",
			"ryzen 7 PRO 3566P",
			"ryzen 7 PRO 3577GE",
			"ryzen 9 PRO3200GE",
			"ryzen 6500 Threadripper",
		)
	default:
		return ""
	}
}

func randomGPUBrand() string {
	return randomStringfromSet("NVIDIA", "AMD", "Intel")
}

func randomGPUName(brand string) string {
	switch brand {
	case "NVIDIA":
		return randomStringfromSet(
			"MX 750ti",
			"GTX 10080",
			"GTX 1080ti",
			"GTX 2070ti",
			"GTX 3050ti",
			"GTX 4080ti",
		)
	case "AMD":
		return randomStringfromSet(
			"RX 580",
			"RX 750",
			"RX 7500-XT",
			"RX Vega-56",
		)
	case "Intel":
		return randomStringfromSet(
			"IRIS 750",
			"VISUAL 550",
			"IRIS 4500 vega",
		)
	default:
		return ""
	}
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomStringfromSet(set ...string) string {
	if len(set) == 0 {
		return ""
	}

	return set[rand.Intn(len(set))]
}

func randomLaptopBrand() string {
	return randomStringfromSet("Dell", "Acer", "Asus", "Apple")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Dell":
		return randomStringfromSet(
			"Vostro",
			"Alienware",
			"Latitude",
		)
	case "Acer":
		return randomStringfromSet(
			"Predator",
			"Aspire",
			"Nitro",
		)
	case "Asus":
		return randomStringfromSet(
			"TUF",
			"ROG",
			"Omen",
		)
	case "Apple":
		return randomStringfromSet(
			"MacBook Air",
			"MacBook Pro",
			"IPad Pro",
		)
	default:
		return ""
	}
}
