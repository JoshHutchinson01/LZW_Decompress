package lf

import (
    "fmt"
    "strconv"
    "strings"
    )
    
/*Recieves a binary string and outputs its character in our dictionary if
 * it exists. If not it outputs an empty string.*/    
func GetCharacter(entry string, characters map[int64]string) string {
    character := ""
    if i,err := strconv.ParseInt(entry,2,64); err == nil{
        if i<256 {
            character = string(i)
        } else if v, found := characters[i]; found {
            character = v
        }
    }
    return character
}


/*Recieves a (potentially large) string and splits it to a list of n-character 
 *strings where n is given by chunk*/
func SplitString(user_bin string, bin_list []string, chunk int) []string {
    var length int = 0
    var count int = 0
    for i := range user_bin {
        if length == chunk {
            bin_list = append(bin_list, user_bin[count:i])
            length = 0
            count = i
        }
        length++
    }
    
    //if the length of the last string is 4 it concatenates it to the previous 
    //one. This is needed for the 12-bit conversion in the case of an 
    //ood number of characters
    if len(user_bin[count:]) == 4 {
        bin_list[len(bin_list)-1] += user_bin[count:]
    } else {  
    bin_list = append(bin_list, user_bin[count:])
    }
    
    return bin_list
}


/*Recieves a list of decimal numbers and converts each decimal number to binary
 *It then concatentates to a large binary string which is returned*/
func ConvertToBinaryString(content []byte) string {
    
    var new_list []string = make([]string, 0)
    
    for _, value := range content {
        n := int64(value)
        s := fmt.Sprintf("%08b", n)
        new_list = append(new_list, s)
    }
    
    new_content := strings.Join(new_list, "")
    return new_content  
}
