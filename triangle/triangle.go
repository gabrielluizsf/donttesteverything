package triangle

func CheckIsTriangle(x,y,z int)bool{return x + y > z && z + y > x && z + x > y;}