package Trignometric

import (
        "math")
// 
func Trignometricidentity(x float64)[3]float64 {
	var result[3]float64;
	result[0]=math.Sin(x);
	result[1]=math.Cos(x);
	result[2]=math.Tan(x);
	return result;
}
// func Trig(){
	// fmt.Println("Enter a value to calculate trignometric identitities:")
	// var x float64
	// fmt.Scanln(&x)
	// Trignometricidentity(x);
// }

      
