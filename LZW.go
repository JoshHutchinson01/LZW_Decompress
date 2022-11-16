package main

import (
    "fmt"
    "io"
    "archive/zip"
    "strings"
    "log"
    
    //my module
    "lf"
    )

/*This function takes the contents of compressed file and outputs its 
 *decompressed content, this enables some library functionality*/
func ProcessFile(content []byte) (string, error) {
    
    //converts the 8-bit output to one long binary string
    new_content := lf.ConvertToBinaryString(content)
        
    //Holds the dictionary of characters which do not have ASCII codes.
    characters := make(map[int64]string)
    
    //Creates a list to store 12-bit binary strings.
    var bin_list []string = make([]string, 0)
    
    //Splits the binary string, new_content, up into 12-bit strings and 
    //stores in the list, bin_list.
    bin_list = lf.SplitString(new_content, bin_list,12)
    
    //initialise some variables
    prev := bin_list[0]
    var hold string
    var output strings.Builder //This type is used to avoid increasing buffer
    output.Grow(10000000) //This number can be modified based on file size.
    //
    var count int64 = 256
    dict_entry := false
     
    for v, entry := range bin_list {
            
        next_char := lf.GetCharacter(entry, characters)
        
        //If we have not already seen the character GetCharacter returns "",  
        //we then work out what the character should be.
        if  next_char == "" {
            new_char := lf.GetCharacter(prev, characters)
            new_char = new_char + hold
                
            characters[int64(count)] = new_char
            count++
                
            output.WriteString(new_char)
            var hold_list []string = make([]string, 0)
            hold = lf.SplitString(next_char, hold_list, 1)[0]
                
        } else {
            output.WriteString(next_char)
                
            //want this to run every time except when our dictionary 
            //is empty. Adds a new character to the dictionary which is 
            //our previous character plus the first symbol of the new one.               
            if dict_entry {
                var hold_list []string = make([]string, 0)
                hold = lf.SplitString(next_char, hold_list, 1)[0]
                new_char := lf.GetCharacter(prev, characters) + hold
                    
                characters[int64(count)] = new_char
                count++
                    
            }
        }
            
        //sets the previous character code to the current one for the
        //next loop round.
        prev = bin_list[v]
        dict_entry = true
            
        //reset the dictionary of characters if we have used them all up
        if count==4095 {
                
            for k := range characters {
                delete(characters, k)
            }
                
            //re initialise these variables
            count = 256
            dict_entry = false
                
        }
    }
    return output.String(), nil
}
        
func main() {
    
    //opens zip folder - choose the name of your desired file
    archive, err := zip.OpenReader("LzwInputData.zip")
    //code appearing multiple times checks for error and quits if one is found
    if err != nil {
        log.Fatal("Error Found check log for details")
        return 
    }
    
    for _, file := range archive.File {
        
        //reads content of an individual file in zip folder
        file_read, err := file.Open()
        if err != nil {
        log.Fatal("Error Found check log for details")
        return 
        }
        content, err := io.ReadAll(file_read)
        if err != nil {
        log.Fatal("Error Found check log for details")
        return 
        }
        
        output, err := ProcessFile(content)
        if err != nil {
        log.Fatal("Error Found check log for details")
        return 
        }
        
        fmt.Println(output)
        fmt.Println("Press the Enter Key to continue")
        fmt.Scanln() // wait for Enter Key
        
    }
    
}
