package bmi

func CalculateBMI(height, weight float64) (float64, string) {
    heightM := height / 100
    bmi := weight / (heightM * heightM)
    
    var category string
    switch {
    case bmi < 16:
        category = "выраженный дефицит массы тела"
    case bmi < 18.5:
        category = "недостаточная масса тела"
    case bmi < 25:
        category = "нормальная масса тела"
    case bmi < 30:
        category = "избыточная масса тела (предожирение)"
    case bmi < 35:
        category = "ожирение I степени"
    case bmi < 40:
        category = "ожирение II степени"
    default:
        category = "ожирение III степени"
    }
    
    return bmi, category
}