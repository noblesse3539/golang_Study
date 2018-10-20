package tempconv

// CToF는 섭씨온도를 화씨온도로 변환한다.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC는 화씨온도를 섭씨온도로 변환한다.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK는 섭씨온도를 켈빈온도로 변환한다.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToC는 켈빈온도를 섭씨온도로 변환한다.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
