package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)


func main() {
  var name string
  var score int = 10
  reader:= bufio.NewReader(os.Stdin)

  fmt.Printf("\n \n \n Welcome to the ultimate quiz game! \n \n \n")

  fmt.Printf("Enter your name : ")

  name, _ = reader.ReadString('\n')

  fmt.Printf("\nEnter your age : ")
  
  _age,_ := reader.ReadString('\n')

	age, err := strconv.Atoi(strings.TrimSpace(_age))


  if(err != nil){
    fmt.Printf("entered value is not valid")
    return
  }
  
  if(age < 13){
    fmt.Printf("%s, you are too young!\n", strings.TrimSpace(name))
    return
  }else{
    fmt.Printf("\n \n \n hey %s welcome to the game \n \n \n", strings.TrimSpace(name))
  }



  // Question 1

  fmt.Printf("1 - Which one is the smallest ocean in the World?  \n\n a) Indian c) Atlantic \n\n b) Pacific d) Arctic\n")

  answer , _ := reader.ReadString('\n')

  if(strings.TrimSpace(answer) == "d"){
    score++
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is correct")
  }else{
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is incorrect")
  }



  // Question 2

  fmt.Printf("2 - Which country gifted the ‘Statue of Liberty’ to USA in 1886?  \n\n a) France c) Brazil \n\n b) Canada d) England\n")

  answer , _ = reader.ReadString('\n')

  if(strings.TrimSpace(answer) == "a"){
    score++
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is correct")
  }else{
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is incorrect")
  }




  // Question 3

  fmt.Printf("3 - Dead Sea is located between which two countries?  \n\n a) Jordan and Sudan c) Turkey and UAE \n\n b) Jordan and Israel d) UAE and Egypt\n")

  answer , _ = reader.ReadString('\n')

  if(strings.TrimSpace(answer) == "b"){
    score++
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is correct")
  }else{
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is incorrect")
  }




  // Question 4

  fmt.Printf("4 - In which ocean ‘Bermuda Triangle’ region is located?  \n\n a) Atlantic c) Pacific \n\n b) Indian d) Arctic\n")

  answer , _ = reader.ReadString('\n')

  if(strings.TrimSpace(answer) == "a"){
    score++
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is correct")
  }else{
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is incorrect")
  }




  // Question 5

  fmt.Printf("5 - Which continent has the highest number of countries?  \n\n a) Asia c) Europe \n\n b) North America d) Africa\n")

  answer , _ = reader.ReadString('\n')

  if(strings.TrimSpace(answer) == "d"){
    score++
    fmt.Printf("the answer " + strings.TrimSpace(answer) + " is correct")
  }else{
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is incorrect")
  }




  // Question 6

  fmt.Printf("6 - Total number of oceans in the World is?  \n\n a) 3 c) 5 \n\n b) 7 d) 12\n")

  answer , _ = reader.ReadString('\n')

  if(strings.TrimSpace(answer) == "c"){
    score++
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is correct")
  }else{
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is incorrect")
  }




  // Question 7

  fmt.Printf("7 - The world’s longest straight road without any corners is located in?  \n\n a) USA c) Australia \n\n b) Saudi Arabia d) China\n")

  answer , _ = reader.ReadString('\n')

  if(strings.TrimSpace(answer) == "b"){
    score++
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is correct")
  }else{
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is incorrect")
  }


    // Question 8

    fmt.Printf("8 - Which one is the biggest island in the World?  \n\n a) Borneo c) Finland \n\n b) Sumatra d) Greenland\n")

    answer , _ = reader.ReadString('\n')
  
    if(strings.TrimSpace(answer) == "d"){
      score++
      fmt.Println("the answer " + strings.TrimSpace(answer) + " is correct")
    }else{
      fmt.Println("the answer " + strings.TrimSpace(answer) + " is incorrect")
    }



      // Question 9

  fmt.Printf("9 - In which year Hong Kong became a part of China after British Rule?  \n\n a) 1982 c) 1989 \n\n b) 1995 d) 1997\n")

  answer , _ = reader.ReadString('\n')

  if(strings.TrimSpace(answer) == "d"){
    score++
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is correct")
  }else{
    fmt.Println("the answer " + strings.TrimSpace(answer) + " is incorrect")
  }



    // Question 10

    fmt.Printf("10 - Which one is the World’s highest-altitude civilian airport?  \n\n a) Daocheng Yading Airport, China c) Kushok Bakula Rimpochhe Airport, Leh \n\n b) Qamdo Bamda Airport, China d) None of the above\n")

    answer , _ = reader.ReadString('\n')
  
    if(strings.TrimSpace(answer) == "a"){
      score++
      fmt.Println("the answer " + strings.TrimSpace(answer) + " is correct")
    }else{
      fmt.Println("the answer " + strings.TrimSpace(answer) + " is incorrect")
    }


    if(score > 7){
      fmt.Println("You win")
      fmt.Println("Do you want try again? [answer with 'Yes' or 'No']")
      answer , _ = reader.ReadString('\n')

      if(strings.TrimSpace(answer) == "Yes"){
        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()

        main()
      }else{
        return
      }
    }else{
      fmt.Println("Nice try try again")
      fmt.Println("Do you want try again? [answer with 'Yes' or 'No']")
      answer , _ = reader.ReadString('\n')

      if(strings.TrimSpace(answer) == "Yes"){
        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()
        
        main()
      }else{
        return
      }
    }


}



