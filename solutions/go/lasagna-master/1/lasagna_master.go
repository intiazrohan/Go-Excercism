package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers [] string, timePerLayer int) int{
    if timePerLayer==0{
        timePerLayer=2
    }
    return len(layers)*timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers [] string) (noodles int,sauce float64){
    noodles=0
    sauce=0.0

    for i:=0; i<len(layers);i++{
        if layers[i] == "sauce"{
            sauce+=0.2
        }else if layers[i] == "noodles"{
            noodles+=50
        }
    }
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient (friendsList, myList [] string){
    length1:= len(friendsList)
    length2:= len(myList)
    myList[length2 -1] = friendsList[length1 -1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe (quantities []float64, portions int) []float64 {
    var res []float64
    for i:=0; i <len(quantities) ; i++{
        res = append(res ,(quantities[i] /2.0) * float64(portions))
    }
    return res
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
