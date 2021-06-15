package main

import "fmt"

type Designation struct {
   EmpId int
   EmpName string
   Profile string
   Experience int
}
func main() {
   var designation1 Designation   
   var designation2 Designation   
 
   /* designation 1 specification */
   designation1.EmpId = 1001
   designation1.EmpName = "Rajnish Kumar"
   designation1.Profile = "Java"
   designation1.Experience = 3

   /* designation 2 specification */
   designation2.EmpId = 1002
   designation2.EmpName = "Manish Kumar"
   designation2.Profile = "Go"
   designation2.Experience = 1
 
   /* print designation1 info */
   fmt.Printf( "designation1.EmpId : %s\n", designation1.EmpId)
   fmt.Printf( "designation1.EmpName : %s\n", designation1.EmpName )
   fmt.Printf( "designation1.Profile : %s\n", designation1.Profile)
   fmt.Printf( "designation1.Experience  : %d\n",designation1.Experience )

   /* print designation2 info */
   fmt.Printf( "designation2.EmpId : %s\n", designation2.EmpId)
   fmt.Printf( "designation2.EmpName : %s\n", designation2.EmpName )
   fmt.Printf( "designation2.Profile : %s\n", designation2.Profile)
   fmt.Printf( "designation2.Experience  : %d\n",designation2.Experience )
}